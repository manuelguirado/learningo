package main
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	dsn := "manudev:270504@tcp(192.168.1.145:3306)/go_db"
	db , err := sql.Open("mysql" , dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM USER ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	insert , err := db.Query("INSERT INTO USER  (nombre,apellido, telefono) VALUES ('manuel','perez',6788908900)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	
}