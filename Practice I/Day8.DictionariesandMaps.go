package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile("./../Lotes de Pruebas/Day8/Test1.txt")

	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	n := int(nTemp)

	input := loadPhoneNumber(n, reader)

	queryPhoneNumber(n, reader, input)
}

func loadPhoneNumber(quantity int, reader *bufio.Reader) map[string]string {
	var i int
	var result map[string]string

	result = make(map[string]string)
	for i = 1; i <= quantity; i++ {
		values := strings.Fields(readLine(reader))
		result[strings.TrimLeft(strings.TrimRight(string(values[0]), " "), " ")] = strings.TrimLeft(strings.TrimRight(string(values[1]), " "), " ")
	}

	return result
}

func queryPhoneNumber(quantity int, reader *bufio.Reader, input map[string]string) {
	var i int

	var input_search string

	for i = 1; i <= quantity; i++ {
		input_search = strings.TrimLeft(strings.TrimRight(readLine(reader), " "), " ")
		value, found := input[input_search]
		if found {
			fmt.Println(input_search + "=" + value)
		} else {
			fmt.Println("Not found")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return string(str)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(pathfile string) {
	dat, err := ioutil.ReadFile(pathfile)
	checkError(err)
	fmt.Print(string(dat))

	f, err := os.Open(pathfile)
	checkError(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}
