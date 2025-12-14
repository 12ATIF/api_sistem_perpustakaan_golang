package repository

import (
	"coba_dulu/internal/entity"

	"gorm.io/gorm"
)

type BukuRepository interface {
	InsertBuku(buku entity.Buku) (entity.Buku, error)
	GetAllBuku(page int, pageSize int, search string) ([]entity.Buku, int64, error)
	FindBukuByID(id uint64) (entity.Buku, error)
	UpdateBuku(buku entity.Buku) (entity.Buku, error)
	DeleteBuku(id uint64) error
}

type bukuConnection struct {
	connection *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &bukuConnection{
		connection: db,
	}
}

func (db *bukuConnection) InsertBuku(buku entity.Buku) (entity.Buku, error) {
	err := db.connection.Save(&buku).Error
	if err != nil {
		return buku, err
	}
	return buku, nil
}

func (db *bukuConnection) GetAllBuku(page int, pageSize int, search string) ([]entity.Buku, int64, error) {
	var allBuku []entity.Buku
	var total int64

	query := db.connection.Model(&entity.Buku{})

	if search != "" {
		query = query.Where("judul LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Preload("Kategori").Preload("Penerbit").Offset(offset).Limit(pageSize).Find(&allBuku).Error
	if err != nil {
		return nil, 0, err
	}

	return allBuku, total, nil
}

func (db *bukuConnection) FindBukuByID(id uint64) (entity.Buku, error) {
	var buku entity.Buku
	err := db.connection.Preload("Kategori").Preload("Penerbit").First(&buku, id).Error
	if err != nil {
		return buku, err
	}
	return buku, nil
}

func (db *bukuConnection) UpdateBuku(buku entity.Buku) (entity.Buku, error) {
	err := db.connection.Save(&buku).Error
	if err != nil {
		return buku, err
	}
	return buku, nil
}

func (db *bukuConnection) DeleteBuku(id uint64) error {
	err := db.connection.Delete(&entity.Buku{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
