package handler

import (
	"RRPC"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) createProduct(c *gin.Context) {
	var product RRPC.Product

	if err := c.BindJSON(&product); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO product (name, creator, description, price, stock, CategoryId) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := h.db.Exec(query, product.Name, product.Creator, product.Description, product.Price, product.Stock, product.CategoryId)
	if err != nil {
		log.Println("Error creating product:", err)
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(200, gin.H{"message": "product created successfully"})

}

func (r *Handler) getAllProduct(c *gin.Context) {
	log.Println("getAllProduct called")

	var product []RRPC.Product

	query := "SELECT * FROM product"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&product, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch product"})
		return
	}

	log.Printf("Fetched %d product\n", len(product))

	if len(product) == 0 {
		c.JSON(200, gin.H{"message": "No product found"})
		return
	}

	c.JSON(200, product)
}

func (r *Handler) getProductById(c *gin.Context) {
	id := c.Param("id")

	var product RRPC.Product
	query := "SELECT * FROM product WHERE id = $1"

	err := r.db.Get(&product, query, id)
	if err != nil {
		log.Println("Error fetching product by ID:", err)
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	c.JSON(200, product)
}

func (r *Handler) getProductByCategory(c *gin.Context) {
	CategoryID := c.Param("CategoryId")

	var product []RRPC.Product
	query := "SELECT * FROM product where categoryid = $1"

	err := r.db.Select(&product, query, CategoryID)
	if err != nil {
		log.Println("Error fetching product by ID:", err)
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	c.JSON(200, product)
}

func (r *Handler) updateProduct(c *gin.Context) {
	id := c.Param("id")

	var input RRPC.Product
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE product
		SET name = $1, creator = $2, description = $3, price = $4, stock = $5, CategoryId = $6
		WHERE id = $7
	`

	_, err := r.db.Exec(query, input.Name, input.Description, input.Price, input.Stock, input.CategoryId, id)
	if err != nil {
		log.Println("Error updating product:", err)
		c.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(200, gin.H{"message": "product updated successfully"})
}

func (r *Handler) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM product WHERE id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting product:", err)
		c.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(200, gin.H{"message": "product deleted successfully"})
}
