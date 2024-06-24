package models

type UserModel struct {
	Id            int64  `json:"id"`
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
	State         string `json:"state"`
	City          string `json:"city"`
	Pincode       int64  `json:"pincode"`
	ProfileImage  string `json:"profileimage"`
	Occupation    string `json:"occupation"`
}

type DataModel struct {
	Id        int64   `form:"id"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
	Image     string  `form:"image"`
	Category  string  `form:"category"`
	Remarks   string  `form:"remarks"`
	Address   string  `form:"address"`
}

type OTPData struct {
	PhoneNumber string `json:"To,omitempty"`
}

type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
