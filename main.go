package main
import "fmt"

func main() {
    var x int
    
    fmt.Print("Enter number x: ")
    fmt.Scan(&x)

    for i := 1; i <= x; i++ {
        for j := 0; j < i; j++ {
            fmt.Print("*")
        }
        fmt.Println() 
    }

    for i := x - 1; i >= 1; i-- {
        for j := 0; j < i; j++ {
            fmt.Print("*")
        }
        fmt.Println() 
    }
}