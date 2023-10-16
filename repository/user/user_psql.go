package userRepository

import (
	"database/sql"
	"errors"

	"yspace.com.br/models"
	"yspace.com.br/utils"
)

type UserRepository struct{}

func (u UserRepository) Signup(db *sql.DB, user models.User) (models.User, error) {

	if utils.IsEmailValid(user.Email) == false {
		err := errors.New("Email invalido")
		return user, err
	}

	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}
