package handler

import (
	"CalendFlowBE/handler/dto"
	"CalendFlowBE/service"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

type ChatbotHandler struct {
	chatbotService *service.ChatbotService
}

func NewChatbotHandler(chatbotService *service.ChatbotService) *ChatbotHandler {
	return &ChatbotHandler{
		chatbotService: chatbotService,
	}
}

func RegisterChatbotHandler(r *router.Router, handler *ChatbotHandler) {
	r.POST("/chatbot/reply", handler.GenerateReply)
}

func (h *ChatbotHandler) GenerateReply(ctx *fasthttp.RequestCtx) {
	var req dto.ChatbotGenerateReplyRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(err.Error())
		return
	}

	chatbotReply, err := h.chatbotService.GenerateReply(ctx, mapDtoChatbotGenerateReplyRequestToParams(req))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
		return
	}

	response := mapChatbotReplyToDtoChatbotGenerateReply(*chatbotReply)

	respData, err := json.Marshal(response)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(err.Error())
		log.Print(err)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(respData)
}
