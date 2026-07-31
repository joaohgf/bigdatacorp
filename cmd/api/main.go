package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaohgf/bigdatacorp/internal/inbound/http"
	"github.com/joaohgf/bigdatacorp/internal/inbound/jsonl"
	"github.com/joaohgf/bigdatacorp/internal/outbound/csv"
	"github.com/joaohgf/bigdatacorp/internal/usecase"
)

func main() {
	router := buildEngine()
	if err := router.Run(":8080"); err != nil {
		log.Default().Println("error running server: %w", err)
	}
}

func buildEngine() *gin.Engine {
	engine := gin.Default()
	group := engine.Group("/api/v1/")
	encoder := csv.NewCSV(csv.NewClubMapper(), csv.NewPlayerMapper())
	usecase := usecase.NewGenerate(encoder)
	decoder := jsonl.NewJSONL(jsonl.NewClubMapper(jsonl.NewPlayerMapper()))
	handler := http.NewHandler(decoder, usecase)
	group.POST("upload", handler.Upload)
	return engine
}
