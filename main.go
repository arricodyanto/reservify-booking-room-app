package main

import (
	"booking-room-app/delivery"
	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
