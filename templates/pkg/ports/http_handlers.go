package ports

import (
	"module_placeholder/pkg/app"
	"module_placeholder/pkg/app/query"
	domain "module_placeholder/pkg/domain/resource"
	"module_placeholder/pkg/shared/parser"
	"module_placeholder/pkg/shared/serializer"
	"net/http"

	"github.com/go-chi/chi"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) CancelResource(w http.ResponseWriter, r *http.Request) {
	h.app.Commands.CancelResource.Handle(r.Context())
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) CreateResource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	payloadResource := &ResourcePayload{}
	err := serializer.SerializePayload(r, payloadResource)
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	domainResource, err := newDomainResourceFromPayloadResource(payloadResource)
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	// CQRS principles don't allow returning a resource, but in our case we return the ID which is then used by the query
	resourceID, err := h.app.Commands.CreateResource.Handle(ctx, domainResource)
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	queryResource, err := h.app.Queries.GetResource.Handle(ctx, resourceID)
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	serializer.SerializeSuccess(w, newQueryResourceToPayloadResource(queryResource))
}

func (h HttpServer) GetResource(w http.ResponseWriter, r *http.Request) {
	resourceID, err := parser.StringToUint64(chi.URLParam(r, "resourceID"))
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	resource, err := h.app.Queries.GetResource.Handle(r.Context(), resourceID)
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	serializer.SerializeSuccess(w, resource)
}

func (h HttpServer) GetAllResources(w http.ResponseWriter, r *http.Request) {
	resources, err := h.app.Queries.AllResources.Handle(r.Context())
	if err != nil {
		serializer.SerializeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	serializer.SerializeSuccess(w, resources)
}

func newDomainResourceFromPayloadResource(resource *ResourcePayload) (*domain.Resource, error) {
	return domain.NewResource(resource.FirstAttribute, resource.SecondAttribute)
}

func newQueryResourceToPayloadResource(resource *query.Resource) *ResourcePayload {
	return &ResourcePayload{
		ID:              resource.ID,
		FirstAttribute:  resource.FirstAttribute,
		SecondAttribute: resource.SecondAttribute,
		IsValid:         resource.IsValid,
	}
}
