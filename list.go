package main

import (
	"container/list"
	"fmt"
)


func main(){

	data := list.New()
	data.PushBack("Insun")
	data.PushBack("Crypto")
	data.PushBack("Master")

	// Cara 1
	for e:= data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}

	// Cara 2
	head := data.Front()
		fmt.Println(head.Value)

	next := head.Next()
		fmt.Println(next.Value)
	
	next2 := next.Next()
		fmt.Println(next2.Value)
	
	

	

}