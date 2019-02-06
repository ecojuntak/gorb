package database

import (
	"github.com/ecojuntak/gorb/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) (err error) {
	db.AutoMigrate(models.User{})

	return
}
