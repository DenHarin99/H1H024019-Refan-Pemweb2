
package model

import "time"

type TokenBlacklist struct {
    ID        uint      `gorm:"primaryKey"`
    TokenHash string    `gorm:"uniqueIndex;size:64;not null"`
    ExpiresAt time.Time `gorm:"not null"`
    CreatedAt time.Time
}
