package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"log"
)

func main() {
	appKey := "application-key"
	token := "token"
	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(trello)
}
