package entity

import "time"

type Anggota struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Nama      string    `gorm:"type:varchar(255)" json:"nama"`
	Alamat    string    `gorm:"type:text" json:"alamat"`
	NoTelp    string    `gorm:"type:varchar(20)" json:"no_telp"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Peminjaman []Peminjaman `gorm:"foreignKey:AnggotaID" json:"-"`
}
