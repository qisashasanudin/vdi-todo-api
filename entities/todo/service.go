package todo

type Service interface {
	FindAll() ([]Todo, error)
	FindById(ID int) (Todo, error)
	Create(todoRequest CreateTodoRequest) (Todo, error)
	Update(ID int, todoRequest UpdateTodoRequest) (Todo, error)
	Delete(ID int) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Todo, error) {
	return s.repository.FindAll()

}

func (s *service) FindById(ID int) (Todo, error) {
	return s.repository.FindById(ID)
}

func (s *service) Create(b CreateTodoRequest) (Todo, error) {

	todo := Todo{
		Title:       *b.Title,
		Description: *b.Description,
		Completed:   false,
	}

	return s.repository.Create(todo)
}

func (s *service) Update(ID int, b UpdateTodoRequest) (Todo, error) {
	todo, err := s.repository.FindById(ID)
	if err != nil {
		return todo, err
	}

	if b.Title != nil {
		todo.Title = *b.Title
	}

	if b.Description != nil {
		todo.Description = *b.Description
	}

	if b.Completed != nil {
		todo.Completed = *b.Completed
	}

	return s.repository.Update(todo)
}

func (s *service) Delete(ID int) (Todo, error) {
	todo, err := s.repository.FindById(ID)
	if err != nil {
		return todo, err
	}

	return s.repository.Delete(todo)
}
