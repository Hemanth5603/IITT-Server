package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func FindUserByPhone(phone string) (user models.UserModel, err error) {
	err = infrastructure.POSTGRES_DB.QueryRow("SELECT id, name, email, password, phone, location, dob, contributions, rank FROM users WHERE phone = $1", phone).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Location, &user.Dob, &user.Contributions, &user.Rank)

	return user, err
}
