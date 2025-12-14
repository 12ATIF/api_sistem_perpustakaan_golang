package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
)

type PeminjamanUsecase interface {
	CreatePeminjaman(peminjaman entity.Peminjaman, peminjamanDetail []entity.PeminjamanDetail) (entity.Peminjaman, error)
	GetAllPeminjaman() ([]entity.Peminjaman, error)
	GetPeminjamanByID(id uint64) (entity.Peminjaman, error)
	UpdatePeminjaman(peminjaman entity.Peminjaman) (entity.Peminjaman, error)
}

type peminjamanUsecase struct {
	peminjamanRepository repository.PeminjamanRepository
}

func NewPeminjamanUsecase(peminjamanRepo repository.PeminjamanRepository) PeminjamanUsecase {
	return &peminjamanUsecase{
		peminjamanRepository: peminjamanRepo,
	}
}

func (uc *peminjamanUsecase) CreatePeminjaman(peminjaman entity.Peminjaman, peminjamanDetail []entity.PeminjamanDetail) (entity.Peminjaman, error) {
	return uc.peminjamanRepository.InsertPeminjaman(peminjaman, peminjamanDetail)
}

func (uc *peminjamanUsecase) GetAllPeminjaman() ([]entity.Peminjaman, error) {
	return uc.peminjamanRepository.GetAllPeminjaman()
}

func (uc *peminjamanUsecase) GetPeminjamanByID(id uint64) (entity.Peminjaman, error) {
	return uc.peminjamanRepository.FindPeminjamanByID(id)
}

func (uc *peminjamanUsecase) UpdatePeminjaman(peminjaman entity.Peminjaman) (entity.Peminjaman, error) {
	return uc.peminjamanRepository.UpdatePeminjaman(peminjaman)
}
