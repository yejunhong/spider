package main

import (
	"fmt"
    "github.com/gin-gonic/gin"
)

func main() {

    router := gin.Default()
    router.Static("/img", "/data/book")

    fmt.Println("开启0.0.0.0:", 7555)
    router.Run(":7555") // listen and serve on 0.0.0.0:8080
}
