package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createReview(c *gin.Context) {
	var review RRPC.Review

	if err := c.BindJSON(&review); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO review (user_id, product_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := h.db.Exec(query, review.UserID, review.ProductID, review.Rating, review.Comment, review.CreatedAt)
	if err != nil {
		log.Println("Error creating review:", err)
		c.JSON(500, gin.H{"error": "Failed to create review"})
		return
	}

	c.JSON(200, gin.H{"message": "review created successfully"})

}

func (r *Handler) getAllReview(c *gin.Context) {
	log.Println("getAllOrderItem called")

	var review []RRPC.Review

	query := "SELECT * FROM review"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&review, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch review"})
		return
	}

	log.Printf("Fetched %d review\n", len(review))

	if len(review) == 0 {
		c.JSON(200, gin.H{"message": "No review found"})
		return
	}

	c.JSON(200, review)
}

func (r *Handler) getReviewById(c *gin.Context) {
	id := c.Param("id")

	var review RRPC.Review
	query := "SELECT * FROM review WHERE id = $1"

	err := r.db.Get(&review, query, id)
	if err != nil {
		log.Println("Error fetching review by ID:", err)
		c.JSON(404, gin.H{"error": "review not found"})
		return
	}

	c.JSON(200, review)
}

func (r *Handler) updateReview(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Review
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE review
		SET user_id = $1, product_id = $2, rating = $3, comment = $4, created_at = $5
		WHERE id = $6
	`

	_, err := r.db.Exec(query, input.UserID, input.ProductID, input.Rating, input.Comment, input.CreatedAt, id)
	if err != nil {
		log.Println("Error updating review:", err)
		c.JSON(500, gin.H{"error": "Failed to update review"})
		return
	}

	c.JSON(200, gin.H{"message": "orderItem updated successfully"})
}

func (r *Handler) deleteReview(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM review WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting review:", err)
		c.JSON(500, gin.H{"error": "Failed to delete review"})
		return
	}

	c.JSON(200, gin.H{"message": "review deleted successfully"})
}
