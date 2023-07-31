package middleware

import (
	"log"
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
		authOK, err := BasicAuth(username, password, c)
		if err != nil {
			// エラーが発生した場合はログ出力など適切な処理を行う
			log.Println("Authentication error:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}
		if !ok || !authOK {
			c.Response().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}
