package repository

import (
	"coba_dulu/internal/entity"

	"gorm.io/gorm"
)

type PeminjamanRepository interface {
	InsertPeminjaman(peminjaman entity.Peminjaman, peminjamanDetail []entity.PeminjamanDetail) (entity.Peminjaman, error)
	GetAllPeminjaman() ([]entity.Peminjaman, error)
	FindPeminjamanByID(id uint64) (entity.Peminjaman, error)
	UpdatePeminjaman(peminjaman entity.Peminjaman) (entity.Peminjaman, error)
}

type peminjamanConnection struct {
	connection *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) PeminjamanRepository {
	return &peminjamanConnection{
		connection: db,
	}
}

func (db *peminjamanConnection) InsertPeminjaman(peminjaman entity.Peminjaman, peminjamanDetail []entity.PeminjamanDetail) (entity.Peminjaman, error) {
	err := db.connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&peminjaman).Error; err != nil {
			return err
		}

		for i := range peminjamanDetail {
			peminjamanDetail[i].PeminjamanID = peminjaman.ID
			if err := tx.Create(&peminjamanDetail[i]).Error; err != nil {
				return err
			}
			var buku entity.Buku
			if err := tx.First(&buku, peminjamanDetail[i].BukuID).Error; err != nil {
				return err
			}
			buku.Stok = buku.Stok - peminjamanDetail[i].Qty
			if err := tx.Save(&buku).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return peminjaman, err
	}
	return peminjaman, nil
}

func (db *peminjamanConnection) GetAllPeminjaman() ([]entity.Peminjaman, error) {
	var allPeminjaman []entity.Peminjaman
	err := db.connection.Preload("Anggota").Find(&allPeminjaman).Error
	if err != nil {
		return nil, err
	}
	return allPeminjaman, nil
}

func (db *peminjamanConnection) FindPeminjamanByID(id uint64) (entity.Peminjaman, error) {
	var peminjaman entity.Peminjaman
	err := db.connection.Preload("Anggota").First(&peminjaman, id).Error
	if err != nil {
		return peminjaman, err
	}
	return peminjaman, nil
}

func (db *peminjamanConnection) UpdatePeminjaman(peminjaman entity.Peminjaman) (entity.Peminjaman, error) {
	err := db.connection.Save(&peminjaman).Error
	if err != nil {
		return peminjaman, err
	}
	return peminjaman, nil
}
