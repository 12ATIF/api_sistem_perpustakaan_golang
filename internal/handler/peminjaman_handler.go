package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PeminjamanHandler struct {
	peminjamanUsecase usecase.PeminjamanUsecase
}

func NewPeminjamanHandler(peminjamanUsecase usecase.PeminjamanUsecase) *PeminjamanHandler {
	return &PeminjamanHandler{
		peminjamanUsecase: peminjamanUsecase,
	}
}

type PeminjamanDetailDTO struct {
	BukuID uint64 `json:"buku_id" binding:"required"`
	Qty    int    `json:"qty" binding:"required"`
}

type PeminjamanCreateDTO struct {
	AnggotaID      uint64                `json:"anggota_id" binding:"required"`
	TanggalPinjam  string                `json:"tanggal_pinjam" binding:"required"`
	TanggalKembali string                `json:"tanggal_kembali" binding:"required"`
	Detail         []PeminjamanDetailDTO `json:"detail" binding:"required"`
}

func (h *PeminjamanHandler) Create(ctx *gin.Context) {
	var peminjamanDTO PeminjamanCreateDTO
	if err := ctx.ShouldBindJSON(&peminjamanDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	tanggalPinjam, err := time.Parse("2006-01-02", peminjamanDTO.TanggalPinjam)
	if err != nil {
		ValidationResponse(ctx, "Invalid tanggal_pinjam format", err.Error())
		return
	}
	tanggalKembali, err := time.Parse("2006-01-02", peminjamanDTO.TanggalKembali)
	if err != nil {
		ValidationResponse(ctx, "Invalid tanggal_kembali format", err.Error())
		return
	}

	peminjaman := entity.Peminjaman{
		AnggotaID:      peminjamanDTO.AnggotaID,
		TanggalPinjam:  tanggalPinjam,
		TanggalKembali: tanggalKembali,
		Status:         "Dipinjam",
	}

	var peminjamanDetail []entity.PeminjamanDetail
	for _, detailDTO := range peminjamanDTO.Detail {
		peminjamanDetail = append(peminjamanDetail, entity.PeminjamanDetail{
			BukuID: detailDTO.BukuID,
			Qty:    detailDTO.Qty,
		})
	}

	createdPeminjaman, err := h.peminjamanUsecase.CreatePeminjaman(peminjaman, peminjamanDetail)
	if err != nil {
		ErrorResponse(ctx, "Failed to create peminjaman", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Peminjaman created successfully", createdPeminjaman)
}
