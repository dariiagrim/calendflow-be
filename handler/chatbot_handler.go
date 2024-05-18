package handler

import (
	"CalendFlowBE/handler/dto"
	"CalendFlowBE/pkg/chatgpt"
	"CalendFlowBE/service"
	"encoding/json"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"log"
	"net/http"
	"os"
)

func init() {
	functions.HTTP("GenerateReply", GenerateReply)
}

func GenerateReply(w http.ResponseWriter, r *http.Request) {
	chatGPTClient := chatgpt.NewClient(http.Client{}, os.Getenv("OPENAI_API_KEY"))
	chatbotService := service.NewChatbotService(chatGPTClient)

	var req dto.ChatbotGenerateReplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chatbotReply, err := chatbotService.GenerateReply(r.Context(), mapDtoChatbotGenerateReplyRequestToParams(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := mapChatbotReplyToDtoChatbotGenerateReply(*chatbotReply)

	respData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}
