package handler

import "C"
import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createCart(c *gin.Context) {
	var cart RRPC.Cart
	id := c.Param("id")

	if err := c.BindJSON(&cart); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO cart (user_id, total_price, created_at) VALUES ($1, $2, $3)`

	_, err := h.db.Exec(query, id, cart.TotalPrice, cart.CreatedAt)
	if err != nil {
		log.Println("Error creating order:", err)
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(200, gin.H{"message": "order created successfully"})

}

func (r *Handler) getAllCart(c *gin.Context) {
	log.Println("getAllOrder called")

	var cart []RRPC.Cart

	query := "SELECT * FROM cart"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&cart, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch order"})
		return
	}

	log.Printf("Fetched %d order\n", len(cart))

	if len(cart) == 0 {
		c.JSON(200, gin.H{"message": "No order found"})
		return
	}

	c.JSON(200, cart)
}

func (r *Handler) getCartById(c *gin.Context) {
	id := c.Param("id")

	var cart RRPC.Cart
	query := "SELECT * FROM cart WHERE id = $1"

	err := r.db.Get(&cart, query, id)
	if err != nil {
		log.Println("Error fetching order by ID:", err)
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	c.JSON(200, cart)
}

func (r *Handler) updateCart(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Cart
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE orders
		SET user_id = $1, status = $2, total_price = $3, created_at = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(query, input.UserID, input.TotalPrice, input.CreatedAt, id)
	if err != nil {
		log.Println("Error updating order:", err)
		c.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(200, gin.H{"message": "order updated successfully"})
}

func (r *Handler) deleteCart(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM cart WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting order:", err)
		c.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(200, gin.H{"message": "order deleted successfully"})
}
