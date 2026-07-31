package http

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaohgf/bigdatacorp/internal/enum"
)

const (
	clubsOutputParam   = "clubs-output"
	playersOutputParam = "players-output"
	archiveOutputParam = "archive-output"
)

// outputName sanitizes a requested output name and applies its required extension.
func outputName(c *gin.Context, parameter string, fallback enum.FileName, fileType enum.FileType) string {
	name := filepath.Base(strings.TrimSpace(c.Query(parameter)))
	if name == "." || name == "" {
		name = string(fallback)
	}
	extension := fmt.Sprintf(".%s", fileType)
	if !strings.EqualFold(filepath.Ext(name), extension) {
		name += extension
	}
	return name
}
