package repository

import (
	todo "ToDo"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgress struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgress {
	return &AuthPostgress{db: db}
}

func (r *AuthPostgress) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, username, passvord_hash) VALUES ($1, $2, $3) RETURNING id`, userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgress) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	fmt.Println("Пароль:  " + password)
	fmt.Println("Юзернейм:   " + username)
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND passvord_hash=$2`, userTable)
	fmt.Println("запрос:   " + query)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
