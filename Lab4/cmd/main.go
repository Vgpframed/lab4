package main

import (
db "lab4/internal/DB"
product "lab4/internal/products"
	"os"
)
var Filename = string(os.Args[1])
func main() {
	var prod product.Product
	products := prod.GetProducts(Filename)
	db.InsertProducts(products)
	db.SelectAllProducts()
	db.UserSelect()
	db.UpdDel()
}