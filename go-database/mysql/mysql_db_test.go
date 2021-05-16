package mysql

import (
	"context"
	"fmt"
	"testing"
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