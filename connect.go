package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func connectHandler(w http.ResponseWriter, r *http.Request) {
	var dbInfo DatabaseInfo
	err := json.NewDecoder(r.Body).Decode(&dbInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	err = connectDatabase(dbInfo)
	if err != nil {
		sendErrorResponse(w, "Error connecting to database:", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func connectDatabase(dbInfo DatabaseInfo) error {
	database, err := sql.Open("sqlite3", dbInfo.Path)
	if err != nil {
		return err
	}
	err = database.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	db = database
	fmt.Printf("Connected to database: %s\n", dbInfo.Name)
	return nil
}
