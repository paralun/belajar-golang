package main

import (
	"fmt"
	"time"
)

const (
	LAYOUTISO = "2006-01-02"
	LAYOUTUS = "January 2, 2006"
)

func main() {
	date := "1999-12-31"
	t, _ := time.Parse(LAYOUTISO, date)
	fmt.Println(t)
	fmt.Println(t.Format(LAYOUTUS))
}
