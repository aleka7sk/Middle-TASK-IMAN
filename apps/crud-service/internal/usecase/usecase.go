package usecase

import (
	"apps/apps/crud-service/internal"
	"apps/apps/models"
	grpcservice "apps/apps/pkg/grpc"
)

type Service struct {
	repository internal.Repository
}

func NewService(repository internal.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetPosts(postsId []int32) ([]*grpcservice.GetPostByIdResponse, error) {
	return s.repository.GetPosts(postsId)
}
func (s *Service) GetPostById(id int) (*models.Data, error) {
	return s.repository.GetPostById(id)
}
func (s *Service) DeletePostById(id int) error {
	return s.repository.DeletePostById(id)
}

func (s *Service) UpdatePostById(id int, userId int, title string, body string) error {
	return s.repository.UpdatePostById(id, userId, title, body)
}
