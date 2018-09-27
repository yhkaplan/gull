package github

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gullEventTypes = []string{
	"IssuesEvent",
	"PullRequestEvent",
	"PullRequestReviewCommentEvent",
	"IssueCommentEvent",
	"CommitCommentEvent",
}

type Client struct {
	owner string
	*github.Client
}

func NewClient(owner, token string) (*Client, error) {
	if len(owner) == 0 {
		return nil, errors.New("missing GitHub Repository owner")
	}
	if len(token) == 0 {
		return nil, errors.New("missing GitHub API token")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &Client{
		owner:  owner,
		Client: client,
	}, nil
}

func (c *Client) GetEventsWithGrouping(ctx context.Context) ([]*github.Event, error) {
	events, _, err := c.Activity.ListEventsPerformedByUser(ctx, c.owner, true, nil)
	if err != nil {
		return nil, err
	}

	dst := make([]*github.Event, 0, len(events))
	for _, event := range events {
		for _, gullEventType := range gullEventTypes {
			if *event.Type == gullEventType {
				// TODO
				// filter from-to time
				dst = append(dst, event)
			}
		}
	}

	return events, nil
}
