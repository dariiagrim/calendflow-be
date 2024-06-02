package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	apiUrl      = "https://api.openai.com/v1/"
	temperature = 0
	topP        = 0.2
)

type Client struct {
	httpClient http.Client
	apiKey     string
}

func NewClient(client http.Client, apiKey string) *Client {
	return &Client{
		httpClient: client,
		apiKey:     apiKey,
	}
}

func (c *Client) CreateChatCompletionWithJSONModeEnabled(
	ctx context.Context,
	messagesToAdd []Message,
	prompt string,
	model string,
) (*Message, error) {
	return c.createChatCompletion(ctx, messagesToAdd, prompt, model, &ResponseFormat{Type: "json_object"})
}

func (c *Client) createChatCompletion(
	ctx context.Context,
	messagesToAdd []Message,
	prompt string,
	model string,
	responseFormat *ResponseFormat,
) (*Message, error) {
	respBytes, err := c.sendRequest(ctx, http.MethodPost, "chat/completions", chatCompletionRequest{
		Model:       model,
		Temperature: temperature,
		TopP:        topP,
		Messages: append([]Message{
			{
				Role:    MessageAuthorRoleSystem,
				Content: prompt,
			},
		}, messagesToAdd...),
		ResponseFormat: responseFormat,
	})
	if err != nil {
		return nil, err
	}

	var resp *ChatCompletionResponse
	if err = json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("no chat completion choices")
	}

	return &resp.Choices[0].Message, nil
}

func (c *Client) sendRequest(ctx context.Context, method string, path string, requestBody interface{}) ([]byte, error) {
	requestPayload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(requestPayload))

	req, err := c.createRequest(ctx, method, path, requestPayload)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return c.createResponse(resp)
}

func (c *Client) createRequest(ctx context.Context, method string, path string, requestPayload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%v%v", apiUrl, path), bytes.NewReader(requestPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	if method != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}

	return req.WithContext(ctx), nil
}

func (c *Client) createResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &apiResponse{}
	if err = json.Unmarshal(respBytes, response); err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(response.Error.Message)
	}

	return respBytes, nil
}
