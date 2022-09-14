package user

import (
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/auth"
	"github.com/fzdwx/open/pkg/cons"
	"github.com/fzdwx/open/pkg/env"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/slog"
	"github.com/spf13/cobra"
	"os"
)

var (
	ghToken  = ""
	username = ""
)

func init() {
	ghToken = env.OrWithFunc(readTokenFromEnv(), func() string {
		host, _ := auth.DefaultHost()
		token, _ := auth.TokenForHost(host)
		return token
	})

	if strutil.IsBlank(ghToken) {
		err := errorx.New("can not get user token")
		slog.ErrorT(err)
		cobra.CheckErr(err)
	}

	slog.Infof("user token %s", ghToken)
}

// Name get user Github name
func Name() string {
	if username == "" {
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

	slog.Infof("username %s", username)

	return username
}

// Token user github token
func Token() string {
	return ghToken
}

// readTokenFromEnv read github token from env
func readTokenFromEnv() string {
	token := os.Getenv(cons.EnvGhToken)

	if token == "" {
		slog.Debug("env token is blank")
	} else {
		slog.Debugf("env token %s", token)
	}

	return token
}
