package types

import "fmt"

type RequestAiChatContent struct {
	RuleName string `json:"ruleName"`
	RuleId   string `json:"ruleId"`
	SearchQL string `json:"searchQL"`
	Content  string `json:"content"`
	Deep     string `json:"deep"`
}

func (a RequestAiChatContent) ValidateParams() error {
	if a.Content == "" {
		return fmt.Errorf("alert event content is empty")
	}

	if a.RuleName == "" {
		return fmt.Errorf("rule name is empty")
	}

	if a.RuleId == "" {
		return fmt.Errorf("rule id is empty")
	}

	return nil
}
