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

	// Bind the request body to the QRCode model
	if err := ctx.ShouldBind(body); err != nil {
		response.BadRequest("QR Code creation failed", "Invalid request payload: "+err.Error())
		return
	}

	// Construct the URL for the QR code
	url := fmt.Sprintf("http://localhost:8080/status/%s", uuid)

	// Generate the QR code
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		response.BadRequest("QR Code creation failed", "Error generating QR code: "+err.Error())
		return
	}

	// Define the filename for the QR code image
	filename := fmt.Sprintf("qrcode_%s.png", uuid) // Use %s to format uuid correctly
	if err := qr.WriteFile(256, filename); err != nil {
		response.BadRequest("QR Code creation failed", "Error saving QR code file: "+err.Error())
		return
	}

	body.QrCodeData = filename // Store the filename in the body

	// Save the QR code information to the database
	if _, err := h.CreateQRCode(body); err != nil { 
		response.BadRequest("QR Code creation failed", "Error saving to database: "+err.Error())
		return
	}

	// Return a success response
	response.Created("QR Code creation success", map[string]string{"filename": filename})
}

