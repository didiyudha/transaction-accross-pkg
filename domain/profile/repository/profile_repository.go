package repository

import (
	"context"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/model"
	"github.com/google/uuid"
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

func (p *profileRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	db := p.GetDB(false)

	var profile model.Profile

	query := `SELECT
				id,
				user_id,
				first_name,
				last_name,
				address,
				created_at,
				updated_at,
				deleted_at
			FROM profiles
			WHERE id = $1 AND deleted_at ISNULL`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Address,
		&profile.CreatedAt,
		&profile.UpdatedAt,
		&profile.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
