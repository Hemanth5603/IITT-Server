package models

type UserModel struct {
	Id       int64  `json:"id"`
	Email    string `json:email`
	Name     string `json:name`
	Password string `json:password`
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
	Email    string `json:email`
	Name     string `json:name`
	Password string `json:password`
}

type LoginInRequest struct {
	Email    string `json:email`
	Password string `json:password`
}
