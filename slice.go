package main

import (
	"fmt"
	"slices"
)


func main(){

	name := []string{"insun", "crypto", "master", "inzun", "hartono"}
	numb := [] int{12,56,98,46,65,45}

	fmt.Println(slices.Max(numb))
	fmt.Println(slices.Min(numb))
	fmt.Println(slices.Contains(name, "insun"))
	fmt.Println(slices.Index(name, "hartono"))
	

}