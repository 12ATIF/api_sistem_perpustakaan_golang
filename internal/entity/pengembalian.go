package entity

import "time"

type Pengembalian struct {
	ID                 uint64     `gorm:"primary_key:auto_increment" json:"id"`
	PeminjamanID       uint64     `gorm:"not null" json:"peminjaman_id"`
	TanggalPengembalian time.Time  `gorm:"type:date" json:"tanggal_pengembalian"`
	Denda              int        `gorm:"type:int" json:"denda"`
	CreatedAt          time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Peminjaman         Peminjaman `gorm:"foreignKey:PeminjamanID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
