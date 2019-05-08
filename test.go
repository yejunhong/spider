package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "golang.org/x/image/webp"
)

func main() {
    var buf bytes.Buffer
    var width, height int
    var data []byte
    var err error

    // Load file data
    if data, err = ioutil.ReadFile("/Volumes/book/cover/C006/ace9b4a7a2a91683aea181806a6ac4e3.jpg"); err != nil {
        fmt.Println(err)
    }

    // GetInfo
    if width, height, _, err = webp.GetInfo(data); err != nil {
        fmt.Println(err)
    }
    fmt.Printf("width = %d, height = %d\n", width, height)

    // GetMetadata
    if metadata, err := webp.GetMetadata(data, "ICCP"); err != nil {
        fmt.Printf("Metadata: err = %v\n", err)
    } else {
        fmt.Printf("Metadata: %s\n", string(metadata))
    }

    // Decode webp
    m, err := webp.Decode(bytes.NewReader(data))
    if err != nil {
        fmt.Println(err)
    }

    // Encode lossless webp
    if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
        fmt.Println(err)
    }
    if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
        fmt.Println(err)
    }
}