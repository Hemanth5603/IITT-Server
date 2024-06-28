package adminutils

import "github.com/Hemanth5603/IITT-Server/infrastructure"

func DBRejectData(dataId int64) error {
	_, err := infrastructure.POSTGRES_DB.Exec("UPDATE data SET is_approved = $1 WHERE data_id = $2", 2, dataId)
	return err
}
