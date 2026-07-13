package handler

import (
	"database/sql"
	"onlineshop/model"

	"github.com/gin-gonic/gin"
)

func ListProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Todo : Ambil dari database
		products, err := model.SelectProducts(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Todo : Berikan respon
		c.JSON(200, gin.H{"products": products})
		return
	}
}

// func GetProduct(c *gin.Context){
// 	// Todo : Baca id dari url
//
// 	// todo : Ambil dari database
//
// 	// Todo : Berikan respon
// }