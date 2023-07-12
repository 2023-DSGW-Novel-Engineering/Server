package initializers

import "github.com/2023-DSGW-Novel-Engineering/cation-backend/models"

func SyncDatabase() {
	DB.AutoMigrate(new(models.User))
	DB.AutoMigrate(new(models.Friend))
}
