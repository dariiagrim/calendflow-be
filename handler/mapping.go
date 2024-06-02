package handler

import (
	"dariiamoisol.com/CalendFlowBE/handler/dto"
	"dariiamoisol.com/CalendFlowBE/service"
)

func MapDtoChatbotGenerateReplyRequestToParams(ChatbotGenerateReplyRequest dto.ChatbotGenerateReplyRequest) service.ChatbotGenerateReplyParams {
	params := service.ChatbotGenerateReplyParams{
		Messages:          mapDtoChatbotMessagesToChatbotMessages(ChatbotGenerateReplyRequest.Messages),
		Events:            mapDtoChatbotEventsDataToChatbotEventsData(ChatbotGenerateReplyRequest.EventsData),
		CalendarsData:     mapDtoChatbotCalendarDataToChatbotCalendarData(ChatbotGenerateReplyRequest.CalendarsData),
		SelectedEventData: mapDtoChatbotEventDataToChatbotEventData(ChatbotGenerateReplyRequest.SelectedEventData),
		CurrentDate:       ChatbotGenerateReplyRequest.CurrentDate,
	}
	return params
}

func MapChatbotReplyToDtoChatbotGenerateReplyResponse(chatbotReply service.ChatbotReply) dto.ChatbotGenerateReplyResponse {
	return dto.ChatbotGenerateReplyResponse{
		EventId:                   chatbotReply.Id,
		CalendarId:                chatbotReply.CalendarId,
		Title:                     chatbotReply.Title,
		StartTime:                 chatbotReply.StartTime,
		EndTime:                   chatbotReply.EndTime,
		Action:                    chatbotReply.Action,
		FurtherClarifyingQuestion: chatbotReply.FurtherClarifyingQuestion,
		ChatbotResponse:           chatbotReply.ChatbotResponse,
	}
}

func mapDtoChatbotMessagesToChatbotMessages(dtoChatbotMessages []dto.ChatbotMessage) []service.ChatbotMessage {
	var chatbotMessages []service.ChatbotMessage
	for _, dtoChatbotMessage := range dtoChatbotMessages {
		chatbotMessage := service.ChatbotMessage{
			Content: dtoChatbotMessage.Content,
			IsBot:   dtoChatbotMessage.IsBot,
		}
		chatbotMessages = append(chatbotMessages, chatbotMessage)
	}
	return chatbotMessages
}

func mapDtoChatbotEventsDataToChatbotEventsData(dtoChatbotEventsData []dto.ChatbotEventData) []service.ChatbotEventData {
	var chatbotEventsData []service.ChatbotEventData
	for _, dtoChatbotEventData := range dtoChatbotEventsData {
		chatbotEventData := mapDtoChatbotEventDataToChatbotEventData(&dtoChatbotEventData)
		chatbotEventsData = append(chatbotEventsData, *chatbotEventData)
	}
	return chatbotEventsData
}

func mapDtoChatbotEventDataToChatbotEventData(dtoChatbotEventData *dto.ChatbotEventData) *service.ChatbotEventData {
	if dtoChatbotEventData == nil {
		return nil
	}

	return &service.ChatbotEventData{
		Id:         dtoChatbotEventData.Id,
		CalendarId: dtoChatbotEventData.CalendarId,
		Title:      dtoChatbotEventData.Title,
		StartTime:  dtoChatbotEventData.StartTime,
		EndTime:    dtoChatbotEventData.EndTime,
	}
}

func mapDtoChatbotCalendarDataToChatbotCalendarData(dtoChatbotCalendarData []dto.ChatbotCalendarData) []service.ChatbotCalendarData {
	var chatbotCalendarData []service.ChatbotCalendarData
	for _, dtoChatbotCalendar := range dtoChatbotCalendarData {
		chatbotCalendar := service.ChatbotCalendarData{
			Id:      dtoChatbotCalendar.CalendarId,
			Summary: dtoChatbotCalendar.CalendarSummary,
		}
		chatbotCalendarData = append(chatbotCalendarData, chatbotCalendar)
	}
	return chatbotCalendarData
}
