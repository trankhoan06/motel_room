package common

import (
	"math/rand"
	"time"
)

func GenerateRandomCode() int {
	rand.Seed(time.Now().UnixNano()) // Khởi tạo seed cho số ngẫu nhiên
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
