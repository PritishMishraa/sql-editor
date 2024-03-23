package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type QueryInfo struct {
	Query string `json:"query"`
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	var queryInfo QueryInfo
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	result, err := queryDatabase(queryInfo)
	if err != nil {
		sendErrorResponse(w, "Error executing query:", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func queryDatabase(queryInfo QueryInfo) (interface{}, error) {
	rows, err := db.Query(queryInfo.Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		rowData := make(map[string]interface{})
		for i, value := range values {
			rowData[columns[i]] = value
		}
		result = append(result, rowData)
	}

	return result, nil
}
