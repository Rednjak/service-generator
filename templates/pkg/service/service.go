package service

import (
	"context"

	"github.com/repo/package/pkg/repository"
)

// Service interface is used to define the available methods
// It's meant to act as a business layer
// It can be renamed to represent a specific set of service methods in case we would like to add multiple service
// abstraction methods. e.g. resourceService, CommentService etc.
type Service interface {
	Get(ctx context.Context, ID uint64) (*repository.Resource, error)
	GetAll(ctx context.Context) ([]*repository.Resource, error)
	Create(ctx context.Context, obj *repository.Resource) (*repository.Resource, error)
	Update(ctx context.Context, obj *repository.Resource) error
	Delete(ctx context.Context, ID uint64) error
}

func (s *service) Get(ctx context.Context, ID uint64) (*repository.Resource, error) {
	obj, err := s.repo.Get(ID)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *service) GetAll(ctx context.Context) ([]*repository.Resource, error) {
	objects, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return objects, nil
}

func (s *service) Create(ctx context.Context, resource *repository.Resource) (*repository.Resource, error) {
	obj, err := s.repo.Create(resource)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *service) Update(ctx context.Context, obj *repository.Resource) error {
	err := s.repo.Update(obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, ID uint64) error {
	err := s.repo.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
