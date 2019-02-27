package main
 
import (
    "fmt"
    "sync"
)

func main(){

    var waitGroup sync.WaitGroup
    waitGroup.Add(1)
    for i := 0; i <= 10; i++ {
        go func(){
            defer waitGroup.Done()
            fmt.Println(i)
        }()
    }
    waitGroup.Wait()
}
