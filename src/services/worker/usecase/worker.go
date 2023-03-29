package usecase

import (
	"ddd/services/worker/domain"
	"fmt"
	"time"
)

type WokerUsecase struct {
	workerRepository domain.WorkerRepository
	authRepository   domain.WorkerAuthInfoRepository
}

// NewWokerUsecaseは後で作成
func NewWorkerUsecase(workerRepository domain.WorkerRepository, authRepository domain.WorkerAuthInfoRepository) *WokerUsecase {
	return &WokerUsecase{
		workerRepository: workerRepository,
		authRepository:   authRepository,
	}
}

func (u WokerUsecase) Add(input AddInput) (AddOutput, error) {
	domainService := domain.NewWorkerDomainService(u.workerRepository, u.authRepository)
	isExists, err := domainService.IsExists(domain.Email(input.Email))
	if isExists {
		return ZeroAddOutput(), fmt.Errorf("既に存在します")
	}
	if err != nil {
		return ZeroAddOutput(), err
	}

	id := domain.NewID()
	createDate := time.Now()
	updateDate := createDate
	workerEntity, err := domain.NewWorker(
		string(id),
		input.Name,
		input.Email,
		int32(domain.WokerStatusValid),
		createDate,
		updateDate,
	)
	if err != nil {
		return ZeroAddOutput(), err
	}

	err = u.workerRepository.Add(workerEntity)
	if err != nil {
		return ZeroAddOutput(), err
	}

	authEntity, err := domain.NewWorkerAuthInfo(
		string(id),
		input.Password,
		int32(0),
		createDate,
		updateDate,
	)
	if err != nil {
		return ZeroAddOutput(), err
	}
	err = u.authRepository.Add(authEntity)
	if err != nil {
		return ZeroAddOutput(), err
	}

	output := AddOutput{
		ID: string(id),
	}
	return output, nil
}
