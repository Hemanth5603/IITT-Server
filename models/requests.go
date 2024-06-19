package models

type SignUpRequest struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Password      string `json:"password"`
	Dob           string `json:"dob"`
	Phone         string `json:"phone"`
	Location      string `json:"location"`
	Contributions int64  `json:"contributions"`
	Rank          int64  `json:"rank"`
	ProfileImage  string `json:"profile_image"`
	State         string `json:"state"`
	City          string `json:"city"`
	Pincode       int64  `json:"pincode"`
	Occupation    string `json:"occupation"`
}

type LoginInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type ProfileUpdateRequest struct {
	Id           int64  `form:"id"`
	FirstName    string `form:"first_name"`
	LastName     string `form:"last_name"`
	ProfileImage string `form:"profile_image"`
	Dob          string `form:"dob"`
	Occupation   string `form:"occupation"`
}

type GetLeaderboardRequest struct {
	Limit    int64  `json:"limit"`
	Category string `json:"category"`
}
