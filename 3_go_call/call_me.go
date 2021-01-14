package main
import (
    "fmt"

    "saumya.learning/gomodule"
)

func main() {
    message := gomodule.Hello("Gladys")
    fmt.Println(message)

    fmt.Println( gomodule.NameIs("Saumya") )
    fmt.Println( gomodule.AgainThis("AgainHappy") )
}