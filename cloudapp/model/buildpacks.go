package model

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"fmt"
)

type App struct {
	Name string
}

func GetListBuildpacks(client *cfclient.Client) []cfclient.Buildpack {

	if client == nil {
		return nil
	}

	buildpacks, err := client.ListBuildpacks()

	if err != nil {
		fmt.Printf("Buildpacks Error: %s", err)
	}

	return buildpacks
}



