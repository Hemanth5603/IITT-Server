package davidutils

import (
	davidmodels "github.com/Hemanth5603/IITT-Server/david_models"
	"github.com/Hemanth5603/IITT-Server/infrastructure"
)

func DBFetchStudentRoll(rollNo string) (studentModel davidmodels.StudentModel, err error) {
	err = infrastructure.POSTGRES_DB.QueryRow("SELECT id, name, roll, phone, branch, academic_year, semester, sec FROM allstudents WHERE roll = $1", rollNo).
		Scan(&studentModel.Id, &studentModel.Name, &studentModel.Roll, &studentModel.Phone, &studentModel.Branch, &studentModel.AcademicYear, &studentModel.Semester, &studentModel.Section)

	return studentModel, err
}
