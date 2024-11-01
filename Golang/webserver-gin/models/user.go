package models

import (
	"errors"

	"example.com/webserver/db"
	"example.com/webserver/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func GetAllUsers() ([]User , error){
	query := `SELECT * FROM users`
	// kalo cm fetch pake query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		// HARUS SESUAI URUTAN DENGAN STRUCT EVENT
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id,password FROM users WHERE email = ?`
	// query row untuk fetch 1 baris saja
	row:= db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err :=row.Scan(&u.ID,&retrievedPassword)
 
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}