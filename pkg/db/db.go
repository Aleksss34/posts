package datab

import (
	"database/sql"
	"fmt"
)

type InfoDatabase struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
	Sslmode  string
}

func NewPostgresDb(info InfoDatabase) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", info.Host, info.Port, info.User, info.Dbname, info.Password, info.Sslmode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
