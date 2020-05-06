package handler

import (
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
	"github.com/google/uuid"
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

func (h *ProfileHandler) FindByID(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")
	id, err := uuid.Parse(paramID)
	if err != nil {
		return err
	}
	profileDetail, err := h.ProfileUseCase.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if profileDetail.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	return c.JSON(http.StatusOK, profileDetail)
}