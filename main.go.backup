package main

import (
	"belajar-gorm/models"
	"belajar-gorm/service"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("A: Connect DB")
	fmt.Println("B: Get All User")
	fmt.Println("C: Insert User")
	fmt.Println("D: Detail User")
	fmt.Println("E: Update User")
	fmt.Println("F: Delete User")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	scan := bufio.NewScanner(os.Stdin)

	switch char {
	case 'A':
		models.ConnectDB()
		log.Println("Connected Database")
		break
	case 'B':
		models.ConnectDB()
		service.GetAllUser()
		break
	case 'C':
		models.ConnectDB()
		fmt.Print("Masukkan Nama: ")
		scan.Scan()
		nama := scan.Text()
		fmt.Print("Masukkan Email: ")
		scan.Scan()
		email := scan.Text()
		service.InsertUser(nama, email)
		break
	case 'D':
		models.ConnectDB()
		fmt.Print("Masukkan ID: ")
		scan.Scan()
		input_user := scan.Text()
		conv_input_user, _ := strconv.ParseUint(input_user, 10, 64)
		user_id := uint(conv_input_user)
		service.DetailUser(user_id)
		break
	case 'E':
		models.ConnectDB()
		fmt.Print("Masukkan ID: ")
		scan.Scan()
		input_user := scan.Text()
		conv_input_user, _ := strconv.ParseUint(input_user, 10, 64)
		user_id := uint(conv_input_user)
		fmt.Print("Masukkan Nama: ")
		scan.Scan()
		nama := scan.Text()
		fmt.Print("Masukkan Email: ")
		scan.Scan()
		email := scan.Text()
		service.UpdateUser(user_id, nama, email)
		break
	case 'F':
		models.ConnectDB()
		fmt.Print("Masukkan ID: ")
		scan.Scan()
		input_user := scan.Text()
		conv_input_user, _ := strconv.ParseUint(input_user, 10, 64)
		user_id := uint(conv_input_user)
		service.DeleteUser(user_id)
		break
	}
}
