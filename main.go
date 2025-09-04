package main

import (
	"log"
	"nhatruong/firstGoBackend/src/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
