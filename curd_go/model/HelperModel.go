package model

type Request struct {
	Uuid string `validate:"required" json:"uuid"`
}
