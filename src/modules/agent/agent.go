package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nightingale-club203/src/common/loggeri"
	"nightingale-club203/src/modules/agent/cache"
	"nightingale-club203/src/modules/agent/config"
	"nightingale-club203/src/modules/agent/core"
	"nightingale-club203/src/modules/agent/http"
	"nightingale-club203/src/modules/agent/log/worker"
	"nightingale-club203/src/modules/agent/report"
	"nightingale-club203/src/modules/agent/statsd"
	"nightingale-club203/src/modules/agent/stra"
	"nightingale-club203/src/modules/agent/sys"
	"nightingale-club203/src/modules/agent/sys/funcs"
	"nightingale-club203/src/modules/agent/sys/plugins"
	"nightingale-club203/src/modules/agent/sys/ports"
	"nightingale-club203/src/modules/agent/sys/procs"
	"nightingale-club203/src/modules/agent/timer"
	"nightingale-club203/src/modules/agent/udp"
	"nightingale-club203/src/toolkits/stats"

	"github.com/toolkits/pkg/logger"
	"github.com/toolkits/pkg/runner"
)

var (
	vers *bool
	help *bool
	conf *string

	version = "No Version Provided"
)

func init() {
	vers = flag.Bool("v", false, "display the version.")
	help = flag.Bool("h", false, "print this help.")
	conf = flag.String("f", "", "specify configuration file.")
	flag.Parse()

	if *vers {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	runner.Init()
	fmt.Println("runner.cwd:", runner.Cwd)
	fmt.Println("runner.hostname:", runner.Hostname)
}

func main() {
	parseConf()

	loggeri.Init(config.Config.Logger)
	stats.Init("agent")

	if config.Config.Enable.Mon {
		monStart()
	}

	if config.Config.Enable.Job {
		jobStart()
	}

	if config.Config.Enable.Report {
		reportStart()
	}

	if config.Config.Enable.Metrics {

		// 初始化 statsd服务
		statsd.Start()

		// 开启 udp监听 和 udp数据包处理进程
		udp.Start()
	}

	core.InitRpcClients()
	http.Start()

	endingProc()
}

func reportStart() {
	if err := report.GatherBase(); err != nil {
		fmt.Println("gatherBase fail: ", err)
		os.Exit(1)
	}

	go report.LoopReport()
}

func jobStart() {
	go timer.Heartbeat()
}

func monStart() {
	sys.Init(config.Config.Sys)
	stra.Init()

	funcs.BuildMappers()
	funcs.Collect()

	//插件采集
	plugins.Detect()

	//进程采集
	procs.Detect()

	//端口采集
	ports.Detect()

	//初始化缓存，用作保存COUNTER类型数据
	cache.Init()

	//日志采集
	go worker.UpdateConfigsLoop()
	go worker.PusherStart()
	go worker.Zeroize()
}

func parseConf() {
	if err := config.Parse(); err != nil {
		fmt.Println("cannot parse configuration file:", err)
		os.Exit(1)
	}
}

func endingProc() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-c:
		fmt.Printf("stop signal caught, stopping... pid=%d\n", os.Getpid())
	}

	logger.Close()
	http.Shutdown()
	fmt.Println("portal stopped successfully")
}
