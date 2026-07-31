package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaohgf/bigdatacorp/internal/inbound/http"
	"github.com/joaohgf/bigdatacorp/internal/inbound/jsonl"
	"github.com/joaohgf/bigdatacorp/internal/outbound/csv"
	"github.com/joaohgf/bigdatacorp/internal/usecase"
)

// main starts the HTTP API server.
func main() {
	router := buildEngine()
	if err := router.Run(":8080"); err != nil {
		log.Printf("error running server: %v", err)
	}
}

// buildEngine wires the HTTP transport and its application dependencies.
func buildEngine() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("/api/v1/")
	encoder := csv.NewCSV(csv.NewClubMapper(), csv.NewPlayerMapper())
	usecase := usecase.NewGenerate(encoder)
	decoder := jsonl.NewJSONL(jsonl.NewClubMapper(jsonl.NewPlayerMapper()))
	handler := http.NewHandler(decoder, usecase)
	group.POST("upload", http.Workspace, handler.Upload)
	return engine
}
