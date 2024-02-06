package main

import (
	"encoding/base64"
	"fmt"
)


func main(){

	data := ("insun crypto 47")
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil{
		fmt.Println(err.Error())
	} else{
		fmt.Println(string(decoded))
	}

}