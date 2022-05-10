package beer

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// UseCase é uma interface de funções que serão usadas no docorrer do projeto.
type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

// NewService função retorna um ponteiro em memória para uma estrutura
func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Beer, error) {
	var result []*Beer

	rows, err := s.DB.Query("select id, name, type, style from beer")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b Beer
		err = rows.Scan(&b.ID, &b.Name, &b.Type, &b.Style)
		if err != nil {
			return nil, err
		}

		result = append(result, &b)
	}

	return result, nil

}

func (s *Service) Get(ID int64) (*Beer, error) {
	var b Beer

	// Utilizamos o Prepare para colocar uma variável na consulta.
	// Observe o "?" abaixo. Nele, será inserido o ID posteriormente.
	stmt, err := s.DB.Prepare("select id, name, type, style from beer where id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&b.ID, &b.Name, &b.Type, &b.Style)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
func (s *Service) Store(b *Beer) error {
	//iniciamos uma transação
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into beer(id, name, type, style) values (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = stmt.Exec(b.ID, b.Name, b.Type, b.Style)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) Update(b *Beer) error {
	if b.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("update beer set name=?, type=?, style=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Name, b.Type, b.Style, b.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from beer where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
