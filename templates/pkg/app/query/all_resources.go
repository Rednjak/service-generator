package query

import "context"

type AllResourcesHandler struct {
	readModel AllResourcesReadModel
}

func NewAllResourceHandler(readModel AllResourcesReadModel) AllResourcesHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllResourcesHandler{
		readModel: readModel,
	}
}

type AllResourcesReadModel interface {
	AllResources() ([]*Resource, error)
}

func (h AllResourcesHandler) Handle(ctx context.Context) ([]*Resource, error) {
	return h.readModel.AllResources()
}
