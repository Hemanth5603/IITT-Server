package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func InsertUser(payload models.SignUpRequest) (int64, error) {
	var id int64

	err := infrastructure.POSTGRES_DB.QueryRow(
		`INSERT INTO users(name, email, password) values ($1, $2, $3) RETURNING id`,
		payload.Name,
		payload.Email,
		payload.Password,
	).Scan(&id)

	return id, err
}
