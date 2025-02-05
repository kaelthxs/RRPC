package RRPC

import (
	"time"
)

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

type Users struct {
	ID            int      `gorm:"primaryKey"`
	Username      string   `gorm:"unique;not null"`
	Email         string   `gorm:"unique;not null"`
	Password_hash string   `gorm:"not null"`
	Role          UserRole `gorm:"type:user_role;default:'customer'"`
	CreatedAt     time.Time
}
