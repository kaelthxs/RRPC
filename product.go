package RRPC

type Product struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Creator     string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Price       int    `gorm:"not null;check:price > 0"`
	Stock       int    `gorm:"default:0;check:stock >= 0"`
	CategoryId  int    `gorm:"not null"`
}
