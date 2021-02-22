package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	for i := 0; i < int(n); i++ {
		strValue := readLine(reader)
		EvenAndOddFunction(strValue)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, error := reader.ReadLine()

	if error == io.EOF {
		return ""
	}

	return string(str)
}

func EvenAndOddFunction(str string) {
	var even string
	var odd string
	for i := 0; i < len(str); i++ {
		if (i % 2) == 0 {
			even += string(str[i])

		} else {
			odd += string(str[i])
		}
	}
	//fmt.Println(even + " " + odd)
	fmt.Printf("%s %s\n", even, odd)
}
