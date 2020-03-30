package internal

import "github.com/parfy-io/users-service/internal/storage"

type Storage interface {
	CreateClient(c storage.Client) error
	CreateUser(u storage.User) (int64, error)
}

type Service struct {
	Storage Storage
}
