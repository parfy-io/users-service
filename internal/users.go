package internal

import (
	"errors"
	"fmt"
	"github.com/parfy-io/users-service/internal/storage"
	"github.com/parfy-io/users-service/internal/web"
)

func (s Service) CreateUser(clientID string, user web.User) (int64, error) {
	id, err := s.Storage.CreateUser(storage.User{
		FullName: user.FullName,
		Names:    user.Names,
		EMail:    user.EMail,
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
