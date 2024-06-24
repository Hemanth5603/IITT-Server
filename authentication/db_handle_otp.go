package authentication

import (
	"fmt"
	"time"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBHandleOTP(token string, otp string) error {
	_, err := infrastructure.POSTGRES_DB.Exec(
		`INSERT INTO verification(token, otp) values($1, $2)`,
		token,
		otp,
	)
	if err != nil {
		return err
	}
	fmt.Println("inserted otp")

	go func() {
		time.Sleep(time.Second * 60)
		_, err = infrastructure.POSTGRES_DB.Exec(
			"DELETE FROM verification WHERE token = $1", token)
		fmt.Println("Delete otp after 60 sec")
	}()

	return err

}
