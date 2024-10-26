package handler

import (
	"fmt"
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
	uuid := ctx.Param("uuid")

	if err := ctx.ShouldBind(body); err != nil {
		response.BadRequest("QR Code creation failed", "Invalid request payload: "+err.Error())
		return
	}

	url := uuid

	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		response.BadRequest("QR Code creation failed", "Error generating QR code: "+err.Error())
		return
	}

	filename := fmt.Sprintf("qrcode_%s.png", uuid)
	if err := qr.WriteFile(256, filename); err != nil {
		response.BadRequest("QR Code creation failed", "Error saving QR code file: "+err.Error())
		return
	}

	body.QrCodeData = filename 

	if _, err := h.CreateQRCode(body); err != nil { 
		response.BadRequest("QR Code creation failed", "Error saving to database: "+err.Error())
		return
	}

	response.Created("QR Code creation success", map[string]string{"filename": filename})
}
