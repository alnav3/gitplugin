package main

import (
	"service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize server
	e := echo.New()
	e.POST("/repo", service.GetStructure)

	// Start server
	e.Logger.Fatal(e.Start(":420"))
}
