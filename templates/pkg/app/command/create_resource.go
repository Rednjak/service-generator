package command

import (
	"context"
	domain "module_placeholder/pkg/domain/resource"
)

type CreateResourceHandler struct {
	repo domain.Repository
}

func NewCreateResourceHandler(repo domain.Repository) CreateResourceHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return CreateResourceHandler{
		repo: repo,
	}
}

func (h CreateResourceHandler) Handle(ctx context.Context, resource *domain.Resource) (uint64, error) {
	resourceID, err := h.repo.CreateResource(resource)
	if err != nil {
		return 0, err
	}
	return resourceID, nil
}
