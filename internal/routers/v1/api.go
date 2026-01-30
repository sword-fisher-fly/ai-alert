package v1

import (
	"strings"

	"github.com/sword-fisher-fly/ai-alert/api"
	"github.com/sword-fisher-fly/ai-alert/internal/static"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	v1 := engine.Group("api/v1")
	{
		api.AiController.API(v1)
	}

	engine.StaticFS("/static", static.GetStaticFileSystem())

	engine.NoRoute(func(ctx *gin.Context) {
		urlPrefix := "/api/v1"
		if strings.HasPrefix(ctx.Request.URL.Path, urlPrefix) {
			ctx.Status(404)
			return
		}
		data, err := static.ReadFile("index.html")
		if err != nil {
			ctx.String(500, "index.html not found")
			return
		}

		html := string(data)

		// configScript := `<script>window.__BASE_URL__ = "` + urlPrefix + `"; window.__API_VERSION__ = "/api/v1";</script>`
		// html = strings.Replace(html, "</head>", configScript+"</head>", 1)

		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Data(200, "text/html; charset=utf-8", []byte(html))
	})
}
