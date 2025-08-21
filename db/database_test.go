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

	script := "SELECT id, name FROM customer where name = 'sandy'"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		t.Errorf("Failed to execute SQL script: %v", err)
		panic(err)
	}

	defer rows.Close()
}
