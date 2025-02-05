package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createCategory(c *gin.Context) {
	var category RRPC.Category

	if err := c.BindJSON(&category); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO category (name) VALUES ($1)`

	_, err := h.db.Exec(query, category.Name)
	if err != nil {
		log.Println("Error creating category:", err)
		c.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(200, gin.H{"message": "category created successfully"})

}

func (r *Handler) getAllCategory(c *gin.Context) {
	log.Println("getAllCategory called")

	var category []RRPC.Category

	query := "SELECT name FROM category"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&category, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch category"})
		return
	}

	log.Printf("Fetched %d category\n", len(category))

	if len(category) == 0 {
		c.JSON(200, gin.H{"message": "No category found"})
		return
	}

	c.JSON(200, category)
}

func (r *Handler) getCategoryById(c *gin.Context) {
	id := c.Param("id")

	var category RRPC.Category
	query := "SELECT * FROM category WHERE id = $1"

	err := r.db.Get(&category, query, id)
	if err != nil {
		log.Println("Error fetching category by ID:", err)
		c.JSON(404, gin.H{"error": "category not found"})
		return
	}

	c.JSON(200, category)
}

func (r *Handler) updateCategory(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Category
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE category
		SET name = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(query, input.Name, id)
	if err != nil {
		log.Println("Error updating category:", err)
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(200, gin.H{"message": "category updated successfully"})
}

func (r *Handler) deleteCategory(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM category WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting category:", err)
		c.JSON(500, gin.H{"error": "Failed to delete caregory"})
		return
	}

	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}
