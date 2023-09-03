package configs

import (
	"ijash-jwt-auth/src/models"
)

func Migration() {

	DB.AutoMigrate(&models.User{})
}
