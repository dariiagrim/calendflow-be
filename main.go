package main

import (
	"CalendFlowBE/handler"
	"CalendFlowBE/pkg/chatgpt"
	"CalendFlowBE/service"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.New()

	chatGPTClient := chatgpt.NewClient(http.Client{}, os.Getenv("OPENAI_API_KEY"))
	chatbotService := service.NewChatbotService(chatGPTClient)
	chatbotHandler := handler.NewChatbotHandler(chatbotService)

	handler.RegisterChatbotHandler(r, chatbotHandler)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
