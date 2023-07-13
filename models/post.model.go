package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title     string    `gorm:"uniqueIndex;not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Image     string    `gorm:"not null" json:"image"`
	User      uuid.UUID `gorm:"not null" json:"user"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type CreatePostRequest struct {
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
	User    string `json:"user"`
}

type UpdatePost struct {
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Image     string    `json:"image,omitempty"`
	User      string    `json:"user,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
