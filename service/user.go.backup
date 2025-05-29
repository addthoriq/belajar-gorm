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
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, *user.Email)
}

func DetailUser(id uint) {
	var user models.User
	result := models.DB.Find(&user, id)
	if result.Error != nil {
		log.Fatalf("Failed to retrieve User: %v", result.Error)
	}
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, *user.Email)
}

func UpdateUser(id uint, nama string, email string) {
	var user models.User
	models.DB.Model(&user).Where("id = ?", id).Updates(models.User{Name: nama, Email: &email})

	fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, user.Name, *user.Email)
}

func DeleteUser(id uint) {
	// models.DB.Where("id = ?", id).Delete(&user)
	models.DB.Delete(models.User{}, id)
	fmt.Println("Deleted Successfully!")
}
