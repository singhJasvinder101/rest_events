package models

import (
	"errors"
	"fmt"

	"github.com/singhJasvinder101/db"
	"github.com/singhJasvinder101/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (u *User) Save() error {
	query := `insert into users (email, password) values (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		fmt.Println("Error executing query in Save:", err)
		return err
	}

	userId, err := res.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) Authenticate() (error){
	query := `select id, email, password from users where email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPaswd string;
	err := row.Scan(&u.ID, &u.Email, &retrievedPaswd)

	if err != nil {
		return err
	}
	
	isPasswordValid := utils.ValidatePassword(retrievedPaswd, u.Password)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	u.Password = retrievedPaswd
	return nil
}


