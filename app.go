package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"apidocserver/action"
	"apidocserver/base"
)

func main() {
	e := echo.New()
	base.Config()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Static("/", "")
	e.GET("/login", action.Login)
	e.GET("/logout", action.Logout)
	e.GET("/project/list", action.ProjectList)
	e.POST("/project/save", action.ProjectSave)
	e.Logger.Fatal(e.Start(":9000"))
}
