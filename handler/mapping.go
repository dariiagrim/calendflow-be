package handler

import (
	"CalendFlowBE/handler/dto"
	"CalendFlowBE/service"
)

func mapDtoChatbotGenerateReplyRequestToParams(ChatbotGenerateReplyRequest dto.ChatbotGenerateReplyRequest) service.ChatbotGenerateReplyParams {
	params := service.ChatbotGenerateReplyParams{
		Messages:        mapDtoChatbotMessagesToChatbotMessages(ChatbotGenerateReplyRequest.Messages),
		TodayEventsData: mapDtoChatbotEventDataToChatbotEventData(ChatbotGenerateReplyRequest.TodayEventsData),
		CalendarsData:   mapDtoChatbotCalendarDataToChatbotCalendarData(ChatbotGenerateReplyRequest.CalendarsData),
	}
	return params
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

func mapDtoChatbotEventDataToChatbotEventData(dtoChatbotEventData []dto.ChatbotEventData) []service.ChatbotEventData {
	var chatbotEventData []service.ChatbotEventData
	for _, dtoChatbotEvent := range dtoChatbotEventData {
		chatbotEvent := service.ChatbotEventData{
			Id:            dtoChatbotEvent.Id,
			CalendarId:    dtoChatbotEvent.CalendarId,
			UserProfileId: dtoChatbotEvent.UserProfileId,
			Title:         dtoChatbotEvent.Title,
			StartTime:     dtoChatbotEvent.StartTime,
			EndTime:       dtoChatbotEvent.EndTime,
		}
		chatbotEventData = append(chatbotEventData, chatbotEvent)
	}
	return chatbotEventData
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

func mapChatbotReplyToDtoChatbotGenerateReply(chatbotReply service.ChatbotReply) dto.ChatbotGenerateReplyResponse {
	return dto.ChatbotGenerateReplyResponse{
		Id:                        chatbotReply.Id,
		CalendarId:                chatbotReply.CalendarId,
		CalendarSummary:           chatbotReply.CalendarSummary,
		UserProfileId:             chatbotReply.UserProfileId,
		Title:                     chatbotReply.Title,
		StartTime:                 chatbotReply.StartTime,
		EndTime:                   chatbotReply.EndTime,
		Action:                    chatbotReply.Action,
		FurtherClarifyingQuestion: chatbotReply.FurtherClarifyingQuestion,
		EditFromDate:              chatbotReply.EditFromDate,
		ActionConfirmed:           chatbotReply.ActionConfirmed,
	}
}
