package controller

import "github.com/AliAkberAakash/auth-with-go-mysql/model"

var Users []model.User

func Register(user model.User) {
	if user.IsValid() {
		Users = append(Users, user)
	}
}
