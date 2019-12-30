package main

import (
	"fmt"
	"github.com/concourse/dutyfree/server"
)

func main() {

	s := server.Server{
		Port:              9090,
		PublicPath:        "../../../web/public",
		ResourceTypesPath: "../resource-types",
	}

	s.Start()
	fmt.Println("Dutyfree server started")
	<-s.Exited
}
