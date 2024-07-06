package utils

import "github.com/Hemanth5603/IITT-Server/infrastructure"

func DBDeleteAccount(Id int64) error {

	_, err := infrastructure.POSTGRES_DB.Exec(`DELETE FROM users WHERE id = $1`, Id)

	return err

}
