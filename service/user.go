package service

import (
	"belajar-gorm/models"
	"fmt"
	"log"
)

func GetAllUser() {
	var users []models.User
	result := models.DB.Find(&users)

	if result.Error != nil {
		log.Fatalf("Failed to retrieve Users: %v", result.Error)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, *user.Email)
	}
}

func InsertUser(nama string, email string) {
	user := models.User{Name: nama, Email: &email}
	models.DB.Create(&user)
}
