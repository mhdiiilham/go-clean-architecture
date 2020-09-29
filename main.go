package main

import (
	"clean/config"
	"clean/domain/entities/book"
	"clean/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.SetupDB()
	port := config.GetPortConnection()
	db := config.GetDBConnection()

	db.AutoMigrate(book.Book{})

	bookRepo := book.NewMySQLRepository(db)
	bookManager := book.NewManager(bookRepo)

	b := r.Group("/books")
	{
		handler.NewBookHandler(b, bookManager)
	}

	r.Run(port)
}
