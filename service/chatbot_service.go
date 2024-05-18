package service

import (
	"context"
	"dariiamoisol.com/CalendFlowBE/pkg/chatgpt"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ChatbotGenerateReplyParams struct {
	Messages        []ChatbotMessage
	TodayEventsData []ChatbotEventData
	CalendarsData   []ChatbotCalendarData
}

type ChatbotMessage struct {
	Content string
	IsBot   bool
}

type ChatbotEventData struct {
	Id            string
	CalendarId    string
	UserProfileId string
	Title         string
	StartTime     time.Time
	EndTime       time.Time
}

type ChatbotCalendarData struct {
	Id      string
	Summary string
}

type ChatbotReply struct {
	Id                        string
	CalendarId                string
	CalendarSummary           string
	UserProfileId             string
	Title                     string
	StartTime                 *time.Time
	EndTime                   *time.Time
	Action                    ChatbotResultAction
	FurtherClarifyingQuestion string
	EditFromDate              *time.Time
	ActionConfirmed           bool
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
	Id                        string              `json:"id,omitempty"`
	CalendarId                string              `json:"calendarId,omitempty"`
	CalendarSummary           string              `json:"calendarSummary,omitempty"`
	UserProfileId             string              `json:"userProfileId,omitempty"`
	Title                     string              `json:"title,omitempty"`
	StartTime                 *time.Time          `json:"startTime,omitempty"`
	EndTime                   *time.Time          `json:"endTime,omitempty"`
	Action                    ChatbotResultAction `json:"action,omitempty"`
	FurtherClarifyingQuestion string              `json:"furtherClarifyingQuestion,omitempty"`
	EditFromDate              *time.Time          `json:"editFromDate,omitempty"`
	ActionConfirmed           bool                `json:"actionConfirmed"`
}

func (s *ChatbotService) GenerateReply(ctx context.Context, params ChatbotGenerateReplyParams) (*ChatbotReply, error) {
	todayEventsJSON, err := json.Marshal(params.TodayEventsData)
	if err != nil {
		return nil, err
	}

	calendarsJSON, err := json.Marshal(params.CalendarsData)
	if err != nil {
		return nil, err
	}

	prompt := "Determine required action by user to make in his Google Calendar according to messages. User today events = %s. User calendars = %s. Possible actions: edit or create event. If some information is unclear ask claryfying question. If you don't know exactly start or end time ask claryfying question. If edit event and you can't find id, ask user date of this event in claryfying questions.  Provide in the form of .json file. In the given .json file, include a fields id, calendarId, calendarSummary, userProfileId, title, startTime (timestamp), endTime(timestamp), action (edit/create), furtherClarifyingQuestion, editFromDate (timestamp), actionConfirmed. Exclude any other objects or arrays. All dates and timestamps in format \"2006-01-02T15:04:05Z07:00\". Do not send empty fields it is very important! if action edit specified or other fields should not be empty, otherwise ask clarifying questions. For create same applies, only id and editFromDate can be empty. Current time: %s. actionConfirmed should be true only if user explicitly confirmed the action after he was asked to confirm it \"Please confirm...\""
	prompt = fmt.Sprintf(prompt, string(todayEventsJSON), string(calendarsJSON), time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"))

	message, err := s.chatGPTClient.CreateChatCompletionWithJSONModeEnabled(
		ctx,
		mapChatbotMessagesToChatGPTMessages(params.Messages),
		prompt,
		"gpt-4-turbo",
	)
	if err != nil {
		return nil, err
	}

	chatGPTResponse, err := s.parseGptGenerateReply(message)
	if err != nil {
		return nil, err
	}

	return &ChatbotReply{
		Id:                        chatGPTResponse.Id,
		CalendarId:                chatGPTResponse.CalendarId,
		CalendarSummary:           chatGPTResponse.CalendarSummary,
		UserProfileId:             chatGPTResponse.UserProfileId,
		Title:                     chatGPTResponse.Title,
		StartTime:                 chatGPTResponse.StartTime,
		EndTime:                   chatGPTResponse.EndTime,
		Action:                    chatGPTResponse.Action,
		FurtherClarifyingQuestion: chatGPTResponse.FurtherClarifyingQuestion,
		EditFromDate:              chatGPTResponse.EditFromDate,
		ActionConfirmed:           chatGPTResponse.ActionConfirmed,
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
