package handler

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PenerbitHandler struct {
	penerbitUsecase usecase.PenerbitUsecase
}

func NewPenerbitHandler(penerbitUsecase usecase.PenerbitUsecase) *PenerbitHandler {
	return &PenerbitHandler{
		penerbitUsecase: penerbitUsecase,
	}
}

type PenerbitCreateDTO struct {
	NamaPenerbit string `json:"nama_penerbit" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
}

func (h *PenerbitHandler) Create(ctx *gin.Context) {
	var penerbitDTO PenerbitCreateDTO
	if err := ctx.ShouldBindJSON(&penerbitDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	penerbit := entity.Penerbit{
		NamaPenerbit: penerbitDTO.NamaPenerbit,
		Alamat:       penerbitDTO.Alamat,
	}

	createdPenerbit, err := h.penerbitUsecase.CreatePenerbit(penerbit)
	if err != nil {
		ErrorResponse(ctx, "Failed to create penerbit", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Penerbit created successfully", createdPenerbit)
}

func (h *PenerbitHandler) GetAll(ctx *gin.Context) {
	allPenerbit, err := h.penerbitUsecase.GetAllPenerbit()
	if err != nil {
		ErrorResponse(ctx, "Failed to get all penerbit", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "All penerbit", allPenerbit)
}

func (h *PenerbitHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	penerbit, err := h.penerbitUsecase.GetPenerbitByID(id)
	if err != nil {
		ErrorResponse(ctx, "Penerbit not found", http.StatusNotFound, err.Error())
		return
	}
	SuccessResponse(ctx, "Penerbit found", penerbit)
}

func (h *PenerbitHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	var penerbitDTO PenerbitCreateDTO
	if err := ctx.ShouldBindJSON(&penerbitDTO); err != nil {
		ValidationResponse(ctx, "Invalid request", SplitError(err))
		return
	}

	penerbit, err := h.penerbitUsecase.GetPenerbitByID(id)
	if err != nil {
		ErrorResponse(ctx, "Penerbit not found", http.StatusNotFound, err.Error())
		return
	}

	penerbit.NamaPenerbit = penerbitDTO.NamaPenerbit
	penerbit.Alamat = penerbitDTO.Alamat

	updatedPenerbit, err := h.penerbitUsecase.UpdatePenerbit(penerbit)
	if err != nil {
		ErrorResponse(ctx, "Failed to update penerbit", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Penerbit updated successfully", updatedPenerbit)
}

func (h *PenerbitHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ErrorResponse(ctx, "Invalid ID", http.StatusBadRequest, err.Error())
		return
	}

	err = h.penerbitUsecase.DeletePenerbit(id)
	if err != nil {
		ErrorResponse(ctx, "Failed to delete penerbit", http.StatusInternalServerError, err.Error())
		return
	}
	SuccessResponse(ctx, "Penerbit deleted successfully", nil)
}
