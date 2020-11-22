package main

import (
	"go-ocr/router"
	"github.com/labstack/echo"
	"github.com/tylerb/graceful"
	"go-ocr/conf"

	"time"
)

func main() {

	conf.InitConfig()
	e := echo.New()

	router.AssignRouting(e)
	e.Server.Addr = conf.Config.Port

	graceful.ListenAndServe(e.Server, 5*time.Second)
}