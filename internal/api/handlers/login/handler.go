package login

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/labstack/echo/v4"

	uErr "monolith/internal/errors"
	"monolith/internal/generated/scheme"
	loginUsecase "monolith/internal/usecase/auth_manager"
	"monolith/internal/utils/validator"
)

type Handler struct {
	login login
}

func New(login login) *Handler {
	return &Handler{
		login: login,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	var req scheme.LoginPostReq
	if err := c.Bind(&req); err != nil || !validator.IsValidUUID(string(req.ID.Value)) {
		return c.JSON(http.StatusBadRequest, scheme.LoginPostBadRequest{})
	}

	token, err := h.login.Login(c.Request().Context(), loginUsecase.LoginRequest{
		UserID:   string(req.ID.Value),
		Password: req.Password.Value,
	})
	if err != nil {
		switch {
		case errors.Is(err, uErr.UserNotFoundErr) ||
			errors.Is(err, uErr.WrongPasswordErr):
			return c.JSON(http.StatusNotFound, scheme.LoginPostNotFound{})
		default:
			return c.JSON(http.StatusInternalServerError, scheme.R5xx{
				Message:   err.Error(),
				RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
				Code:      scheme.NewOptInt(500),
			})
		}
	}

	return c.JSON(http.StatusOK, scheme.LoginPostOK{
		Token: scheme.NewOptString(token),
	})
}
