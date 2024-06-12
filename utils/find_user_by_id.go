package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func FindUserById(id int64) (user models.UserModel, err error) {
	err = infrastructure.POSTGRES_DB.QueryRow("SELECT id, name, email, password, phone, location, dob, contributions, rank FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Location, &user.Dob, &user.Contributions, &user.Rank)

	return user, err
}
