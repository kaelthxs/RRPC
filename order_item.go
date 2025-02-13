package RRPC

type Order_Items struct {
	ID        int `gorm:"primaryKey"`
	OrderID   int `gorm:"not null"`
	ProductID int `gorm:"not null"`
}
