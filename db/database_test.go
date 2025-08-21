package db

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := `
		-- Your SQL script here
		INSERT INTO customer (user_name, email, name) VALUES ('sandy', 'sandy@go-lesson.com', 'Sandy');
	`
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		t.Errorf("Failed to execute SQL script: %v", err)
		panic(err)
	}

	fmt.Println("SQL script executed successfully")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		t.Errorf("Failed to execute SQL script: %v", err)
		panic(err)
	}
	results := make(chan struct {
		id, name string
		err      error
	})

	go func() {
		for rows.Next() {
			var id, name string
			err := rows.Scan(&id, &name)
			results <- struct {
				id, name string
				err      error
			}{id, name, err}
		}
		close(results)
	}()

	for result := range results {
		if result.err != nil {
			t.Errorf("Failed to scan row: %v", result.err)
			continue
		}
		fmt.Printf("ID: %s, Name: %s\n", result.id, result.name)
	}
	defer rows.Close()
}
