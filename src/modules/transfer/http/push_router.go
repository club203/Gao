package http

import (
	"nightingale-club203/src/common/dataobj"
	"nightingale-club203/src/modules/transfer/rpc"
	"nightingale-club203/src/toolkits/http/render"
	"nightingale-club203/src/toolkits/stats"

	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/errors"
)

func PushData(c *gin.Context) {
	if c.Request.ContentLength == 0 {
		render.Message(c, "blank body")
		return
	}

	recvMetricValues := make([]*dataobj.MetricValue, 0)
	errors.Dangerous(c.ShouldBindJSON(&recvMetricValues))

	errCount, errMsg := rpc.PushData(recvMetricValues)
	stats.Counter.Set("http.points.in.err", errCount)
	if errMsg != "" {
		render.Message(c, errMsg)
		return
	}

	render.Data(c, "ok", nil)
}
