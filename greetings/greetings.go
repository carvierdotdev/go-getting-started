package greetings

import "fmt"

// Greet returns a greeting for the named person.
func Greet(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message

	/*
	The long way of doing it:
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
		return message
	*/
}