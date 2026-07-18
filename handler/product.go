package handler

import (
	"database/sql"
	"log"
	"onlineshop/model"

	"github.com/google/uuid"

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
		c.JSON(200, gin.H{
			"status":        "success",
			"response_code": 200,
			"message":       "Berhasil mengambil daftar produk",
			"data":          products,
		})
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
		c.JSON(200, gin.H{
			"status":        "success",
			"response_code": 200,
			"message":       "Berhasil mengambil detail produk",
			"data":          product,
		})
		return
	}
}

func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Product
		if err := c.BindJSON(&product); err != nil {
			log.Printf("Terjadi Kesalahan saat membaca requests body: %v", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		product.ID = uuid.New().String()

		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("Gagal membuat product: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{
			"status":        "success",
			"response_code": 201,
			"message":       "Produk berhasil ditambahkan",
			"data":          product,
		})
		return
	}
}

func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var productReq model.Product
		if err := c.BindJSON(&productReq); err != nil {
			log.Printf("Terjadi Kesalahan saat membaca requests body: %v", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		product, err := model.SelectProductByID(db, id)

		if err != nil {
			log.Printf("Terjadi Kesalahan dalam mengambil product: %v", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		if productReq.Name != " " {
			product.Name = productReq.Name
		}

		if productReq.Price != 0 {
			product.Price = productReq.Price
		}

		if err := model.UpdateProduct(db, product); err != nil {
			log.Printf("Gagal memperbaharui product: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"status":        "success",
			"response_code": 200,
			"message":       "Produk berhasil diperbarui",
			"data":          product,
		})
		return
	}
}

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := model.DeleteProduct(db, id); err != nil {
			log.Printf("Terjadi Kesalahan dalam menghapus product: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"status":        "success",
			"response_code": 200,
			"message":       "Product berhasil dihapus",
			"data":          nil,
		})
		return
	}
}
