package cli

import (
	"fmt"
	"go-trial-class/helpers"
	"os"
)

func MainMenu() {
	helpers.ClearScreen()
	fmt.Println("Selamat datang di Mini Ecommerce")
	fmt.Println("-----------------------------------")

	var input string
	fmt.Println("Tekan (1) untuk melanjutkan ke list produk")
	fmt.Println("Tekan (2) untuk melanjutkan ke list order ")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi ")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "1":
		ListProduct()
	case "2":
		ListOrder()
	case "q":
		fmt.Println("Terimakasih")
		os.Exit(1)
	default:
		MainMenu()
	}
}

func LoginMenu() {
	helpers.ClearScreen()
	fmt.Println("---Main Menu---")

	var input string
	fmt.Println("1. Register User")
	fmt.Println("2. Login")
	fmt.Println("Ketik 1 atau 2 : ")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "1":
		Register()
	case "2":
		Login()
	default:
		fmt.Println("Terimakasih")
		os.Exit(1)
	}
}
