package handler

import (
	"gilangaryap/gym-buddy/internal/repository"
	"gilangaryap/gym-buddy/pkg"

	"github.com/gin-gonic/gin"
)

type StatusHandler struct {
	repository.StatusRepositoryInterface
}

func NewStatusHandler(repo repository.StatusRepositoryInterface) *StatusHandler {
	return &StatusHandler{repo}
}

func (h *StatusHandler) CekStatusByUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		response.BadRequest("UUID is required", "Error")
		return
	}

	result, err := h.GetDataByUser(uuid)
	if err != nil {
		response.InternalServerError("Failed to check status for user", err.Error())
		return
	}

	response.Success("Status check successful", result)
}