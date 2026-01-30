package services

import (
	"github.com/sword-fisher-fly/ai-alert/internal/ctx"
)

var (
	AiService InterAiService
)

func NewServices(ctx *ctx.Context) {
	AiService = newInterAiService(ctx)
}
