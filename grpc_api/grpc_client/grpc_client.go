package main

import (
	"context"
	"log"
	"time"

	email "github.com/younesious/mailinglist/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("myserver:8081", grpc.WithInsecure()) // if you wanna run with docker-compose.yml and the server container name is myserver
	// conn, err := grpc.Dial(":8081", grpc.WithInsecure())  # comment out if you wanna run in your local
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := email.NewEmailServiceClient(conn)

	addEmail(client, "younesious@gmail.com")
	getEmail(client, "younesious@gmail.com")
	updateEmail(client, 1, time.Now().Unix(), true)
	getEmail(client, "younesious@gmail.com")
	deleteEmail(client, "younesious@gmail.com")
	getEmailBatch(client, 1, 10)
}

func logResponse(res interface{}, err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("response: %v", res)
}

func addEmail(client email.EmailServiceClient, addr string) {
	log.Println("Adding email")
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	res, err := client.AddEmail(ctx, &email.AddEmailRequest{Email: addr})
	logResponse(res, err)
}

func getEmail(client email.EmailServiceClient, addr string) {
	log.Println("Getting email")
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	res, err := client.GetEmail(ctx, &email.GetEmailRequest{Email: addr})
	logResponse(res, err)
}

func updateEmail(client email.EmailServiceClient, id int64, confirmedAt int64, optOut bool) {
	log.Println("Updating email")
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	entry := &email.EmailEntry{
		Id:          id,
		ConfirmedAt: confirmedAt,
		OptOut:      optOut,
	}

	res, err := client.UpdateEmail(ctx, &email.UpdateEmailRequest{Entry: entry})
	logResponse(res, err)
}

func deleteEmail(client email.EmailServiceClient, addr string) {
	log.Println("Deleting email")
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	res, err := client.DeleteEmail(ctx, &email.DeleteEmailRequest{Email: addr})
	logResponse(res, err)
}

func getEmailBatch(client email.EmailServiceClient, page int32, count int32) {
	log.Println("Getting email batch")
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	res, err := client.GetEmailBatch(ctx, &email.GetEmailBatchRequest{Page: page, Count: count})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, entry := range res.Entries {
		log.Printf("ID: %d, Email: %s, ConfirmedAt: %d, OptOut: %v", entry.Id, entry.Email, entry.ConfirmedAt, entry.OptOut)
	}
}
