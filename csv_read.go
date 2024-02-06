package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)


func main(){

	dataCsv := "insun, Crypto, 47 \n" + "master, good, taichi \n" + "rizal, hartono, 47"

	readCsv := csv.NewReader(strings.NewReader(dataCsv))

	for {
		record, err := readCsv.Read()
		if err == io.EOF{
			break
		} 
		fmt.Println(record)
	} 

}