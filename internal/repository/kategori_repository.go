package repository

import (
	"coba_dulu/internal/entity"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	InsertKategori(kategori entity.Kategori) (entity.Kategori, error)
	GetAllKategori() ([]entity.Kategori, error)
	FindKategoriByID(id uint64) (entity.Kategori, error)
	UpdateKategori(kategori entity.Kategori) (entity.Kategori, error)
	DeleteKategori(id uint64) error
}

type kategoriConnection struct {
	connection *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) KategoriRepository {
	return &kategoriConnection{
		connection: db,
	}
}

func (db *kategoriConnection) InsertKategori(kategori entity.Kategori) (entity.Kategori, error) {
	err := db.connection.Save(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (db *kategoriConnection) GetAllKategori() ([]entity.Kategori, error) {
	var allKategori []entity.Kategori
	err := db.connection.Find(&allKategori).Error
	if err != nil {
		return nil, err
	}
	return allKategori, nil
}

func (db *kategoriConnection) FindKategoriByID(id uint64) (entity.Kategori, error) {
	var kategori entity.Kategori
	err := db.connection.First(&kategori, id).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (db *kategoriConnection) UpdateKategori(kategori entity.Kategori) (entity.Kategori, error) {
	err := db.connection.Save(&kategori).Error
	if err != nil {
		return kategori, err
	}
	return kategori, nil
}

func (db *kategoriConnection) DeleteKategori(id uint64) error {
	err := db.connection.Delete(&entity.Kategori{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
