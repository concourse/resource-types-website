package main

import (
	"fmt"
	"github.com/concourse/dutyfree/server"
)

func main() {

	s := server.Server {
		Port: 9090,
	}

	s.Start()
	fmt.Println("Dutyfree server started")
	<- s.Exited
}
