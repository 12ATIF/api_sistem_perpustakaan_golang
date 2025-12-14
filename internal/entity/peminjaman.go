package entity

import "time"

type Peminjaman struct {
	ID             uint64    `gorm:"primary_key:auto_increment" json:"id"`
	AnggotaID      uint64    `gorm:"not null" json:"anggota_id"`
	TanggalPinjam  time.Time `gorm:"type:date" json:"tanggal_pinjam"`
	TanggalKembali time.Time `gorm:"type:date" json:"tanggal_kembali"`
	Status         string    `gorm:"type:varchar(50)" json:"status"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Anggota        Anggota   `gorm:"foreignKey:AnggotaID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
