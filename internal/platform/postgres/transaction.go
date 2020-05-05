package postgres

import (
	"database/sql"
	"github.com/pkg/errors"
	"sync"
)

type CtxTrxKey string

var TrxKeyContext CtxTrxKey = "contextTransactionKey"

type Transactioner interface {
	Commit() error
	Rollback() error
}

type SQL interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Tx struct {
	Tx *sql.Tx
	sync.Mutex
}

func (t *Tx) checkTx() error {
	if t == nil {
		return errors.New("transaction is not created yet")
	}
	if t.Tx == nil {
		return errors.New("transaction is not started yet")
	}
	return nil
}

func (t *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	if err := t.checkTx(); err != nil {
		return nil, err
	}
	return t.Tx.Exec(query, args)
}

func (t *Tx) Prepare(query string) (*sql.Stmt, error) {
	if err := t.checkTx(); err != nil {
		return nil, err
	}
	return t.Tx.Prepare(query)
}

func (t *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if err := t.checkTx(); err != nil {
		return nil, err
	}
	return t.Tx.Query(query, args)
}

func (t *Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.Tx.QueryRow(query, args)
}

func (t *Tx) Commit() error {
	if err := t.checkTx(); err != nil {
		return err
	}
	t.Lock()
	defer t.Unlock()
	err := t.Tx.Commit()
	if err != nil {
		t.Tx.Rollback()
	}
	t.Tx = nil
	return err
}

func (t *Tx) Rollback() error {
	if err := t.checkTx(); err != nil {
		return err
	}
	t.Lock()
	defer t.Unlock()
	err := t.Tx.Rollback()
	t.Tx = nil
	return err
}


