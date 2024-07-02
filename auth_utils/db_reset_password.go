package auth_utils

import "github.com/Hemanth5603/IITT-Server/infrastructure"

func DBResetPassword(id int64, newPassword string) error {
	_, err := infrastructure.POSTGRES_DB.Exec(`UPDATE users SET password = $1 WHERE id = $2`,
		newPassword,
		id,
	)
	return err

}
