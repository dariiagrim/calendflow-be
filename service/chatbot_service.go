package service

import (
	"context"
	"dariiamoisol.com/CalendFlowBE/model"
	"dariiamoisol.com/CalendFlowBE/pkg/chatgpt"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ChatbotGenerateReplyParams struct {
	Messages          []ChatbotMessage
	Events            []ChatbotEventData
	CalendarsData     []ChatbotCalendarData
	SelectedEventData *ChatbotEventData
	CurrentDate       time.Time
}

type ChatbotMessage struct {
	Content string
	IsBot   bool
}

type ChatbotEventData struct {
	Id         string
	CalendarId string
	Title      string
	StartTime  time.Time
	EndTime    time.Time
}

type ChatbotCalendarData struct {
	Id      string
	Summary string
}

type ChatbotReply struct {
	Id                        string
	CalendarId                string
	Title                     string
	StartTime                 *time.Time
	EndTime                   *time.Time
	Action                    ChatbotResultAction
	FurtherClarifyingQuestion string
	ChatbotResponse           string
}

type ChatbotResultAction string

const (
	ChatbotResultActionEdit   ChatbotResultAction = "edit"
	ChatbotResultActionCreate ChatbotResultAction = "create"
)

type ChatbotService struct {
	chatGPTClient *chatgpt.Client
}

func NewChatbotService(chatGPTClient *chatgpt.Client) *ChatbotService {
	return &ChatbotService{
		chatGPTClient: chatGPTClient,
	}
}

type ChatGPTResponse struct {
	EventId                   string              `json:"eventId,omitempty"`
	CalendarId                string              `json:"calendarId,omitempty"`
	Title                     string              `json:"title,omitempty"`
	EventStartTime            *time.Time          `json:"eventStartTime,omitempty"`
	EventEndTime              *time.Time          `json:"eventEndTime,omitempty"`
	Action                    ChatbotResultAction `json:"action,omitempty"`
	FurtherClarifyingQuestion string              `json:"furtherClarifyingQuestion,omitempty"`
	ChatbotResponse           string              `json:"chatbotResponse,omitempty"`
}

func (s *ChatbotService) GenerateReply(ctx context.Context, params ChatbotGenerateReplyParams) (*ChatbotReply, error) {
	prompt := model.NewChatbotPrompt(
		mapChatbotCalendarDataToModelChatbotPromptCalendars(params.CalendarsData),
		mapChatbotEventDataToModelChatbotPromptEvents(params.Events),
		mapChatbotEventDataToModelChatbotPromptEvent(params.SelectedEventData),
		params.CurrentDate,
	)

	promptMessage, err := prompt.String()
	if err != nil {
		return nil, err
	}

	message, err := s.chatGPTClient.CreateChatCompletionWithJSONModeEnabled(
		ctx,
		mapChatbotMessagesToChatGPTMessages(params.Messages),
		promptMessage,
		"gpt-4o",
	)
	if err != nil {
		return nil, err
	}

	chatGPTResponse, err := s.parseGptGenerateReply(message)
	if err != nil {
		return nil, err
	}

	return &ChatbotReply{
		Id:                        chatGPTResponse.EventId,
		CalendarId:                chatGPTResponse.CalendarId,
		Title:                     chatGPTResponse.Title,
		StartTime:                 chatGPTResponse.EventStartTime,
		EndTime:                   chatGPTResponse.EventEndTime,
		Action:                    chatGPTResponse.Action,
		FurtherClarifyingQuestion: chatGPTResponse.FurtherClarifyingQuestion,
		ChatbotResponse:           chatGPTResponse.ChatbotResponse,
	}, nil
}

func (s *ChatbotService) parseGptGenerateReply(
	m *chatgpt.Message,
) (*ChatGPTResponse, error) {
	chatGPTResponse := &ChatGPTResponse{}

	if err := json.Unmarshal([]byte(m.Content), chatGPTResponse); err != nil {
		log.Print(err)
		return nil, err
	}

	fmt.Println(m.Content)

	return chatGPTResponse, nil
}
