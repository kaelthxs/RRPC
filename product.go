package RRPC

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null;check:price > 0"`
	Stock       int       `gorm:"default:0;check:stock >= 0"`
	CategoryID  int       `gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
