package http

import (
	"apps/apps/api/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc service.UseCase) {
	h := NewHandler(uc)
	v1 := router.Group("/")
	{
		v1.POST("parse", h.ParsePosts)
		v1.POST("get-posts", h.GetPosts)
		v1.POST("get-post-by-id", h.GetPostById)
		v1.POST("delete-post", h.DeletePost)
		v1.POST("update-post", h.UpdatePost)
	}
}
