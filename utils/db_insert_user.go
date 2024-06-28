package utils

import (
	"fmt"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertUser(payload models.SignUpRequest) (int64, error) {
	var id int64

	var existingPhone string
	err := infrastructure.POSTGRES_DB.QueryRow(
		`SELECT phone FROM users WHERE phone = $1`, payload.Phone,
	).Scan(&existingPhone)
	fmt.Print(existingPhone)

	if err == nil || err.Error() == "no rows in result set" {

		return 0, err

	} else {
		err = infrastructure.POSTGRES_DB.QueryRow(
			`INSERT INTO users(name, email, password, location, dob, phone, contributions, rank, profileimage, firstname, lastname, state, city, pincode, occupation) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id`,
			payload.Name,
			payload.Email,
			payload.Password,
			payload.Location,
			payload.Dob,
			payload.Phone,
			payload.Contributions,
			payload.Rank,
			payload.ProfileImage,
			payload.FirstName,
			payload.LastName,
			payload.State,
			payload.City,
			payload.Pincode,
			payload.Occupation,
		).Scan(&id)
		return id, err
	}

}
