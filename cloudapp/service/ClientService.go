package service

import (
       "github.com/cloudfoundry-community/go-cfclient"
       "fmt"
)

func GetClient() *cfclient.Client {

       c := &cfclient.Config{
              ApiAddress:   "https://api.13.74.251.125.xip.io",
              Username:     "moriar",
              Password:     "moriar",
              SkipSslValidation: true, // window is better without ssl validation
       }

       client, err := cfclient.NewClient(c)

       if err != nil {
              fmt.Printf("ClientService Error: %s", err)
       }

       return client
}
