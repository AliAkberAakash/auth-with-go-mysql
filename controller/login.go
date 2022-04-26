package controller

import (
	"errors"

	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"golang.org/x/crypto/bcrypt"
)

func Login(lr model.LoginRequest) (bool, error) {
	if lr.IsValid() {
		for _, user := range Users {
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
