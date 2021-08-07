package source_code_group

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"github.com/tranHieuDev23/IdeTwo/controllers/proxies/faktory_proxy"
	"github.com/tranHieuDev23/IdeTwo/models/daos/execution_dao"
	"github.com/tranHieuDev23/IdeTwo/models/daos/source_code_dao"
	"github.com/tranHieuDev23/IdeTwo/models/execution"
	"github.com/tranHieuDev23/IdeTwo/models/source_code"
)

func SourceCodeGroup(base gin.RouterGroup) gin.RouterGroup {
	sourceDao := source_code_dao.GetInstance()
	executionDao := execution_dao.GetInstance()
	proxy := faktory_proxy.GetInstance()

	group := base.Group("/source_codes")
	{
		group.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			source := sourceDao.GetSourceCode(id)
			if source == nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}
			c.JSON(http.StatusOK, source)
		})

		group.POST("/", func(c *gin.Context) {
			type SourceCodeBase struct {
				Content  string                          `json:"content"`
				Language source_code.ProgrammingLanguage `json:"language"`
			}
			base := SourceCodeBase{}
			if err := c.ShouldBindJSON(&base); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			source := source_code.SourceCode{
				Id:       xid.New().String(),
				Content:  base.Content,
				Language: base.Language,
				Input:    "",
			}
			if _, err := govalidator.ValidateStruct(source); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			sourceDao.CreateSourceCode(source)
			c.Header("Location", fmt.Sprintf("%s%s", c.FullPath(), source.Id))
			c.JSON(http.StatusCreated, source)
		})

		group.PATCH("/:id", func(c *gin.Context) {
			id := c.Param("id")
			source := sourceDao.GetSourceCode(id)
			if source == nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}
			type SourceCodeBase struct {
				Language source_code.ProgrammingLanguage `json:"language" valid:"range(0|4)"`
				Content  string                          `json:"content" valid:"length(0|8192)"`
				Input    string                          `json:"input" valid:"length(0|8192),optional"`
			}
			base := SourceCodeBase{}
			if err := c.ShouldBindJSON(&base); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			source.Content = base.Content
			source.Language = base.Language
			source.Input = base.Input
			newSource := sourceDao.UpdateSourceCode(*source)
			if newSource == nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}
			c.JSON(http.StatusOK, newSource)
		})

		group.POST("/:id/execute", func(c *gin.Context) {
			id := c.Param("id")
			source := sourceDao.GetSourceCode(id)
			if source == nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}
			exec := execution.Execution{
				Id:             xid.New().String(),
				OfSourceCodeId: source.Id,
				Timestamp:      time.Now().UnixNano(),
				Status:         execution.NotExecuted,
			}
			executionDao.CreateExecution(exec)
			proxy.PushExecuteJob(exec.Id)
			c.JSON(http.StatusOK, exec)
		})
	}
	return *group
}
