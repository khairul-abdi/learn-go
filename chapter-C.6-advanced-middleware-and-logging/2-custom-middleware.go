package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func middlewareSomething(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from middleware something")
		next.ServeHTTP(w, r)
	})
}

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(middlewareSomething))

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
