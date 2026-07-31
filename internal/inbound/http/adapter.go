package http

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joaohgf/bigdatacorp/internal/enum"
	"github.com/joaohgf/bigdatacorp/internal/port"
	"github.com/joaohgf/bigdatacorp/internal/usecase/domain"
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

func (h *Handler) Upload(c *gin.Context) {
	workspace := c.GetString(workspaceKey)
	inputPath := filepath.Join(workspace, "input.jsonl")
	status, err := h.saveUpload(c, inputPath)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	clubName := outputName(c, clubsOutputParam, enum.ClubFileName, enum.CSVType)
	playerName := outputName(c, playersOutputParam, enum.PlayerFileName, enum.CSVType)
	archiveName := outputName(c, archiveOutputParam, enum.ArchiveFileName, enum.ZIPType)
	ctx := context.WithValue(c.Request.Context(), enum.ClubFileName, filepath.Join(workspace, clubName))
	ctx = context.WithValue(ctx, enum.PlayerFileName, filepath.Join(workspace, playerName))
	files, err := h.usecase.Generate(ctx, h.decoder.Decode(inputPath))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("generate CSV files: %s", err)})
		return
	}
	archivePath := filepath.Join(workspace, archiveName)
	if err := createArchive(archivePath, files); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("create result archive: %s", err)})
		return
	}
	c.FileAttachment(archivePath, archiveName)
}
