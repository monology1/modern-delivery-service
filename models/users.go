package models

import (
	"errors"
	"log"
	"modern-delivery-service/db"
	"modern-delivery-service/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error Prepare: %v\n", err)
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		log.Printf("Error Exec: %v\n", err)
		return err
	}
	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		log.Printf("Error Scan: %v\n", err)
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid password")
	}

	return nil
}
