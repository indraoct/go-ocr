package handler

import (
	"github.com/labstack/echo"
	"github.com/otiai10/gosseract"
)

type Handler struct {
	Gosseract 	*gosseract.Client
}

type ResponsePayload struct {
	Error      int     		`json:"error"`
	Message    string  		`json:"message"`
	Data 	   interface{} 	`json:"data"`
}


func (h *Handler) getErrorResponse(ctx echo.Context,errorCode int, errorMsg string) ResponsePayload {
	s := ResponsePayload{}
	s.Error 	= errorCode
	s.Message   = errorMsg
	return s
}

func (h *Handler) getSuccessResponse(ctx echo.Context,errorCode int, data interface{}, total int, params []string) ResponsePayload {
	s := ResponsePayload{}
	s.Error 	= errorCode
	s.Message   = "Success"
	s.Data 		= data

	return s
}