package command

import (
	"context"
	domain "module_placeholder/pkg/domain/resource"
)

type CancelResourceHandler struct {
	repo domain.Repository
}

func NewCancelResourceHandler(repo domain.Repository) CancelResourceHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return CancelResourceHandler{
		repo: repo,
	}
}

func (h CancelResourceHandler) Handle(ctx context.Context) error {
	return nil
}
