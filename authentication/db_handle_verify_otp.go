package authentication

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBHandleVerifyOtp(token string, otp string) (int, error) {
	var tokenExist string
	var actualOtp string
	err := infrastructure.POSTGRES_DB.QueryRow(
		`SELECT token FROM verification WHERE token = $1`, token,
	).Scan(&tokenExist)
	fmt.Println(tokenExist)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return 404, nil
		}
		return 0, err
	} else {
		err = infrastructure.POSTGRES_DB.QueryRow(
			"SELECT otp FROM verification WHERE token = $1", token,
		).Scan(&actualOtp)
		println(actualOtp)
		if err != nil {
			println("Error 0 Occoured")
			return 0, err
		}

		if otp != actualOtp {
			return 400, err
		}

	}
	return 200, nil
}
