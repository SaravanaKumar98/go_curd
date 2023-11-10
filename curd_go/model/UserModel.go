package model

// type User struct {
// 	Uuid   string `validate:"required",json:"uuid"`
// 	Name   string `validate:"required",json:"name"`
// 	Email  string `validate:"required",json:"eamil"`
// 	Mobile string `validate:"required,len=10",json:"mobile"`
// 	DOB    string `validate:"required",json:"dob"`
// 	Age    int    `validate:"required,numeric",json:"age"`
// 	Gender string `validate:"required",json:"gender"`
// }

type User struct {
	Uuid   string `validate:"required" json:"uuid"`
	Name   string `validate:"required" json:"name"`
	Email  string `validate:"required" json:"email"` // Fixed typo here
	Mobile string `validate:"required,len=10" json:"mobile"`
	DOB    string `validate:"required" json:"dob"`
	Age    int    `validate:"required,numeric" json:"age"`
	Gender string `validate:"required" json:"gender"`
}
