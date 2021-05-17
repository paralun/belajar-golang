package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestOpenConnection(t *testing.T)  {
	db := GetConnectionMysql()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func TestExecSql(t *testing.T)  {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO task(id, item) VALUES('t001', 'Belajar Golang')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert")
}

func TestQuerySql(t *testing.T)  {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, item FROM task"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, item string
		err := rows.Scan(&id, &item)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Item:", item)
	}
}

func TestQueryNoCtx(t *testing.T)  {
	db := GetConnectionMysql()
	defer db.Close()

	script := "SELECT id, item FROM task"
	rows, err := db.Query(script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, item string
		err := rows.Scan(&id, &item)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Item:", item)
	}
}

func TestQueryWithTypeData(t *testing.T)  {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		/**
		var id, name, email string // varchar, char
		var balance int64 // int, bigint
		var rating float64 // float, double
		var birthDate time.Time // date, datetime
		var createdAt time.Time // time, timestamp
		var married bool //boolean
		**/

		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate sql.NullTime
		var married sql.NullBool
		var cratedAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &cratedAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("================")
		fmt.Println("ID:", id)
		fmt.Println("NAME:", name)
		if email.Valid {
			fmt.Println("EMAIL:", email.String)
		}

		if balance.Valid {
			fmt.Println("BALANCE:", balance.Int32)
		}

		if rating.Valid {
			fmt.Println("RATING:", rating.Float64)
		}

		if birthDate.Valid {
			fmt.Println("BIRTH_DATE:", birthDate.Time)
		}

		if married.Valid {
			fmt.Println("MARRIED:", married.Bool)
		}

		fmt.Println("CREATED_AT:", cratedAt)

	}
}