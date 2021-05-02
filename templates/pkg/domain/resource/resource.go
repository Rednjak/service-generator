package domain

import (
	"errors"

	"github.com/google/jsonapi"
)

// Resource is used internally to represent the domain model and should be renamed to
// a proper domain name after creating the service
type Resource struct {
	firstAttribute  string
	secondAttribute string
}

func NewResource(firstAttribute, secondAttribute string) (*Resource, error) {
	if firstAttribute == "" {
		return nil, &custom_errors.ClientError{
			JSONAPIError: []*jsonapi.ErrorObject{
				&jsonapi.ErrorObject{
					Detail: "Empty firstAttribute",
				}},
		}
	}

	if secondAttribute == "" {
		return nil, &custom_errors.ClientError{
			JSONAPIError: []*jsonapi.ErrorObject{
				&jsonapi.ErrorObject{
					Detail: "Empty secondAttribute",
				}},
		}
	}

	return &Resource{
		firstAttribute:  firstAttribute,
		secondAttribute: secondAttribute,
	}, nil
}

func (r *Resource) SomeBusinessRule() error {
	return nil
}

func (r *Resource) SomeBusinessRule2() error {
	if r.firstAttribute != r.secondAttribute {
		return errors.New("invalid values")
	}
	return nil
}
