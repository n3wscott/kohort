package main

import (
	"fmt"
	"github.com/n3wscott/kohort/cmd/hello/resources"
	"github.com/n3wscott/kohort/pkg/kohort"
)

func main() {
	kohort.RunMaybe(resources.Resource)

	fmt.Println("Hello, World!")

}
