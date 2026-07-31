package main

import (
	"log"

	"github.com/joaohgf/bigdatacorp/internal/inbound/cli"
	"github.com/joaohgf/bigdatacorp/internal/inbound/jsonl"
	"github.com/joaohgf/bigdatacorp/internal/outbound/csv"
	"github.com/joaohgf/bigdatacorp/internal/usecase"
	"github.com/spf13/cobra"
)

// main starts cli process
func main() {
	cmd := buildRootCommand()
	if err := cmd.Execute(); err != nil {
		log.Default().Println("error executing root command")
		return
	}
}

func buildRootCommand() *cobra.Command {
	encoder := csv.NewCSV(csv.NewClubMapper(), csv.NewPlayerMapper())
	usecase := usecase.NewGenerate(encoder)
	decoder := jsonl.NewJSONL(jsonl.NewClubMapper(jsonl.NewPlayerMapper()))
	handler := cli.NewHandler(decoder, usecase)
	cmd := &cobra.Command{
		Use:     "",
		Short:   "",
		Long:    "",
		Version: "",
		RunE:    handler.Run,
	}
	return cmd
}
