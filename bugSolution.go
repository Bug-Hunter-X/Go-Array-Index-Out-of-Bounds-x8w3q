```go
func main() {
    var a [10]int
    for i := 0; i < 10; i++ { //Corrected loop condition
        a[i] = i
    }
    fmt.Println(a)
}
```