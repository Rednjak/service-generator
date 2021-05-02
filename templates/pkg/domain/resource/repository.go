package domain

// Repository represents the resource and the methods that can be executed in this domain
// The methods can be renamed to represent the proper actions e.g GetUser, Update User
type Repository interface {
	GetResource(ID uint64) (*query.Resource, error)
	AllResources() ([]*query.Resource, error)
	CreateResource(obj *Resource) (uint64, error)
	UpdateResource(obj *Resource) error
	DeleteResource(ID uint64) error
}
