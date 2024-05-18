package service

import "dariiamoisol.com/CalendFlowBE/pkg/chatgpt"

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
