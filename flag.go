package main

import (
	"flag"
	"fmt"
)


func main(){

	host := flag.String("host", "0.0.0.0", "Put your database host")
	username := flag.String("username", "root", "Put your database host")
	password := flag.String("password", "root", "Put your database host")
	port := flag.Int("port", 0, "Put your database host")

	flag.Parse()

	fmt.Println(*host, *username, *password, *port)

}

// pada saat menjalankan perintah perlu diperhatikan spasi agar tidak terjadi hal dibawah ini

// benar
// go run flag.go -host=host -username=insun -password=insun123 -port=8080
// host insun insun123 8080

// salah
// go run flag.go -host=host -username = insun -password = insun123 -port = 8080
// host = root 0

//salah
// go run flag.go -host = host -username = insun -password = insun123 -port = 8080
// = root root 0