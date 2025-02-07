package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (r *Handler) getAllUser(c *gin.Context) {
	log.Println("getAllUser called")

	var users []RRPC.Users

	query := "SELECT * FROM users"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&users, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d clients\n", len(users))

	if len(users) == 0 {
		c.JSON(200, gin.H{"message": "No clients found"})
		return
	}

	c.JSON(200, users)
}

func (r *Handler) getUserById(c *gin.Context) {
	id := c.Param("id")

	var users RRPC.Users
	query := "SELECT * FROM users WHERE id = $1"

	err := r.db.Get(&users, query, id)
	if err != nil {
		log.Println("Error fetching client by ID:", err)
		c.JSON(404, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(200, users)
}

func (r *Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")

	var users RRPC.Users
	query := "SELECT * FROM users WHERE username = $1"

	err := r.db.Get(&users, query, username)
	if err != nil {
		log.Println("Error fetching client by username:", err)
		c.JSON(404, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(200, users)
}

func (r *Handler) updateUser(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Users
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE users
		SET username = $1, email = $2, password_hash = $3, role = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(query, input.Username, input.Email, input.Password_hash, input.Role, id)
	if err != nil {
		log.Println("Error updating client:", err)
		c.JSON(500, gin.H{"error": "Failed to update client"})
		return
	}

	c.JSON(200, gin.H{"message": "Client updated successfully"})
}

func (r *Handler) deleteUser(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM users WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting client:", err)
		c.JSON(500, gin.H{"error": "Failed to delete client"})
		return
	}

	c.JSON(200, gin.H{"message": "Client deleted successfully"})
}
