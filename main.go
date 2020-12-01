package main

import (
	"github.com/otiai10/gosseract"
	"go-ocr/router"
	"github.com/labstack/echo"
	"github.com/tylerb/graceful"
	"go-ocr/conf"

	"time"
)

func main() {

	conf.InitConfig()
	e := echo.New()

	gosseract := gosseract.NewClient()

	router.AssignRouting(e, gosseract)
	e.Server.Addr = conf.Config.Port

	graceful.ListenAndServe(e.Server, 5*time.Second)
}