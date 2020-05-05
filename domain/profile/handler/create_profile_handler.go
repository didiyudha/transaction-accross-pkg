package handler

import (
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *ProfileHandler) CreateProfile(c echo.Context) error {
	var payload model.CreateProfileReq
	if err := c.Bind(&payload); err != nil {
		return err
	}
	ctx := c.Request().Context()
	profileDetail, err := h.ProfileUseCase.Save(ctx, payload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, profileDetail)
}
