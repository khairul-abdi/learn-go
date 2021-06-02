package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmtInsert, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
	stmtInsert.Exec("G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("insert success!")

	stmtUpdate, err := db.Prepare("update tb_student set age = ? where id = ?")
	stmtUpdate.Exec(28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	stmtDelete, err := db.Prepare("delete from tb_student where id = ?")
	stmtDelete.Exec("G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}

func main() {
	sqlExec()
}
