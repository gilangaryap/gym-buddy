package routers

import (
	handler "gilangaryap/gym-buddy/internal/handlers"
	"gilangaryap/gym-buddy/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func statushRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/scan")

	var statusRepo repository.StatusRepositoryInterface = repository.NewStatusRepository(d)
	handler := handler.NewScanHandler(statusRepo)

	router.POST("/", handler.ScanQRCodeHandler)
}