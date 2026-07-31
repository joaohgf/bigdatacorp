package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const workspaceKey = "request-workspace"

// Workspace creates and cleans up a request-scoped temporary directory.
func Workspace(c *gin.Context) {
	path, err := os.MkdirTemp("", "bigdatacorp-api-*")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("create request workspace: %s", err),
		})
		return
	}
	defer os.RemoveAll(path)
	c.Set(workspaceKey, path)
	c.Next()
}
