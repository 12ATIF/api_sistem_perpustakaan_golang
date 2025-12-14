package repository

import (
	"coba_dulu/internal/entity"

	"gorm.io/gorm"
)

type PengembalianRepository interface {
	InsertPengembalian(pengembalian entity.Pengembalian) (entity.Pengembalian, error)
	GetPeminjamanDetail(peminjamanID uint64) ([]entity.PeminjamanDetail, error)
}

type pengembalianConnection struct {
	connection *gorm.DB
}

func NewPengembalianRepository(db *gorm.DB) PengembalianRepository {
	return &pengembalianConnection{
		connection: db,
	}
}

func (db *pengembalianConnection) InsertPengembalian(pengembalian entity.Pengembalian) (entity.Pengembalian, error) {
	err := db.connection.Transaction(func(tx *gorm.DB) error {
		var peminjaman entity.Peminjaman
		if err := tx.First(&peminjaman, pengembalian.PeminjamanID).Error; err != nil {
			return err
		}

		// Calculate fine
		denda := 0
		if pengembalian.TanggalPengembalian.After(peminjaman.TanggalKembali) {
			diff := pengembalian.TanggalPengembalian.Sub(peminjaman.TanggalKembali).Hours() / 24
			denda = int(diff) * 1000 // Example: 1000 per day
		}
		pengembalian.Denda = denda

		if err := tx.Save(&pengembalian).Error; err != nil {
			return err
		}

		peminjaman.Status = "Dikembalikan"
		if err := tx.Save(&peminjaman).Error; err != nil {
			return err
		}

		var peminjamanDetail []entity.PeminjamanDetail
		if err := tx.Where("peminjaman_id = ?", pengembalian.PeminjamanID).Find(&peminjamanDetail).Error; err != nil {
			return err
		}

		for _, detail := range peminjamanDetail {
			var buku entity.Buku
			if err := tx.First(&buku, detail.BukuID).Error; err != nil {
				return err
			}
			buku.Stok += detail.Qty
			if err := tx.Save(&buku).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return pengembalian, err
	}
	return pengembalian, nil
}

func (db *pengembalianConnection) GetPeminjamanDetail(peminjamanID uint64) ([]entity.PeminjamanDetail, error) {
	var peminjamanDetail []entity.PeminjamanDetail
	err := db.connection.Where("peminjaman_id = ?", peminjamanID).Find(&peminjamanDetail).Error
	if err != nil {
		return nil, err
	}
	return peminjamanDetail, nil
}
