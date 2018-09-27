package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Usage = "GitHub activity dashboard"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "activity",
			Usage: "Shows GitHub activities",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "from, f",
					Usage: "From `date`",
				},
				cli.StringFlag{
					Name:  "to, t",
					Usage: "To `date`",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("Show activities: from %v, to %v\n", c.String("from"), c.String("to"))
				return nil
			},
		},
		{
			Name:  "dashboard",
			Usage: "Visualizes GitHub activities",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "from, f",
					Usage: "From `date`",
				},
				cli.StringFlag{
					Name:  "to, t",
					Usage: "To `date`",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("Visualizes activities: from %v, to %v\n", c.String("from"), c.String("to"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
