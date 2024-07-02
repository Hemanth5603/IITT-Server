package auth_utils

import "github.com/Hemanth5603/IITT-Server/infrastructure"

func DBHandleExpiredOtp(token string) error {
	_, err := infrastructure.POSTGRES_DB.Exec("DELETE FROM verification WHERE token = $1", token)
	return err
}
