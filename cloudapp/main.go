package main

import (
	"github.com/go-martini/martini"
	//"cloudapp/service"
	//"cloudapp/model"
	//"cloudapp/helpers"
)

func main() {

	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello!"
	})

	m.Run()

}
