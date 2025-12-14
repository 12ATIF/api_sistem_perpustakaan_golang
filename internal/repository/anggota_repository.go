package repository

import (
	"coba_dulu/internal/entity"

	"gorm.io/gorm"
)

type AnggotaRepository interface {
	InsertAnggota(anggota entity.Anggota) (entity.Anggota, error)
	GetAllAnggota(page int, pageSize int, search string) ([]entity.Anggota, int64, error)
	FindAnggotaByID(id uint64) (entity.Anggota, error)
	UpdateAnggota(anggota entity.Anggota) (entity.Anggota, error)
	DeleteAnggota(id uint64) error
}

type anggotaConnection struct {
	connection *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) AnggotaRepository {
	return &anggotaConnection{
		connection: db,
	}
}

func (db *anggotaConnection) InsertAnggota(anggota entity.Anggota) (entity.Anggota, error) {
	err := db.connection.Save(&anggota).Error
	if err != nil {
		return anggota, err
	}
	return anggota, nil
}

func (db *anggotaConnection) GetAllAnggota(page int, pageSize int, search string) ([]entity.Anggota, int64, error) {
	var allAnggota []entity.Anggota
	var total int64

	query := db.connection.Model(&entity.Anggota{})

	if search != "" {
		query = query.Where("nama LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&allAnggota).Error
	if err != nil {
		return nil, 0, err
	}

	return allAnggota, total, nil
}

func (db *anggotaConnection) FindAnggotaByID(id uint64) (entity.Anggota, error) {
	var anggota entity.Anggota
	err := db.connection.First(&anggota, id).Error
	if err != nil {
		return anggota, err
	}
	return anggota, nil
}

func (db *anggotaConnection) UpdateAnggota(anggota entity.Anggota) (entity.Anggota, error) {
	err := db.connection.Save(&anggota).Error
	if err != nil {
		return anggota, err
	}
	return anggota, nil
}

func (db *anggotaConnection) DeleteAnggota(id uint64) error {
	err := db.connection.Delete(&entity.Anggota{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
