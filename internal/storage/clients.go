package storage

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
)

var ErrClientAlreadyExists = errors.New("client with the given name already exists")

type Client struct {
	Name string
}

func (s Storage) CreateClient(c Client) error {
	_, err := s.db.Exec("INSERT INTO clients (name) VALUES($1);", c.Name)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Constraint == "name_unique" {
			return ErrClientAlreadyExists
		}
		return fmt.Errorf("failed to exec stmt: %w", err)
	}

	return nil
}
