package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", code)
}
