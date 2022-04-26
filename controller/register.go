package controller

import (
	"errors"

	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"golang.org/x/crypto/bcrypt"
)

var Users []model.User

func Register(user model.User) error {
	if user.IsValid() {

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
		Users = append(Users, user)
		return nil
	}

	return errors.New("Invalid user data")
}
