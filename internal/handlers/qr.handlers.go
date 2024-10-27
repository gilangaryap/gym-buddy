package handler

import (
	"context"
	"fmt"
	"gilangaryap/gym-buddy/internal/models"
	"gilangaryap/gym-buddy/internal/repository"
	"gilangaryap/gym-buddy/pkg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/skip2/go-qrcode"
)

type QrHandler struct {
	repository.QrRepositoryInterface
	repository.SubRepositoryInterface
	db *sqlx.DB
}

func NewQrHandler(repoQr repository.QrRepositoryInterface , repoSub repository.SubRepositoryInterface  ,db *sqlx.DB) *QrHandler{
	return &QrHandler{repoQr , repoSub , db}
}

func (h *QrHandler) CreateQRCodeHandler(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	uuid := ctx.Param("uuid")
	bodySub := &models.Subs{}
	bodyQr := &models.QRCode{}

	if err := ctx.ShouldBindJSON(&bodySub); err != nil {
		response.BadRequest("Invalid input", "Error binding subscription data: "+err.Error())
		return 
	}

	tx, err := h.db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Println("db.BeginTx", err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	subscriptionID, err := h.CreateDataByUser(uuid, bodySub)
	if err != nil {
	    tx.Rollback()
	    response.BadRequest("QR Code creation failed", "Error creating subscription: "+err.Error())
	    return 
	}

	log.Printf("Created subscription with ID: %s", subscriptionID)

	qr, err := qrcode.New(uuid, qrcode.Medium)
	if err != nil {
		tx.Rollback()
		response.BadRequest("QR Code creation failed", "Error generating QR code: "+err.Error())
		return 
	}

	filename := fmt.Sprintf("qrcode_%s.png", uuid)
	if err := qr.WriteFile(256, filename); err != nil {
		tx.Rollback()
		response.BadRequest("QR Code creation failed", "Error saving QR code file: "+err.Error())
		return 
	}

	bodyQr.QrCodeData = filename
	bodyQr.SubscriptionID = subscriptionID

	if _, err := h.CreateQRCode(bodyQr); err != nil {
		tx.Rollback()
		response.BadRequest("QR Code creation failed", "Error saving to database: "+err.Error())
		return 
	}

	if err := tx.Commit(); err != nil {
		response.BadRequest("QR Code creation failed", "Error committing transaction: "+err.Error())
		return 
	}

	response.Success("QR Code created successfully", bodyQr)
}

/* func (h *QrHandler) CreateQRCodeHandlera(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	bodySub := &models.Subs{}
	bodyQr := &models.QRCode{}
	uuid := ctx.Param("uuid")

	if err := ctx.ShouldBind(bodySub); err != nil {
		response.BadRequest("QR Code creation failed", "Invalid request payload: "+err.Error())
		return
	}

	if _, err := h.CreateDataByUser(uuid , bodySub); err != nil { 
		response.BadRequest("")
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

	bodyQr.QrCodeData = filename 

	if _, err := h.CreateQRCode(bodyQr); err != nil { 
		response.BadRequest("QR Code creation failed", "Error saving to database: "+err.Error())
		return
	}

	response.Created("QR Code creation success", map[string]string{"filename": filename})
} */