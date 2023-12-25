package main

import (
	"booking-room-app/delivery"
	"fmt"
)

func main() {
	delivery.NewServer().Run()
}

func CalculateArea(panjang int, lebar int) string {
	if panjang > 0 && lebar > 0 {
		if panjang > 100 && lebar > 100 {
			area := panjang * lebar
			return fmt.Sprintf("luas persegi panjang huge: %d", area)
		}
		area := panjang * lebar
		return fmt.Sprintf("luas persegi panjang normal: %d", area)
	}

	return "invalid panjang or lebar"
}
