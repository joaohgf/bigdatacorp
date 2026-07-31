package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const maxUploadSize = int64(1 << 30)

// saveUpload locates the multipart file field and persists it at path.
func (h *Handler) saveUpload(c *gin.Context, path string) (int, error) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)
	reader, err := c.Request.MultipartReader()
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("read multipart request: %w", err)
	}
	for {
		part, nextErr := reader.NextPart()
		if errors.Is(nextErr, io.EOF) {
			return http.StatusBadRequest, errors.New("multipart field file is required")
		}
		if nextErr != nil {
			return uploadErrorStatus(nextErr), fmt.Errorf("read multipart part: %w", nextErr)
		}
		if part.FormName() != "file" {
			part.Close()
			continue
		}
		return copyUpload(part, path)
	}
}

// copyUpload streams an uploaded part into a local file.
func copyUpload(source io.ReadCloser, path string) (int, error) {
	defer source.Close()
	target, err := os.Create(path)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("create uploaded file: %w", err)
	}
	_, copyErr := io.Copy(target, source)
	closeErr := target.Close()
	if copyErr != nil {
		return uploadErrorStatus(copyErr), fmt.Errorf("save uploaded file: %w", copyErr)
	}
	if closeErr != nil {
		return http.StatusInternalServerError, fmt.Errorf("close uploaded file: %w", closeErr)
	}
	return 0, nil
}

// uploadErrorStatus maps upload failures to an HTTP response status.
func uploadErrorStatus(err error) int {
	var maxBytesError *http.MaxBytesError
	if errors.As(err, &maxBytesError) {
		return http.StatusRequestEntityTooLarge
	}
	return http.StatusBadRequest
}
