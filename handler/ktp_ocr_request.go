package handler

import (
	"go-ocr/conf"
	"go-ocr/lib"
	"github.com/labstack/echo"
	"github.com/otiai10/gosseract"
	"io"
	"net/http"
	"os"
)

func (h *Handler) KtpOcrRequest(ctx echo.Context) (err error){
	var(
		client 			*gosseract.Client
		text 			string
		errorCode		int
		ktp 			lib.Ktp
		filename 		string
		res             ResponsePayload
	)

	file, err :=ctx.FormFile("file")
	if err != nil{
		res = h.getErrorResponse(ctx,401,err.Error())
		return ctx.JSON(http.StatusBadRequest,res)
	}

	src, err := file.Open()
	if err != nil{
		res = h.getErrorResponse(ctx,401,err.Error())
		return ctx.JSON(http.StatusBadRequest,res)
	}
	defer src.Close()

	//default path
	filename = conf.Config.PathFile+"/"+file.Filename

	//Destination
	dst, err := os.Create(filename)
	if err != nil{
		res = h.getErrorResponse(ctx,401,err.Error())
		return ctx.JSON(http.StatusBadRequest,res)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		res = h.getErrorResponse(ctx,401,err.Error())
		return ctx.JSON(http.StatusBadRequest,res)
	}

	//validate image ktp
	err = lib.ValidateImageKtp(dst.Name())
	if err != nil{
		res = h.getErrorResponse(ctx,401,err.Error())
		return ctx.JSON(http.StatusBadRequest,res)
	}


	//Go Serract
	client = gosseract.NewClient()
	defer client.Close()

	//image input
	client.SetImage(filename)

	//define language
	client.SetLanguage("eng")

	//generate text from image
	text, err = client.Text()

	if err != nil{
		res = h.getErrorResponse(ctx,401,err.Error())
		return ctx.JSON(http.StatusBadRequest,res)
	}

	ktp,err = lib.FormatDataKTP(text)
	if err!=nil{
		res = h.getErrorResponse(ctx,401,"Failed to formating text!")
		return ctx.JSON(http.StatusBadRequest,res)
	}

	//delete file
	os.Remove(filename)

	res = h.getSuccessResponse(ctx,errorCode,ktp,1,[]string{})
	return  ctx.JSON(http.StatusOK, res)
}