package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func DBUpdateProfile(updateValues models.ProfileUpdateRequest, profileImage string) error {

	fullName := updateValues.FirstName + " " + updateValues.LastName

	_, err := infrastructure.POSTGRES_DB.Exec(`UPDATE users SET first_name = $1, last_name = $2, dob = $3, occupation = $4, profile_image = $5, name = $6 WHERE id = $7`,
		updateValues.FirstName,
		updateValues.LastName,
		updateValues.Dob,
		updateValues.Occupation,
		profileImage,
		fullName,
		updateValues.Id,
	)

	return err
}
