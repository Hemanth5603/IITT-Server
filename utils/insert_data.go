package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertData(data models.DataModel) error {

	_, err := infrastructure.POSTGRES_DB.Exec(
		`INSERT INTO data(id, latitude, longitude, image, category) values($1, $2, $3, $4, $5)`,
		data.Id,
		data.Latitude,
		data.Longitude,
		data.Image,
		data.Category,
	)

	return err

}
