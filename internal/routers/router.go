package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	authRouter(router, db)
	qrhRouter(router,db)
	statushRouter(router , db)
	
	return router
}