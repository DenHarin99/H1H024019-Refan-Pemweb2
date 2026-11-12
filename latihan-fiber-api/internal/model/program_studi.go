package model

import "time"

type ProgramStudi struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Kode      string    `gorm:"size:10;uniqueIndex" json:"kode"`
    Nama      string    `gorm:"size:100" json:"nama"`
    Jenjang   string    `gorm:"size:10" json:"jenjang"`
    CreatedAt time.Time `json:"dibuat_pada"`
    UpdatedAt time.Time `json:"diubah_pada"`
}