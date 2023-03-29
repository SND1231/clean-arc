package domain

type WorkerDomainService struct {
	workerRepository WorkerRepository
	authRepository   WorkerAuthInfoRepository
}

func NewWorkerDomainService(workerRepository WorkerRepository, authRepository WorkerAuthInfoRepository) *WorkerDomainService {
	return &WorkerDomainService{
		workerRepository: workerRepository,
		authRepository:   authRepository,
	}
}

func (s *WorkerDomainService) IsExists(email Email) (bool, error) {
	return s.workerRepository.IsExists(email)
}
