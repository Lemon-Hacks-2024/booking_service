package pkg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(host, port, user, pass, dbname, sslmode string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, pass, sslmode))
	if err != nil {
		return nil, err
	}
	//
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// Создание пула соединений
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	//
	fmt.Println("The connection to postgresql was successful")
	return db, nil
}
