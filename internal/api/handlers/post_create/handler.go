package post_create

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
	"monolith/internal/models"
)

type Handler struct {
	postCreator postCreator
}

func New(postCreator postCreator) *Handler {
	return &Handler{
		postCreator: postCreator,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	var req scheme.PostCreatePostReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	postID, err := h.postCreator.Create(c.Request().Context(), models.Post{
		UserID: userID,
		Text:   string(req.Text),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	return c.JSON(http.StatusOK, scheme.PostId(postID))
}
