package handler

import (
	"clean/domain/entities/book"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BookHandler struct {
	m book.Manager
}

func NewBookHandler(r *gin.RouterGroup, s book.Manager) {
	handler := &BookHandler{
		m: s,
	}
	r.POST("/", handler.createBook)
}

func (b *BookHandler) createBook(c *gin.Context) {
	var book book.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		log.Info("Something happened.", err)
	}

	id, err := b.m.Create(&book)

	if err != nil {
		log.Fatal("waah")
	}

	c.JSON(200, gin.H{
		"id": id,
	})

}
