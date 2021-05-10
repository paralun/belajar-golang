package main

import (
	"fmt"
	"time"
)

func main() {
	waktu := time.Now()

	fmt.Println(waktu.Local())
	fmt.Println(waktu.Year())
	fmt.Println(waktu.Month())
	fmt.Println(waktu.Day())
	fmt.Println(waktu.Hour())
	fmt.Println(waktu.Minute())
	fmt.Println(waktu.Second())
	fmt.Println(waktu.Nanosecond())
	fmt.Println(waktu.Weekday())
	fmt.Println("")

	//utc := time.Date(2020, 8, 1, 12, 30, 40, 100, time.UTC)
	utc := time.Date(2020, time.August, 1, 12, 30, 40, 100, time.UTC)
	fmt.Println(utc)
	fmt.Println("")

	//layout := "2006-01-02"
	//parse, _ := time.Parse(layout, "2021-02-23")
	parse, _ := time.Parse(time.RFC3339, "2021-02-23T14:04:05Z")
	fmt.Println(parse)
}
