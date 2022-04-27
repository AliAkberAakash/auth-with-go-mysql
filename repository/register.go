package repository

import "github.com/AliAkberAakash/auth-with-go-mysql/model"

func Register(user model.User) {
	Users = append(Users, user)
}
