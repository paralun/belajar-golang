package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestOpenConnection(t *testing.T) {
	db := GetConnectionMysql()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func TestExecSql(t *testing.T) {
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

func TestQuerySql(t *testing.T) {
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

func TestQueryNoCtx(t *testing.T) {
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

func TestQueryWithTypeData(t *testing.T) {
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

func TestSqlParameter(t *testing.T) {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	user := "admin"
	pass := "pas1234"
	script := "INSERT INTO user (username, password) VALUES (?, ?)"

	_, err := db.ExecContext(ctx, script, user, pass)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Success")
}

func TestQueryParameter(t *testing.T) {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	user := "admin"
	pass := "pas1234"
	script := "SELECT username FROM user WHERE username = ? AND password = ?"

	rows, err := db.QueryContext(ctx, script, user, pass)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestGetAutoIncrement(t *testing.T) {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	user := "guest"
	pass := "guest123"
	script := "INSERT INTO user (username, password) VALUES (?, ?)"

	result, err := db.ExecContext(ctx, script, user, pass)
	if err != nil {
		panic(err)
	}

	insertId, err2 := result.LastInsertId()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Insert Success -> ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO user (username, password) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		user := "guest" + strconv.Itoa(i)
		pass := "passs" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, user, pass)
		if err != nil {
			panic(err)
		}

		insertId, err2 := result.LastInsertId()
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("Insert Success -> ID:", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnectionMysql()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO user (username, password) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		user := "roll" + strconv.Itoa(i)
		pass := "back" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, user, pass)
		if err != nil {
			panic(err)
		}

		insertId, err2 := result.LastInsertId()
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("Insert Success -> ID:", insertId)
	}

	//err3 := tx.Commit()
	err3 := tx.Rollback()
	if err3 != nil {
		panic(err3)
	}
}
