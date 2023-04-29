package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files working")

	content := "this needs to go in a file - LearnCodeonline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("lenght is:", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data insdie the file is \n", databyte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
