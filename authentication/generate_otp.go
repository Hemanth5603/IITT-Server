package authentication

import (
	"fmt"
	"math/rand"
	"time"
)

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	return otp
}
