package github

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

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
