package main

import (
	"fmt"
	"log"
	"github.com/slyang-git/GoTutorial/greetings"
	"github.com/slyang-git/GoTutorial/world"
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

	greeting, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}
