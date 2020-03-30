package storage

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
)

var ErrChannelAlreadyExists = errors.New("channel with the given id already exists")

type Channel struct {
	Name string
}

func (s Storage) CreateChannel(c Channel) error {
	_, err := s.db.Exec("INSERT INTO channels (type) VALUES($1);", c.Name)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Constraint == "channels_type_unique" {
			return ErrChannelAlreadyExists
		}
		return fmt.Errorf("failed to exec stmt: %w", err)
	}

	return nil
}
