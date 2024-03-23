package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type TableInfo struct {
	Name   string   `json:"table_name"`
	Schema []Column `json:"schema"`
}

type Column struct {
	Name     string `json:"column_name"`
	DataType string `json:"data_type"`
	Primary  bool   `json:"primary"`
}

func listTablesHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		http.Error(w, "Database not connected", http.StatusBadRequest)
		return
	}

	tables, err := getTables()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting tables: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tables)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error serializing data to JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(data)
}

func getTables() ([]TableInfo, error) {
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type = 'table'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []TableInfo
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}

		tableInfo, err := getTableSchema(tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableInfo)
	}

	return tables, nil
}

func getTableSchema(tableName string) (TableInfo, error) {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(%s)", tableName))
	if err != nil {
		return TableInfo{}, err
	}
	defer rows.Close()

	var tableInfo TableInfo
	tableInfo.Name = tableName

	for rows.Next() {
		var cid int
		var name string
		var dataType string
		var notnull string
		var dflt_value sql.NullString
		var pk string
		if err := rows.Scan(&cid, &name, &dataType, &notnull, &dflt_value, &pk); err != nil {
			return TableInfo{}, err
		}

		column := Column{
			Name:     name,
			DataType: dataType,
			Primary:  pk == "1",
		}
		tableInfo.Schema = append(tableInfo.Schema, column)
	}

	return tableInfo, nil
}
