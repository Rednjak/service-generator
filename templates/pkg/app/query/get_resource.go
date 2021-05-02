package query

import "context"

type GetResourceHandler struct {
	readModel GetResourcesReadModel
}

func NewGetResourceHandler(readModel GetResourcesReadModel) GetResourceHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetResourceHandler{
		readModel: readModel,
	}
}

type GetResourcesReadModel interface {
	GetResource(ID uint64) (*Resource, error)
}

func (h GetResourceHandler) Handle(ctx context.Context, ID uint64) (*Resource, error) {
	queryResource, err := h.readModel.GetResource(ID)
	if err != nil {
		return nil, err
	}

	queryResource.IsValid = true
	return queryResource, nil
}
