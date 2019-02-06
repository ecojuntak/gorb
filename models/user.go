package models

import "time"

type User struct {
	ID        uint64     `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id" faker:"-"`
	Name      string     `gorm:"column:name; type:varchar(100)" json:"name" faker:"name"`
	Email     string     `gorm:"column:email; type:varchar(100); unique_index" json:"email" faker:"email"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at" faker:"timestamp"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at" faker:"timestamp"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" faker:"-"`
}
