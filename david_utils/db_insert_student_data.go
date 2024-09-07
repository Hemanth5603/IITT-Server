package davidutils

import (
	davidmodels "github.com/Hemanth5603/IITT-Server/david_models"
	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBInsertStudent(payload davidmodels.RegisterStudentRequest) (int64, error) {
	var id int64

	err := infrastructure.POSTGRES_DB.QueryRow(
		`INSERT INTO allstudents(name, roll, phone, branch, academic_year, semester, sec) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		payload.Name,
		payload.Roll,
		payload.Phone,
		payload.Branch,
		payload.AcademicYear,
		payload.Semester,
		payload.Section,
	).Scan(&id)

	return id, err
}
