package quilltypes

import "gorm.io/gorm"

type Note struct {
	gorm.Model

	ID       int    `json:"id"`
	Etag     string `json:"etag"`
	Readonly bool   `json:"readonly"`
	Modified int    `json:"modified"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Favorite bool   `json:"favorite"`

	// hidden fields
	UserID uint
	User   User
}

type Notes []Note
