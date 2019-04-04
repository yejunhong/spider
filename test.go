package main
 
import (
    "fmt"
    "github.com/satori/go.uuid"
    "reflect"
)

func main(){
    var a = uuid.NewV4().String()
    fmt.Println(reflect.TypeOf(a))
}