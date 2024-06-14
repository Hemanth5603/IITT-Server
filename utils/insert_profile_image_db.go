package utils

import "github.com/Hemanth5603/IITT-Server/infrastructure"

func InsertProfileImageDB(Id int64, profileImage string) error {

	_, err := infrastructure.POSTGRES_DB.Exec(`UPDATE users SET profile_image = $1 WHERE id = $2`, profileImage, Id)

	return err
}
