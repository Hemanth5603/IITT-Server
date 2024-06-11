package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func FindUserByEmail(email string) (user models.UserModel, err error) {
	err = infrastructure.POSTGRES_DB.QueryRow("SELECT id, name, email, password, phone, location, dob, contributions, rank FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Location, &user.Dob, &user.Contributions, &user.Rank)

	return user, err
}
