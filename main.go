package main

import (
	"fmt"
	"log"

	"github.com/slyang-git/go-tutorial/greetings"
	"github.com/slyang-git/go-tutorial/world"
	"rsc.io/quote"
)

// 基本数据类型

//
func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	fmt.Println(world.Country)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	greeting, err := greetings.Hello("xx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	fmt.Println(messages)
}
