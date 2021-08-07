package execution_group

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tranHieuDev23/IdeTwo/models/daos/execution_dao"
)

func ExecutionGroup(base gin.RouterGroup) gin.RouterGroup {
	executionDao := execution_dao.GetInstance()

	group := base.Group("/executions")
	{
		group.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			exec := executionDao.GetExecution(id)
			if exec == nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}
			c.JSON(http.StatusOK, exec)
		})
	}
	return *group
}
