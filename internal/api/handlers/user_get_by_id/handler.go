package user_get_by_id

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/labstack/echo/v4"

	uErr "monolith/internal/errors"
	"monolith/internal/generated/scheme"
	"monolith/internal/utils/validator"
)

type Handler struct {
	userGetter userGetter
}

func New(userGetter userGetter) *Handler {
	return &Handler{
		userGetter: userGetter,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	userID := c.Param("id")
	if !validator.IsValidUUID(userID) {
		return c.JSON(http.StatusBadRequest, scheme.UserGetIDGetBadRequest{})
	}

	user, err := h.userGetter.GetByID(c.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, uErr.UserNotFoundErr):
			return c.JSON(http.StatusNotFound, scheme.UserGetIDGetNotFound{})
		default:
			return c.JSON(http.StatusInternalServerError, scheme.R5xx{
				Message:   err.Error(),
				RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
				Code:      scheme.NewOptInt(500),
			})
		}
	}

	return c.JSON(http.StatusOK, scheme.User{
		ID:         scheme.NewOptUserId(scheme.UserId(user.ID)),
		FirstName:  scheme.NewOptString(user.FirstName),
		SecondName: scheme.NewOptString(user.SecondName),
		Birthdate:  scheme.NewOptBirthDate(scheme.BirthDate(user.Birthday)),
		Biography:  scheme.NewOptString(user.Biography),
		City:       scheme.NewOptString(user.City),
	})
}
