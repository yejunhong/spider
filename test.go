package main
import (
  "fmt"
)
func main() {
    var next chan int = make(chan int, 5)
    for index := 0; index < 10; index++ {
        next <- 1
        go func(v int){
            fmt.Println(v)
            <-next
        }(index)
    }

    for {

    }
}