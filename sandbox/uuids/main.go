package main

import (
	"fmt"
	"log"

	"github.com/falbanese9484/rum"
)

func main() {
	uuid, err := rum.NewUUID()
	if err != nil {
		log.Fatalf("We fucked up!")
	}
	fmt.Printf("%s\n", uuid.String())
}
