package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/sword-fisher-fly/ai-alert/config"
	"github.com/sword-fisher-fly/ai-alert/internal/ctx"
	"github.com/sword-fisher-fly/ai-alert/internal/global"
	"github.com/sword-fisher-fly/ai-alert/internal/middleware"
	"github.com/sword-fisher-fly/ai-alert/internal/repo"
	v1 "github.com/sword-fisher-fly/ai-alert/internal/routers/v1"
	"github.com/sword-fisher-fly/ai-alert/internal/services"
	"github.com/zeromicro/go-zero/core/logc"
)

var Version string

func main() {
	global.Config = config.InitConfig()
	InitRoute()
}

func InitRoute() {
	logc.Info(context.Background(), "Launching...")

	mode := global.Config.Server.Mode
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	ginEngine := gin.New()
	ginEngine.MaxMultipartMemory = 10 << 20
	ginEngine.Use(
		middleware.Cors(),
		middleware.GinZapLogger(),
		gin.Recovery(),
		middleware.LoggingMiddleware(),
	)
	allRouter(ginEngine)

	global.Config = config.InitConfig()

	dbRepo := repo.NewRepoEntry()
	services.NewServices(ctx.NewContext(context.TODO(), dbRepo))
	err := ginEngine.Run(":" + global.Config.Server.Port)
	if err != nil {
		logc.Error(context.Background(), "Launch failed:", err)
		return
	}
}

func allRouter(engine *gin.Engine) {
	v1.Router(engine)
}
