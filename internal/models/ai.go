package models

const (
	AI_URL     = "https://api.minimaxi.com/v1/chat/completions"
	AI_API_KEY = "sk-cp-A89utIGsWsKMutmVNAjvlLtjxRkjF4b4iVBODiEqi9_O07izAlmmdNvxQrOxe7iB23gX6TWPE18GQ_D8qPSE2n8E_j4xj6VVz8lTis4tX5POM_QOsoZ1Uhw"
	AI_MODEL   = "MiniMax-M2"
)

type AiContentRecord struct {
	RuleId  string `json:"RuleId" form:"ruleId"`
	Content string `json:"content" form:"content"`
}

type AiConfig struct {
	// Enable *bool `json:"enable"`
	//Type      string `json:"type"` // OpenAi, DeepSeek
	Url       string `json:"url"`
	AppKey    string `json:"appKey"`
	Model     string `json:"model"`
	Timeout   int    `json:"timeout"`
	MaxTokens int    `json:"maxTokens"`
	Prompt    string `json:"prompt"`
}

var DefaultAiConfig = &AiConfig{
	Url:       AI_URL,
	AppKey:    AI_API_KEY,
	Model:     AI_MODEL,
	Timeout:   10000,
	MaxTokens: 10000,
	// TODO:
	Prompt: `
规则名称: {{RULE_NAME}}
`,
}

func (a AiContentRecord) TableName() string {
	return "ai_content_record"
}
