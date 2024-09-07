package davidmodels

type StudentModel struct {
	Name         string `json:"name"`
	Roll         string `json:"roll"`
	Branch       string `json:"branch"`
	Phone        string `json:"phone"`
	AcademicYear int64  `json:"academicyear"`
	Semester     int64  `json:"semester"`
	Section      string `json:"section"`
}
