package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type hostnameResponse struct {
	Hostname string `json:"hostname"`
}

func RegisterHostnameHandler(router EchoRouter) {
	router.GET("/hostname", func(ctx echo.Context) error {
		hostname, _ := os.Hostname()

		resp := hostnameResponse{Hostname: hostname}

		return ctx.JSON(http.StatusOK, resp)
	})
}
