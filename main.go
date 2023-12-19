package main

import (
	"booking-room-app/delivery"
	"fmt"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("init run")
}

func main() {
	delivery.NewServer().Run()
}