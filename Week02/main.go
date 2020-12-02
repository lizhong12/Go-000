package main

import (
	"log"
	"service"
)

func main()  {
	err := service.Service()
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
}
