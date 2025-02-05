package RRPC

import "time"

type Review struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	User      Users     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ProductID int       `gorm:"not null"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Rating    int       `gorm:"not null;check:rating BETWEEN 1 AND 5"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
