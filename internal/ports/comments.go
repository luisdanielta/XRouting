package ports

import (
	"context"
	"xrouting/internal/domain/entities"
	"xrouting/internal/domain/repositories"
)

type CommentService interface {
	RegisterComment(ctx context.Context, tableName string, comment *entities.Comment) error
	FindComment(ctx context.Context, tableName string, commentID string) (*entities.Comment, error)
	RemoveComment(ctx context.Context, tableName string, commentID string) error
}

type commentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) CommentService {
	return &commentService{commentRepo: commentRepo}
}

func (s *commentService) RegisterComment(ctx context.Context, tableName string, comment *entities.Comment) error {
	return s.commentRepo.CreateComment(ctx, tableName, comment)
}

func (s *commentService) FindComment(ctx context.Context, tableName string, commentID string) (*entities.Comment, error) {
	return s.commentRepo.GetComment(ctx, tableName, commentID)
}

func (s *commentService) RemoveComment(ctx context.Context, tableName string, commentID string) error {
	return s.commentRepo.DeleteComment(ctx, tableName, commentID)
}
