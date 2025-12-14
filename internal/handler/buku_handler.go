package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BukuHandler struct {
	bukuUsecase usecase.BukuUsecase
}

func NewBukuHandler(bukuUsecase usecase.BukuUsecase) *BukuHandler {
	return &BukuHandler{
		bukuUsecase: bukuUsecase,
	}
}

type BukuCreateDTO struct {
	Judul      string `json:"judul" binding:"required"`
	Penulis    string `json:"penulis" binding:"required"`
	TahunTerbit int    `json:"tahun_terbit" binding:"required"`
	Stok       int    `json:"stok" binding:"required"`
	KategoriID uint64 `json:"kategori_id" binding:"required"`
	PenerbitID uint64 `json:"penerbit_id" binding:"required"`
}

func (h *BukuHandler) Create(ctx *gin.Context) {
	var bukuDTO BukuCreateDTO
	if err := ctx.ShouldBindJSON(&bukuDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	buku := entity.Buku{
		Judul:      bukuDTO.Judul,
		Penulis:    bukuDTO.Penulis,
		TahunTerbit: bukuDTO.TahunTerbit,
		Stok:       bukuDTO.Stok,
		KategoriID: bukuDTO.KategoriID,
		PenerbitID: bukuDTO.PenerbitID,
	}

	createdBuku, err := h.bukuUsecase.CreateBuku(buku)
	if err != nil {
		ErrorResponse(ctx, "Failed to create buku", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Buku created successfully", createdBuku)
}

func (h *BukuHandler) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	search := ctx.Query("search")

	allBuku, total, err := h.bukuUsecase.GetAllBuku(page, pageSize, search)
	if err != nil {
		ErrorResponse(ctx, "Failed to get all buku", http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"data":       allBuku,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	}

	SuccessResponse(ctx, "All buku", response)
}

func (h *BukuHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	buku, err := h.bukuUsecase.GetBukuByID(id)
	if err != nil {
		ErrorResponse(ctx, "Buku not found", http.StatusNotFound, err.Error())
		return
	}
	SuccessResponse(ctx, "Buku found", buku)
}

func (h *BukuHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	var bukuDTO BukuCreateDTO
	if err := ctx.ShouldBindJSON(&bukuDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	buku, err := h.bukuUsecase.GetBukuByID(id)
	if err != nil {
		ErrorResponse(ctx, "Buku not found", http.StatusNotFound, err.Error())
		return
	}

	buku.Judul = bukuDTO.Judul
	buku.Penulis = bukuDTO.Penulis
	buku.TahunTerbit = bukuDTO.TahunTerbit
	buku.Stok = bukuDTO.Stok
	buku.KategoriID = bukuDTO.KategoriID
	buku.PenerbitID = bukuDTO.PenerbitID

	updatedBuku, err := h.bukuUsecase.UpdateBuku(buku)
	if err != nil {
		ErrorResponse(ctx, "Failed to update buku", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Buku updated successfully", updatedBuku)
}

func (h *BukuHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	err = h.bukuUsecase.DeleteBuku(id)
	if err != nil {
		ErrorResponse(ctx, "Failed to delete buku", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Buku deleted successfully", nil)
}
