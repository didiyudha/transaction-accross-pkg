package repository

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/model"
	"github.com/google/uuid"
)

func (u *userRepository) Save(ctx context.Context, user *model.User) error {
	db := u.GetDB(true)
	query := `INSERT INTO users (id, username, email, created_at, updated_at, deleted_at) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	return err
}

func (u *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	db := u.GetDB(false)

	var usr model.User

	query := `SELECT
				id,
				username,
				email,
				created_at,
				updated_at,
				deleted_at
			FROM users
			WHERE id = $1 AND deleted_at ISNULL`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&usr.ID,
		&usr.Username,
		&usr.Email,
		&usr.CreatedAt,
		&usr.UpdatedAt,
		&usr.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}