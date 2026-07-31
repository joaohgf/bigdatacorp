package cli

import (
	"context"
	"fmt"

	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
	"github.com/spf13/cobra"
)

type Handler struct {
	decoder port.Decoder[*domain.Club]
	usecase port.Usecase[*domain.Club, *domain.File]
}

func NewHandler(
	decoder port.Decoder[*domain.Club],
	usecase port.Usecase[*domain.Club, *domain.File],
) *Handler {
	target := new(Handler)
	target.decoder = decoder
	target.usecase = usecase
	return target
}

func (h *Handler) Run(cmd *cobra.Command, args []string) error {
	input := NewCommandInput()
	if len(args) > 0 {
		input.FilePath = args[0]
	}
	input.ClubOutput, _ = cmd.Flags().GetString(ClubOutputFlag)
	input.PlayerOutput, _ = cmd.Flags().GetString(PlayerOutputFlag)
	ctx := context.WithValue(cmd.Context(), enum.ClubFileName, input.ClubOutput)
	ctx = context.WithValue(ctx, enum.PlayerFileName, input.PlayerOutput)
	clubs := h.decoder.Decode(input.FilePath)
	if _, err := h.usecase.Generate(ctx, clubs); err != nil {
		return fmt.Errorf("error generating files: %w", err)
	}
	return nil
}
