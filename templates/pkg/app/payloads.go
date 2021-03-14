package app

import "github.com/repo/module/pkg/repository"

// ResourcePayload is esentially a DTO because we don't want to use the database models as the source for our communication with the client.
// We could argue that this is premature optimization since the microservice shouldn't be too big and the database model probably won't get polluted with to many annotations
type ResourcePayload struct {
	ID              uint32 `jsonapi:"primary,resources"`
	FirstAttribute  string `jsonapi:"attr,first_attribute"`
	SecondAttribute string `jsonapi:"attr,second_atttribute"`
}

func NewResourcePayload(obj repository.Resource) *ResourcePayload {
	return &ResourcePayload{
		ID:              obj.ID,
		FirstAttribute:  obj.FirstAttribute,
		SecondAttribute: obj.SecondAttribute,
	}
}

func (up *ResourcePayload) toModel() *repository.Resource {
	return &repository.Resource{
		ID:              up.ID,
		FirstAttribute:  up.FirstAttribute,
		SecondAttribute: up.SecondAttribute,
	}
}
