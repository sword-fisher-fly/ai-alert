package ai

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/sword-fisher-fly/ai-alert/internal/models"
	"github.com/bytedance/sonic"

	"github.com/sword-fisher-fly/ai-alert/pkg/tools"
)

func NewAiClient(config *models.AiConfig) (AiClient, error) {
	var ai *AiConfig
	if config != nil {
		ai = &AiConfig{
			Url:       config.Url,
			ApiKey:    config.AppKey,
			MaxTokens: config.MaxTokens,
			Model:     config.Model,
			Timeout:   config.Timeout,
		}
	} else {
		ai = &AiConfig{
			Url:       models.DefaultAiConfig.Url,
			ApiKey:    models.DefaultAiConfig.AppKey,
			MaxTokens: models.DefaultAiConfig.MaxTokens,
			Model:     models.DefaultAiConfig.Model,
			Timeout:   models.DefaultAiConfig.Timeout,
		}
	}

	fmt.Println("---------------------------")
	fmt.Printf("Model config: %#v\n", ai)
	fmt.Println("---------------------------")

	err := ai.Check(context.Background())
	if err != nil {
		return nil, err
	}

	return ai, nil
}

func (o *AiConfig) ChatCompletion(_ context.Context, prompt string) (string, error) {
	fmt.Println("@@@@@@@@@@@@@@@  ChatCompletion  called  @@@@@@@@@@@@")

	// 构建请求参数
	reqParams := Request{
		Model: o.Model,
		Messages: []*Message{
			{
				Role:    "system",
				Content: "您是站点可靠性工程 (SRE) 可观测性监控专家、资深 DevOps 工程师、资深运维专家",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Stream:    false,
		MaxTokens: o.MaxTokens,
	}

	bodyBytes, _ := sonic.Marshal(reqParams)
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + o.ApiKey
	response, err := tools.Post(headers, o.Url, bytes.NewReader(bodyBytes), o.Timeout)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(response.Body)
		var errResp Response
		_ = sonic.Unmarshal(errorBody, &errResp)
		return "", fmt.Errorf("API request error: %d - %s", response.StatusCode, errResp.Error.Message)
	}

	var result Response
	err = tools.ParseReaderBody(response.Body, &result)
	if err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("invalid reponse")
	}

	return result.Choices[0].Message.Content, nil
}

func (o *AiConfig) StreamCompletion(ctx context.Context, prompt string) (<-chan string, error) {
	reqParams := Request{
		Model: o.Model,
		Messages: []*Message{
			{Role: "user", Content: prompt},
		},
		Stream:    true,
		MaxTokens: o.MaxTokens,
	}

	bodyBytes, _ := sonic.Marshal(reqParams)
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + o.ApiKey

	response, err := tools.Post(headers, o.Url, bytes.NewReader(bodyBytes), o.Timeout)
	if err != nil {
		return nil, fmt.Errorf("stream request failed: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(response.Body)
		var errResp Response
		_ = sonic.Unmarshal(errorBody, &errResp)
		return nil, fmt.Errorf("OpenAI API错误: %d - %s", response.StatusCode, errResp.Error.Message)
	}

	// create stream channel
	streamChan := make(chan string)

	go func() {
		defer close(streamChan)
		defer response.Body.Close()
		select {
		case <-ctx.Done():
			return
		default:
			scanner := bufio.NewScanner(response.Body)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "data: ") {
					content := strings.TrimPrefix(line, "data: ")
					if strings.TrimSpace(content) == "[DONE]" {
						continue
					}

					var chunk StreamChunk
					if err := sonic.Unmarshal([]byte(content), &chunk); err != nil {
						log.Printf("paser error: %v | content: %s", err, content)
						continue
					}

					// send the response content into the stream channel
					if len(chunk.Choices) > 0 {
						streamChan <- chunk.Choices[0].Delta.Content
					}
				}
			}
		}
	}()

	return streamChan, nil
}

func (o *AiConfig) Check(_ context.Context) error {
	if o.Url == "" || o.ApiKey == "" {
		return fmt.Errorf("model url or api key is empty")
	}

	if o.Timeout == 0 {
		return fmt.Errorf("model timeout is not set")
	}
	return nil
}
