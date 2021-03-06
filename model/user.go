package model

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u User) IsValid() bool {
	return u.Email != "" && u.Password != ""
}
