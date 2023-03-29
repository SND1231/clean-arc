package domain

type WorkerRepository interface {
	Add(*WorkerEntity) error
	FineByID(ID) (*WorkerEntity, error)
	IsExists(Email) (bool, error)
}

type WorkerAuthInfoRepository interface {
	Add(*WorkerAuthInfoEntity) error
	Login(*WorkerAuthInfoEntity) (bool, error)
	FineByID(ID) (*WorkerAuthInfoEntity, error)
	Update(*WorkerAuthInfoEntity) error
}
