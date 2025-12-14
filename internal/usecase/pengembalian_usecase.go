package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
)

type PengembalianUsecase interface {
	CreatePengembalian(pengembalian entity.Pengembalian) (entity.Pengembalian, error)
}

type pengembalianUsecase struct {
	pengembalianRepository repository.PengembalianRepository
}

func NewPengembalianUsecase(pengembalianRepo repository.PengembalianRepository) PengembalianUsecase {
	return &pengembalianUsecase{
		pengembalianRepository: pengembalianRepo,
	}
}

func (uc *pengembalianUsecase) CreatePengembalian(pengembalian entity.Pengembalian) (entity.Pengembalian, error) {
	return uc.pengembalianRepository.InsertPengembalian(pengembalian)
}
