package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Content string    `json:"content"`
	PostID  uuid.UUID `json:"post_id"`
	UserID  uuid.UUID `json:"user_id"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	comment.ID = uuid.New()
	return
}
