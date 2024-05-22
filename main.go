package main

import (
	"database/sql"
	"log"
	"net/http"
	"sync"

	"github.com/younesious/mailinglist/db"
	"github.com/younesious/mailinglist/jsonapi"

	_ "github.com/mattn/go-sqlite3"
)

func runServer(addr string, database *sql.DB) {
	http.HandleFunc("/email/add", jsonapi.AddEmailHandler(database))
	http.HandleFunc("/email/get", jsonapi.GetEmailHandler(database))
	http.HandleFunc("/email/update", jsonapi.UpdateEmailHandler(database))
	http.HandleFunc("/email/delete", jsonapi.DeleteEmailHandler(database))
	http.HandleFunc("/email/batch", jsonapi.GetEmailBatchHandler(database))

	log.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// we can hard code pass mailinglist.db file for this small project or read from env file for more complex project
	database, err := sql.Open("sqlite3", "mailinglist.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := db.InitializeDB(database); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		runServer(":8080", database)
	}()

	wg.Wait()
}
