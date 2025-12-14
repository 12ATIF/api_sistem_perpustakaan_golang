package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PengembalianHandler struct {
	pengembalianUsecase usecase.PengembalianUsecase
}

func NewPengembalianHandler(pengembalianUsecase usecase.PengembalianUsecase) *PengembalianHandler {
	return &PengembalianHandler{
		pengembalianUsecase: pengembalianUsecase,
	}
}

type PengembalianCreateDTO struct {
	PeminjamanID       uint64 `json:"peminjaman_id" binding:"required"`
	TanggalPengembalian string `json:"tanggal_pengembalian" binding:"required"`
}

func (h *PengembalianHandler) Create(ctx *gin.Context) {
	var pengembalianDTO PengembalianCreateDTO
	if err := ctx.ShouldBindJSON(&pengembalianDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	tanggalPengembalian, err := time.Parse("2006-01-02", pengembalianDTO.TanggalPengembalian)
	if err != nil {
		ValidationResponse(ctx, "Invalid tanggal_pengembalian format", err.Error())
		return
	}

	pengembalian := entity.Pengembalian{
		PeminjamanID:       pengembalianDTO.PeminjamanID,
		TanggalPengembalian: tanggalPengembalian,
	}

	createdPengembalian, err := h.pengembalianUsecase.CreatePengembalian(pengembalian)
	if err != nil {
		ErrorResponse(ctx, "Failed to create pengembalian", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Pengembalian created successfully", createdPengembalian)
}
