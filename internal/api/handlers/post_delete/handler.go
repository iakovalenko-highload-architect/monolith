package post_delete

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
	"monolith/internal/models"
	"monolith/internal/utils/validator"
)

type Handler struct {
	postDeleter postDeleter
}

func New(postDeleter postDeleter) *Handler {
	return &Handler{
		postDeleter: postDeleter,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	postID := c.Param("id")
	if !validator.IsValidUUID(postID) {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}
	if err := h.postDeleter.Delete(c.Request().Context(), models.Post{
		ID:     postID,
		UserID: userID,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	return nil
}
