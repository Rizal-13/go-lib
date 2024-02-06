package main

import (
	"fmt"
	"strings"
)


func main(){

	fmt.Println(strings.Contains("insun crypto", "insun"))
	fmt.Println(strings.Split("insun crypto", " "))
	fmt.Println(strings.ToLower("Insun Crypto Insun"))
	fmt.Println(strings.ToUpper("insun crypto insun"))
	fmt.Println(strings.Trim("      insun crypto  ", " "))
	fmt.Println(strings.ReplaceAll("insun crypto", "insun", "master"))

}