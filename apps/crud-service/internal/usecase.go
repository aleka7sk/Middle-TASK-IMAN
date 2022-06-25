package internal

import (
	"task/models"
	grpcservice "task/pkg/grpc"
)

type UseCase interface {
	GetPosts(postsId []int32) ([]*grpcservice.GetPostByIdResponse, error)
	GetPostById(id int) (*models.Data, error)
	DeletePostById(id int) error
	UpdatePostById(id int, userId int, title string, body string) error
}
