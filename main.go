package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"log"
	"os"
)

func main() {
	appKey := os.Getenv("TRELLO_DEVELOPER_PUBLIC_KEY")
	token := os.Getenv("TRELLO_MEMBER_TOKEN")
	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	user, err := trello.Member("me")
	if err != nil {
		log.Fatal(err)
	}

	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range boards {
		if b.Closed {
			continue
		}

		lists, err := b.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, l := range lists {
			section := false

			if l.Name == "Do" {
				cards, _ := l.Cards()
				for _, c := range cards {
					if section == false {
						fmt.Println(b.Name)
						section = true
					}

					fmt.Println("- " + c.Name)
				}
			}
		}
	}
}
