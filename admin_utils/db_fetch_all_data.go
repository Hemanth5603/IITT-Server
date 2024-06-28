package adminutils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func DBFetchUnApprovedData() ([]models.DataModel, error) {
	var dataList []models.DataModel

	rows, err := infrastructure.POSTGRES_DB.Query(
		"SELECT id, latitude, longitude, image, category, remarks, address, date, time, is_approved FROM data WHERE is_approved = $1",
		0,
	)
	if err != nil {
		return dataList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id          int64
			latitude    float64
			longitude   float64
			image       string
			category    string
			remarks     string
			address     string
			date        string
			time        string
			is_approved int64
		)
		if err := rows.Scan(&id, &latitude, &longitude, &image, &category, &remarks, &address, &date, &time, &is_approved); err != nil {
			return dataList, err
		}
		dataList = append(dataList, models.DataModel{
			Id:         id,
			Latitude:   latitude,
			Longitude:  longitude,
			Image:      image,
			Category:   category,
			Remarks:    remarks,
			Address:    address,
			Date:       date,
			Time:       time,
			IsApproved: is_approved,
		})

	}
	if len(dataList) == 0 {
		return []models.DataModel{}, nil
	}
	return dataList, nil
}
