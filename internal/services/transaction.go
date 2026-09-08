package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GetBalance() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	balance := rand.Float64() * 1000
	return balance, nil
}