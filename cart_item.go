package RRPC

type CartItem struct {
	ID              int     `gorm:"primaryKey"`
	CartID          int     `gorm:"not null"`
	ProductID       int     `gorm:"not null"`
	Product         Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Quantity        int     `gorm:"not null;check:quantity > 0"`
	PriceAtPurchase float64 `gorm:"not null;check:price_at_purchase > 0"`
}
