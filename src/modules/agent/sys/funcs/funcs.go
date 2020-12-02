package funcs

import (
	"log"

	"nightingale-club203/src/common/dataobj"
	"nightingale-club203/src/modules/agent/sys"
)

type FuncsAndInterval struct {
	Fs       []func() []*dataobj.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

func BuildMappers() {
	interval := sys.Config.Interval
	if sys.Config.Enable {
		log.Println("sys collect enable is true.")
		Mappers = []FuncsAndInterval{
			{
				Fs: []func() []*dataobj.MetricValue{
					CollectorMetrics,
					CpuMetrics,
					MemMetrics,
					NetMetrics,
					LoadAvgMetrics,
					IOStatsMetrics,
					NfMetrics,
					FsKernelMetrics,
					FsRWMetrics,
					ProcsNumMetrics,
					EntityNumMetrics,
					NtpOffsetMetrics,
					SocketStatSummaryMetrics,
					UdpMetrics,
					TcpMetrics,
				},
				Interval: interval,
			},
			{
				Fs: []func() []*dataobj.MetricValue{
					DeviceMetrics,
				},
				Interval: interval,
			},
		}
	} else {
		log.Println("sys collect enable is false.")
		Mappers = []FuncsAndInterval{
			{
				Fs: []func() []*dataobj.MetricValue{
					CollectorMetrics,
				},
				Interval: interval,
			},
		}
	}
}
