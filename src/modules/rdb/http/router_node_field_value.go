package http

import (
	"github.com/gin-gonic/gin"
	"nightingale-club203/src/models"
)

func nodeFieldGets(c *gin.Context) {
	lst, err := models.NodeFieldValueGets(urlParamInt64(c, "id"))
	renderData(c, lst, err)
}

func nodeFieldPuts(c *gin.Context) {
	var objs []models.NodeFieldValue
	bind(c, &objs)

	id := urlParamInt64(c, "id")
	node := Node(id)

	loginUser(c).CheckPermByNode(node, "rdb_node_modify")

	renderMessage(c, models.NodeFieldValuePuts(id, objs))
}
