package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = ""
	port     = 0
	user     = ""
	password = ""
	db       = ""
)

//Connect crea conexion a la base de datos
func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

//Disconnect cierra conexion a la base de datos
func Disconnect(db *sql.DB) {
	db.Close()
}

var DB = Connect()
