package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
	"os"
	"time"
)

func ListProduct() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)

	var products []entity.Product
	err := config.DB.Find(&products).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("---List Produk---")
	for _, product := range products {
		product.PrintDetail()
	}

	var input string
	fmt.Println("Masukkan ID Produk untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, err = consoleReader.ReadString('\n')

	switch input {
	case "m\n":
		MainMenu()
	case "q\n":
		fmt.Println("Terimaksih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		OrderProduct(input)
	}

}

func OrderProduct(id string) {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)

	var product entity.Product
	err := config.DB.Where("ID = ?", id).First(&product).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	product.PrintDetail()

	var input string
	fmt.Println("Tekan (y) untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, err = consoleReader.ReadString('\n')

	switch input {
	case "y\n":
		CreateOrder(product)
	case "m\n":
		MainMenu()
	case "q\n":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		OrderProduct(id)
	}

}

func CreateOrder(product entity.Product) {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)
	var email string
	var address string

	fmt.Println("Masukkan email anda: ")
	email, _ = consoleReader.ReadString('\n')

	fmt.Println("Masukkan alamat anda: ")
	address, _ = consoleReader.ReadString('\n')

	order := entity.Order{
		ProductId:    int(product.ID),
		BuyerEmail:   email,
		BuyerAddress: address,
		OrderDate:    time.Now(),
	}

	err := config.DB.Create(&order).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("Order berhasil dibuat")

	var input string
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err = fmt.Scanln(&input)
	if err != nil {
		MainMenu()
	}

	switch input {
	case "q\n":
		fmt.Println("Terimaksih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu()
	}

}
