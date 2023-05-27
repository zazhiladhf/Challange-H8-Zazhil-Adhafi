package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
	"os"
)

func Register() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)
	var username string
	var email string
	var password string

	fmt.Println("Masukkan username anda: ")
	username, _ = consoleReader.ReadString('\n')

	fmt.Println("Masukkan password anda: ")
	password, _ = consoleReader.ReadString('\n')

	user := entity.User{
		UserName: username,
		Email:    email,
		Password: password,
	}

	err := config.DB.Create(&user).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("Registrasi Berhasil")
}

func Login() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)
	var Username string
	var email string
	var password string

	fmt.Println("Masukkan username anda: ")
	Username, _ = consoleReader.ReadString('\n')

	fmt.Println("Masukkan password anda: ")
	password, _ = consoleReader.ReadString('\n')

	user := entity.User{
		UserName: Username,
		Email:    email,
		Password: password,
	}

	err := config.DB.Where("user_name = ? AND password = ?", Username, password).First(&user).Error
	if err != nil {
		ErrorLoginHandler("Silakan Login Kembali")
	}

	fmt.Println("Selamat datang " + Username)

	var input string
	fmt.Println("Tekan (m) untuk ke menu utama")
	fmt.Println("Tekan (l) untuk logout")

	_, err = fmt.Scanln(&input)
	if err != nil {
		MainMenu()
	}

	switch input {
	case "m":
		MainMenu()
	case "l":
		LoginMenu()
	default:
		fmt.Println("Terimaksih telah menggunakan aplikasi ini")
		os.Exit(1)
	}

}
