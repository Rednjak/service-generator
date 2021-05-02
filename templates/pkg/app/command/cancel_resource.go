package command

import (
	"context"
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
