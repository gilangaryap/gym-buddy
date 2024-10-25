package handler

import (
	"encoding/base64"
	"gilangaryap/gym-buddy/internal/models"
	"gilangaryap/gym-buddy/internal/repository"
	"gilangaryap/gym-buddy/pkg"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type QrHandler struct {
	repository.QrRepositoryInterface
}

func NewQrHandler(repo repository.QrRepositoryInterface) *QrHandler {
	return &QrHandler{repo}
}

func (h *QrHandler) CreateQRCodeHandler(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := &models.QRCode{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("QR Code creation failed", "Invalid request payload")
		return
	}

	data := ".png"
	
	png, err := qrcode.Encode(data, qrcode.Medium, 256)
    if err != nil {
        response.BadRequest("QR Code creation failed", "Error generating QR code: "+err.Error())
        return
    }

	

	body.QrCodeData = base64.StdEncoding.EncodeToString(png)

	if _, err := h.CreateQRCode(body); err != nil {
		response.BadRequest("QR Code creation failed", "Error saving to database: "+err.Error())
		return
	}

	filename := "qrcode.png" 
	response.Created("QR Code creation success", map[string]string{"filename": filename})
}
