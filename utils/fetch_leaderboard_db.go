package utils

import (
	"database/sql"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func FetchLeaderBoardFromDB(limit int64, category string) ([]models.UserModel, error) {
	println(category)
	println(limit)

	var userList []models.UserModel
	var query string
	var rows *sql.Rows
	var err error

	if category == "Default" && limit == 0 {
		query = "SELECT id, name, email, phone, dob, location, contributions, rank, profile_image FROM users ORDER BY contributions DESC LIMIT 100"
		rows, err = infrastructure.POSTGRES_DB.Query(query)
	} else if category != "Default" && limit == 0 {
		query = `
			SELECT u.id, name, email, phone, dob, location, contributions, rank, profile_image
			FROM users u
			JOIN (
				SELECT id, COUNT(id) AS frequency
				FROM data
				WHERE category = $1
				GROUP BY id
			) d ON u.id = d.id
			ORDER BY d.frequency DESC LIMIT 100;`
		rows, err = infrastructure.POSTGRES_DB.Query(query, category)
	} else if category == "Default" && limit != 0 {
		query = "SELECT id, name, email, phone, dob, location, contributions, rank, profile_image FROM users ORDER BY contributions DESC LIMIT $1"
		rows, err = infrastructure.POSTGRES_DB.Query(query, limit)
	} else if category != "Default" && limit != 0 {
		query = `
			SELECT u.id, name, email, phone, dob, location, contributions, rank, profile_image
			FROM users u
			JOIN (
				SELECT id, COUNT(id) AS frequency
				FROM data
				WHERE category = $1
				GROUP BY id
			) d ON u.id = d.id
			ORDER BY d.frequency DESC LIMIT $2;`
		rows, err = infrastructure.POSTGRES_DB.Query(query, category, limit)
	}
	print(query)

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
	if err = rows.Err(); err != nil {
		return userList, err
	}

	if len(userList) == 0 {
		return []models.UserModel{}, nil
	}

	return userList, nil
}
