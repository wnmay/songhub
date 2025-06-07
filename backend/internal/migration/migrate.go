package migration

import (
	"log"

	"github.com/wnmay/songhub/backend/internal/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    err := db.AutoMigrate(
        &entities.User{},
        &entities.Artist{},
        &entities.Listener{},
        &entities.Song{},
    )
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    log.Println("Migration completed successfully.")
}
