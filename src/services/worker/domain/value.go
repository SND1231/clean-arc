package domain

import (
	"ddd/lib"
	"errors"
	"time"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID() ID {
	ulid := ulid.Make()
	return ID(ulid.String())
}

type Name string
type Email string

func (n Name) valid() error {
	var err error
	size := utf8.RuneCountInString(string(n))
	if size <= 0 || 255 < size {
		err = errors.New("文字数えラー 1~255文字にしてください")
	}
	return err
}

type WokerStatus int32

const (
	WokerStatusInValid WokerStatus = 0
	WokerStatusValid   WokerStatus = 1
	WokerStatusSuspend WokerStatus = 2
)

func (s WokerStatus) valid() error {
	var err error
	if !s.includStatus() {
		err = errors.New("ステータスが間違っています")
	}
	return err
}

func (s WokerStatus) includStatus() bool {
	var statuses []WokerStatus = []WokerStatus{
		WokerStatusValid,
		WokerStatusInValid,
	}
	for _, v := range statuses {
		if s == v {
			return true
		}
	}
	return false
}

type Date time.Time

type Password string

func NewPassword(password string) (Password, error) {
	hash, err := lib.EncryptPassword(password)
	if err != nil {
		return Password(""), err
	}
	return Password(hash), nil
}

type CanChangePasswordCount int32
type FailAuthCount int32
type Session string
