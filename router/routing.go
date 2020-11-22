package router

import (
	"go-ocr/conf"
	"go-ocr/handler"
	"github.com/labstack/echo"
)

func AssignRouting(e *echo.Echo){

	//api  handlers
	api := handler.Handler{
	}

	//ktp routing
	ktpGroup := e.Group(conf.Config.RootUrl + "/v1/ktp")
	ktpGroup.POST("/request",api.KtpOcrRequest)

}
