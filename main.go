package main

import (
	"unitag/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Define route
	e.GET("/:id", handlers.RedirectHandler)

	e.Logger.Fatal(e.Start(":3200"))
}
