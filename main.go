package main

import (
	_ "CalendFlowBE/handler"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"os"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
