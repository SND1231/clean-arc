package usecase

import (
	"net/http"

	"github.com/mholt/binding"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginInput) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&req.Email:    "email",
		&req.Password: "password",
	}
}

type AddInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (req AddInput) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&req.Email:    "email",
		&req.Password: "password",
		&req.Name:     "name",
	}
}
