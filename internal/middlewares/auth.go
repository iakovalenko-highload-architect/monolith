package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"monolith/internal/generated/scheme"
)

func (m *Middlewares) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		vals := strings.Fields(c.Request().Header.Get("Authorization"))
		if len(vals) != 2 {
			return c.JSON(http.StatusUnauthorized, scheme.R401{})
		}

		userID, err := m.tokenManager.ExtractUserID(vals[1])
		if err != nil {
			return c.JSON(http.StatusUnauthorized, scheme.R401{})
		}

		c.Set("userID", userID)
		return next(c)
	}
}
