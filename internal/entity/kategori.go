package entity

import "time"

type Kategori struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	NamaKategori string    `gorm:"type:varchar(255)" json:"nama_kategori"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Buku         []Buku    `gorm:"foreignKey:KategoriID" json:"-"`
}
