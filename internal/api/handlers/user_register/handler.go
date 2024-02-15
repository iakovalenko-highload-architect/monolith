package user_register

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
	"monolith/internal/models"
)

type Handler struct {
	authManager authManager
}

func New(authManager authManager) *Handler {
	return &Handler{
		authManager: authManager,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	var req scheme.UserRegisterPostReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, scheme.UserRegisterPostBadRequest{})
	}

	userID, err := h.authManager.Register(c.Request().Context(), models.User{
		Password:   req.Password.Value,
		FirstName:  req.FirstName.Value,
		SecondName: req.SecondName.Value,
		Birthday:   time.Time(req.Birthdate.Value),
		City:       req.City.Value,
		Biography:  req.Biography.Value,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(""),
			Code:      scheme.NewOptInt(500),
		})
	}

	return c.JSON(http.StatusOK, scheme.UserRegisterPostOK{
		UserID: scheme.NewOptString(userID),
	})
}
