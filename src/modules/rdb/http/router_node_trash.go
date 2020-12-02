package http

import (
	"github.com/gin-gonic/gin"
	"nightingale-club203/src/models"
)

func nodeTrashGets(c *gin.Context) {
	limit := queryInt(c, "limit", 20)
	query := queryStr(c, "query", "")

	total, err := models.NodeTrashTotal(query)
	dangerous(err)

	list, err := models.NodeTrashGets(query, limit, offset(c, limit))
	dangerous(err)

	renderData(c, gin.H{
		"list":  list,
		"total": total,
	}, nil)
}

func nodeTrashRecycle(c *gin.Context) {
	var f idsForm
	bind(c, &f)
	renderMessage(c, models.NodeTrashRecycle(f.Ids))
}
