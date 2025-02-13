package RRPC

import (
	"time"
)

type Cart struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	User       Users     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	TotalPrice int       `gorm:"not null;check:total_price >= 0"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
