package repository

// Storage represents the resource and the methods that can be executes
// It can be renamed to represent the resource that the service will work with
// e.g UserStorage, CommentStorage etc.
type Storage interface {
	Get()
	GetAll()
	Create()
	Update()
	Delete()
}

func (s *storage) Get() {
}

func (s *storage) GetAll() {
}

func (s *storage) Create() {
}

func (s *storage) Update() {
}

func (s *storage) Delete() {
}
