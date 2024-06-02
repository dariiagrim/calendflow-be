package model

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"
)

type ChatbotPrompt struct {
	Calendars     []ChatbotPromptCalendar
	Events        []ChatbotPromptEvent
	SelectedEvent *ChatbotPromptEvent
	CurrentTime   time.Time
}

type ChatbotPromptCalendar struct {
	Id      string `json:"id"`
	Summary string `json:"summary"`
}

type ChatbotPromptEvent struct {
	Id         string    `json:"id"`
	CalendarId string    `json:"calendarId"`
	Title      string    `json:"title"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

func NewChatbotPrompt(
	calendars []ChatbotPromptCalendar,
	events []ChatbotPromptEvent,
	selectedEvent *ChatbotPromptEvent,
	currentTime time.Time,
) *ChatbotPrompt {
	return &ChatbotPrompt{
		Calendars:     calendars,
		Events:        events,
		SelectedEvent: selectedEvent,
		CurrentTime:   currentTime,
	}
}

//go:embed files/prompt_template.txt
var promptTemplate string

func (p *ChatbotPrompt) String() (string, error) {
	marshaledCalendars, err := json.Marshal(p.Calendars)
	if err != nil {
		return "", err
	}

	marshaledEvents, err := json.Marshal(p.Events)
	if err != nil {
		return "", err
	}

	marshaledSelectedEvent, err := json.Marshal(p.SelectedEvent)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		promptTemplate,
		string(marshaledCalendars),
		string(marshaledEvents),
		string(marshaledSelectedEvent),
		p.CurrentTime.Format(time.RFC3339),
	), nil
}
