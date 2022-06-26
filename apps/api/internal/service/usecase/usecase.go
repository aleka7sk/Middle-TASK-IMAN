package usecase

import (
	"apps/apps/pkg/grpc"
	grpcservice "apps/apps/pkg/grpc"
	"context"
)

type Service struct {
	parserUseCase grpc.CreatorClient
	crudUseCase   grpc.EditorClient
}

func NewService(pu grpc.CreatorClient, cu grpc.EditorClient) *Service {
	return &Service{
		parserUseCase: pu,
		crudUseCase:   cu,
	}
}

func (s *Service) Parse(ctx context.Context, url string) (string, error) {
	status, err := s.parserUseCase.Parse(ctx, &grpcservice.Request{Url: url})
	if err != nil {
		return "", err
	}
	return status.GetStatus(), nil
}
func (s *Service) GetPosts(ctx context.Context, postsId []int32) ([]*grpcservice.GetPostByIdResponse, error) {
	posts, err := s.crudUseCase.GetPosts(ctx, &grpcservice.GetPostsRequest{PostsId: postsId})
	if err != nil {
		return nil, err
	}
	return posts.GetPosts(), nil
}
func (s *Service) GetPostById(ctx context.Context, id int32) (*grpc.GetPostByIdResponse, error) {
	post, err := s.crudUseCase.GetPostById(ctx, &grpcservice.GetPostByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (s *Service) DeletePost(ctx context.Context, id int32) (string, error) {
	status, err := s.crudUseCase.DeletePost(ctx, &grpcservice.DeletePostRequest{Id: id})
	if err != nil {
		return status.GetStatus(), err
	}
	return status.GetStatus(), nil
}
func (s *Service) UpdatePost(ctx context.Context, id int32, userId int32, title string, body string) (string, error) {
	status, err := s.crudUseCase.UpdatePost(ctx, &grpcservice.UpdatePostRequest{Id: id, UserId: userId, Title: title, Body: body})
	if err != nil {
		return status.GetStatus(), err
	}
	return status.GetStatus(), err
}
