package sayhello

import "fmt"

func Sayhello(name string) string {
	message := fmt.Sprintf("Hi, %v", name)
	return message
}
