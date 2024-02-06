package main

import (
	"encoding/csv"
	"os"
)


func main(){

writer := csv.NewWriter(os.Stdout)

_ = writer.Write([]string{"insun", "Crypto", "47"})
_ = writer.Write([]string{"master", "good", "taichi"})
_ = writer.Write([]string{"rizal", "hartono", "47"})

writer.Flush()

}