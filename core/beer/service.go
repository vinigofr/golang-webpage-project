package beer

// UseCase é uma interface de funções que serão usadas no docorrer do projeto.
type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(b *Beer) error
}

type Service struct{}

// NewService função retorna um ponteiro em memória para uma estrutura
func NewService() *Service {
	return &Service{}
}

//vamos implementar as funções na próxima etapa
func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int64) (*Beer, error) {
	return nil, nil
}
func (s *Service) Store(b *Beer) error {
	return nil
}
func (s *Service) Update(b *Beer) error {
	return nil
}
func (s *Service) Remove(ID int64) error {
	return nil
}
