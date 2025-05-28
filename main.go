package main

import (
	"belajar-gorm/models"
	"belajar-gorm/service"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("A: Connect DB")
	fmt.Println("B: Get All User")
	fmt.Println("C: Insert User")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	scan := bufio.NewScanner(os.Stdin)

	switch char {
	case 'A':
		models.ConnectDB()
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
	}

}
