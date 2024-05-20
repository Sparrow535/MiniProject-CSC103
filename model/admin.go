package model

import "myapp/dataBase/postgres"

type Admin struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
}

const queryInsertAdmin = "INSERT INTO admin(fname, lname, email, password, gender) VALUES ($1, $2, $3, $4, $5) RETURNING email;"

func (admin *Admin) CreateAdmin() error {
	row := postgres.Db.QueryRow(queryInsertAdmin, admin.FirstName, admin.LastName, admin.Email, admin.Password, admin.Gender)
	err := row.Scan(&admin.Email)
	return err
}

const querySelectAdmin = "SELECT email, password, gender FROM admin WHERE email=$1 and password=$2;"

func (admin *Admin) GetAdmin() error {
	return postgres.Db.QueryRow(querySelectAdmin, admin.Email, admin.Password).Scan(&admin.Email, &admin.Password, &admin.Gender)
}
