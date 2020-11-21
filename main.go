package main

import (
	"fmt"
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
	fmt.Printf(greetings.Hello("yangshuanglong"))
}
