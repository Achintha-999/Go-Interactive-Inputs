package main
import "fmt"
func main2() {

	var name string

var age uint8

var height float32

print("Enter your name: ")
fmt.Scan(&name)

print("Enter your age: ")
fmt.Scan(&age)

print("Enter your height in meters: ")
fmt.Scan(&height)


fmt.Printf("Hello, %s! You are %d years old and %.2f meters tall.\n", name, age, height)

}