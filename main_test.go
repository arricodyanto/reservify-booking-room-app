package main

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestCalculateArea_Success(t *testing.T) {
	_ = CalculateArea(4, 5)
	_ = CalculateArea(400, 500)
	_ = CalculateArea(0, 0)
	//assert.Equal(t, "luas persegi panjang huge: 20")
}
