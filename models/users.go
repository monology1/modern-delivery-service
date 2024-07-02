package models

import (
	"log"
	"modern-delivery-service/db"
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
	_, err = stmt.Exec(user.Email, user.Password)

	if err != nil {
		log.Printf("Error Exec: %v\n", err)
		return err
	}
	return err
}
