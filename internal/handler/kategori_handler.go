package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KategoriHandler struct {
	kategoriUsecase usecase.KategoriUsecase
}

func NewKategoriHandler(kategoriUsecase usecase.KategoriUsecase) *KategoriHandler {
	return &KategoriHandler{
		kategoriUsecase: kategoriUsecase,
	}
}

type KategoriCreateDTO struct {
	NamaKategori string `json:"nama_kategori" binding:"required"`
}

func (h *KategoriHandler) Create(ctx *gin.Context) {
	var kategoriDTO KategoriCreateDTO
	if err := ctx.ShouldBindJSON(&kategoriDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	kategori := entity.Kategori{
		NamaKategori: kategoriDTO.NamaKategori,
	}

	createdKategori, err := h.kategoriUsecase.CreateKategori(kategori)
	if err != nil {
		ErrorResponse(ctx, "Failed to create kategori", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Kategori created successfully", createdKategori)
}

func (h *KategoriHandler) GetAll(ctx *gin.Context) {
	allKategori, err := h.kategoriUsecase.GetAllKategori()
	if err != nil {
		ErrorResponse(ctx, "Failed to get all kategori", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "All kategori", allKategori)
}

func (h *KategoriHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	kategori, err := h.kategoriUsecase.GetKategoriByID(id)
	if err != nil {
		ErrorResponse(ctx, "Kategori not found", http.StatusNotFound, err.Error())
		return
	}
	SuccessResponse(ctx, "Kategori found", kategori)
}

func (h *KategoriHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	var kategoriDTO KategoriCreateDTO
	if err := ctx.ShouldBindJSON(&kategoriDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	kategori, err := h.kategoriUsecase.GetKategoriByID(id)
	if err != nil {
		ErrorResponse(ctx, "Kategori not found", http.StatusNotFound, err.Error())
		return
	}

	kategori.NamaKategori = kategoriDTO.NamaKategori

	updatedKategori, err := h.kategoriUsecase.UpdateKategori(kategori)
	if err != nil {
		ErrorResponse(ctx, "Failed to update kategori", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Kategori updated successfully", updatedKategori)
}

func (h *KategoriHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	err = h.kategoriUsecase.DeleteKategori(id)
	if err != nil {
		ErrorResponse(ctx, "Failed to delete kategori", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Kategori deleted successfully", nil)
}
