package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReadingList struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"image_url"`
	Posts       []*Post    `gorm:"many2many:reading_list_posts;" json:"posts"`
	UserID      uuid.UUID  `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (readingList *ReadingList) BeforeCreate(tx *gorm.DB) (err error) {
	readingList.ID = uuid.New()
	return
}
