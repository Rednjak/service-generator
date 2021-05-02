package serializer

import (
	"log"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

func SerializeSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", jsonapi.MediaType)

	if err := jsonapi.MarshalPayload(w, data); err != nil {
		SerializeError(w, errors.Wrap(err, "Method: jsonapi.MarshalPayload"))
	}
}

func SerializeError(w http.ResponseWriter, err error) {
	switch e := errors.Cause(err).(type) {
	case custom_errors.ClientError:
		// In case the status code is not provided we default to 400
		if e.StatusCode == 0 {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(e.StatusCode)
		}

		jsonapi.MarshalErrors(w, e.JSONAPIError)
	default:
		// Any other error will contain a stacktrace because we will use errors.Wrap method from
		// https://github.com/pkg/errors
		log.Printf("Error message: %s, Stacktrace: %+v \n", e.Error(), e)

		w.WriteHeader(500)
		jsonapi.MarshalErrors(w,
			[]*jsonapi.ErrorObject{
				&jsonapi.ErrorObject{
					Detail: "There was a server error. Please try again.",
				}})
	}
}
