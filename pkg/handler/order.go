package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createOrder(c *gin.Context) {
	var order RRPC.Orders
	id := c.Param("id")
	log.Println(id)

	if err := c.BindJSON(&order); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	log.Println(order)

	var newID int
	query := `INSERT INTO orders (UserID, TotalPrice, CreatedAt) VALUES ($1, $2, $3) RETURNING id`
	row := h.db.QueryRow(query, id, order.TotalPrice, order.CreatedAt)
	if err := row.Scan(&newID); err != nil {
		log.Println("Error creating order:", err)
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}
	log.Println(newID)
	c.JSON(200, gin.H{
		"message": "order created successfully",
		"id":      newID,
	})

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

func (r *Handler) getOrdersByUserId(c *gin.Context) {
	id := c.Param("id")

	log.Println("ID пользователя чтобы найти его заказы", id)

	var orders []RRPC.Orders
	query := "SELECT * FROM orders WHERE UserID = $1"

	err := r.db.Select(&orders, query, id)
	if err != nil {
		log.Println("Error fetching orders by user ID:", err)
		c.JSON(404, gin.H{"error": "orders not found"})
		return
	}

	log.Println("Это все заказы пользователя:", orders)
	c.JSON(200, orders) // Возвращаем массив заказов!
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
		SET UserID = $1, TotalPrice = $2, CreatedAt = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(query, input.UserID, input.TotalPrice, input.CreatedAt, id)
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
