package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DB_URI"))
	
	if err != nil {
		fmt.Printf("Gagal Membuat koneksi ke database: %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("Gagal melakukan ping ke database: %v\n", err)
		os.Exit(1)
	}

	if _, err := migrate(db); err != nil {
		fmt.Printf("Gagal melakukan migrasi database: %v\n", err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	if err = server.ListenAndServe(); err != nil {
		fmt.Printf("Gagal menjalankan server: %v\n", err)
		os.Exit(1)
	}
}
