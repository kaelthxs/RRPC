package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createOrderItem(c *gin.Context) {
	var orderItem RRPC.Order_Items

	if err := c.BindJSON(&orderItem); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO order_item (order_id, product_id, quantity, price_at_purchase) VALUES ($1, $2, $3, $4)`

	_, err := h.db.Exec(query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.PriceAtPurchase)
	if err != nil {
		log.Println("Error creating orderItem:", err)
		c.JSON(500, gin.H{"error": "Failed to create orderItem"})
		return
	}

	c.JSON(200, gin.H{"message": "orderItem created successfully"})

}

func (r *Handler) getAllOrderItem(c *gin.Context) {
	log.Println("getAllOrderItem called")

	var orderItem []RRPC.Order_Items

	query := "SELECT * FROM order_item"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&orderItem, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch orderItem"})
		return
	}

	log.Printf("Fetched %d orderItem\n", len(orderItem))

	if len(orderItem) == 0 {
		c.JSON(200, gin.H{"message": "No orderItem found"})
		return
	}

	c.JSON(200, orderItem)
}

func (r *Handler) getOrderItemById(c *gin.Context) {
	id := c.Param("id")

	var orderItem RRPC.Order_Items
	query := "SELECT * FROM order_item WHERE id = $1"

	err := r.db.Get(&orderItem, query, id)
	if err != nil {
		log.Println("Error fetching orderItem by ID:", err)
		c.JSON(404, gin.H{"error": "orderItem not found"})
		return
	}

	c.JSON(200, orderItem)
}

func (r *Handler) updateOrderItem(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Order_Items
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE order_item
		SET order_id = $1, product_id = $2, quantity = $3, price_at_purchase = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(query, input.OrderID, input.ProductID, input.Quantity, input.PriceAtPurchase, id)
	if err != nil {
		log.Println("Error updating orderItem:", err)
		c.JSON(500, gin.H{"error": "Failed to update orderItem"})
		return
	}

	c.JSON(200, gin.H{"message": "orderItem updated successfully"})
}

func (r *Handler) deleteOrderItem(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM order_item WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting orderItem:", err)
		c.JSON(500, gin.H{"error": "Failed to delete orderItem"})
		return
	}

	c.JSON(200, gin.H{"message": "orderItem deleted successfully"})
}
