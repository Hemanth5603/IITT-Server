package models

type SignUpRequest struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Dob           string `json:"dob"`
	Phone         string `json:"phone"`
	Location      string `json:"location"`
	Contributions int64  `json:"contributions"`
	Rank          int64  `json:"rank"`
	ProfileImage  string `json:"profile_image"`
}

type LoginInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type ProfileUploadRequest struct {
	Id           int64  `form:"id"`
	ProfileImage string `form:"profile_image"`
}

type GetLeaderboardRequest struct {
	Limit    int64  `json:"limit"`
	Category string `json:"category"`
}
