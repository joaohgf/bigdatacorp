package main

import (
	"log"
	"os"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/inbound/cli"
	"github.com/joaohgf/bigdatacorp/internal/inbound/jsonl"
	"github.com/joaohgf/bigdatacorp/internal/outbound/csv"
	"github.com/joaohgf/bigdatacorp/internal/sample"
	"github.com/joaohgf/bigdatacorp/internal/usecase"
	"github.com/spf13/cobra"
)

// main starts cli process
// main builds and executes the batch CLI.
func main() {
	cmd := buildRootCommand()
	if err := cmd.Execute(); err != nil {
		log.Printf("error executing root command: %v", err)
		os.Exit(1)
	}
}

// buildRootCommand wires CLI adapters and declares root command options.
func buildRootCommand() *cobra.Command {
	encoder := csv.NewCSV(csv.NewClubMapper(), csv.NewPlayerMapper())
	usecase := usecase.NewGenerate(encoder)
	decoder := jsonl.NewJSONL(jsonl.NewClubMapper(jsonl.NewPlayerMapper()))
	handler := cli.NewHandler(decoder, usecase)
	cmd := &cobra.Command{
		Use:   "bigdatacorp [input.jsonl]",
		Short: "Generate club and player CSV files from JSONL input",
		Args:  cobra.ExactArgs(1),
		RunE:  handler.Run,
	}
	cmd.Flags().String(cli.ClubOutputFlag, string(enum.ClubFileName), "clubs CSV output path")
	cmd.Flags().String(cli.PlayerOutputFlag, string(enum.PlayerFileName), "players CSV output path")
	cmd.AddCommand(cli.NewSampleCommand(sample.NewGenerator()))
	return cmd
}
