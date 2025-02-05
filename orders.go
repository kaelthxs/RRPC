package RRPC

import (
	"time"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderShipped   OrderStatus = "shipped"
	OrderDelivered OrderStatus = "delivered"
	OrderCancelled OrderStatus = "cancelled"
)

type Orders struct {
	ID         int         `gorm:"primaryKey"`
	UserID     int         `gorm:"not null"`
	User       Users       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Status     OrderStatus `gorm:"type:order_status;default:'pending'"`
	TotalPrice float64     `gorm:"not null;check:total_price >= 0"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
}
