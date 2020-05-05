package repository

import (
	"context"
	"database/sql"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/model"
	"github.com/didiyudha/transaction-accross-pkg/internal/platform/postgres"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	StartTx(ctx context.Context) (UserRepository, error)
	Commit() error
	Rollback() error
	WithTransaction(ctx context.Context) (UserRepository, error)
	Context() context.Context
}

type userRepository struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
	postgres.SQL
	postgres.Transactioner
	TrxCtx context.Context
}


func (u *userRepository) Context() context.Context {
	return u.TrxCtx
}

func (u *userRepository) WithTransaction(ctx context.Context) (UserRepository, error) {
	trx, ok := ctx.Value(postgres.TrxKeyContext).(*sql.Tx)
	if !ok {
		return nil, errors.New("unable to get transaction in context")
	}
	pqTx := postgres.Tx{
		Tx:    trx,
	}
	return &userRepository{
		SQL: &pqTx,
		Transactioner: &pqTx,
	}, nil
}


func (uq *userRepository) StartTx(ctx context.Context) (UserRepository, error) {
	trx, err := uq.dbWrite.Begin()
	if err != nil {
		return nil, err
	}
	pqTx := postgres.Tx{
		Tx:    trx,
	}
	trxCtx := context.WithValue(ctx, postgres.TrxKeyContext, trx)
	return &userRepository{
		SQL: &pqTx,
		Transactioner: &pqTx,
		TrxCtx: trxCtx,
	}, nil
}

func (uq *userRepository) GetDB(isWrite bool) postgres.SQL {
	if uq.SQL != nil {
		return uq.SQL
	}
	if isWrite {
		return uq.dbWrite
	}
	return uq.dbRead
}

func (uq *userRepository) Commit() error {
	if uq.Transactioner == nil {
		return nil
	}
	if err := uq.Transactioner.Commit(); err != nil {
		uq.Transactioner.Rollback()
		return err
	}
	return nil
}

func (uq *userRepository) Rollback() error {
	if uq.Transactioner == nil {
		return nil
	}
	if err := uq.Transactioner.Rollback(); err != nil {
		return err
	}
	return nil
}
