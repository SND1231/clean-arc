package domain

import (
	"time"
)

type WorkerEntity struct {
	ID        ID
	Email     Email
	Name      Name
	Status    WokerStatus
	CreatedAt Date
	UpdatedAt Date
}

func NewWorker(id string, name string, email string, status int32, createdDate time.Time, updatedDate time.Time) (*WorkerEntity, error) {
	var worker *WorkerEntity = &WorkerEntity{
		ID:        ID(id),
		Name:      Name(name),
		Email:     Email(email),
		Status:    WokerStatus(status),
		CreatedAt: Date(createdDate),
		UpdatedAt: Date(updatedDate),
	}
	err := worker.valid()
	return worker, err
}

func (w *WorkerEntity) valid() error {
	var err error
	err = w.Name.valid()
	if err != nil {
		return err
	}
	err = w.Status.valid()
	return err
}

type WorkerAuthInfoEntity struct {
	ID                     ID
	Password               Password
	CanChangePasswordCount CanChangePasswordCount
	FailAuthCount          FailAuthCount
	CreatedAt              Date
	UpdatedAt              Date
}

func NewWorkerAuthInfo(id string, password string, changePasswordCount int32, createdDate time.Time, updatedDate time.Time) (*WorkerAuthInfoEntity, error) {
	encryptPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}
	var workerAuthInfo *WorkerAuthInfoEntity = &WorkerAuthInfoEntity{
		ID:                     ID(id),
		Password:               encryptPassword,
		CanChangePasswordCount: CanChangePasswordCount(changePasswordCount),
		CreatedAt:              Date(createdDate),
		UpdatedAt:              Date(updatedDate),
	}
	return workerAuthInfo, nil
}
