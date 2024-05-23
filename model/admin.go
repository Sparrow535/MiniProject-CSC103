package model

import (
	"myapp/dataBase/postgres"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
}

const queryInsertAdmin = "INSERT INTO admin(fname, lname, email, password, gender) VALUES ($1, $2, $3, $4, $5) RETURNING email;"

func (admin *Admin) CreateAdmin() error {
	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)

	row := postgres.Db.QueryRow(queryInsertAdmin, admin.FirstName, admin.LastName, admin.Email, admin.Password, admin.Gender)
	err = row.Scan(&admin.Email)
	return err
}

const querySelectAdmin = "SELECT email, password, gender FROM admin WHERE email=$1;"

func (admin *Admin) GetAdmin() error {
	var hashedPassword string
	err := postgres.Db.QueryRow(querySelectAdmin, admin.Email).Scan(&admin.Email, &hashedPassword, &admin.Gender)
	if err != nil {
		return err
	}

	// Compare the hashed password with the password provided
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(admin.Password))
	if err != nil {
		return err
	}

	// Set the password to the hashed password after verification
	admin.Password = hashedPassword
	return nil
}
