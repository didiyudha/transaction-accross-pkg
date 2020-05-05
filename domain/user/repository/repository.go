package repository

import (
	"context"
	"database/sql"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/model"
	"github.com/didiyudha/transaction-accross-pkg/internal/platform/postgres"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
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


func (u *userRepository) GetDB(isWrite bool) postgres.SQL {
	if u.SQL != nil {
		return u.SQL
	}
	if isWrite {
		return u.dbWrite
	}
	return u.dbRead
}

func (u *userRepository) Commit() error {
	if u.Transactioner == nil {
		return nil
	}
	if err := u.Transactioner.Commit(); err != nil {
		u.Transactioner.Rollback()
		return err
	}
	return nil
}

func (u *userRepository) Rollback() error {
	if u.Transactioner == nil {
		return nil
	}
	if err := u.Transactioner.Rollback(); err != nil {
		return err
	}
	return nil
}