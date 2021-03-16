package database

import (
	"github.com/jinzhu/gorm"
	"github.com/tutorialedge/production-ready-api/internal/comment"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
