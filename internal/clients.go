package internal

import (
	"errors"
	"fmt"
	"github.com/parfy-io/users-service/internal/storage"
)

var ErrClientAlreadyExists = errors.New("client with the given id already exists")
var ErrClientDoesntExists = errors.New("client with the given id does not exists")

func (s Service) CreateClient(id string) error {
	err := s.Storage.CreateClient(storage.Client{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, storage.ErrClientAlreadyExists) {
			return ErrClientAlreadyExists
		}
		return fmt.Errorf("failed to create client: %w", err)
	}

	return nil
}
