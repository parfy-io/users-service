package internal

import (
	"errors"
	"fmt"
	"github.com/parfy-io/users-service/internal/storage"
)

type User struct {
	ID       int64
	FullName string
	Names    []string
}

func (s Service) CreateUser(clientID string, user User) (int64, error) {
	id, err := s.Storage.CreateUser(storage.User{
		FullName: user.FullName,
		Names:    user.Names,
		ClientID: clientID,
	})
	if err != nil {
		if errors.Is(err, storage.ErrClientDoesntExists) {
			return 0, ErrClientDoesntExists
		}
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}
