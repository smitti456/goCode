package main

import (
	"fmt"
	"gocode/controllers"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("Hello!")
	fmt.Printf("hello, world\n")

	controllers.Hi()

	s := "Hello World"
	sr := strings.Split(s, " ") // create substring
	fmt.Println(strings.Contains(s, sr[0]))

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

}
