package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID                uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	ImageUrl          string         `json:"image_url"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	Content           string         `json:"content"`
	AuthorID          uuid.UUID      `json:"author_id"`
	Categories        []*Category    `gorm:"many2many:post_categories;"`
	Comments          []Comment      `json:"comments"`
	ReadingList       []*ReadingList `gorm:"many2many:reading_list_posts;" json:"reading_list"`
	BookmarkedByUsers []*User        `gorm:"many2many:bookmarks;" json:"bookmarked_by_users"`
	IsActive          bool           `json:"is_active" gorm:"default:true"`
	CreatedAt         *time.Time     `json:"created_at"`
	UpdatedAt         *time.Time     `json:"updated_at"`
	DeletedAt         *time.Time     `json:"deleted_at"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.New()
	return
}
