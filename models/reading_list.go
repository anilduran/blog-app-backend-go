package models

import "time"

type ReadingList struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Posts       []*Post    `gorm:"many2many:reading_list_posts;" json:"posts"`
	UserID      uint       `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
