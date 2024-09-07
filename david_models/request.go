package davidmodels

type FindFaceRequest struct {
	Image string `form:"image"`
}

type RegisterStudentRequest struct {
	Roll         string `json:"roll"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Branch       string `json:"branch"`
	AcademicYear int    `json:"academicyear"`
	Semester     int    `json:"semester"`
	Section      string `json:"section"`
}
