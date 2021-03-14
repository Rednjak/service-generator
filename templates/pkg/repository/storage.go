package repository

// Storage represents the resource and the methods that can be executes
// It can be renamed to represent the resource that the service will work with
// e.g UserStorage, CommentStorage etc.
type Storage interface {
	Get(ID uint64) (*Resource, error)
	GetAll() ([]*Resource, error)
	Create(obj *Resource) (*Resource, error)
	Update(obj *Resource) error
	Delete(ID uint64) error
}

func (s *storage) Get(ID uint64) (*Resource, error) {
	return nil, nil
}

func (s *storage) GetAll() ([]*Resource, error) {
	return nil, nil
}

func (s *storage) Create(obj *Resource) (*Resource, error) {
	return nil, nil
}

func (s *storage) Update(obj *Resource) error {
	return nil
}

func (s *storage) Delete(ID uint64) error {
	return nil
}
