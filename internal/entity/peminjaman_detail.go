package entity

import "time"

type PeminjamanDetail struct {
	ID           uint64     `gorm:"primary_key:auto_increment" json:"id"`
	PeminjamanID uint64     `gorm:"not null" json:"peminjaman_id"`
	BukuID       uint64     `gorm:"not null" json:"buku_id"`
	Qty          int        `gorm:"type:int" json:"qty"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Peminjaman   Peminjaman `gorm:"foreignKey:PeminjamanID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Buku         Buku       `gorm:"foreignKey:BukuID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
