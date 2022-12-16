package user

import (
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/auth"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/env"
	"github.com/fzdwx/x/strx"
	"github.com/gookit/slog"
	"github.com/spf13/cobra"
	"os"
)

var (
	ghToken  = ""
	username = ""
	initFlag = false
)

func checkInit() {
	if initFlag {
		return
	}

	initFlag = true

	ghToken = env.OrWithFunc(readTokenFromEnv(), func() string {
		host, _ := auth.DefaultHost()
		token, _ := auth.TokenForHost(host)

		slog.Infof("github cli token is %s", token)
		return token
	})

	if strx.IsBlank(ghToken) {
		slog.Fatal("can not get user token")
	}

}

// Name get user Github name
func Name() string {
	checkInit()

	if strx.IsBlank(username) {
		username = func() string {
			client, err := gh.RESTClient(nil)
			cobra.CheckErr(err)
			response := struct{ Login string }{}
			err = client.Get("user", &response)
			cobra.CheckErr(err)
			username = response.Login

			return username
		}()
	}

	slog.Infof("username is %s", username)
	return username
}

// Token user github token
func Token() string {
	checkInit()

	return ghToken
}

// readTokenFromEnv read github token from env
func readTokenFromEnv() string {
	token := os.Getenv(cons.EnvGhToken)

	if strx.IsBlank(token) {
		slog.Debug("env token is blank")
	} else {
		slog.Debugf("env token is %s", token)
	}

	return token
}
