package handler

import (
	"fmt"
	"gilangaryap/gym-buddy/internal/middlewares"
	"gilangaryap/gym-buddy/internal/models"
	"gilangaryap/gym-buddy/internal/repository"
	"gilangaryap/gym-buddy/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository.UserRepositoryInterface
	repository.AuthRepositoryInterface
}

func NewAuthHandler(userRepo repository.UserRepositoryInterface , authRepo repository.AuthRepositoryInterface) *AuthHandler {
	return &AuthHandler{userRepo , authRepo}
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

	validatedEmail, err := middlewares.Register(body.Email)
    if err != nil {
        response.BadRequest("Register failed", "Email validation error: "+err.Error())
        return
    }
    
    body.Email = validatedEmail

	result, err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("Register failed", "Error")
		return
	}

	response.Created("Register success", result)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Auth{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Invalid request format", err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		response.BadRequest("Validation failed", err.Error())
		return
	}

	result, err := h.GetByEmail(body.Email)
	if err != nil {
		response.BadRequest("Login failed", "Error")
		return
	}
	fmt.Printf("Executing query: %s with email: %s\n", result.PasswordHash, body.PasswordHash)

	err = pkg.VerifyPassword(result.PasswordHash, body.PasswordHash)
	if err != nil {
		response.Unauthorized("Incorrect password", "Error")
		return
	}

	jwt := pkg.NewJWT(result.Id, result.Email, result.Role)
	token, err := jwt.GenerateToken()
	if err != nil {
		response.InternalServerError("Failed to generate token", err.Error())
		return
	}

	response.Success("Login successful", gin.H{
		"token": token,
		"user": gin.H{
			"email": result.Email,
			"role":  result.Role,
			"id":    result.Id,
		},
	})
}
