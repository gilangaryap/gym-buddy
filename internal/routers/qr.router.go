package routers

import (
	handler "gilangaryap/gym-buddy/internal/handlers"
	"gilangaryap/gym-buddy/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func qrhRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/qr")

	var qrRepo repository.QrRepositoryInterface = repository.NewQrRepository(d)
	var subRepo repository.SubRepositoryInterface = repository.NewSubRepository(d)
	handler := handler.NewQrHandler(qrRepo, subRepo , d)

	router.POST("/:uuid", handler.CreateQRCodeHandler)
}
