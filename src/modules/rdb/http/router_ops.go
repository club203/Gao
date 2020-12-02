package http

import (
	"github.com/gin-gonic/gin"
	"nightingale-club203/src/modules/rdb/config"
)

func globalOpsGet(c *gin.Context) {
	renderData(c, config.GlobalOps, nil)
}

func localOpsGet(c *gin.Context) {
	renderData(c, config.LocalOps, nil)
}
