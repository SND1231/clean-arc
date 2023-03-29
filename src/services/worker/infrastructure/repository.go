package infrastructure

import (
	"context"
	"database/sql"
	"ddd/infrastructure/models"
	"ddd/services/worker/domain"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type WorkerRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func NewWorkerRepositoryImpl(ctx context.Context, tx *sql.Tx) WorkerRepositoryImpl {
	return WorkerRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}

func (r *WorkerRepositoryImpl) Add(entity *domain.WorkerEntity) error {
	worker := models.Worker{
		WorkerID:  string(entity.ID),
		Email:     string(entity.Email),
		Name:      string(entity.Name),
		Status:    int(entity.Status),
		CreatedAt: time.Time(entity.CreatedAt),
		UpdatedAt: time.Time(entity.UpdatedAt),
	}
	err := worker.Insert(r.ctx, r.tx, boil.Infer())
	return err
}

func (r *WorkerRepositoryImpl) IsExists(email domain.Email) (bool, error) {
	return models.Workers(qm.Where("email=?", string(email))).Exists(r.ctx, r.tx)
}

func (r *WorkerRepositoryImpl) FineByID(id domain.ID) (*domain.WorkerEntity, error) {
	return nil, nil
}

type WorkerAuthInfoRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func NewWorkerAuthInfoRepositoryImpl(ctx context.Context, tx *sql.Tx) WorkerAuthInfoRepositoryImpl {
	return WorkerAuthInfoRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}

func (r *WorkerAuthInfoRepositoryImpl) Add(entity *domain.WorkerAuthInfoEntity) error {
	auth := models.Auth{
		WorkerID:            string(entity.ID),
		Password:            string(entity.Password),
		Session:             "",
		ChangePasswordCount: int(entity.CanChangePasswordCount),
		FailAuthCount:       int(entity.FailAuthCount),
	}
	err := auth.Insert(r.ctx, r.tx, boil.Infer())
	return err
}

func (r *WorkerAuthInfoRepositoryImpl) Login(entity *domain.WorkerAuthInfoEntity) (bool, error) {
	return true, nil
}

func (r *WorkerAuthInfoRepositoryImpl) FineByID(id domain.ID) (*domain.WorkerAuthInfoEntity, error) {
	return nil, nil
}

func (r *WorkerAuthInfoRepositoryImpl) Update(entity *domain.WorkerAuthInfoEntity) error {
	return nil
}
