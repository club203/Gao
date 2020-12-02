package funcs

import (
	"github.com/toolkits/pkg/logger"
	"github.com/toolkits/pkg/nux"

	"nightingale-club203/src/common/dataobj"
	"nightingale-club203/src/modules/agent/core"
)

func LoadAvgMetrics() []*dataobj.MetricValue {
	load, err := nux.LoadAvg()
	if err != nil {
		logger.Error(err)
		return nil
	}

	return []*dataobj.MetricValue{
		core.GaugeValue("cpu.loadavg.1", load.Avg1min),
		core.GaugeValue("cpu.loadavg.5", load.Avg5min),
		core.GaugeValue("cpu.loadavg.15", load.Avg15min),
	}
}
