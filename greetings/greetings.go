package greetings

import "fmt"

// Hello returns a greeting for the named person.

func Hello(name string) string {
	// Return a greeting tht embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // := declares and initialize a variable in one line
	return message
}
