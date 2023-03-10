package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const ResponseOk = "OK"

type healthResponse struct {
	Status string `json:"status" default:"OK"`
}

func RegisterHealthHandler(router EchoRouter) {
	router.GET("/health", func(ctx echo.Context) error {
		resp := healthResponse{Status: ResponseOk}

		return ctx.JSON(http.StatusOK, resp)
	})
}
