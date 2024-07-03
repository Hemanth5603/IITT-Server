package utils

import "github.com/Hemanth5603/IITT-Server/infrastructure"

func DBDeleteData(dataId int64) error {
	_, err := infrastructure.POSTGRES_DB.Exec(
		"DELETE FROM data WHERE data_id = $1", dataId)

	return err
}
