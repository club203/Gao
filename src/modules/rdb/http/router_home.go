package http

import (
	"github.com/gin-gonic/gin"
	"nightingale-club203/src/modules/rdb/config"
)

func ldapUsed(c *gin.Context) {
	renderData(c, config.Config.LDAP.DefaultUse, nil)
}
