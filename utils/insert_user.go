package utils

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertUser(payload models.SignUpRequest) (int64, error) {
	var id int64

	var existingEmail string
	err := infrastructure.POSTGRES_DB.QueryRow(
		`SELECT email FROM users WHERE email = $1`, payload.Email,
	).Scan(&existingEmail)
	fmt.Print(existingEmail)

	if err == nil || err.Error() == "no rows in result set" {

		return 0, err

	} else {
		err = infrastructure.POSTGRES_DB.QueryRow(
			`INSERT INTO users(name, email, password, location, dob, phone, contributions, rank, profile_image) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
			payload.Name,
			payload.Email,
			payload.Password,
			payload.Location,
			payload.Dob,
			payload.Phone,
			payload.Contributions,
			payload.Rank,
			payload.ProfileImage,
		).Scan(&id)
		return id, err
	}

}
