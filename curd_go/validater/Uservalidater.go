package validater

import (
	"curd/model"

	"github.com/go-playground/validator/v10"
)

func UserAdd(user *model.User) error {
	validate := validator.New()
	err := validate.StructExcept(user, "Age", "Uuid")
	return err
}

func UserUpdate(user *model.User) error {
	validate := validator.New()
	err := validate.StructExcept(user, "Age")
	return err
}

func UserReq(req *model.Request) error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
