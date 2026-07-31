package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type SampleGenerator interface {
	Generate(context.Context, string, int, int) error
}

func NewSampleCommand(generator SampleGenerator) *cobra.Command {
	var output string
	var clubs int
	var players int
	command := &cobra.Command{
		Use:   "generate-sample",
		Short: "Generate a large JSONL sample for load and robustness checks",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generator.Generate(cmd.Context(), output, clubs, players); err != nil {
				return fmt.Errorf("generate sample: %w", err)
			}
			return nil
		},
	}
	command.Flags().StringVar(&output, "output", "example/sample_clubes_bigger.jsonl", "output JSONL path")
	command.Flags().IntVar(&clubs, "clubs", 250_000, "number of valid club records")
	command.Flags().IntVar(&players, "players", 4, "players generated per club")
	return command
}
