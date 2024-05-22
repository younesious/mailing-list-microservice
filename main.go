package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	// Adjust the import path to your actual project structure

	_ "github.com/mattn/go-sqlite3"
	"github.com/younesious/mailinglist/db"
)

func main() {
	// Open the database connection
	database, err := sql.Open("sqlite3", "mailinglist.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Initialize the database
	if err := db.InitializeDB(database); err != nil {
		log.Fatal(err)
	}

	// Create a context with a timeout for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure resources are cleaned up

	// Add a new email entry
	emailID, err := db.AddEmailEntry(ctx, database, "test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Added Email Entry with ID:", emailID)
	// Retrieve the email entry
	emailEntry, err := db.GetEmailEntry(ctx, database, "test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Retrieved Email Entry:", emailEntry)

	// Update the email entry
	emailEntry.OptOut = true
	err = db.UpdateEmailEntry(ctx, database, *emailEntry)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email entry was updated successfully")

	// Retrieve the email entry
	emailEntry, err = db.GetEmailEntry(ctx, database, "test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Retrieved Email Entry:", emailEntry)

	// Delete the email entry
	err = db.DeleteEmailEntry(ctx, database, "test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email entry deleted successfully")

	// Add a new email entry
	emailID, err = db.AddEmailEntry(ctx, database, "test4@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Added Email Entry with ID:", emailID)

	// Add a new email entry
	emailID, err = db.AddEmailEntry(ctx, database, "test2@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Added Email Entry with ID:", emailID)

	// Add a new email entry
	emailID, err = db.AddEmailEntry(ctx, database, "test3@example.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Added Email Entry with ID:", emailID)

	// Retrieve a batch of email entries
	params := db.GetEmailBatchQueryParam{
		Page:  1,
		Count: 5,
	}
	emailBatch, err := db.GetEmailBatch(ctx, database, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Retrieved Email Batch:", emailBatch)
}
