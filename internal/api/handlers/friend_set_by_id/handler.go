package friend_set_by_id

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
	"monolith/internal/utils/validator"
)

type Handler struct {
	friendSetter friendSetter
}

func New(friendSetter friendSetter) *Handler {
	return &Handler{
		friendSetter: friendSetter,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	friendID := c.Param("user_id")
	if !validator.IsValidUUID(friendID) {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	err := h.friendSetter.SetFriend(c.Request().Context(), userID, friendID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	return nil
}
