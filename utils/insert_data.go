package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertData(data models.DataModel) error {

	_, err := infrastructure.POSTGRES_DB.Exec(
		`INSERT INTO data(id, latitude, longitude, image, category, remarks) values($1, $2, $3, $4, $5, $6)`,
		data.Id,
		data.Latitude,
		data.Longitude,
		data.Image,
		data.Category,
		data.Remarks,
	)
	if err != nil {
		return err
	}
	// Updating Contributions Count

	_, err = infrastructure.POSTGRES_DB.Exec(`UPDATE users SET contributions = contributions + 1 WHERE id = $1`, data.Id)

	return err

}
