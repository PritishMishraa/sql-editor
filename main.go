package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sql-editor/web"
)

var db *sql.DB

func main() {
	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/connect", connectHandler)
	http.HandleFunc("/tables", listTablesHandler)
	http.HandleFunc("/query", queryHandler)

	web.Serve()
}

func sendErrorResponse(w http.ResponseWriter, errorMessage string, statusCode int, err error) {
	response := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Error:   http.StatusText(statusCode),
		Message: errorMessage,
	}
	if err != nil {
		response.Message += " " + err.Error()
	}
	jsonBytes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
