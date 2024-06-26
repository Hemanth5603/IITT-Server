package utils

import (
	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/Hemanth5603/IITT-Server/models"
)

func FindUserById(id int64) (user models.UserModel, rank int64, err error) {
	err = infrastructure.POSTGRES_DB.QueryRow("SELECT id, name, firstname, lastname, email, password, phone, location, dob, contributions, rank, profileimage, state, city, pincode, occupation FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Phone, &user.Location, &user.Dob, &user.Contributions, &user.Rank, &user.ProfileImage, &user.State, &user.City, &user.Pincode, &user.Occupation)

	if err != nil {
		return models.UserModel{}, 0, err
	}

	query := `WITH RankedUsers AS (
    			SELECT 
        			id,
        			RANK() OVER (ORDER BY contributions DESC) AS rank
    			FROM 
        			users
			)
				SELECT 
    				id,
    				rank
				FROM 
    				RankedUsers
				WHERE 
    				id = $1`

	err = infrastructure.POSTGRES_DB.QueryRow(query, id).Scan(&id, &rank)

	return user, rank, err
}
