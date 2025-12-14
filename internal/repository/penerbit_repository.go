package repository

import (
	"coba_dulu/internal/entity"

	"gorm.io/gorm"
)

type PenerbitRepository interface {
	InsertPenerbit(penerbit entity.Penerbit) (entity.Penerbit, error)
	GetAllPenerbit() ([]entity.Penerbit, error)
	FindPenerbitByID(id uint64) (entity.Penerbit, error)
	UpdatePenerbit(penerbit entity.Penerbit) (entity.Penerbit, error)
	DeletePenerbit(id uint64) error
}

type penerbitConnection struct {
	connection *gorm.DB
}

func NewPenerbitRepository(db *gorm.DB) PenerbitRepository {
	return &penerbitConnection{
		connection: db,
	}
}

func (db *penerbitConnection) InsertPenerbit(penerbit entity.Penerbit) (entity.Penerbit, error) {
	err := db.connection.Save(&penerbit).Error
	if err != nil {
		return penerbit, err
	}
	return penerbit, nil
}

func (db *penerbitConnection) GetAllPenerbit() ([]entity.Penerbit, error) {
	var allPenerbit []entity.Penerbit
	err := db.connection.Find(&allPenerbit).Error
	if err != nil {
		return nil, err
	}
	return allPenerbit, nil
}

func (db *penerbitConnection) FindPenerbitByID(id uint64) (entity.Penerbit, error) {
	var penerbit entity.Penerbit
	err := db.connection.First(&penerbit, id).Error
	if err != nil {
		return penerbit, err
	}
	return penerbit, nil
}

func (db *penerbitConnection) UpdatePenerbit(penerbit entity.Penerbit) (entity.Penerbit, error) {
	err := db.connection.Save(&penerbit).Error
	if err != nil {
		return penerbit, err
	}
	return penerbit, nil
}

func (db *penerbitConnection) DeletePenerbit(id uint64) error {
	err := db.connection.Delete(&entity.Penerbit{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
