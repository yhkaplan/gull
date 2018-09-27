package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Usage = "GitHub activity dashboard"
	app.Version = "0.0.1"

	dateFormat := "2006-01-02"
	today := time.Now().Format(dateFormat)
	lastMonth := time.Now().AddDate(0, -1, 0).Format(dateFormat)

	app.Commands = []cli.Command{
		{
			Name:  "activity",
			Usage: "Shows GitHub activities",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "from, f",
					Value: today,
					Usage: "From `date`",
				},
				cli.StringFlag{
					Name:  "to, t",
					Value: today,
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
					Value: lastMonth,
					Usage: "From `date`",
				},
				cli.StringFlag{
					Name:  "to, t",
					Value: today,
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
