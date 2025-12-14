package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
)

type PenerbitUsecase interface {
	CreatePenerbit(penerbit entity.Penerbit) (entity.Penerbit, error)
	GetAllPenerbit() ([]entity.Penerbit, error)
	GetPenerbitByID(id uint64) (entity.Penerbit, error)
	UpdatePenerbit(penerbit entity.Penerbit) (entity.Penerbit, error)
	DeletePenerbit(id uint64) error
}

type penerbitUsecase struct {
	penerbitRepository repository.PenerbitRepository
}

func NewPenerbitUsecase(penerbitRepo repository.PenerbitRepository) PenerbitUsecase {
	return &penerbitUsecase{
		penerbitRepository: penerbitRepo,
	}
}

func (uc *penerbitUsecase) CreatePenerbit(penerbit entity.Penerbit) (entity.Penerbit, error) {
	return uc.penerbitRepository.InsertPenerbit(penerbit)
}

func (uc *penerbitUsecase) GetAllPenerbit() ([]entity.Penerbit, error) {
	return uc.penerbitRepository.GetAllPenerbit()
}

func (uc *penerbitUsecase) GetPenerbitByID(id uint64) (entity.Penerbit, error) {
	return uc.penerbitRepository.FindPenerbitByID(id)
}

func (uc *penerbitUsecase) UpdatePenerbit(penerbit entity.Penerbit) (entity.Penerbit, error) {
	return uc.penerbitRepository.UpdatePenerbit(penerbit)
}

func (uc *penerbitUsecase) DeletePenerbit(id uint64) error {
	return uc.penerbitRepository.DeletePenerbit(id)
}
