package RRPC

import (
	"time"
)

type Orders struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	TotalPrice int       `gorm:"not null;check:total_price >= 0"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
