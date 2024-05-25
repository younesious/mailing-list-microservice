package grpcapi

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/younesious/mailinglist/db"
	email "github.com/younesious/mailinglist/proto"
)

type EmailServiceServer struct {
	email.UnimplementedEmailServiceServer
	db *sql.DB
}

func NewEmailServiceServer(db *sql.DB) *EmailServiceServer {
	return &EmailServiceServer{db: db}
}

func dbToProto(entry *db.EmailEntry) *email.EmailEntry {
	confirmedAt := int64(0)
	if entry.ConfirmedAt != nil {
		confirmedAt = entry.ConfirmedAt.Unix()
	}

	return &email.EmailEntry{
		Id:          entry.ID,
		Email:       entry.Email,
		ConfirmedAt: confirmedAt,
		OptOut:      entry.OptOut,
	}
}

func protoToDb(entry *email.EmailEntry) db.EmailEntry {
	confirmedAt := time.Unix(entry.ConfirmedAt, 0)

	return db.EmailEntry{
		ID:          entry.Id,
		Email:       entry.Email,
		ConfirmedAt: &confirmedAt,
		OptOut:      entry.OptOut,
	}
}

func (s *EmailServiceServer) AddEmail(ctx context.Context, req *email.AddEmailRequest) (*email.AddEmailResponse, error) {
	id, err := db.AddEmailEntry(ctx, s.db, req.Email)
	if err != nil {
		return nil, err
	}
	return &email.AddEmailResponse{Id: id}, nil
}

func (s *EmailServiceServer) GetEmail(ctx context.Context, req *email.GetEmailRequest) (*email.GetEmailResponse, error) {
	entry, err := db.GetEmailEntry(ctx, s.db, req.Email)
	if err != nil {
		return nil, err
	}
	return &email.GetEmailResponse{
		Entry: dbToProto(entry),
	}, nil
}

func (s *EmailServiceServer) UpdateEmail(ctx context.Context, req *email.UpdateEmailRequest) (*email.UpdateEmailResponse, error) {
	entry := protoToDb(req.Entry)
	log.Printf("Updating email entry: %+v\n", entry)
	err := db.UpdateEmailEntry(ctx, s.db, entry)
	if err != nil {
		return nil, err
	}
	return &email.UpdateEmailResponse{Message: "email updated successfully"}, nil
}

func (s *EmailServiceServer) DeleteEmail(ctx context.Context, req *email.DeleteEmailRequest) (*email.DeleteEmailResponse, error) {
	err := db.DeleteEmailEntry(ctx, s.db, req.Email)
	if err != nil {
		return nil, err
	}
	return &email.DeleteEmailResponse{Message: "email deleted successfully"}, nil
}

func (s *EmailServiceServer) GetEmailBatch(ctx context.Context, req *email.GetEmailBatchRequest) (*email.GetEmailBatchResponse, error) {
	params := db.GetEmailBatchQueryParam{
		Page:  int(req.Page),
		Count: int(req.Count),
	}

	entries, err := db.GetEmailBatch(ctx, s.db, params)
	if err != nil {
		return nil, err
	}

	grpcEntries := make([]*email.EmailEntry, len(entries))
	for i, entry := range entries {
		grpcEntries[i] = dbToProto(&entry)
	}

	return &email.GetEmailBatchResponse{Entries: grpcEntries}, nil
}
