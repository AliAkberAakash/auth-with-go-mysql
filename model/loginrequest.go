package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr LoginRequest) IsValid() bool {
	return lr.Email != "" && lr.Password != ""
}
