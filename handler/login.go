package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AliAkberAakash/auth-with-go-mysql/controller"
	"github.com/AliAkberAakash/auth-with-go-mysql/model"
	"github.com/AliAkberAakash/auth-with-go-mysql/response"
)

type loginHandler struct{}

func newLoginHandler() http.Handler {
	return &loginHandler{}
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var lr model.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&lr)

	if err != nil {
		json.NewEncoder(w).Encode(
			response.GetCommonResponse(
				http.StatusBadRequest,
				"Invalid body",
				nil,
			),
		)
		return
	}

	_, err = controller.Login(lr)

	if err != nil {
		json.NewEncoder(w).Encode(
			response.GetCommonResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			),
		)
		return
	}

	json.NewEncoder(w).Encode(
		response.GetCommonResponse(
			http.StatusOK,
			"Logged in successfully",
			nil,
		),
	)
	return

}

func GetLoginHandler() http.Handler {
	return GetCommonHeader(newLoginHandler())
}
