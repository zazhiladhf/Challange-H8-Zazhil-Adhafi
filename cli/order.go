package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
	"os"
)

func ListOrder() {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)
	var orders []entity.Order

	err := config.DB.Preload("Product").Find(&orders).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("---List Order---")
	for _, order := range orders {
		order.PrintDetail()
	}

	var input string
	fmt.Println("Tekan (any key) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, _ = consoleReader.ReadString('\n')

	switch input {
	case "q\n":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu()
	}
}
