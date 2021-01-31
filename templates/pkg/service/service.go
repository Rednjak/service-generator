package service

// Service interface is used to define the available methods
// It's meant to act as a business layer
// It can be renamed to represent a specific set of service methods in case we would like to add multiple service
// abstraction methods. e.g. UserService, CommentService etc.
type Service interface {
	Get()
	GetAll()
	Create()
	Update()
	Delete()
}

func (s *service) Get() {
	s.Get()
}

func (s *service) GetAll() {
	s.GetAll()
}

func (s *service) Create() {
	s.Create()
}

func (s *service) Update() {
	s.Update()
}

func (s *service) Delete() {
	s.Delete()
}
