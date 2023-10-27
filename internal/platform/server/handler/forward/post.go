package forward

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckHandler returns an HTTP handler to perform forwarding connections.
func ForwardPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok!")
	}
}
