package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createOrder(c *gin.Context) {
	var order RRPC.Orders

	if err := c.BindJSON(&order); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO orders (user_id, status, total_price, created_at) VALUES ($1, $2, $3, $4)`

	_, err := h.db.Exec(query, order.UserID, order.Status, order.TotalPrice, order.CreatedAt)
	if err != nil {
		log.Println("Error creating order:", err)
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(200, gin.H{"message": "order created successfully"})

}

func (r *Handler) getAllOrder(c *gin.Context) {
	log.Println("getAllOrder called")

	var order []RRPC.Orders

	query := "SELECT * FROM orders"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&order, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch order"})
		return
	}

	log.Printf("Fetched %d order\n", len(order))

	if len(order) == 0 {
		c.JSON(200, gin.H{"message": "No order found"})
		return
	}

	c.JSON(200, order)
}

func (r *Handler) getOrderById(c *gin.Context) {
	id := c.Param("id")

	var order RRPC.Orders
	query := "SELECT * FROM category WHERE id = $1"

	err := r.db.Get(&order, query, id)
	if err != nil {
		log.Println("Error fetching order by ID:", err)
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	c.JSON(200, order)
}

func (r *Handler) updateOrder(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Orders
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

	_, err := r.db.Exec(query, input.UserID, input.Status, input.TotalPrice, input.CreatedAt, id)
	if err != nil {
		log.Println("Error updating order:", err)
		c.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(200, gin.H{"message": "order updated successfully"})
}

func (r *Handler) deleteOrder(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM orders WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting order:", err)
		c.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(200, gin.H{"message": "order deleted successfully"})
}
