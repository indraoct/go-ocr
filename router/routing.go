package router

import (
	"github.com/otiai10/gosseract"
	"go-ocr/conf"
	"go-ocr/handler"
	"github.com/labstack/echo"
)

func AssignRouting(e *echo.Echo, gosseractClient *gosseract.Client){

	//api  handlers
	api := handler.Handler{
		Gosseract: gosseractClient,
	}

	//ktp routing
	ktpGroup := e.Group(conf.Config.RootUrl + "/v1/ktp")
	ktpGroup.POST("/request",api.KtpOcrRequest)

}
