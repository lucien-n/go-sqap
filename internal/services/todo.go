package services

import (
	"context"
	"sqap/internal/models"
	"sqap/internal/repositories"

	"github.com/google/uuid"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) Get() (*models.Todo, error) {
	todo, err := s.repo.Get(context.Background())
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetTodo(uid string) (*models.Todo, error) {
	todo, err := s.repo.GetTodo(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) PostTodo(todoReq *models.TodoRequest) error {
	todo := &models.Todo{
		UID:        uuid.NewString(),
		Content:    todoReq.Content,
		IsComplete: todoReq.IsComplete,
	}

	err := s.repo.CreateTodo(context.Background(), todo)

	if err != nil {
		return err
	}

	return nil
}
