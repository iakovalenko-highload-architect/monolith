package post_feed

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
)

type Handler struct {
	feedGetter feedGetter
}

func New(feedGetter feedGetter) *Handler {
	return &Handler{
		feedGetter: feedGetter,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	var limit int64
	var err error
	limitString := c.Param("limit")
	if limitString != "" {
		limit, err = strconv.ParseInt(limitString, 10, 64)
	} else {
		limit = 10
	}

	if limit < 1 || err != nil {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	var offset int64
	offsetString := c.Param("offset")
	if offsetString != "" {
		offset, err = strconv.ParseInt(offsetString, 10, 64)
	}
	if offset < 0 || err != nil {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	posts, err := h.feedGetter.Get(c.Request().Context(), userID, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	res := make([]scheme.Post, 0, len(posts))
	for _, post := range posts {
		res = append(res, scheme.Post{
			ID:           scheme.NewOptPostId(scheme.PostId(post.ID)),
			Text:         scheme.NewOptPostText(scheme.PostText(post.Text)),
			AuthorUserID: scheme.NewOptUserId(scheme.UserId(post.UserID)),
		})
	}
	return c.JSON(http.StatusOK, res)
}
