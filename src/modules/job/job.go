package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"github.com/toolkits/pkg/logger"
	"github.com/toolkits/pkg/runner"

	"nightingale-club203/src/common/identity"
	"nightingale-club203/src/common/loggeri"
	"nightingale-club203/src/models"
	"nightingale-club203/src/modules/job/config"
	"nightingale-club203/src/modules/job/http"
	"nightingale-club203/src/modules/job/rpc"
	"nightingale-club203/src/modules/job/timer"
	"nightingale-club203/src/toolkits/i18n"
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

func checkIdentity() {
	ip, err := identity.GetIP()
	if err != nil {
		fmt.Println("cannot get ip:", err)
		os.Exit(1)
	}

	fmt.Println("ip:", ip)

	if ip == "127.0.0.1" {
		fmt.Println("identity: 127.0.0.1, cannot work")
		os.Exit(2)
	}
}

func main() {
	parseConf()

	loggeri.Init(config.Config.Logger)

	checkIdentity()

	// 初始化数据库和相关数据
	models.InitMySQL("rdb", "job")

	go timer.Heartbeat()
	go timer.Schedule()
	go timer.CleanLong()

	// 将task_host_doing表缓存到内存里，减少DB压力
	timer.CacheHostDoing()

	i18n.Init(config.Config.I18n)

	go rpc.Start()
	http.Start()

	endingProc()
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
	fmt.Println("process stopped successfully")
}
