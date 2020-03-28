package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dataBase = "root:11111@(127.0.0.1:3306)/?loc=Local&parseTime=true"

func Init() {
	var err error

	DB, err = sql.Open("mysql", dataBase)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}

	DB.SetMaxOpenConns(2)

	err = DB.Ping()

	if err != nil {
		log.Fatalln("ping db fail:", err)
	}
}

func main () {
	Init()

	// 开启20个goroutine

	for i :=0; i <20; i++ {
		go oneWorker(i)
	}

	select {

	}
}

func oneWorker(i int) {
	var connection_id int
	err := DB.QueryRow("select CONNECTION_ID()").Scan(&connection_id)
	if err != nil {
		log.Println("query connection id failed:", err)
		return
	}

	log.Println("worker:", i, ", connection id:", connection_id)

	var result int
	err = DB.QueryRow("select sleep(10)").Scan(&result)
	if err != nil {
		log.Println("query sleep connection id faild:", err)
		return
	}
}