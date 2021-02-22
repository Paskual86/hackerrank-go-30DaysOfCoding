#Comandos

printLn

Uso:

Para el redondeo de numeros a enteros utilizar

math.Round()

Por otro lado para castear valores, se puede utilizar

- int32()

- float64()

Leer valores desde la linea de comandos:

reader := bufio.NewReaderSize(os.Stdin, 1024*1024)


Una funcion para leer lineas:

```go
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()

	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
```go

#GO Commands

https://golangbot.com/structs-instead-of-classes/

**Struct**
Structs are user-defined data types/structures.


Example:

```go
type SuperAnimals struct { 
    locomotion string
} 
```

**Interfaces**

In Go the names of the interfaces must end **"er"** because is a Go convention.

**Example**

```go
type Animaler interface {
	Eat()
	Move()
	Speak()
	Error()
}
```

**Functions**

Functions are generally the block of codes or statements in a program that gives the user ability to reuse the same code which ultimately use of memory, acts as a time saver and more importantly, provides better readability of the code.

*Sintax*

```go
func function_name(parameter-list)(return_type){
	// Function body
}
```

