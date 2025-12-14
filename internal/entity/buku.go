package entity

import "time"

type Buku struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Judul      string `gorm:"type:varchar(255)" json:"judul"`
	Penulis    string `gorm:"type:varchar(255)" json:"penulis"`
	TahunTerbit int    `gorm:"type:int" json:"tahun_terbit"`
	Stok       int    `gorm:"type:int" json:"stok"`
	KategoriID uint64 `gorm:"not null" json:"kategori_id"`
	PenerbitID uint64 `gorm:"not null" json:"penerbit_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Kategori   Kategori `gorm:"foreignKey:KategoriID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Penerbit   Penerbit `gorm:"foreignKey:PenerbitID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
