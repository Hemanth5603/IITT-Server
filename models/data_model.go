package models

type UserModel struct {
	Id            int64  `json:"id"`
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

type DataModel struct {
	Id        int64   `form:"id"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
	Image     string  `form:"image"`
	Category  string  `form:"category"`
	Remarks   string  `form:"remarks"`
}

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
	Email    string `json:email`
	Password string `json:password`
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type ProfileUploadRequest struct {
	Id           int64  `form:"id"`
	ProfileImage string `form:"profile_image"`
}
