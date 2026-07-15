package handler

import (
	"database/sql"
	"log"
	"onlineshop/model"

	"github.com/gin-gonic/gin"
)

func ListProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Todo : Ambil dari database
		products, err := model.SelectProducts(db)
		if err != nil {
			log.Printf("Gagal query data product: %v", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Todo : Berikan respon
		c.JSON(200, gin.H{"products": products})
		return
	}
}

func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Todo : Baca id dari url
		id := c.Param("id")
		// todo : Ambil dari database
		product, err := model.SelectProductByID(db, id)
		if err != nil {
			log.Printf("Gagal query data product: %v", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}
		// Todo : Berikan respon
		c.JSON(200, gin.H{"product": product})
		return
	}
}
