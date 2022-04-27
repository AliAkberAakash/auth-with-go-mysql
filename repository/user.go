package repository

import "github.com/AliAkberAakash/auth-with-go-mysql/model"

var Users []model.User

func GetAllUsers() []model.User {
	return Users
}
