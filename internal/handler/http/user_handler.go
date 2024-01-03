package http

import (
	"bphn/artikel-hukum/api/v1"
	pkgerror "bphn/artikel-hukum/internal/errors"
	"bphn/artikel-hukum/internal/ito"
	"bphn/artikel-hukum/internal/service"
	"bphn/artikel-hukum/internal/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserManagementHandler struct {
	*Handler
	userService service.UserService
}

func NewUserManagementHandler(handler *Handler, userService service.UserService) *UserManagementHandler {
	return &UserManagementHandler{handler, userService}
}

func (h *UserManagementHandler) List(ctx echo.Context) error {

	var request ito.ListQuery
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	users, err := h.userService.List(ctx.Request().Context(), request)

	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, users)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserManagementHandler) Create(ctx echo.Context) error {
	var createUserRequest v1.CreateUserRequest

	if err := ctx.Bind(&createUserRequest); err != nil {
		h.Logger().Debug(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(createUserRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.userService.Create(ctx.Request().Context(), &createUserRequest); err != nil {
		// TODO Error Struct
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, v1.CommonResponse{
		Code:    0,
		Message: "success",
		Data:    createUserRequest,
	})
}

func (h *UserManagementHandler) Update(ctx echo.Context) error {

	userId := ctx.Param("id")
	ID, err := strconv.Atoi(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var updateRequest v1.UpdateUserRequest
	if err := ctx.Bind(&updateRequest); err != nil {
		h.Logger().Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateRequest.Id = uint(ID)

	if err := ctx.Validate(updateRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.userService.Update(ctx.Request().Context(), &updateRequest); err != nil {
		if errors.Is(err, pkgerror.ErrUserDoesNotExists) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)

}

func (h *UserManagementHandler) Delete(ctx echo.Context) error {
	var userId = ctx.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := h.userService.Delete(ctx.Request().Context(), uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)

}

func (h *UserManagementHandler) ChangePasswordByNonAdmin(ctx echo.Context) error {

	userId := utils.GetUserIdFromCtx(ctx)
	var request v1.ChangePasswordRequest

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, err)
	}

	if err := ctx.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	changePassword := ito.ChangePasswordQuery{
		UserId:   userId,
		Password: request.NewPassword,
	}

	if err := h.userService.ChangePasswordByNonAdmin(ctx.Request().Context(), changePassword); err != nil {
		return err
	}

	return nil

}

func (h *UserManagementHandler) ForgotPassword(ctx echo.Context) error {

	var forgotRequest v1.ForgotPasswordRequest
	if err := ctx.Bind(&forgotRequest); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, err)
	}

	if err := ctx.Validate(forgotRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.userService.ForgotPassword(ctx.Request().Context(), forgotRequest); err != nil {
		if errors.Is(err, pkgerror.ErrUserDoesNotExists) {
			return echo.NewHTTPError(http.StatusNotAcceptable, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)
}
