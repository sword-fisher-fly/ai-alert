package api

import (
	"github.com/sword-fisher-fly/ai-alert/internal/middleware"
	"github.com/sword-fisher-fly/ai-alert/internal/services"
	"github.com/sword-fisher-fly/ai-alert/internal/types"
	"github.com/gin-gonic/gin"
)

type aiController struct{}

var AiController = new(aiController)

func (aiController aiController) API(gin *gin.RouterGroup) {
	a := gin.Group("ai")
	a.Use(
		middleware.Cors(),
		middleware.Auth(),
	)
	{
		a.POST("chat", aiController.Chat)
	}
}

func (aiController aiController) Chat(ctx *gin.Context) {
	r := new(types.RequestAiChatContent)
	r.Content = ctx.PostForm("content")
	r.RuleId = ctx.PostForm("rule_id")
	r.RuleName = ctx.PostForm("rule_name")
	r.Deep = ctx.PostForm("deep")
	r.SearchQL = ctx.PostForm("search_ql")

	Service(ctx, func() (interface{}, interface{}) {
		return services.AiService.Chat(r)
	})
}
