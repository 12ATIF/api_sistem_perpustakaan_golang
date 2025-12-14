package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
)

type AnggotaUsecase interface {
	CreateAnggota(anggota entity.Anggota) (entity.Anggota, error)
	GetAllAnggota(page int, pageSize int, search string) ([]entity.Anggota, int64, error)
	GetAnggotaByID(id uint64) (entity.Anggota, error)
	UpdateAnggota(anggota entity.Anggota) (entity.Anggota, error)
	DeleteAnggota(id uint64) error
}

type anggotaUsecase struct {
	anggotaRepository repository.AnggotaRepository
}

func NewAnggotaUsecase(anggotaRepo repository.AnggotaRepository) AnggotaUsecase {
	return &anggotaUsecase{
		anggotaRepository: anggotaRepo,
	}
}

func (uc *anggotaUsecase) CreateAnggota(anggota entity.Anggota) (entity.Anggota, error) {
	return uc.anggotaRepository.InsertAnggota(anggota)
}

func (uc *anggotaUsecase) GetAllAnggota(page int, pageSize int, search string) ([]entity.Anggota, int64, error) {
	return uc.anggotaRepository.GetAllAnggota(page, pageSize, search)
}

func (uc *anggotaUsecase) GetAnggotaByID(id uint64) (entity.Anggota, error) {
	return uc.anggotaRepository.FindAnggotaByID(id)
}

func (uc *anggotaUsecase) UpdateAnggota(anggota entity.Anggota) (entity.Anggota, error) {
	return uc.anggotaRepository.UpdateAnggota(anggota)
}

func (uc *anggotaUsecase) DeleteAnggota(id uint64) error {
	return uc.anggotaRepository.DeleteAnggota(id)
}
