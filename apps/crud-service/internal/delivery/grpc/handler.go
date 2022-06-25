package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"task/crud/internal"
	grpcservice "task/pkg/grpc"
)

type GRPCServer struct {
	usecase internal.UseCase
}

func NewGRPCServer(server *grpc.Server, usecase internal.UseCase) {
	grpcServer := &GRPCServer{
		usecase: usecase,
	}
	grpcservice.RegisterEditorServer(server, grpcServer)
}

type Post struct {
	Id     int32
	UserId int32
	Title  string
	Body   string
}

func (s *GRPCServer) GetPosts(ctx context.Context, req *grpcservice.GetPostsRequest) (*grpcservice.GetPostsResponse, error) {
	postsId := req.GetPostsId()
	if len(postsId) == 0 {
		return nil, errors.New("posts Id is empty")
	}
	result, err := s.usecase.GetPosts(postsId)
	if err != nil {
		return &grpcservice.GetPostsResponse{}, err
	}
	return &grpcservice.GetPostsResponse{Posts: result}, nil
}
func (s *GRPCServer) GetPostById(ctx context.Context, req *grpcservice.GetPostByIdRequest) (*grpcservice.GetPostByIdResponse, error) {
	id := int(req.GetId())
	if id < 0 {
		return nil, errors.New("post with this id doesn't exist")
	}
	post, err := s.usecase.GetPostById(id)
	if err != nil {
		return nil, err
	}
	return &grpcservice.GetPostByIdResponse{Id: int32(post.ID), UserId: int32(post.UserID), Title: post.Title, Body: post.Body}, nil
}

func (s *GRPCServer) DeletePost(ctx context.Context, req *grpcservice.DeletePostRequest) (*grpcservice.DeletePostResponse, error) {
	err := s.usecase.DeletePostById(int(req.GetId()))
	if err != nil {
		return &grpcservice.DeletePostResponse{Status: "Bad Request"}, err
	}
	return &grpcservice.DeletePostResponse{Status: "OK!"}, err
}

func (s *GRPCServer) UpdatePost(ctx context.Context, req *grpcservice.UpdatePostRequest) (*grpcservice.UpdatePostResponse, error) {
	err := s.usecase.UpdatePostById(int(req.GetId()), int(req.GetUserId()), req.Title, req.Body)
	if err != nil {
		return &grpcservice.UpdatePostResponse{Status: "Bad Request"}, err
	}
	return &grpcservice.UpdatePostResponse{Status: "OK!"}, err
}
