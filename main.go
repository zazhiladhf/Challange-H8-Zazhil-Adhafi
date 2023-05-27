package main

import (
	"go-trial-class/cli"
	"go-trial-class/config"
)

func main() {
	config.DBConnect()
	// defer cli.MainMenu()
	cli.LoginMenu()
}
