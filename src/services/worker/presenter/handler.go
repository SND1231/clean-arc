package presenter

import (
	"context"
	"ddd/infrastructure/db"
	"ddd/infrastructure/setting"
	"ddd/services/worker/infrastructure"
	"ddd/services/worker/usecase"
	"encoding/json"
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response usecase.LoginOutput
	defer func() {
		// JSONデータを書き込む
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		res, _ := json.Marshal(response)
		if err != nil || berr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		w.Write(res)
	}()

	/*
		var request usecase.LoginInput
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
	*/
}

func (h *AuthHandler) Add(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response usecase.AddOutput
	defer func() {
		// JSONデータを書き込む
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		res, _ := json.Marshal(response)

		if err != nil || berr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		w.Write(res)
	}()

	var request usecase.AddInput
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

	workerRepository := infrastructure.NewWorkerRepositoryImpl(ctx, tx)
	workerAuthRepository := infrastructure.NewWorkerAuthInfoRepositoryImpl(ctx, tx)

	workerUsecase := usecase.NewWorkerUsecase(&workerRepository, &workerAuthRepository)
	response, err = workerUsecase.Add(request)
	if err != nil {
		log.Println(err)
	}
}
