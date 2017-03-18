package main

import (
	"github.com/go-martini/martini"
	"time"
	"fmt"
	"os"
	"html/template"
	"cloudapp/service"
	"cloudapp/helpers"
	"cloudapp/model"
)

var currentRoot, _ = os.Getwd()

type Welcome struct {
	Date   string
}

const welcomeTemplate = `Hello!!! :)
				The date is: {{.Date}}.`

func main() {

	m := martini.Classic()
	fmt.Println(currentRoot)

	//stOp := martini.StaticOptions{Prefix: currentRoot, IndexFile:"main.html", Fallback: "main.go"}
	//m.Use(martini.Static(currentRoot, stOp))

	c := service.GetClient()

	m.Group("/buildpacks", func(r martini.Router) {
		listBuildpacks := model.GetListBuildpacks(c)

		r.Get("", func() string {
			return helpers.GetConvertToJson(listBuildpacks)
		})

		r.Get("/sort/(?P<name>[a-zA-Z]+)", func(params martini.Params) string {
			switch params["name"] {
				case "byName":
					helpers.GetSortListByName(listBuildpacks)
					return helpers.GetConvertToJson(listBuildpacks)
				case "shuffle":
					helpers.Shuffle(listBuildpacks)
					return helpers.GetConvertToJson(listBuildpacks)
			}

			return "Sort type doesn't exist"
		})
	})

	m.Get("/", func() string {
		date := time.Now()
		strTime := date.Format("15:04:05 2006-01-02")

		t := template.New("Welcome template")
		t, err := t.Parse(welcomeTemplate)
		err = t.Execute(os.Stdout, Welcome{strTime})

		if err != nil {
			fmt.Printf("welcome template Error: %s", err)
		}

		return "Hello! :) The date is:" + strTime
	})


	m.Run()

}