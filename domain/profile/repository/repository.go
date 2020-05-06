package repository

import (
	"context"
	"database/sql"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
	"github.com/didiyudha/transaction-accross-pkg/internal/platform/postgres"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ProfileRepository interface {
	Save(ctx context.Context, profile * model.Profile) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Profile, error)
	StartTx(ctx context.Context) (ProfileRepository, error)
	Commit() error
	Rollback() error
	WithTransaction(ctx context.Context) (ProfileRepository, error)
	Context() context.Context
}

type profileRepository struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
	postgres.SQL
	postgres.Transactioner
	TrxCtx context.Context
}

func NewProfileRepository(dbWrite, dbRead *sql.DB) ProfileRepository {
	return &profileRepository{
		dbRead:        dbRead,
		dbWrite:       dbWrite,
	}
}

func (uq *profileRepository) StartTx(ctx context.Context) (ProfileRepository, error) {
	trx, err := uq.dbWrite.Begin()
	if err != nil {
		return nil, err
	}
	pqTx := postgres.Tx{
		Tx:    trx,
	}
	trxCtx := context.WithValue(ctx, postgres.TrxKeyContext, trx)
	return &profileRepository{
		SQL: &pqTx,
		Transactioner: &pqTx,
		TrxCtx: trxCtx,
	}, nil
}

func (uq *profileRepository) Context() context.Context {
	return uq.TrxCtx
}

func (uq *profileRepository) WithTransaction(ctx context.Context) (ProfileRepository, error) {
	trx, ok := ctx.Value(postgres.TrxKeyContext).(*sql.Tx)
	if !ok {
		return nil, errors.New("unable to get transaction in context")
	}
	pqTx := postgres.Tx{
		Tx:    trx,
	}
	return &profileRepository{
		SQL: &pqTx,
		Transactioner: &pqTx,
	}, nil
}


func (uq *profileRepository) GetDB(isWrite bool) postgres.SQL {
	if uq.SQL != nil {
		return uq.SQL
	}
	if isWrite {
		return uq.dbWrite
	}
	return uq.dbRead
}

func (uq *profileRepository) Commit() error {
	if uq.Transactioner == nil {
		return nil
	}
	if err := uq.Transactioner.Commit(); err != nil {
		uq.Transactioner.Rollback()
		return err
	}
	return nil
}

func (uq *profileRepository) Rollback() error {
	if uq.Transactioner == nil {
		return nil
	}
	if err := uq.Transactioner.Rollback(); err != nil {
		return err
	}
	return nil
}
