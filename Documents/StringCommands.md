```go
func main() {
    value := "cat"

    // Take length of string with len.
    length := len(value)
    fmt.Println(length)

    // Loop over the string with len.
    for i := 0; i < len(value); i++ {
        fmt.Println(string(value[i]))
    }
}
```