package serializer

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

func SerializePayload(r *http.Request, object interface{}) error {
	if err := jsonapi.UnmarshalPayload(r.Body, object); err != nil {
		return errors.Wrap(err, "Method: PayloadSerializer")
	}

	return nil
}
