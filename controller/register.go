package controller

import (
	"errors"

	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"github.com/AliAkberAakash/auth-with-go-mysql/repository"
	"golang.org/x/crypto/bcrypt"
)

func Register(user model.User) error {
	if user.IsValid() {

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
		repository.Register(user)
		return nil
	}

	return errors.New("Invalid user data")
}
