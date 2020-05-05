package usecase

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/model"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/repository"
	"github.com/google/uuid"
)

type UserUseCase interface {
	FindByID(ctx context.Context, id uuid.UUID) (*model.User, error)
}

type userUseCase struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		UserRepository:userRepository,
	}
}