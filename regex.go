package main

import (
	"fmt"
	"regexp"
)


func main(){

	regex := regexp.MustCompile(`ins([a,i,u,e,o])n`)

	fmt.Println(regex.MatchString("insun"))
	fmt.Println(regex.MatchString("insoen"))
	fmt.Println(regex.MatchString("inzun"))
	fmt.Println(regex.MatchString("insan"))

	fmt.Println(regex.FindAllString("insun inson insan insoen sasi isna nszun", 10))

}