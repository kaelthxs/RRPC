package RRPC

type Order_Items struct {
	ID              int     `gorm:"primaryKey"`
	OrderID         int     `gorm:"not null"`
	Order           Orders  `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID       int     `gorm:"not null"`
	Product         Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Quantity        int     `gorm:"not null;check:quantity > 0"`
	PriceAtPurchase float64 `gorm:"not null;check:price_at_purchase > 0"`
}
