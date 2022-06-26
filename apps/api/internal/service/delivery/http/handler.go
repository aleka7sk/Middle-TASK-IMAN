package http

import (
	"apps/apps/api/internal/service"
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase service.UseCase
}

func NewHandler(uc service.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) ParsePosts(c *gin.Context) {
	status, err := h.useCase.Parse(context.Background(), "https://gorest.co.in/public/v1/posts")
	if err != nil {
		c.JSON(401, gin.H{"Status": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Status": status})
}

type Post struct {
	Id     int32  `json:"id"`
	UserId int32  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostsId struct {
	PostsId []int32 `json:"posts_id"`
}

func (h *Handler) GetPosts(c *gin.Context) {
	postsId := PostsId{}
	if c.BindJSON(&postsId) == nil {
		posts, err := h.useCase.GetPosts(context.Background(), postsId.PostsId)
		if err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Posts": posts})
		return
	}
	c.JSON(400, gin.H{"Error": "Bad Request"})
}

func (h *Handler) GetPostById(c *gin.Context) {
	post := Post{}
	if c.BindJSON(&post) == nil {
		post, err := h.useCase.GetPostById(context.Background(), post.Id)
		if err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Post": post})
		return
	}
	c.JSON(400, gin.H{"Status": "Bad Request"})

}

func (h *Handler) DeletePost(c *gin.Context) {
	post := Post{}
	if c.BindJSON(&post) == nil {
		status, err := h.useCase.DeletePost(context.Background(), post.Id)
		if err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Status": status})
		return
	}
	c.JSON(400, gin.H{"Status": "Bad Request"})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	post := Post{}
	if c.BindJSON(&post) == nil {
		result, err := h.useCase.UpdatePost(context.Background(), post.Id, post.UserId, post.Title, post.Body)
		if err != nil {
			c.JSON(400, gin.H{"Error": err.Error()})
		}
		c.JSON(200, gin.H{"Status": result})
		return
	}
	c.JSON(400, gin.H{"Status": "Bad Request"})
}
