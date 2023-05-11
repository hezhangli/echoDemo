package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "你好")
}
func main() {

	//得到实例
	echo := echo.New()
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
	//注册路由
	echo.GET("/", hello)
	//开启server
	echo.Logger.Fatal(echo.Start(":1234"))
}
