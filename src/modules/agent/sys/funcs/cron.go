package funcs

import (
	"time"

	"nightingale-club203/src/common/dataobj"
	"nightingale-club203/src/modules/agent/config"
	"nightingale-club203/src/modules/agent/core"
	"nightingale-club203/src/modules/agent/sys"
)

func Collect() {
	go PrepareCpuStat()
	go PrepareDiskStats()

	for _, v := range Mappers {
		for _, f := range v.Fs {
			go collect(int64(v.Interval), f)
		}
	}
}

func collect(sec int64, fn func() []*dataobj.MetricValue) {
	t := time.NewTicker(time.Second * time.Duration(sec))
	defer t.Stop()

	ignoreMetrics := sys.Config.IgnoreMetricsMap

	for {
		<-t.C

		metricValues := []*dataobj.MetricValue{}
		now := time.Now().Unix()

		items := fn()
		if items == nil || len(items) == 0 {
			continue
		}

		for _, item := range items {
			if _, exists := ignoreMetrics[item.Metric]; exists {
				continue
			}

			item.Step = sec
			item.Endpoint = config.Endpoint
			item.Timestamp = now
			metricValues = append(metricValues, item)
		}
		core.Push(metricValues)
	}
}
