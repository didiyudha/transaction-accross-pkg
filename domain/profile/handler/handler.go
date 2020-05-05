package handler

import "github.com/didiyudha/transaction-accross-pkg/domain/profile/usecase"

type ProfileHandler struct {
	ProfileUseCase usecase.ProfileUseCase
}

func NewProfileHandler(profileUseCase usecase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{
		ProfileUseCase:profileUseCase,
	}
}