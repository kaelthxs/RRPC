package RRPC

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
