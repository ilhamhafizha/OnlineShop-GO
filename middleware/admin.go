package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: ambil header authoratization
		keys := os.Getenv("ADMIN_SECRET")

		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{
				"status":        "error",
				"response_code": 401,
				"message":       "Authorization header tidak ditemukan",
				"data":          nil,
			})
			c.Abort()
			return
		}

		// TODO : validasi header sesuai dengan kata sandi admin
		if auth != keys {
			c.JSON(401, gin.H{
				"status":        "error",
				"response_code": 401,
				"message":       "Akses tidak diizinkan",
				"data":          nil,
			})
			c.Abort()
			return
		}

		// TODO : Lanjutkan proses jika berhasil ke handler
		c.Next()
	}
}
