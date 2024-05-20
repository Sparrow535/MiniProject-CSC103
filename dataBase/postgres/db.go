package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "dpg-cp5hiun79t8c73eue3c0-a.singapore-postgres.render.com"
	port     = 5432
	user     = "postgres_admin"
	password = "BJCc5QXHxzBFehK85U46RYTMcfwfJ2S5"
	dbname   = "my_db_0pgs"
)

var Db *sql.DB

func init() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	var err error
	Db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database successfully configured")

}
