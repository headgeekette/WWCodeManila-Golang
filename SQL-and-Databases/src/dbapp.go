package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := ConnectToDb()
	// defer db.Close()
	CheckConnection(db)

	CreateTable(db)
	InsertNumbers(db)
	RetrieveNumbers(db)
	UpdateRow(db)
	DeleteData(db)
	DropTable(db)

	db.Close()
}

func ConnectToDb() *sql.DB {
	db, err := sql.Open("mysql", "gouser:gopasswd@tcp(mariadb:3306)/sakila")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CheckConnection(db *sql.DB) {
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

func CreateTable(db *sql.DB) {
	// Create the database handle, confirm that the driver is present
	// db, err := sql.Open("mysql", "user:password@tcp(address:port)/dbname")
	// db, err := sql.Open("mysql", "test:@tcp(localhost:3306)/sakila")
	tableCreate := `CREATE TABLE IF NOT EXISTS number(number integer NOT NULL, property text NOT NULL)`
	_, err := db.Exec(tableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table was successfully created")
	}
}

func DropTable(db *sql.DB) {
	stmt := `DROP TABLE number`
	_, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table was successfully dropped.")
	}
}

func InsertNumbers(db *sql.DB) {
	var prop string
	insert, err := db.Prepare("INSERT INTO number VALUES(?,?)")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "EVEN"
		} else {
			prop = "ODD"
		}
		_, err = insert.Exec(i, prop)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The number:", i, "is:", prop)
		}
	}
	insert.Close()
	fmt.Printf("The Numbers are ready.")
}

func RetrieveNumbers(db *sql.DB) {
	var number int
	var property string
	rows, err := db.Query("SELECT * from number")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&number, &property)
		if err != nil {
			panic(err)
		}
		fmt.Println(number, property)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	rows.Close()
}

func UpdateRow(db *sql.DB) {
	num := 0
	ppt := "ODD"
	stmt := `UPDATE number SET property = ? WHERE number = ?`

	// Form 1
	// result, err := db.Exec(stmt, ppt, num)
	// if err != nil {
	// 	panic(err)
	// }

	// Form 2
	ps, err := db.Prepare(stmt)
	if err != nil {
		panic(err)
	}
	result, err := ps.Exec(ppt, num)
	if err != nil {
		panic(err)
	}

	upd, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of records updated:", upd)
}

func DeleteData(db *sql.DB) {
	// This can also be done using truncate if you intend to delete all the table's contents.
	stmt := `TRUNCATE TABLE number`
	// stmt := `DELETE FROM number`
	ps, err := db.Prepare(stmt)
	if err != nil {
		panic(err)
	}
	result, err := ps.Exec()
	if err != nil {
		panic(err)
	}
	upd, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of records updated:", upd)
}
