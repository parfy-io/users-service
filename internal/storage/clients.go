package storage

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
)

var ErrClientAlreadyExists = errors.New("client with the given id already exists")
var ErrClientDoesntExists = errors.New("client with the given id does not exists")

type Client struct {
	ID string
}

func (s Storage) CreateClient(c Client) error {
	_, err := s.db.Exec("INSERT INTO clients (id) VALUES($1);", c.ID)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Constraint == "clients_id_unique" {
			return ErrClientAlreadyExists
		}
		return fmt.Errorf("failed to exec stmt: %w", err)
	}

	return nil
}
