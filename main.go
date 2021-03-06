package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
	"github.com/yhkaplan/gull/date"
	"github.com/yhkaplan/gull/github"
	"github.com/yhkaplan/gull/view"
)

const (
	defaultBaseURL = "https://api.github.com/"

	envGitHubToken = "GITHUB_TOKEN"
	envGitHubAPI   = "GITHUB_API"
)

func main() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Usage = "GitHub activity dashboard"
	app.Version = "0.1.0"

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
				cli.BoolFlag{
					Name:  "eventType",
					Usage: "Show event type along with output",
				},
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Get activities of specified `username`",
				},
				cli.BoolFlag{
					Name:  "comment, c",
					Usage: "Show events that are PR, issue, and other kinds of comments",
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

				activities, err := parseEvents(c, from, to)
				if err != nil {
					return err
				} else if activities == nil {
					return errors.New("Activities is nil")
				}

				for _, a := range activities {
					_, isCommentType := github.CommentEventTypes[a.EventType]
					if !c.Bool("comment") && isCommentType {
						continue
					}

					if c.Bool("eventType") {
						fmt.Printf("- [%s](%s): %s\n", a.Title, a.Link, a.EventType)
					} else {
						fmt.Printf("- [%s](%s)\n", a.Title, a.Link)
					}
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

				activities, err := parseEvents(c, from, to)
				if err != nil {
					return err
				} else if activities == nil {
					return errors.New("Activities is nil")
				}

				if err := view.New(activities).Run(); err != nil {
					return err
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseEvents(c *cli.Context, from, to time.Time) ([]github.GitHubActivity, error) {
	baseURLStr := defaultBaseURL
	if urlStr := os.Getenv(envGitHubAPI); urlStr != "" {
		baseURLStr = urlStr
	}
	client, err := github.NewClient(c.String("user"), os.Getenv(envGitHubToken), baseURLStr)
	if err != nil {
		return nil, err
	}

	events, err := client.GetEventsWithGrouping(context.Background(), from, to)
	if err != nil {
		return nil, err
	}

	activities := make([]github.GitHubActivity, 0)

	for _, event := range events {
		var eventType, title, link string

		eventType = *event.Type
		switch eventType {
		case "IssuesEvent":
			issuesEvent, err := github.GetIssuesEventFromRaw(event)
			if err != nil {
				return nil, err
			}
			link = *issuesEvent.Issue.HTMLURL
			title = *issuesEvent.Issue.Title
		case "PullRequestEvent":
			pullRequestEvent, err := github.GetPullRequestEventFromRaw(event)
			if err != nil {
				return nil, err
			}
			link = *pullRequestEvent.PullRequest.HTMLURL
			title = *pullRequestEvent.PullRequest.Title
		case "PullRequestReviewCommentEvent":
			pullRequestReviewCommentEvent, err := github.GetPullRequestReviewCommentEventFromRaw(event)
			if err != nil {
				return nil, err
			}
			link = *pullRequestReviewCommentEvent.Comment.HTMLURL
			title = *pullRequestReviewCommentEvent.PullRequest.Title + " (comment)"
		case "IssueCommentEvent":
			issueCommentEvent, err := github.GetIssueCommentEventFromRaw(event)
			if err != nil {
				return nil, err
			}
			link = *issueCommentEvent.Comment.HTMLURL
			title = *issueCommentEvent.Issue.Title + " (comment)"
		case "CommitCommentEvent":
			commitCommentEvent, err := github.GetCommitCommentEventFromRaw(event)
			if err != nil {
				return nil, err
			}
			link = *commitCommentEvent.Comment.HTMLURL
			title = *commitCommentEvent.Repo.HTMLURL + " (comment)"
		default:
			return nil, errors.New("invalid event type")
		}

		title = strings.Replace(title, " (comment)", "", 1)

		activity := github.GitHubActivity{
			Link:      link,
			Title:     title,
			EventType: eventType,
		}
		activities = append(activities, activity)
	}

	return activities, nil
}
