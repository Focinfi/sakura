package migrations

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/db"
)

func init() {
	migrate(
		&models.User{},
		&models.ActivityNote{},
		&models.Tag{},
		&models.ActivityNoteTag{},
		&models.ActivityCategory{},
		&models.ActivityNoteCategory{},
		&models.Feeling{},
		&models.UserFeeling{},
	)
}

func migrate(tables ...interface{}) {
	for _, table := range tables {
		db.DB.DropTableIfExists(table)
		db.DB.AutoMigrate(table)
	}
}
