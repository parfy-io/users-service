package internal

import (
	"errors"
	"fmt"
	"github.com/parfy-io/users-service/internal/storage"
)

var ErrClientAlreadyExists = errors.New("client with the given name already exists")

func (s Service) CreateClient(name string) error {
	err := s.Storage.CreateClient(storage.Client{
		Name: name,
	})
	if err != nil {
		if errors.Is(err, storage.ErrClientAlreadyExists) {
			return ErrClientAlreadyExists
		}
		return fmt.Errorf("failed to create client: %w", err)
	}

	return nil
}
