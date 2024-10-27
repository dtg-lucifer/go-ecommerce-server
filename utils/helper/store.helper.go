package helper

import (
	"database/sql"

	"github.com/dtg-lucifer/go-backend/typedef"
)

func ScanRowIntoUser(row *sql.Rows) (*typedef.User, error) {
	u := new(typedef.User)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
