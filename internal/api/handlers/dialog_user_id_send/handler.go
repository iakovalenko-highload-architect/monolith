package dialog_user_id_send

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	"monolith/internal/generated/scheme"
	"monolith/internal/models"
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
		"request_id": uuid.New().String(),
	}))

	userID := c.Get("userID").(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, scheme.R401{})
	}

	toID := c.Param("user_id")
	if !validator.IsValidUUID(toID) {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	var req scheme.DialogUserIDSendPostReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, scheme.R400{})
	}

	if err := h.dialogManager.SendMessage(ctx, models.Message{
		FromID: c.Get("userID").(string),
		ToID:   toID,
		Text:   string(req.Text),
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, scheme.R5xx{
			Message:   err.Error(),
			RequestID: scheme.NewOptString(""),
			Code:      scheme.NewOptInt(500),
		})
	}

	return nil
}
