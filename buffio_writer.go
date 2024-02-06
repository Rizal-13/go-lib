package main

import (
	"bufio"
	"os"
)


func main(){

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello world\n")
	writer.WriteString("apa kabar\n")
	writer.Flush()
}