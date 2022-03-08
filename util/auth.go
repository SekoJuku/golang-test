package util

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(os.Getenv("SESSION_USER_KEY"))
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
