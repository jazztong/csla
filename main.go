package main

import (
	"fmt"
	"log"

	"github.com/golang/example/stringutil"
)

func main() {
	log.Println("print log")
	fmt.Println("hello world")
	fmt.Println(stringutil.Reverse("hello"))
}
