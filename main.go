package main

import (
	"fmt"
	"github.com/BigGroupProgramming/discord-clone/app"
)

func main() {
	s, err := app.Init()
	if err != nil {
		panic(fmt.Sprintf("error starting server | %v", err))
	}

	router, err := app.Router(s)
	if err != nil {
		panic(fmt.Sprintf("error starting router | %v", err))
	}

	if err := router.Run(":3001"); err != nil {
		panic("error running server")
	}

	fmt.Println("hello")
}
