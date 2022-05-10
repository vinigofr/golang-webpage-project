package beer_test

import (
	"database/sql"
	"github.com/vinigofr/golang-webpage-project/core/beer"
	"testing"
)

func TestService_Store(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Skol",
		Style: beer.StylePilsner,
		Type:  beer.TypeLager,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")

	if err != nil {
		t.Fatalf("Erro ao conectar-se com o banco de dados: %s", err.Error())
	}

	defer db.Close()
	clearDB(db)
	if err != nil {
		t.Fatalf("Erro ao limpar o banco de dados: %s", err.Error())
	}

	service := beer.NewService(db)

	err = service.Store(b)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("drop table beer")
	tx.Commit()
	return err
}
