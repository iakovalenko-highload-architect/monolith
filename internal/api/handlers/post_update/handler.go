package post_update

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
	"monolith/internal/models"
)

type Handler struct {
	postUpdater postUpdater
}

func New(postUpdater postUpdater) *Handler {
	return &Handler{
		postUpdater: postUpdater,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	var req scheme.PostUpdatePutReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	if err := h.postUpdater.Update(c.Request().Context(), models.Post{
		ID:     string(req.ID),
		UserID: userID,
		Text:   string(req.Text),
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	return nil
}
