package authentication

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBHandleVerifyOtp(token string, otp string) (error, int) {
	var tokenExist string
	var actualOtp string
	err := infrastructure.POSTGRES_DB.QueryRow(
		`SELECT token FROM verification WHERE token = $1`, token,
	).Scan(&tokenExist)
	fmt.Println(tokenExist)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, 404
		}
		return err, 0
	} else {
		err = infrastructure.POSTGRES_DB.QueryRow(
			"SELECT otp FROM verification WHERE token = $1", token,
		).Scan(&actualOtp)
		println(actualOtp)
		if err != nil {
			println("Error 0 Occoured")
			return err, 0
		}

		if otp != actualOtp {
			return err, 400
		}

	}
	return nil, 200
}
