package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "example"
	app.Usage = "Usage example"

	app.Commands = []*cli.Command{
		{
			Name:  "greet",
			Usage: "greeting",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name, n",
					Value:    "",
					Usage:    "who are you?",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("Hello %s\n", c.String("name"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
