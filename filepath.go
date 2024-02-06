package main

import (
	"fmt"
	"path/filepath"
)


func main(){

	fmt.Println(filepath.Dir("insun/insun.go"))
	fmt.Println(filepath.Base("insun/insun.go"))
	fmt.Println(filepath.Ext("insun/insun.go"))
	fmt.Println(filepath.Ext("insun/insun.css"))
	fmt.Println(filepath.Join("insun", "file", "insun.go"))

}