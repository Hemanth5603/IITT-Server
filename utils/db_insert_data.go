package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertData(data models.DataModelRequest) error {
	//currentTime := time.Now()

	// Format the time in HH:MM:SS format
	//time := currentTime.Format("15:04:05")
	//date := currentTime.Format("2006-01-02")

	_, err := infrastructure.POSTGRES_DB.Exec(
		`INSERT INTO data(id, latitude, longitude, image, category, remarks, address, date, time, is_approved) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		data.Id,
		data.Latitude,
		data.Longitude,
		data.Image,
		data.Category,
		data.Remarks,
		data.Address,
		data.Date,
		data.Time,
		data.IsApproved,
	)
	if err != nil {
		return err
	}
	// Updating Contributions Count

	_, err = infrastructure.POSTGRES_DB.Exec(`UPDATE users SET contributions = contributions + 1 WHERE id = $1`, data.Id)

	return err

}
