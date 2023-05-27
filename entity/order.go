package entity

import (
	"fmt"
	"time"
)

type Order struct {
	ID           int
	ProductId    int
	BuyerEmail   string
	BuyerAddress string
	OrderDate    time.Time
	Product      Product
}

func (o *Order) PrintDetail() {
	fmt.Println("ID : ", o.ID)
	fmt.Println("Produk : ", o.Product.Name)
	fmt.Println("Email : ", o.BuyerEmail)
	fmt.Println("Address : ", o.BuyerAddress)
	fmt.Println("Order Date : ", o.OrderDate)
}
