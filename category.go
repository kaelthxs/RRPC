package RRPC

type Category struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
