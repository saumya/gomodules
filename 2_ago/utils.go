//
package gomodule

import "fmt"


func NameIs(name string) string {
    //message := fmt.Sprintf("My name is, %v!", name)
    //return message
    return Hello(name)
}

func AgainThis(name string) string {
    message := fmt.Sprintf("Again this is %v!", name)
    return message
}