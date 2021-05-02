package errors

import (
	"github.com/google/jsonapi"
)

// ClientError is used for errors where the client will display some custom message to the user
// e.g payload has an error, social auth token exchange has failed etc.
// We wrap the JSONAPI error and add a StackTraceError which is used to log the error details
type ClientError struct {
	StatusCode      int
	StackTraceError error
	JSONAPIError    []*jsonapi.ErrorObject
}

func (err ClientError) Error() string {
	return "There was a ClientError"
}
