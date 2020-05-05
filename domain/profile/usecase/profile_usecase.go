package usecase

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
	usermodel "github.com/didiyudha/transaction-accross-pkg/domain/user/model"
	"github.com/google/uuid"
	"time"
)

func (p *profileUseCase) Save(ctx context.Context, req model.CreateProfileReq) (*model.ProfileDetail, error) {
	userTrx, err := p.UserRepository.StartTx(ctx)
	if err != nil {
		return nil, err
	}
	trxContext := userTrx.Context()
	user := usermodel.User{
		ID:        uuid.New(),
		Username:  req.User.Username,
		Email:     req.User.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	if err := userTrx.Save(ctx, &user); err != nil {
		return nil, err
	}
	profile := model.Profile{
		ID:        uuid.New(),
		UserID:    user.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	profileTrx, err := p.ProfileRepository.WithTransaction(trxContext)
	if err != nil {
		userTrx.Rollback()
		return nil, err
	}
	if err := profileTrx.Save(ctx, &profile); err != nil {
		profileTrx.Rollback()
		return nil, err
	}
	profileDetail := model.ProfileDetail{
		ID:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Address:   profile.Address,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
		DeletedAt: nil,
		User:      user,
	}
	return &profileDetail, nil
}

func (p *profileUseCase) FindByID(ctx context.Context, id uuid.UUID) (*model.ProfileDetail, error) {
	profile, err := p.ProfileRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	user, err := p.UserRepository.FindByID(ctx, profile.UserID)
	if err != nil {
		return nil, err
	}
	profileDetail := model.ProfileDetail{
		ID:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Address:   profile.Address,
		CreatedAt: profile.CreatedAt,
		UpdatedAt: profile.UpdatedAt,
		DeletedAt: profile.DeletedAt,
		User:      *user,
	}
	return &profileDetail, nil
}
