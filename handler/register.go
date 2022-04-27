package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AliAkberAakash/auth-with-go-mysql/controller"
	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"github.com/AliAkberAakash/auth-with-go-mysql/response"
)

type registerHandler struct{}

func newRegisterHandler() http.Handler {
	return &registerHandler{}
}

func (h *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userCredential model.User

	err := json.NewDecoder(r.Body).Decode(&userCredential)

	if err != nil {
		json.NewEncoder(w).Encode(
			response.GetCommonResponse(
				http.StatusBadRequest,
				"Invalid Request Body",
				nil,
			),
		)
		return
	}

	controller.Register(userCredential)

	json.NewEncoder(w).Encode(
		response.GetCommonResponse(
			http.StatusCreated,
			"User  Created Successfully",
			userCredential,
		),
	)
	return
}

func GetRegisterHandler() http.Handler {
	return GetCommonHeader(newRegisterHandler())
}
