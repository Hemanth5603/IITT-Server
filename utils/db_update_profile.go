package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func DBUpdateProfile(updateValues models.ProfileUpdateRequest, profileImage string) error {

	fullName := updateValues.FirstName + " " + updateValues.LastName

	_, err := infrastructure.POSTGRES_DB.Exec(`UPDATE users SET firstname = $1, lastname = $2, dob = $3, occupation = $4, profileimage = $5, name = $6, phone = $7 WHERE id = $8`,
		updateValues.FirstName,
		updateValues.LastName,
		updateValues.Dob,
		updateValues.Occupation,
		profileImage,
		fullName,
		updateValues.Phone,
		updateValues.Id,
	)

	return err
}
