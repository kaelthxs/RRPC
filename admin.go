package RRPC

import "github.com/lib/pq"

type Admin struct {
	ID          int            `gorm:"primaryKey"`
	UserID      int            `gorm:"unique;not null"`
	User        Users          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Permissions pq.StringArray `gorm:"type:jsonb;not null"`
}
