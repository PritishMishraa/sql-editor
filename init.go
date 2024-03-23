package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func initHandler(w http.ResponseWriter, r *http.Request) {
	rootDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Error getting current directory", http.StatusInternalServerError)
		return
	}

	databases := searchSQLiteDatabases(rootDir)

	data, err := json.Marshal(databases)
	if err != nil {
		http.Error(w, "Error serializing data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(data)
}