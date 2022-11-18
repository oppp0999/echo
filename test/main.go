package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//파일을 밖에 빼내어 html 파일로 띄우고싶을 때
/*func handleHome(c echo.Context) error {
	return c.File("home.html")
}*/

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	/*	e := echo.New()
		e.GET("/", handleHome)
		e.Logger.Fatal(e.Start(":1323"))*/

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
