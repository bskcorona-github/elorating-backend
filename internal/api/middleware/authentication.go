package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// サンプルのユーザー情報
var sampleUser = struct {
	Username string
	Password string
}{
	Username: "sampleuser",
	Password: "samplepassword",
}

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == sampleUser.Username && password == sampleUser.Password {
		return true, nil
	}
	return false, nil
}

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, password, ok := c.Request().BasicAuth()
		if !ok || !BasicAuth(username, password, c) {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}
