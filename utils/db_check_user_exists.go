package utils

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBCheckUserExists(email string) (int64, int64, error) {

	var existingEmail string
	var id int64
	err := infrastructure.POSTGRES_DB.QueryRow(
		`SELECT email, id FROM users WHERE email = $1`, email,
	).Scan(&existingEmail, &id)
	fmt.Print(existingEmail)

	if err == nil || err.Error() == "no rows in result set" {

		return id, 0, err

	}
	return 0, 404, nil
}
