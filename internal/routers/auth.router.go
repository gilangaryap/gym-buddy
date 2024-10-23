package routers

import (
	handler "gilangaryap/gym-buddy/internal/handlers"
	"gilangaryap/gym-buddy/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func authRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/user")

	var userRepo repository.UserRepositoryInterface = repository.NewUserRepository(d)
	handler := handler.NewAuthHandler(userRepo)

	router.POST("/register", handler.Register)
}