package app

import (
	"log"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

func (s *server) payloadSerializer(r *http.Request, object interface{}) error {
	if err := jsonapi.UnmarshalPayload(r.Body, object); err != nil {
		return errors.Wrap(err, "Method: payloadSerializer")
	}

	return nil
}

func (s *server) successSerializer(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", jsonapi.MediaType)

	if err := jsonapi.MarshalPayload(w, data); err != nil {
		s.errorSerializer(w, errors.Wrap(err, "Method: jsonapi.MarshalPayload"))
	}
}

func (s *server) errorSerializer(w http.ResponseWriter, err error) {

	switch e := errors.Cause(err).(type) {
	case ClientError:
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
