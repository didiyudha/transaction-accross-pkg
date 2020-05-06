package usecase

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/repository"
	userrepo "github.com/didiyudha/transaction-accross-pkg/domain/user/repository"
	"github.com/google/uuid"
)

type ProfileUseCase interface {
	Save(ctx context.Context, req model.CreateProfileReq) (*model.ProfileDetail, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.ProfileDetail, error)
}

type profileUseCase struct {
	ProfileRepository repository.ProfileRepository
	UserRepository    userrepo.UserRepository
}

func NewProfileRepository(profileRepository repository.ProfileRepository, userRepository userrepo.UserRepository) ProfileUseCase {
	return &profileUseCase{
		ProfileRepository: profileRepository,
		UserRepository:    userRepository,
	}
}