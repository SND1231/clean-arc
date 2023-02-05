package handler

import (
	"context"
	"ddd/infrastructure/db"
	"ddd/infrastructure/models"
	"ddd/infrastructure/setting"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mholt/binding"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	Session string `json:"session"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response LoginResponse
	defer func() {
		// JSONデータを書き込む
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		res, _ := json.Marshal(response)
		w.Write(res)
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

	fmt.Println("認証開始")
	ctx := context.Background()
	db, err := db.GetDBconnection(h.settings.DB)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(request)
	authes, err := models.Auths(models.AuthWhere.WorkerID.EQ(request.ID), models.AuthWhere.Password.EQ(request.Password)).All(ctx, db)
	if err != nil {
		log.Println(err)
		return
	}
	if len(authes) == 0 {
		err = fmt.Errorf("ID or Password not match")
		log.Println(err)
		return
	}
	dt := time.Now()
	unix := dt.Unix()
	var data string = fmt.Sprintf("%s:%s:%d", request.ID, request.Password, unix)
	session := base64.StdEncoding.EncodeToString([]byte(data))
	response.Session = session

	auth := authes[0]
	auth.Session = session
	auth.UpdatedAt = time.Now()
	auth.Update(ctx, db, boil.Infer())
}

type AddRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (req AddRequest) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&req.ID:       "id",
		&req.Password: "password",
		&req.Name:     "name",
	}
}

type AddResponse struct {
	ID string `json:"id"`
}

func (h *AuthHandler) Add(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response AddResponse
	defer func() {
		// JSONデータを書き込む
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		res, _ := json.Marshal(response)
		w.Write(res)
		if err != nil || berr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}()

	var request AddRequest
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		return
	}

	ctx := context.Background()
	db, err := db.GetDBconnection(h.settings.DB)
	if err != nil {
		log.Println(err)
		return
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	dt := time.Now()
	worker := models.Worker{
		WorkerID:  request.ID,
		Name:      request.Name,
		Status:    1,
		CreatedAt: dt,
		UpdatedAt: dt,
	}
	fmt.Println(worker)
	err = worker.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return
	}

	auth := models.Auth{
		WorkerID:            request.ID,
		Password:            request.Password,
		Session:             "",
		ChangePasswordCount: 0,
		FailAuthCount:       0,
		Status:              1,
		CreatedAt:           dt,
		UpdatedAt:           dt,
	}
	err = auth.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return
	}
	response.ID = request.ID
}
