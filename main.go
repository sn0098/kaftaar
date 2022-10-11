package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	cols, err := fetchCols("cols.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	query := getQuery(cols)

	fmt.Println(query)
}

func fetchCols(addr string) ([]column, error) {
	var cols []column

	content, err := os.ReadFile(addr)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &cols)
	if err != nil {
		return nil, err
	}

	return cols, nil
}

func getQuery(cols []column) string {
	query := "CREATE TABLE table_name("
	query += fmt.Sprintf("%v %v", cols[0].Name, cols[0].Type)

	for _, col := range cols[1:] {
		query += fmt.Sprintf(", %v %v", col.Name, col.Type)
	}

	query += ")"

	return query
}
