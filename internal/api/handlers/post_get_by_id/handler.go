package post_get_by_id

import (
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
	"monolith/internal/utils/validator"
)

type Handler struct {
	postGetter postGetter
}

func New(postGetter postGetter) *Handler {
	return &Handler{
		postGetter: postGetter,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	postID := c.Param("id")
	if !validator.IsValidUUID(postID) {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	p, err := h.postGetter.GetByID(c.Request().Context(), postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	post := pointer.Get(p)
	return c.JSON(http.StatusOK, scheme.Post{
		ID:           scheme.NewOptPostId(scheme.PostId(post.ID)),
		Text:         scheme.NewOptPostText(scheme.PostText(post.Text)),
		AuthorUserID: scheme.NewOptUserId(scheme.UserId(post.UserID)),
	})
}
