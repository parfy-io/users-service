package storage

import (
	"fmt"
	"github.com/lib/pq"
)

type User struct {
	ID       int64
	FullName string
	Names    []string
	ClientID string
}

func (s Storage) CreateUser(u User) (int64, error) {
	var id int64
	err := s.db.QueryRow(
		"INSERT INTO users (full_name, names, client_id) VALUES($1, $2, $3) RETURNING id;",
		u.FullName, pq.Array(u.Names), u.ClientID).Scan(&id)

	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Constraint == "users_clients_fkey" {
			return 0, ErrClientDoesntExists
		}
		return 0, fmt.Errorf("failed to exec stmt: %w", err)
	}

	return id, nil
}
