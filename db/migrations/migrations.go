package migrations

import (
	"github.com/Focinfi/sakura/app/models"
	"github.com/Focinfi/sakura/db"
)

func init() {
	migrate(&models.User{})
}

func migrate(tables ...interface{}) {
	for _, table := range tables {
		db.DB.DropTableIfExists(table)
		db.DB.AutoMigrate(table)
	}
}
