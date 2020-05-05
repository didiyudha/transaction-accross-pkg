package repository

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
)

func (p *profileRepository) Save(ctx context.Context, profile * model.Profile) error {
	db := p.GetDB(true)
	query := `INSERT INTO profiles (id, user_id, first_name, last_name, address, created_at, updated_at, deleted_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(profile.ID, profile.UserID, profile.FirstName, profile.LastName, profile.Address, profile.CreatedAt,
		profile.UpdatedAt, profile.DeletedAt)
	return err
}
