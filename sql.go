package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type musicItem struct {
    id int
    external_ids sql.NullString
    Name string
    Artist string
    price float64
    seller sql.NullString
    note sql.NullString
    purchase_date sql.NullString
}

func openDB(DBName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBName+".db")
	if err != nil {
		panic("couldn't open db")
	}

	return db, nil
}

func loadMusic() []musicItem {
	db, err := openDB("walletdrain")
	if err != nil {
		panic("couldn't open db")
	}
	defer db.Close()

	query := "SELECT * FROM music"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var results []musicItem
	for rows.Next() {
		var result musicItem
		err := rows.Scan(
			&result.id,
			&result.external_ids,
			&result.Name,
			&result.Artist,
			&result.price,
			&result.seller,
			&result.note,
			&result.purchase_date)
		if err != nil {
			log.Print(err)
		}

		results = append(results, result)
	}
	return results
}

