package dialog_user_id_list

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"monolith/internal/generated/scheme"
	"monolith/internal/utils/validator"
)

type Handler struct {
	dialogManager dialogManager
}

func New(dialogManager dialogManager) *Handler {
	return &Handler{
		dialogManager: dialogManager,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	ctx := metadata.NewOutgoingContext(c.Request().Context(), metadata.New(map[string]string{
		"request_id": c.Request().Header.Get(echo.HeaderXRequestID),
	}))

	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	toID := c.Param("user_id")
	if !validator.IsValidUUID(toID) {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	res, err := h.dialogManager.GetDialog(ctx, userID, toID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(c.Request().Header.Get(echo.HeaderXRequestID)),
			Code:      scheme.NewOptInt(500),
		})
	}

	messages := make([]scheme.DialogMessage, 0, len(res))
	for _, msg := range res {
		messages = append(messages, scheme.DialogMessage{
			From: scheme.UserId(msg.FromID),
			To:   scheme.UserId(msg.ToID),
			Text: scheme.DialogMessageText(msg.Text),
		})
	}

	return c.JSON(http.StatusOK, messages)
}
