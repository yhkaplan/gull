package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/yhkaplan/gull/date"
	"github.com/yhkaplan/gull/github"
)

func main() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Usage = "GitHub activity dashboard"
	app.Version = "0.0.1"

	dateFormat := "2006-01-02"
	today := time.Now().Format(dateFormat)
	lastWeek := time.Now().AddDate(0, 0, -7).Format(dateFormat)
	lastMonth := time.Now().AddDate(0, -1, 0).Format(dateFormat)

	app.Commands = []cli.Command{
		{
			Name:  "activity",
			Usage: "Shows GitHub activities",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "from, f",
					Value: lastWeek,
					Usage: "From `date`",
				},
				cli.StringFlag{
					Name:  "to, t",
					Value: today,
					Usage: "To `date`",
				},
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Get activities of specified `username`",
				},
			},
			Action: func(c *cli.Context) error {
				from, err := date.Parse(c.String("from"))
				if err != nil {
					return err
				}
				to, err := date.Parse(c.String("to"))
				if err != nil {
					return err
				}
				to = date.EndOfDay(to)

				fmt.Printf("Show activities: from %v, to %v\n", from, to)

				// TODO
				// Sample
				client, err := github.NewClient(c.String("user"), os.Getenv("GITHUB_TOKEN"))
				if err != nil {
					return err
				}

				events, err := client.GetEventsWithGrouping(context.Background(), from, to)
				if err != nil {
					return err
				}
				for _, event := range events {
					var (
						eventType, title, link string
					)

					title = *event.Type     // TODO
					eventType = *event.Type // TODO
					switch title {
					case "IssuesEvent":
						issuesEvent, err := github.GetIssuesEventFromRaw(event)
						if err != nil {
							return err
						}
						link = *issuesEvent.Issue.HTMLURL
					case "PullRequestEvent":
						pullRequestEvent, err := github.GetPullRequestEventFromRaw(event)
						if err != nil {
							return err
						}
						link = *pullRequestEvent.PullRequest.HTMLURL
					case "PullRequestReviewCommentEvent":
						pullRequestReviewCommentEvent, err := github.GetPullRequestReviewCommentEventFromRaw(event)
						if err != nil {
							return err
						}
						link = *pullRequestReviewCommentEvent.Comment.HTMLURL
					case "IssueCommentEvent":
						issueCommentEvent, err := github.GetIssueCommentEventFromRaw(event)
						if err != nil {
							return err
						}
						link = *issueCommentEvent.Comment.HTMLURL
					case "CommitCommentEvent":
						commitCommentEvent, err := github.GetCommitCommentEventFromRaw(event)
						if err != nil {
							return err
						}
						link = *commitCommentEvent.Comment.HTMLURL
					default:
						return errors.New("invalid event type")
					}

					fmt.Printf("- [%s](%s: %s)\n", eventType, link, title)
				}

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
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Get activities of specified `username`",
				},
			},
			Action: func(c *cli.Context) error {
				from, err := date.Parse(c.String("from"))
				if err != nil {
					return err
				}
				to, err := date.Parse(c.String("to"))
				if err != nil {
					return err
				}
				to = date.EndOfDay(to)

				fmt.Printf("Visualizes activities: from %v, to %v\n", from, to)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
