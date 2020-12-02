package main

import (
	"log"
	"studygolang/demo/service"
)

func main()  {
	err := service.Service()
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
}