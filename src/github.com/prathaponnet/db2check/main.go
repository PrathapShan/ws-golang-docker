package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	_ "github.com/asifjalil/cli"
)

func main() {
	fmt.Println("starting http server")
	dataSource := "DATABASE=" + "DBNAME" + "; HOSTNAME=" + "148.171.229.2" + "; PORT=" + "7260" + "; PROTOCOL=" + "TCPIP" + "; UID=" + "name" + "; PWD=" + "password" + ";"
	db, err := sql.Open("cli", dataSource)
	fmt.Println(err)
	fmt.Println(db.Ping())		
	fmt.Println("After Ping")	
	rows, err := db.Query("SELECT COUNT(1) FROM CE1.T5073")

	defer rows.Close()

	var count int

	if err != nil {
		fmt.Println("error executing query ", err)
	} else {

		for rows.Next() {
			if err := rows.Scan(&count); err != nil {
				fmt.Println("error finding count ", err)
			}

		}
	}

	fmt.Println("Number of rows : ", count)
	defer db.Close()
	
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
	
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello from a docker container")
}
