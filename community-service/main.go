package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.File("/", "index.html")

	// assets
	e.Static("/assets", "./assets")
	e.Static("/css", "./css")
	e.Static("/js", "./js")

	e.Logger.Fatal(e.Start(":" + port()))
}

func port() (p string) {
	if p = os.Getenv("PORT"); p != "" {
		return
	}
	return "8081"
}
