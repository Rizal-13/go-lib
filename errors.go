package main

import (
	"errors"
	"fmt"
)

var(
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error{
	if id == ""{
		return ValidationError
	} 
	
	if id !="insun" {
		return NotFoundError
	}

	return nil
}

func main(){

err := GetById("insaun")
if err != nil{
	if errors.Is(err, ValidationError){
		fmt.Println("validation error")
	} else if errors.Is(err, NotFoundError){
		fmt.Println("not Found Error")
	} else{
		fmt.Println("unknown error")
	}

}

}