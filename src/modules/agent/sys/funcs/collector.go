package funcs

import (
	"nightingale-club203/src/common/dataobj"
	"nightingale-club203/src/modules/agent/core"
)

func CollectorMetrics() []*dataobj.MetricValue {
	return []*dataobj.MetricValue{
		core.GaugeValue("proc.agent.alive", 1),
	}
}
