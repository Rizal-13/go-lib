package main

import (
	"fmt"
	"time"
)


func main(){

	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(1997, time.June, 20, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(parse)
}