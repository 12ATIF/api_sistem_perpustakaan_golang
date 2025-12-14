package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnggotaHandler struct {
	anggotaUsecase usecase.AnggotaUsecase
}

func NewAnggotaHandler(anggotaUsecase usecase.AnggotaUsecase) *AnggotaHandler {
	return &AnggotaHandler{
		anggotaUsecase: anggotaUsecase,
	}
}

type AnggotaCreateDTO struct {
	Nama   string `json:"nama" binding:"required"`
	Alamat string `json:"alamat" binding:"required"`
	NoTelp string `json:"no_telp" binding:"required"`
}

func (h *AnggotaHandler) Create(ctx *gin.Context) {
	var anggotaDTO AnggotaCreateDTO
	if err := ctx.ShouldBindJSON(&anggotaDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	anggota := entity.Anggota{
		Nama:   anggotaDTO.Nama,
		Alamat: anggotaDTO.Alamat,
		NoTelp: anggotaDTO.NoTelp,
	}

	createdAnggota, err := h.anggotaUsecase.CreateAnggota(anggota)
	if err != nil {
		ErrorResponse(ctx, "Failed to create anggota", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Anggota created successfully", createdAnggota)
}

func (h *AnggotaHandler) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	search := ctx.Query("search")

	allAnggota, total, err := h.anggotaUsecase.GetAllAnggota(page, pageSize, search)
	if err != nil {
		ErrorResponse(ctx, "Failed to get all anggota", http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"data":       allAnggota,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	}

	SuccessResponse(ctx, "All anggota", response)
}

func (h *AnggotaHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	anggota, err := h.anggotaUsecase.GetAnggotaByID(id)
	if err != nil {
		ErrorResponse(ctx, "Anggota not found", http.StatusNotFound, err.Error())
		return
	}
	SuccessResponse(ctx, "Anggota found", anggota)
}

func (h *AnggotaHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	var anggotaDTO AnggotaCreateDTO
	if err := ctx.ShouldBindJSON(&anggotaDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	anggota, err := h.anggotaUsecase.GetAnggotaByID(id)
	if err != nil {
		ErrorResponse(ctx, "Anggota not found", http.StatusNotFound, err.Error())
		return
	}

	anggota.Nama = anggotaDTO.Nama
	anggota.Alamat = anggotaDTO.Alamat
	anggota.NoTelp = anggotaDTO.NoTelp

	updatedAnggota, err := h.anggotaUsecase.UpdateAnggota(anggota)
	if err != nil {
		ErrorResponse(ctx, "Failed to update anggota", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Anggota updated successfully", updatedAnggota)
}

func (h *AnggotaHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	err = h.anggotaUsecase.DeleteAnggota(id)
	if err != nil {
		ErrorResponse(ctx, "Failed to delete anggota", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Anggota deleted successfully", nil)
}
