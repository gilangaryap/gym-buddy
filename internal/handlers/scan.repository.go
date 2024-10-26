package handler

import (
	"gilangaryap/gym-buddy/internal/repository"
	"gilangaryap/gym-buddy/pkg"
	"image"

	"github.com/gin-gonic/gin"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

type ScanHandler struct {
	repository.StatusRepositoryInterface
}

func NewScanHandler(repo repository.StatusRepositoryInterface) *ScanHandler {
	return &ScanHandler{repo}
}

func (h *ScanHandler) ScanQRCodeHandler(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	file, err := ctx.FormFile("image")
	if err != nil {
		response.BadRequest("Error getting image file", err.Error())
		return
	}

	img, err := file.Open()
	if err != nil {
		response.BadRequest("Error opening image file", err.Error())
		return
	}
	defer img.Close()

	decodedImg, _, err := image.Decode(img)
	if err != nil {
		response.BadRequest("Error decoding image", err.Error())
		return
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(decodedImg)
	if err != nil {
		response.BadRequest("Error creating binary bitmap", err.Error())
		return
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		response.BadRequest("Error decoding QR code", err.Error())
		return
	}

	data := result.GetText()

	userData, err := h.GetDataByUser(data)
    if err != nil {
        response.BadRequest("Register failed", "Error")
        return
    }

	response.Success("QR Code scanned successfully", userData)
}