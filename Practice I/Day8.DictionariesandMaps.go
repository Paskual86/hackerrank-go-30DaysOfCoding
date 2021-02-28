package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	n := int(nTemp)

	input := loadPhoneNumber(n, reader)

	queryPhoneNumber(n, reader, input)
	/*var input_search string

	input_search = strings.TrimLeft(strings.TrimRight(readLine(reader), " "), " ")

	for ok := true; ok; ok = (input_search != "") {
		value, found := input[input_search]
		if found {
			fmt.Println(input_search, "=", value)
		} else {
			fmt.Println("Not found")
		}
		input_search = strings.TrimLeft(strings.TrimRight(readLine(reader), " "), " ")
	}*/
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
	var result []string

	for i = 1; i <= quantity; i++ {
		input_search = strings.TrimLeft(strings.TrimRight(readLine(reader), " "), " ")
		value, found := input[input_search]
		if found {
			result = append(result, input_search+"="+value)
		} else {
			result = append(result, "Not found")
		}
	}

	for i = 0; i < len(result); i++ {
		fmt.Println(result[i])
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
