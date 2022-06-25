package http

import (
	"github.com/gin-gonic/gin"
	"task/pkg/grpc"
)

func RegisterHTTPEndpoints(router *gin.Engine, parserUsecase grpc.CreatorClient, crudUsecase grpc.EditorClient) {
	h := NewHandler(parserUsecase, crudUsecase)
	v1 := router.Group("/")
	{
		v1.POST("parse", h.ParsePosts)
		v1.POST("get-posts", h.GetPosts)
		v1.POST("get-post-by-id", h.GetPostById)
		v1.POST("delete-post", h.DeletePost)
		v1.POST("update-post", h.UpdatePost)
	}
}
