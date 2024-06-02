package service

import (
	"dariiamoisol.com/CalendFlowBE/model"
	"dariiamoisol.com/CalendFlowBE/pkg/chatgpt"
)

func mapChatbotMessagesToChatGPTMessages(chatbotMessages []ChatbotMessage) []chatgpt.Message {
	var chatGPTMessages []chatgpt.Message
	for _, chatbotMessage := range chatbotMessages {
		role := chatgpt.MessageAuthorRoleUser
		if chatbotMessage.IsBot {
			role = chatgpt.MessageAuthorRoleAssistant
		}

		chatGPTMessage := chatgpt.Message{
			Content: chatbotMessage.Content,
			Role:    role,
		}
		chatGPTMessages = append(chatGPTMessages, chatGPTMessage)
	}
	return chatGPTMessages
}

func mapChatbotCalendarDataToModelChatbotPromptCalendars(chatbotCalendarData []ChatbotCalendarData) []model.ChatbotPromptCalendar {
	var modelChatbotPromptCalendars []model.ChatbotPromptCalendar
	for _, chatbotCalendar := range chatbotCalendarData {
		modelChatbotPromptCalendar := model.ChatbotPromptCalendar{
			Id:      chatbotCalendar.Id,
			Summary: chatbotCalendar.Summary,
		}
		modelChatbotPromptCalendars = append(modelChatbotPromptCalendars, modelChatbotPromptCalendar)
	}
	return modelChatbotPromptCalendars
}

func mapChatbotEventDataToModelChatbotPromptEvents(chatbotEventData []ChatbotEventData) []model.ChatbotPromptEvent {
	var modelChatbotPromptEvents []model.ChatbotPromptEvent
	for _, chatbotEvent := range chatbotEventData {
		modelChatbotPromptEvent := mapChatbotEventDataToModelChatbotPromptEvent(&chatbotEvent)
		modelChatbotPromptEvents = append(modelChatbotPromptEvents, *modelChatbotPromptEvent)
	}
	return modelChatbotPromptEvents
}

func mapChatbotEventDataToModelChatbotPromptEvent(chatbotEventData *ChatbotEventData) *model.ChatbotPromptEvent {
	if chatbotEventData == nil {
		return nil
	}

	return &model.ChatbotPromptEvent{
		Id:         chatbotEventData.Id,
		CalendarId: chatbotEventData.CalendarId,
		Title:      chatbotEventData.Title,
		StartTime:  chatbotEventData.StartTime,
		EndTime:    chatbotEventData.EndTime,
	}
}
