package model

import (
	_ "embed"
	"testing"
	"time"
)

//go:embed files/expected_prompt.txt
var expectedPrompt string

func TestChatbotPrompt_String(t *testing.T) {
	chatbotCalendars := []ChatbotPromptCalendar{
		{
			Id:      "calendar-1-id",
			Summary: "Calendar 1 summary",
		},
		{
			Id:      "calendar-2-id",
			Summary: "Calendar 2 summary",
		},
		{
			Id:      "calendar-3-id",
			Summary: "",
		},
	}

	event1StartTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	event1EndTime, err := time.Parse(time.RFC3339, "2006-01-02T16:04:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	event2StartTime, err := time.Parse(time.RFC3339, "2006-01-03T17:04:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	event2EndTime, err := time.Parse(time.RFC3339, "2006-01-03T20:04:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	event3StartTime, err := time.Parse(time.RFC3339, "2006-01-01T08:04:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	event3EndTime, err := time.Parse(time.RFC3339, "2006-01-01T08:34:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	chatbotEvents := []ChatbotPromptEvent{
		{
			Id:         "event-1-id",
			CalendarId: "calendar-1-id",
			Title:      "Event 1 title",
			StartTime:  event1StartTime,
			EndTime:    event1EndTime,
		},
		{
			Id:         "event-2-id",
			CalendarId: "calendar-2-id",
			Title:      "Event 2 title",
			StartTime:  event2StartTime,
			EndTime:    event2EndTime,
		},
		{
			Id:         "event-3-id",
			CalendarId: "calendar-2-id",
			Title:      "Event 3 title",
			StartTime:  event3StartTime,
			EndTime:    event3EndTime,
		},
	}

	selectedEvent := &ChatbotPromptEvent{
		Id:         "event-3-id",
		CalendarId: "calendar-2-id",
		Title:      "Event 3 title",
		StartTime:  event3StartTime,
		EndTime:    event3EndTime,
	}

	currentTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+03:00")
	if err != nil {
		t.Fatal(err)
	}

	chatbotPrompt := NewChatbotPrompt(chatbotCalendars, chatbotEvents, selectedEvent, currentTime)

	actualPrompt, err := chatbotPrompt.String()
	if err != nil {
		t.Fatal(err)
	}

	if actualPrompt != expectedPrompt {
		t.Errorf("invalid prompt, expected = %s; actual %s", expectedPrompt, actualPrompt)
	}
}
