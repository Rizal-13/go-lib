package main

import (
	"fmt"
	"path"
)


func main(){

	fmt.Println(path.Dir("insun/insun.go"))
	fmt.Println(path.Base("insun/insun.go"))
	fmt.Println(path.Ext("insun/insun.go"))
	fmt.Println(path.Ext("insun/insun.css"))
	fmt.Println(path.Join("insun", "file", "insun.go"))

}