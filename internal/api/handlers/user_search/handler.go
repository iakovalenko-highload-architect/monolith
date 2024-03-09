package user_search

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
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
	params := c.QueryParams()
	users, err := h.userGetter.Search(
		c.Request().Context(),
		params.Get("first_name"),
		params.Get("last_name"),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	res := make([]scheme.User, 0, len(users))
	for _, user := range users {
		res = append(res, scheme.User{
			ID:         scheme.NewOptUserId(scheme.UserId(user.ID)),
			FirstName:  scheme.NewOptString(user.FirstName),
			SecondName: scheme.NewOptString(user.SecondName),
			Birthdate:  scheme.NewOptBirthDate(scheme.BirthDate(user.Birthday)),
			Biography:  scheme.NewOptString(user.Biography),
			City:       scheme.NewOptString(user.City),
		})
	}
	return c.JSON(http.StatusOK, res)
}
