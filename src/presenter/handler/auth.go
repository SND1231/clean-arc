package handler

import (
	"context"
	"ddd/infrastructure/db"
	"ddd/infrastructure/models"
	"ddd/infrastructure/setting"
	"fmt"
	"log"
	"net/http"

	"github.com/mholt/binding"
)

type AuthHandler struct {
	settings setting.Setting
}

func NewAuthHandler(settings setting.Setting) *AuthHandler {
	return &AuthHandler{
		settings: settings,
	}
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (req LoginRequest) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&req.ID:       "id",
		&req.Password: "password",
	}
}

type LoginResponse struct {
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	defer func() {
		if err != nil || berr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	}()
	var request LoginRequest
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		return
	}

	fmt.Println("getUsers 動いてるでー")
	// ここにsqlboilerをベタガキ
	ctx := context.Background()
	db, err := db.GetDBconnection(h.settings.DB)
	if err != nil {
		log.Println(err)
		return
	}
	isExists, err := models.Auths(models.AuthWhere.WorkerID.EQ(request.ID), models.AuthWhere.Password.EQ(request.Password)).Exists(ctx, db)
	if err != nil {
		log.Println("ここにあたる？")
		log.Println(err)
		return
	}
	if !isExists {
		err = fmt.Errorf("ID or Password not match")
		log.Println(err)
		return
	}

}
