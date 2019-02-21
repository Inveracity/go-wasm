package main

import (
    "fmt"
    "syscall/js"
    )

func main() {
    fmt.Println("it's working")
    doc := js.Global().Get("document")
    body := doc.Call("getElementById", "paragraph")
    body.Set("innerHTML", "good stuff")
}
