package seed

import (
	"log"
	"url-shortner/src/config"
	userModel "url-shortner/src/modules/user/model"
)

func MigrateDB() {
	if config.DB == nil {
		log.Fatal("Database connection is nil")
		return
	}

	err := config.DB.AutoMigrate(
		&userModel.User{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
