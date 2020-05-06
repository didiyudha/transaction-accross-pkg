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

func (p *profileRepository) StartTx(ctx context.Context) (ProfileRepository, error) {
	trx, err := p.dbWrite.Begin()
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

func (p *profileRepository) Context() context.Context {
	return p.TrxCtx
}

func (p *profileRepository) WithTransaction(ctx context.Context) (ProfileRepository, error) {
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


func (p *profileRepository) GetDB(isWrite bool) postgres.SQL {
	if p.SQL != nil {
		return p.SQL
	}
	if isWrite {
		return p.dbWrite
	}
	return p.dbRead
}

func (p *profileRepository) Commit() error {
	if p.Transactioner == nil {
		return nil
	}
	if err := p.Transactioner.Commit(); err != nil {
		p.Transactioner.Rollback()
		return err
	}
	return nil
}

func (p *profileRepository) Rollback() error {
	if p.Transactioner == nil {
		return nil
	}
	if err := p.Transactioner.Rollback(); err != nil {
		return err
	}
	return nil
}
