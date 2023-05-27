package cli

import (
	"fmt"
	"os"
)

func ErrorHandler(msg string) {
	fmt.Println("Terjadi kesalahan dalam aplikasi")
	fmt.Println(msg)

	var input string
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "m":
		MainMenu()
	case "q":
		fmt.Println("Terimakasih telah mengguanakan aplikasi ini")
		os.Exit(1)
	default:
		ErrorHandler(msg)
	}
}

func ErrorLoginHandler(msg string) {
	fmt.Println("Username atau Password anda salah")
	// fmt.Println(msg)

	var input string
	fmt.Println("Tekan (1) untuk Login kembali")
	fmt.Println("Tekan (any key) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "1":
		Login()
	// case "q":
	// 	fmt.Println("Terimakasih telah mengguanakan aplikasi ini")
	// 	os.Exit(1)
	default:
		fmt.Println("Terimakasih telah mengguanakan aplikasi ini")
		os.Exit(1)
		// ErrorHandler(msg)
	}
}
