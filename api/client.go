package api

import (
	"context"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/auth"
	"github.com/google/go-github/v46/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func init() {
	host, _ := auth.DefaultHost()

	token, _ = auth.TokenForHost(host)
}

func fetchUsername() string {
	client, err := gh.RESTClient(nil)
	cobra.CheckErr(err)

	response := struct{ Login string }{}
	err = client.Get("user", &response)
	cobra.CheckErr(err)

	return response.Login
}

func Token() string {
	return token
}

func Get(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func UserName() string {
	if userName == "" {
		userName = fetchUsername()
	}

	return userName
}

var (
	userName = ""
	token    = ""
)
