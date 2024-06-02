package dto

import (
	"dariiamoisol.com/CalendFlowBE/service"
	"time"
)

type ChatbotGenerateReplyRequest struct {
	Messages          []ChatbotMessage
	EventsData        []ChatbotEventData
	CalendarsData     []ChatbotCalendarData
	SelectedEventData *ChatbotEventData
	CurrentDate       time.Time
}

type ChatbotMessage struct {
	Content string `json:"content"`
	IsBot   bool   `json:"isBot"`
}

type ChatbotEventData struct {
	Id         string    `json:"id"`
	CalendarId string    `json:"calendarId"`
	Title      string    `json:"title"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

type ChatbotCalendarData struct {
	CalendarId      string `json:"calendarId"`
	CalendarSummary string `json:"calendarSummary"`
}

type ChatbotGenerateReplyResponse struct {
	EventId                   string                      `json:"eventId,omitempty"`
	CalendarId                string                      `json:"calendarId,omitempty"`
	Title                     string                      `json:"title,omitempty"`
	StartTime                 *time.Time                  `json:"startTime,omitempty"`
	EndTime                   *time.Time                  `json:"endTime,omitempty"`
	Action                    service.ChatbotResultAction `json:"action,omitempty"`
	FurtherClarifyingQuestion string                      `json:"furtherClarifyingQuestion,omitempty"`
	ChatbotResponse           string                      `json:"chatbotResponse,omitempty"`
}
