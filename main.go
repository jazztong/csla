package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"log"
)

func main() {
	log.Println("print log")
	fmt.Println("hello world")
	fmt.Println(stringutil.Reverse("hello"))
}
