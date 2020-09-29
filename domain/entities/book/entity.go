package book

import "time"

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
