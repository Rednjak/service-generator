package service

import "github.com/repo/package/pkg/repository"

// Holds all the dependencies that are being used by the servie layer
type service struct {
	repo repository.Storage
}

// SetupService returns a service struct that can be used to call service methods
func SetupService(repo repository.Storage) *service {
	return &service{
		repo: repo,
	}
}
