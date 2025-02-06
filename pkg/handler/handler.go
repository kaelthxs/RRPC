package handler

import (
	"RRPC/pkg/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	services *service.Service
	db       *sqlx.DB
}

func NewHandler(services *service.Service, db *sqlx.DB) *Handler {
	return &Handler{
		services: services,
		db:       db,
	}
}
func (r *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	fmt.Println("Setting up CORS middleware...")

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	fmt.Println("CORS middleware applied!")

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Status(204)
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", r.signUp)
		auth.POST("/sign-in", r.signIn)
	}

	api := router.Group("/api", r.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", r.getAllUser)
			user.GET("/:id", r.getUserById)
			user.PUT("/:id", r.updateUser)
			user.DELETE("/:id", r.deleteUser)
		}

		category := api.Group("/category")
		{
			category.POST("/", r.createCategory)
			category.GET("/", r.getAllCategory)
			category.GET("/:id", r.getCategoryById)
			category.PUT("/:id", r.updateCategory)
			category.DELETE("/:id", r.deleteCategory)
		}

		product := api.Group("/product")
		{
			product.POST("/", r.createProduct)
			product.GET("/", r.getAllProduct)
			product.GET("/:id", r.getProductById)
			product.PUT("/:id", r.updateProduct)
			product.DELETE("/:id", r.deleteProduct)
		}

		orders := api.Group("/orders")
		{
			orders.POST("/", r.createOrder)
			orders.GET("/", r.getAllOrder)
			orders.GET("/:id", r.getOrderById)
			orders.PUT("/:id", r.updateOrder)
			orders.DELETE("/:id", r.deleteOrder)
		}

		orderItem := api.Group("/orderItem")
		{
			orderItem.POST("/", r.createOrderItem)
			orderItem.GET("/", r.getAllOrderItem)
			orderItem.GET("/:id", r.getOrderItemById)
			orderItem.PUT("/:id", r.updateOrderItem)
			orderItem.DELETE("/:id", r.deleteOrderItem)
		}

		review := api.Group("/review")
		{
			review.POST("/", r.createReview)
			review.GET("/", r.getAllReview)
			review.GET("/:id", r.getReviewById)
			review.PUT("/:id", r.updateReview)
			review.DELETE("/:id", r.deleteReview)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/", r.createAdmin)
			admin.GET("/", r.getAllAdmin)
			admin.GET("/:id", r.getAdminById)
			admin.PUT("/:id", r.updateAdmin)
			admin.DELETE("/:id", r.deleteAdmin)
		}
		payment := api.Group("/pay")
		{
			payment.POST("/", CreatePaymentIntent)
		}

	}
	return router
}
