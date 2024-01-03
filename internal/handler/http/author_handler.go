package http

import (
	v1 "bphn/artikel-hukum/api/v1"
	errors2 "bphn/artikel-hukum/internal/errors"
	"bphn/artikel-hukum/internal/ito"
	"bphn/artikel-hukum/internal/service"
	"bphn/artikel-hukum/internal/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthorManagementHandler struct {
	*Handler
	authorService service.AuthorService
}

func NewAuthorManagementHandler(handler *Handler, authorService service.AuthorService) *AuthorManagementHandler {
	return &AuthorManagementHandler{
		Handler:       handler,
		authorService: authorService,
	}
}

func (h *AuthorManagementHandler) Register(ctx echo.Context) error {

	var registrationRequest v1.AuthorRegistrationRequest

	if err := ctx.Bind(&registrationRequest); err != nil {
		h.Logger().Info(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(registrationRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.authorService.Register(ctx.Request().Context(), registrationRequest); err != nil {
		if errors.Is(err, errors2.ErrEmailAlreadyExists) {
			return ctx.JSON(http.StatusConflict, errors2.ErrEmailAlreadyExists)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, &v1.CommonResponse{
		Code:    0,
		Message: "Success",
	})
}

func (h *AuthorManagementHandler) GetProfile(ctx echo.Context) error {
	userId := utils.GetUserIdFromCtx(ctx)
	profileResponse, err := h.authorService.Profile(ctx.Request().Context(), userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, profileResponse)
}

func (h *AuthorManagementHandler) UpdateProfile(ctx echo.Context) error {
	var updateRequest v1.UpdateAuthorProfileRequest
	if err := ctx.Bind(&updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(updateRequest); err != nil {
		return ctx.JSON(http.StatusNotAcceptable, err)
	}

	if err := h.authorService.UpdateProfile(ctx.Request().Context(), updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *AuthorManagementHandler) List(ctx echo.Context) error {
	var listQuery ito.ListQuery

	if err := ctx.Bind(&listQuery); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	authors, err := h.authorService.List(ctx.Request().Context(), listQuery)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, authors)
}
