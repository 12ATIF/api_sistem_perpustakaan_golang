package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
)

type BukuUsecase interface {
	CreateBuku(buku entity.Buku) (entity.Buku, error)
	GetAllBuku(page int, pageSize int, search string) ([]entity.Buku, int64, error)
	GetBukuByID(id uint64) (entity.Buku, error)
	UpdateBuku(buku entity.Buku) (entity.Buku, error)
	DeleteBuku(id uint64) error
}

type bukuUsecase struct {
	bukuRepository repository.BukuRepository
}

func NewBukuUsecase(bukuRepo repository.BukuRepository) BukuUsecase {
	return &bukuUsecase{
		bukuRepository: bukuRepo,
	}
}

func (uc *bukuUsecase) CreateBuku(buku entity.Buku) (entity.Buku, error) {
	return uc.bukuRepository.InsertBuku(buku)
}

func (uc *bukuUsecase) GetAllBuku(page int, pageSize int, search string) ([]entity.Buku, int64, error) {
	return uc.bukuRepository.GetAllBuku(page, pageSize, search)
}

func (uc *bukuUsecase) GetBukuByID(id uint64) (entity.Buku, error) {
	return uc.bukuRepository.FindBukuByID(id)
}

func (uc *bukuUsecase) UpdateBuku(buku entity.Buku) (entity.Buku, error) {
	return uc.bukuRepository.UpdateBuku(buku)
}

func (uc *bukuUsecase) DeleteBuku(id uint64) error {
	return uc.bukuRepository.DeleteBuku(id)
}
