```go
func main() {
    var a [10]int
    for i := 0; i <= 10; i++ { //Error: Index out of bounds
        a[i] = i
    }
    fmt.Println(a)
}
```