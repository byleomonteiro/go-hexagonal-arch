package main

import (
	"database/sql"

	"github.com/byleomonteiro/hexagonal-arch/adapters/db"
	"github.com/byleomonteiro/hexagonal-arch/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connection, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db.NewProductDb(connection)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)

}
