package db

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

const (
	errScanRowMsg       = "Failed to scan row: %v"
	errExecSqlQueryMsg  = "Failed to execute SQL query: %v"
	selectCustomerQuery = "SELECT id, name FROM customer"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := `
		-- Your SQL script here
		INSERT INTO customer (user_name, email, name) VALUES ('hesa', 'hesa@go-lesson.com', 'Hesa');
	`
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		t.Errorf("Failed to execute SQL script: %v", err)
		panic(err)
	}
	fmt.Println("SQL script executed successfully")
}

func TestQueryDefaultSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, selectCustomerQuery)
	if err != nil {
		t.Errorf(errExecSqlQueryMsg, err)
		panic(err)
	}
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Errorf(errScanRowMsg, err)
			return
		}
	}
	rows.Close()
	fmt.Println("Query executed successfully")
}

func TestQueryWGSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, selectCustomerQuery)
	if err != nil {
		t.Errorf(errExecSqlQueryMsg, err)
		panic(err)
	}
	// var wg sync.WaitGroup
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		// wg.Add(1)
		go func(id, name string, err error) {
			// defer wg.Done()
			if err != nil {
				t.Errorf(errScanRowMsg, err)
				return
			}
		}(id, name, err)
	}
	// wg.Wait()
	rows.Close()
	fmt.Println("Query executed successfully")
}

func BenchmarkQueryWGSql(b *testing.B) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, selectCustomerQuery)
	if err != nil {
		b.Errorf(errExecSqlQueryMsg, err)
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for rows.Next() {
			var id, name string
			err := rows.Scan(&id, &name)
			wg.Add(1)
			go func(id, name string, err error) {
				defer wg.Done()
				if err != nil {
					b.Errorf(errScanRowMsg, err)
					return
				}
			}(id, name, err)
		}
		wg.Wait()
		rows.Close()
	}
}

func BenchmarkQueryDefaultSql(b *testing.B) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, selectCustomerQuery)
	if err != nil {
		b.Errorf(errExecSqlQueryMsg, err)
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		for rows.Next() {
			var id, name string
			err := rows.Scan(&id, &name)
			if err != nil {
				b.Errorf(errScanRowMsg, err)
				return
			}
			_ = id
			_ = name
		}
		rows.Close()
	}
}

func TestSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	name := "Sandy"
	script := `SELECT id, name FROM customer WHERE name = $1`
	rows, err := db.QueryContext(ctx, script, name)
	if err != nil {
		t.Errorf(errExecSqlQueryMsg, err)
		panic(err)
	}
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Errorf(errScanRowMsg, err)
			return
		}
	}
	rows.Close()
	fmt.Println("Query executed successfully")
}
