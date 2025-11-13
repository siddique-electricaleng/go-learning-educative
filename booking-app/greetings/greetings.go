package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hola, %s you are who again?", name)
}
