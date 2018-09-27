package github

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	perPage = 100
)

var EventTypes = []string{
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

// GetEventsWithGrouping lists the events received by a user
// and filter the events with specified event types.
func (c *Client) GetEventsWithGrouping(ctx context.Context, from, to time.Time) ([]*github.Event, error) {
	opt := &github.ListOptions{PerPage: perPage}
	events, _, err := c.Activity.ListEventsPerformedByUser(ctx, c.owner, true, opt)
	if err != nil {
		return nil, err
	}

	dst := make([]*github.Event, 0, len(events))
	for _, event := range events {
		for _, gullEventType := range EventTypes {
			if *event.Type == gullEventType {
				if event.CreatedAt.After(from) && event.CreatedAt.Before(to) {
					dst = append(dst, event)
				}
			}
		}
	}

	return dst, nil
}

// GetIssuesEventFromRaw parse `RawPayload` in github.Event`
// Ref: https://godoc.org/github.com/google/go-github/github#Event
// Ref: https://developer.github.com/v3/activity/events/types/#issuesevent
func GetIssuesEventFromRaw(event *github.Event) (*github.IssuesEvent, error) {
	dst := &github.IssuesEvent{}
	err := json.Unmarshal(event.GetRawPayload(), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// GetPullRequestEventFromRaw parse `RawPayload` in github.Event`
// Ref: https://godoc.org/github.com/google/go-github/github#Event
// Ref: https://developer.github.com/v3/activity/events/types/#pullrequestevent
func GetPullRequestEventFromRaw(event *github.Event) (*github.PullRequestEvent, error) {
	dst := &github.PullRequestEvent{}
	err := json.Unmarshal(event.GetRawPayload(), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// GetPullRequestReviewCommentEventFromRaw parse `RawPayload` in github.Event`
// Ref: https://godoc.org/github.com/google/go-github/github#Event
// Ref: https://developer.github.com/v3/activity/events/types/#pullrequestreviewcommentevent
func GetPullRequestReviewCommentEventFromRaw(event *github.Event) (*github.PullRequestReviewCommentEvent, error) {
	dst := &github.PullRequestReviewCommentEvent{}
	err := json.Unmarshal(event.GetRawPayload(), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// GetIssueCommentEventFromRaw parse `RawPayload` in github.Event`
// Ref: https://godoc.org/github.com/google/go-github/github#Event
// Ref: https://developer.github.com/v3/activity/events/types/#issuecommentevent
func GetIssueCommentEventFromRaw(event *github.Event) (*github.IssueCommentEvent, error) {
	dst := &github.IssueCommentEvent{}
	err := json.Unmarshal(event.GetRawPayload(), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// GetCommitCommentEventFromRaw parse `RawPayload` in github.Event`
// Ref: https://godoc.org/github.com/google/go-github/github#Event
// Ref: https://developer.github.com/v3/activity/events/types/#commitcommentevent
func GetCommitCommentEventFromRaw(event *github.Event) (*github.CommitCommentEvent, error) {
	dst := &github.CommitCommentEvent{}
	err := json.Unmarshal(event.GetRawPayload(), dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}
