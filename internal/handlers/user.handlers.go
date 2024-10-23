package handler

import (
	"gilangaryap/gym-buddy/internal/models"
	"gilangaryap/gym-buddy/internal/repository"
	"gilangaryap/gym-buddy/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository.UserRepositoryInterface
}

func NewAuthHandler(userRepo repository.UserRepositoryInterface) *AuthHandler {
	return &AuthHandler{userRepo}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Register failed", "Error")
		return
	}

	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("Register failed", "Error")
		return
	}

	body.PasswordHash, err = pkg.HashPassword(body.PasswordHash)
	if err != nil {
		response.BadRequest("Register failed", "Error")
		return
	}

	result, err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("Register failed", "Error")
		return
	}

	response.Created("Register success", result)
}