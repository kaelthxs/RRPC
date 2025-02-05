package service

import (
	"RRPC"
	"RRPC/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(client RRPC.Users) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Users interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type Category interface {
	createCategory(c *gin.Context)
	getAllCategory(c *gin.Context)
	getCategoryById(c *gin.Context)
	updateCategory(c *gin.Context)
	deleteCategory(c *gin.Context)
}

type Orders interface {
	createOrder(c *gin.Context)
	getAllOrder(c *gin.Context)
	getOrderById(c *gin.Context)
	updateOrder(c *gin.Context)
	deleteOrder(c *gin.Context)
}

type Order_item interface {
	createReview(c *gin.Context)
	getAllReview(c *gin.Context)
	getReviewById(c *gin.Context)
	updateReview(c *gin.Context)
	deleteReview(c *gin.Context)
}

type Product interface {
	createReview(c *gin.Context)
	getAllReview(c *gin.Context)
	getReviewById(c *gin.Context)
	updateReview(c *gin.Context)
	deleteReview(c *gin.Context)
}

type Review interface {
	createReview(c *gin.Context)
	getAllReview(c *gin.Context)
	getReviewById(c *gin.Context)
	updateReview(c *gin.Context)
	deleteReview(c *gin.Context)
}

type Admin interface {
	createReview(c *gin.Context)
	getAllReview(c *gin.Context)
	getReviewById(c *gin.Context)
	updateReview(c *gin.Context)
	deleteReview(c *gin.Context)
}

type Service struct {
	Authorization
	Users
	Category
	Orders
	Order_item
	Product
	Review
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
