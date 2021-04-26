package di

import (
	"fmt"
	"io"
)

const helloFormat = "Hello, %s"

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, helloFormat, name)
	// return helloPrefix + name
}
