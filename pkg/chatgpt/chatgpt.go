package chatgpt

type chatCompletionRequest struct {
	Model          string          `json:"model"`
	Temperature    float64         `json:"temperature"`
	TopP           float64         `json:"top_p"`
	Messages       []Message       `json:"messages"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type Message struct {
	Role    MessageAuthorRole `json:"role"`
	Content string            `json:"content"`
}

type apiResponse struct {
	Error *apiError `json:"error"`
}

type apiError struct {
	Message string `json:"message"`
}

type MessageAuthorRole string

const (
	MessageAuthorRoleSystem    MessageAuthorRole = "system"
	MessageAuthorRoleUser      MessageAuthorRole = "user"
	MessageAuthorRoleAssistant MessageAuthorRole = "assistant"
)
