package http

import (
  "apps/apps/pkg/grpc"
  grpcservice "apps/apps/pkg/grpc"
  "context"
  "github.com/gin-gonic/gin"
)

type Handler struct {
  parserUsecase grpc.CreatorClient
  crudUsecase   grpc.EditorClient
}

func NewHandler(parserUsecase grpc.CreatorClient, crudUsecase grpc.EditorClient) *Handler {
  return &Handler{
    parserUsecase: parserUsecase,
    crudUsecase:   crudUsecase,
  }
}

func (h *Handler) ParsePosts(c *gin.Context) {
  res, err := h.parserUsecase.Parse(context.Background(), &grpcservice.Request{Url: "https://gorest.co.in/public/v1/posts"})
  if err != nil {
    c.JSON(401, gin.H{"Status": err.Error()})
    return
  }
  c.JSON(200, gin.H{"Status": res.GetStatus()})
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
    result, err := h.crudUsecase.GetPosts(context.Background(), &grpcservice.GetPostsRequest{PostsId: postsId.PostsId})
    if err != nil {
      c.JSON(400, gin.H{"Error": err.Error()})
      return
    }
    c.JSON(200, gin.H{"Status": result})
    return
  }
  c.JSON(400, gin.H{"Error": "Bad Request"})
}

func (h *Handler) GetPostById(c *gin.Context) {
  post := Post{}
  if c.BindJSON(&post) == nil {
    result, err := h.crudUsecase.GetPostById(context.Background(), &grpcservice.GetPostByIdRequest{Id: post.Id})
    if err != nil {
      c.JSON(400, gin.H{"Error": err.Error()})
      return
    }
    c.JSON(200, gin.H{"Status": result})
    return
  }
  c.JSON(400, gin.H{"Status": "Bad Request"})

}

func (h *Handler) DeletePost(c *gin.Context) {
  post := Post{}
  if c.BindJSON(&post) == nil {
    result, err := h.crudUsecase.DeletePost(context.Background(), &grpcservice.DeletePostRequest{Id: post.Id})
    if err != nil {
      c.JSON(400, gin.H{"Error": err.Error()})
      return
    }
    c.JSON(200, gin.H{"Status": result})
    return
  }
  c.JSON(400, gin.H{"Status": "Bad Request"})
}

func (h *Handler) UpdatePost(c *gin.Context) {
  post := Post{}
  if c.BindJSON(&post) == nil {
    result, err := h.crudUsecase.UpdatePost(context.Background(), &grpcservice.UpdatePostRequest{Id: post.Id, UserId: post.UserId, Title: post.Title, Body: post.Body})
    if err != nil {
      c.JSON(400, gin.H{"Error": err.Error()})
    }
    c.JSON(200, gin.H{"Status": result})
    return
  }
  c.JSON(400, gin.H{"Status": "Bad Request"})
}
