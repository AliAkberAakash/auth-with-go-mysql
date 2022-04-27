package controller

import (
	"errors"

	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"github.com/AliAkberAakash/auth-with-go-mysql/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(lr model.LoginRequest) (bool, error) {
	if lr.IsValid() {
		users := repository.GetAllUsers()
		for _, user := range users {
			err := bcrypt.CompareHashAndPassword(
				[]byte(user.Password),
				[]byte(lr.Password),
			)

			if err == nil {
				return true, nil
			}
		}
	}
	return false, errors.New("Invalid credentials")
}
