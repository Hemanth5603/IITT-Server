package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func FetchLeaderBoardFromDB() ([]models.UserModel, error) {

	var userList []models.UserModel

	rows, err := infrastructure.POSTGRES_DB.Query(
		"SELECT id, name, email, phone, dob, location, contributions, rank, profile_image FROM users ORDER BY contributions DESC")

	if err != nil {
		return userList, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id            int64
			name          string
			email         string
			phone         string
			dob           string
			location      string
			contributions int64
			rank          int64
			profile_image string
		)
		if err := rows.Scan(&id, &name, &email, &phone, &dob, &location, &contributions, &rank, &profile_image); err != nil {
			return userList, err
		}

		userList = append(userList, models.UserModel{
			Id:            id,
			Name:          name,
			Email:         email,
			Phone:         phone,
			Dob:           dob,
			Location:      location,
			Contributions: contributions,
			Rank:          rank,
			ProfileImage:  profile_image,
		})

	}
	if len(userList) == 0 {
		return []models.UserModel{}, nil
	}

	return userList, nil
}
