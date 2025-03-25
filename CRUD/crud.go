package main

import (
	
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
func insertData(db *sql.DB, id int, nombre string, stock int) {

	_, err := db.Exec("INSERT INTO PRODUCTOS (ID,NOMBRE,STOCK) VALUES (?, ?, ?)", id, nombre, stock)
	if err != nil {
		panic(err)
	}
	
}
func updateData(db *sql.DB, id int, nombre string, stock int) {
	_, err := db.Exec("UPDATE PRODUCTOS SET NOMBRE = ?, STOCK = ? WHERE ID = ?", nombre, stock, id)
	if err != nil {
		panic(err)
	}
}
func deletedata(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM PRODUCTOS WHERE ID = ?", id)
	if err != nil {
		panic(err)
	}

}
func selectData(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM PRODUCTOS")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre string
		var stock int
		err := rows.Scan(&id, &nombre, &stock)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, nombre, stock)
	}
}
	
func main() {
	fmt.Println("Conectando a la base de datos")
	dsn := "manudev:270504@tcp(192.168.1.145:3306)/example"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Print("que operacion desea realizar: 1. Insertar 2. Actualizar 3. Eliminar 4. Consultar")
	var operacion int
	fmt.Scanf("%d" , &operacion)
	switch operacion {
	case 1 :
		fmt.Println("Insertar")
		insertData(db, 1, "manzana", 10)
		insertData(db, 2, "pera", 20)
	case 2:
		fmt.Println("Actualizar")
		updateData(db, 1, "manzana", 20)
	case 3:
		fmt.Println("Eliminar")
		deletedata(db, 2)
	case 4:
		fmt.Println("Consultar")
		selectData(db)
	default:
		fmt.Println("Operacion no valida")
	}
	


}