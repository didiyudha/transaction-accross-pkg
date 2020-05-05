package repository

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/user/model"
)

func (u *userRepository) Save(ctx context.Context, user *model.User) error {
	db := u.GetDB(true)
	query := `INSERT INTO users (id, username, email, created_at, updated_at, deleted_at) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	return err
}
