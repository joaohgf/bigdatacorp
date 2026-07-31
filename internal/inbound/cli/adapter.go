package cli

import (
	"fmt"

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
	clubs, err := h.decoder.Decode(input.FilePath)
	if err != nil {
		return fmt.Errorf("error decoding: %w", err)
	}
	_, err = h.usecase.Generate(cmd.Context(), clubs...)
	if err != nil {
		return fmt.Errorf("error generating files: %w", err)
	}
	return nil
}
