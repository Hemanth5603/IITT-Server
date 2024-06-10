package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func GetUploadsByUserID(user_id string) ([]models.DataModel, error) {
	var dataList []models.DataModel

	rows, err := infrastructure.POSTGRES_DB.Query(
		"SELECT id, latitude, longitude, image, category, remarks FROM data WHERE id = $1",
		user_id,
	)
	if err != nil {
		return dataList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        int64
			latitude  float64
			longitude float64
			image     string
			category  string
			remarks   string
		)
		if err := rows.Scan(&id, &latitude, &longitude, &image, &category, &remarks); err != nil {
			return dataList, err
		}
		dataList = append(dataList, models.DataModel{
			Id:        id,
			Latitude:  latitude,
			Longitude: longitude,
			Image:     image,
			Category:  category,
			Remarks:   remarks,
		})

	}
	if len(dataList) == 0 {
		return []models.DataModel{}, nil
	}
	return dataList, nil
}
