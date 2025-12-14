package entity

import "time"

type Penerbit struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	NamaPenerbit string    `gorm:"type:varchar(255)" json:"nama_penerbit"`
	Alamat       string    `gorm:"type:text" json:"alamat"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Buku         []Buku    `gorm:"foreignKey:PenerbitID" json:"-"`
}
