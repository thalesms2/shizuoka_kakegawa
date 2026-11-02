package server

import (
	"github.com/gin-gonic/gin"

	"shizuoka/kakegawa/internal/client"
)

type Handlers struct {
	Client *client.Handler
}

func RegisterRoutes(r *gin.Engine, h Handlers) {
	api := r.Group("/api")

	h.Client.RegisterRoutes(api)
}
