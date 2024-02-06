package main

import (
	"fmt"
	"time"
)


func main(){

	var duration1 time.Duration = time.Second * 100
	duration2 := time.Second * 10
	duration3 := duration1-duration2
	
	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)



}