package greetings

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
