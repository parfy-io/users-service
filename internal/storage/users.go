package storage

import (
	"fmt"
	"github.com/lib/pq"
)

type User struct {
	ID       int64
	FullName string
	Names    []string
	EMail    string
	ClientID string
}

func (s Storage) CreateUser(u User) (int64, error) {
	resp, err := s.db.Exec("INSERT INTO users (full_name, names, email, client_id) VALUES($1, $2, $3, $4);", u.FullName, u.Names, u.EMail, u.ClientID)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Constraint == "users_clients_fkey" {
			return 0, ErrClientDoesntExists
		}
		return 0, fmt.Errorf("failed to exec stmt: %w", err)
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last-inserted-id: %w", err)
	}

	return id, nil
}
