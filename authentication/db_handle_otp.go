package authentication

import (
	"fmt"
	"time"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBHandleOTP(token string, otp string) error {
	_, err := infrastructure.POSTGRES_DB.Exec(
		"INSERT INTO verification(token, otp) values($1, $2)",
		token,
		otp,
	)
	if err != nil {
		return err
	}
	fmt.Println("inserted otp")

	// Start a new goroutine to delete the OTP after 60 seconds
	go deleteOTPAfter60Seconds(token)

	return nil
}

func deleteOTPAfter60Seconds(token string) {
	time.Sleep(time.Second * 60)
	_, err := infrastructure.POSTGRES_DB.Exec(
		"DELETE FROM verification WHERE token = $1", token)
	if err != nil {
		fmt.Println("Error deleting OTP:", err)
	} else {
		fmt.Println("Deleted OTP after 60 sec")
	}
}
