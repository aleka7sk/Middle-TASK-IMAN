package service

import (
	"apps/apps/pkg/grpc"
	grpcservice "apps/apps/pkg/grpc"
	"context"
)

type UseCase interface {
	Parse(ctx context.Context, url string) (string, error)
	GetPosts(ctx context.Context, postsId []int32) ([]*grpcservice.GetPostByIdResponse, error)
	GetPostById(ctx context.Context, id int32) (*grpc.GetPostByIdResponse, error)
	DeletePost(ctx context.Context, id int32) (string, error)
	UpdatePost(ctx context.Context, id int32, userId int32, title string, body string) (string, error)
}
