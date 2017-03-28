package main

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "trello-fetcher"
	app.Usage = ""
	app.Action = current

	app.Run(os.Args)
}

func current(c *cli.Context) error {
	appKey := os.Getenv("TRELLO_DEVELOPER_PUBLIC_KEY")
	token := os.Getenv("TRELLO_MEMBER_TOKEN")
	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
		return err
	}

	user, err := trello.Member("me")
	if err != nil {
		log.Fatal(err)
		return err
	}

	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, b := range boards {
		if b.Closed {
			continue
		}

		lists, err := b.Lists()
		if err != nil {
			log.Fatal(err)
			return err
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

	return nil
}
