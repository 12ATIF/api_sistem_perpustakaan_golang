package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
)

type KategoriUsecase interface {
	CreateKategori(kategori entity.Kategori) (entity.Kategori, error)
	GetAllKategori() ([]entity.Kategori, error)
	GetKategoriByID(id uint64) (entity.Kategori, error)
	UpdateKategori(kategori entity.Kategori) (entity.Kategori, error)
	DeleteKategori(id uint64) error
}

type kategoriUsecase struct {
	kategoriRepository repository.KategoriRepository
}

func NewKategoriUsecase(kategoriRepo repository.KategoriRepository) KategoriUsecase {
	return &kategoriUsecase{
		kategoriRepository: kategoriRepo,
	}
}

func (uc *kategoriUsecase) CreateKategori(kategori entity.Kategori) (entity.Kategori, error) {
	return uc.kategoriRepository.InsertKategori(kategori)
}

func (uc *kategoriUsecase) GetAllKategori() ([]entity.Kategori, error) {
	return uc.kategoriRepository.GetAllKategori()
}

func (uc *kategoriUsecase) GetKategoriByID(id uint64) (entity.Kategori, error) {
	return uc.kategoriRepository.FindKategoriByID(id)
}

func (uc *kategoriUsecase) UpdateKategori(kategori entity.Kategori) (entity.Kategori, error) {
	return uc.kategoriRepository.UpdateKategori(kategori)
}

func (uc *kategoriUsecase) DeleteKategori(id uint64) error {
	return uc.kategoriRepository.DeleteKategori(id)
}
