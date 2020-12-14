package models

import (
	"time"
)

type Book struct {
	ID            int
	ISBN          string    `json:"isbn"`
	Publisher     string    `json:"publisher"`
	Price         int       `json:"price"`
	Title         string    `json:"title"`
	Year          uint8     `json:"year"`
	Author        string    `json:"author"`
	CoverImage    string    `json:"cover_image"`
	Description   string    `json:"description"`
	PublishedDate time.Time `json:"published_date"`
	CategoryID    int       `json:"category_id"`
	Category      Category  `json:"category"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
