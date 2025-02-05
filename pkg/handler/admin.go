package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createAdmin(c *gin.Context) {
	var admin RRPC.Admin

	if err := c.BindJSON(&admin); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO admin (user_id, permissions) VALUES ($1, $2)`

	_, err := h.db.Exec(query, admin.UserID, admin.Permissions)
	if err != nil {
		log.Println("Error creating admin:", err)
		c.JSON(500, gin.H{"error": "Failed to create admin"})
		return
	}

	c.JSON(200, gin.H{"message": "admin created successfully"})

}

func (r *Handler) getAllAdmin(c *gin.Context) {
	log.Println("getAllAdmin called")

	var admin []RRPC.Admin

	query := "SELECT * FROM admin"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&admin, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch admin"})
		return
	}

	log.Printf("Fetched %d admin\n", len(admin))

	if len(admin) == 0 {
		c.JSON(200, gin.H{"message": "No admin found"})
		return
	}

	c.JSON(200, admin)
}

func (r *Handler) getAdminById(c *gin.Context) {
	id := c.Param("id")

	var admin RRPC.Admin
	query := "SELECT * FROM admin WHERE id = $1"

	err := r.db.Get(&admin, query, id)
	if err != nil {
		log.Println("Error fetching review by ID:", err)
		c.JSON(404, gin.H{"error": "review not found"})
		return
	}

	c.JSON(200, admin)
}

func (r *Handler) updateAdmin(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Admin
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE admin
		SET user_id = $1, permissions = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(query, input.UserID, input.Permissions, id)
	if err != nil {
		log.Println("Error updating admin:", err)
		c.JSON(500, gin.H{"error": "Failed to update admin"})
		return
	}

	c.JSON(200, gin.H{"message": "admin updated successfully"})
}

func (r *Handler) deleteAdmin(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM admin WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting admin:", err)
		c.JSON(500, gin.H{"error": "Failed to delete admin"})
		return
	}

	c.JSON(200, gin.H{"message": "admin deleted successfully"})
}
