package services

import (
	"fmt"
	"strings"

	"github.com/sword-fisher-fly/ai-alert/internal/ctx"
	"github.com/sword-fisher-fly/ai-alert/internal/models"
	"github.com/sword-fisher-fly/ai-alert/internal/types"
	"github.com/sword-fisher-fly/ai-alert/pkg/ai"
	"gorm.io/gorm"
)

type (
	aiService struct {
		ctx      *ctx.Context
		aiClient ai.AiClient
	}

	InterAiService interface {
		Chat(req interface{}) (interface{}, interface{})
	}
)

func newInterAiService(ctx *ctx.Context) InterAiService {
	aiClient, err := ai.NewAiClient(nil)
	if err != nil {
		panic(err)
	}
	return &aiService{
		ctx:      ctx,
		aiClient: aiClient,
	}
}

func (a aiService) Chat(req interface{}) (interface{}, interface{}) {
	r := req.(*types.RequestAiChatContent)
	fmt.Println("##############################")
	fmt.Printf("AI request: %#v\n", r)
	fmt.Println("##############################")

	err := r.ValidateParams()
	if err != nil {
		return nil, err
	}

	aiClient := a.aiClient

	prompt := `
规则名称: {{ RuleName }}
告警内容: {{ Content }}
搜索QL: {{ SearchQL }}
}	
`
	prompt = strings.ReplaceAll(prompt, "{{ RuleName }}", r.RuleName)
	prompt = strings.ReplaceAll(prompt, "{{ Content }}", r.Content)
	prompt = strings.ReplaceAll(prompt, "{{ SearchQL }}", r.SearchQL)
	r.Content = prompt

	switch r.Deep {
	case "true":
		r.Content = fmt.Sprintf("注意, 请深度思考下面的问题!\n%s", r.Content)
		completion, err := aiClient.ChatCompletion(a.ctx.Ctx, r.Content)
		if err != nil {
			return "", err
		}
		err = a.ctx.DB.Ai().Update(models.AiContentRecord{
			RuleId:  r.RuleId,
			Content: completion,
		})
		if err != nil {
			return nil, err
		}

		return completion, nil

	default:
		data, exist, err := a.ctx.DB.Ai().Get(r.RuleId)
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", err
		}
		if exist {
			return data.Content, nil
		}

		completion, err := aiClient.ChatCompletion(a.ctx.Ctx, r.Content)
		if err != nil {
			return "", err
		}

		err = a.ctx.DB.Ai().Create(models.AiContentRecord{
			RuleId:  r.RuleId,
			Content: completion,
		})
		if err != nil {
			return nil, err
		}

		return completion, nil
	}
}
