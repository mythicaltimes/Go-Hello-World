package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//Load a csv
	f, _ := os.Open("12-09-2018.csv")

	r := csv.NewReader(bufio.NewReader(f))

	//Connect to Server
	db, err := sql.Open("mysql", "username:password@tcp(cloud.local.riedstra.us:53306)/brad")

	if err != nil {
		log.Fatal(err)
	}
	tx, _ := db.Begin()
	stmt, err := tx.Prepare("INSERT INTO WeeklyInformation(UserName, Role, Donated, Received, WPR, FBWR, MFB, MCB) VALUES(?,?,?,?,?,?,?,?)")

	for {
		record, err := r.Read()
		//Check for EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}

		stmt.Exec(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7])
	}
	tx.Commit()
}
