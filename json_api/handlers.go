package jsonapi

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/younesious/mailinglist/db"
)

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json; charset=utf8")
}

func fromJson[T any](body io.Reader, target T) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	if err != nil {
		return err
	}
	json.Unmarshal(buf.Bytes(), target)

	return nil
}

func toJson[T any](w http.ResponseWriter, withDate func() (T, error)) {
	setJsonHeader(w)

	data, serverErr := withDate()
	if serverErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		serverErrJson, err := json.Marshal(map[string]string{"error": serverErr.Error()})
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(serverErrJson)
		return
	}

	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(dataJson)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	toJson(w, func() (interface{}, error) {
		w.WriteHeader(code)
		return map[string]string{"error": err.Error()}, nil
	})
}

func AddEmailHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			returnErr(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
			return
		}

		email := db.EmailEntry{}
		if err := fromJson(r.Body, &email); err != nil {
			returnErr(w, err, http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		id, err := db.AddEmailEntry(ctx, database, email.Email)
		if err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}

		toJson(w, func() (interface{}, error) {
			return map[string]int64{"id": id}, nil
		})
	}
}

func GetEmailHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			returnErr(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
			return
		}

		email := r.URL.Query().Get("email")
		if email == "" {
			returnErr(w, errors.New("email is required"), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		entry, err := db.GetEmailEntry(ctx, database, email)
		if err != nil {
			if err == sql.ErrNoRows {
				returnErr(w, errors.New("email not found"), http.StatusNotFound)
				return
			}
			returnErr(w, err, http.StatusInternalServerError)
			return
		}

		toJson(w, func() (interface{}, error) {
			return entry, nil
		})
	}
}

func UpdateEmailHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			returnErr(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
			return
		}

		var entry db.EmailEntry
		if err := fromJson(r.Body, &entry); err != nil {
			returnErr(w, err, http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		if err := db.UpdateEmailEntry(ctx, database, entry); err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}

		toJson(w, func() (interface{}, error) {
			return map[string]string{"message": "email updated successfully"}, nil
		})
	}
}

func DeleteEmailHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			returnErr(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
			return
		}

		email := r.URL.Query().Get("email")
		if email == "" {
			returnErr(w, errors.New("email is required"), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		if err := db.DeleteEmailEntry(ctx, database, email); err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}

		toJson(w, func() (interface{}, error) {
			return map[string]string{"message": "email deleted successfully"}, nil
		})
	}
}

func GetEmailBatchHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			returnErr(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
			return
		}

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil || count <= 0 {
			count = 10
		}

		params := db.GetEmailBatchQueryParam{
			Page:  page,
			Count: count,
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		batch, err := db.GetEmailBatch(ctx, database, params)
		if err != nil {
			returnErr(w, err, http.StatusInternalServerError)
			return
		}

		toJson(w, func() (interface{}, error) {
			return batch, nil
		})
	}
}
