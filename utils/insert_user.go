package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertUser(payload models.SignUpRequest) (int64, error) {
	var id int64

	// Check if the user already exists with the given email
	var existingEmail string
	err := infrastructure.POSTGRES_DB.QueryRow(
		`SELECT email FROM users WHERE email = $1`, payload.Email,
	).Scan(&existingEmail)

	// If an error occurred during the query or an email is found, exit early
	if err == nil || err.Error() == "no rows in result set" {
		// User already exists
		return 0, nil
	} else if err != nil {
		// Error querying the database
		return 0, err
	}

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
