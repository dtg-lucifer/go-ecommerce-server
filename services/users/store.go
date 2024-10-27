package users

import (
	"database/sql"

	"github.com/dtg-lucifer/go-backend/typedef"
	"github.com/dtg-lucifer/go-backend/utils/helper"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetUserByEmail(email string) (*typedef.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(typedef.User)
	for rows.Next() {
		u, err = helper.ScanRowIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, nil
	}
	return u, nil
}
