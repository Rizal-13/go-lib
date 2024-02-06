package main

import (
	"bufio"
	"io"
	"os"
)

// membuat file
func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

// membaca file
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil{
		return "", err
	}
	defer file.Close()

	reader:= bufio.NewReader(file)

	var message string
	for{
		line, _, err := reader.ReadLine()
		message += string(line)
		if err == io.EOF{
			break
		}
	}
	return message, nil

}

func addToFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil{
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {

	// createNewFile("index.html", "<h1>Hello World</h1>" )

	// result, _ := readFile("index.html")
	// fmt.Println(result)

	addToFile("index.html", "\n<br>")

}
