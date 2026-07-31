package http

import (
	"fmt"
	"io/fs"
	"net/http"

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
	path, err := h.readFile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	clubs, err := h.decoder.Decode(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error decoding file: %s", err.Error()),
		})
		return
	}
	files, err := h.usecase.Generate(c, clubs...)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.attachFile(c, files...)
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) readFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("error getting file: %s", err.Error())
	}
	path := fmt.Sprintf("./%s/%s", enum.UploadFileName, file.Filename)
	if err = c.SaveUploadedFile(file, path, fs.ModePerm); err != nil {
		return "", fmt.Errorf("error uploading file locally: %w", err)
	}
	return path, nil
}

func (h *Handler) attachFile(c *gin.Context, files ...*domain.File) {
	for _, file := range files {
		c.FileAttachment(file.Name, file.Name)
	}
}
