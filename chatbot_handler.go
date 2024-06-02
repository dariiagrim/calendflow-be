package CalendFlowBE

import (
	"dariiamoisol.com/CalendFlowBE/handler"
	"dariiamoisol.com/CalendFlowBE/handler/dto"
	"dariiamoisol.com/CalendFlowBE/pkg/chatgpt"
	"dariiamoisol.com/CalendFlowBE/service"
	"encoding/json"
	"fmt"
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

	fmt.Println(req)

	chatbotReply, err := chatbotService.GenerateReply(r.Context(), handler.MapDtoChatbotGenerateReplyRequestToParams(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handler.MapChatbotReplyToDtoChatbotGenerateReplyResponse(*chatbotReply)

	respData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	fmt.Println(string(respData))

	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}
