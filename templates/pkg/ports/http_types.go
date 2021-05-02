package ports

// ResourcePayload is a separate model used to marshal request payloads. We don't want to add unnecessary tags to
// the domain model
type ResourcePayload struct {
	ID              uint64 `jsonapi:"primary,resources"`
	FirstAttribute  string `jsonapi:"attr,first_attribute"`
	SecondAttribute string `jsonapi:"attr,second_attribute"`
	IsValid         bool   `jsonapi:"attr,is_valid"`
}
