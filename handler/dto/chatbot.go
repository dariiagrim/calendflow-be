package dto

import (
	"CalendFlowBE/service"
	"time"
)

type ChatbotGenerateReplyRequest struct {
	Messages        []ChatbotMessage
	TodayEventsData []ChatbotEventData
	CalendarsData   []ChatbotCalendarData
}

type ChatbotGenerateReplyResponse struct {
	Id                        string                      `json:"id,omitempty"`
	CalendarId                string                      `json:"calendarId,omitempty"`
	CalendarSummary           string                      `json:"calendarSummary,omitempty"`
	UserProfileId             string                      `json:"userProfileId,omitempty"`
	Title                     string                      `json:"title,omitempty"`
	StartTime                 *time.Time                  `json:"startTime,omitempty"`
	EndTime                   *time.Time                  `json:"endTime,omitempty"`
	Action                    service.ChatbotResultAction `json:"action,omitempty"`
	FurtherClarifyingQuestion string                      `json:"furtherClarifyingQuestion,omitempty"`
	EditFromDate              *time.Time                  `json:"editFromDate,omitempty"`
	ActionConfirmed           bool                        `json:"actionConfirmed"`
}

type ChatbotMessage struct {
	Content string `json:"content"`
	IsBot   bool   `json:"isBot"`
}

type ChatbotEventData struct {
	Id            string    `json:"id"`
	CalendarId    string    `json:"calendarId"`
	UserProfileId string    `json:"userProfileId"`
	Title         string    `json:"title"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}

type ChatbotCalendarData struct {
	CalendarId      string `json:"calendarId"`
	CalendarSummary string `json:"calendarSummary"`
}
