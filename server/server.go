package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/younesious/mailinglist/db"
	grpcapi "github.com/younesious/mailinglist/grpc_api"
	jsonapi "github.com/younesious/mailinglist/json_api"
	"github.com/younesious/mailinglist/proto/email"
	"google.golang.org/grpc"

	_ "github.com/mattn/go-sqlite3"
)

func runHTTPServer(addr string, database *sql.DB) {
	http.HandleFunc("/email/add", jsonapi.AddEmailHandler(database))
	http.HandleFunc("/email/get", jsonapi.GetEmailHandler(database))
	http.HandleFunc("/email/update", jsonapi.UpdateEmailHandler(database))
	http.HandleFunc("/email/delete", jsonapi.DeleteEmailHandler(database))
	http.HandleFunc("/email/batch", jsonapi.GetEmailBatchHandler(database))

	log.Printf("Starting HTTP server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func runGRPCServer(addr string, database *sql.DB) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	email.RegisterEmailServiceServer(s, grpcapi.NewEmailServiceServer(database))

	log.Printf("Starting gRPC server on %s...\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
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
	wg.Add(2)

	go func() {
		defer wg.Done()
		runHTTPServer(":8080", database)
	}()

	go func() {
		defer wg.Done()
		runGRPCServer(":8081", database)
	}()

	wg.Wait()
}
