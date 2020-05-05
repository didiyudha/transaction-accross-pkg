package usecase

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/model"
	"github.com/google/uuid"
)

func (u *userUseCase) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return u.UserRepository.FindByID(ctx, id)
}
