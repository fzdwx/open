package gh

import (
	"fmt"
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/x/strx"
	"github.com/spf13/cobra"
)

var (
	gh = &cobra.Command{
		Use:   "gh [search keyword] | [subcommand]",
		Short: "open github in browser",
		Example: `$ open gh        -> open https://github.com
$ open gh fzdwx  -> open https://github.com/search?q=fzdwx
$ open gh -l java sky -> open https://github.com/search?q=sky&l=java
$ open gh fzdwx -u -> https://github.com/search?q=fzdwx&type=users`,
		Run: func(cmd *cobra.Command, args []string) {
			url := cons.GithubUrl

			if len(args) < 1 {
				cobra.CheckErr(browser.Open(url))
				return
			}

			url = fmt.Sprintf("%s/search?q=%s", url, args[0])

			searchType := getType()
			if strx.IsNotBlank(searchType) {
				url += "&type=" + searchType
			}

			if statusClosed {
				url += "&state=closed"
			}
			if statusOpen {
				url += "&state=open"
			}

			if strx.IsNotBlank(languages) {
				url += "&l=" + languages
			}

			cobra.CheckErr(browser.Open(url))
		},
	}

	typeRepoFlag        bool
	typeCommitsFlag     bool
	typeIssuesFlag      bool
	typePrFlag          bool
	typeDiscussionsFlag bool
	typePackagesFlag    bool
	typeMarketplaceFlag bool
	typeTopicsFlag      bool
	typeWikisFlag       bool
	typeUsersFlag       bool

	statusClosed bool
	statusOpen   bool

	languages string
)

func Command() *cobra.Command {
	gh.AddCommand(profile)
	gh.AddCommand(repo)

	gh.PersistentFlags().BoolVarP(&typeRepoFlag, "repo", "r", typeRepoFlag, "set search type is repositories")
	gh.PersistentFlags().BoolVarP(&typeIssuesFlag, "issues", "i", typeIssuesFlag, "set search type is issues")
	gh.PersistentFlags().BoolVarP(&typePrFlag, "pr", "p", typePrFlag, "set search type is pull requests")
	gh.PersistentFlags().BoolVarP(&typeDiscussionsFlag, "discussions", "s", typeDiscussionsFlag, "set search type is discussions")
	gh.PersistentFlags().BoolVarP(&typeUsersFlag, "users", "u", typeUsersFlag, "set search type is users")
	gh.PersistentFlags().BoolVarP(&typeCommitsFlag, "commits", "c", typeCommitsFlag, "set search type is commits")
	gh.PersistentFlags().BoolVarP(&typePackagesFlag, "packages", "g", typePackagesFlag, "set search type is packages")
	gh.PersistentFlags().BoolVarP(&typeWikisFlag, "wikis", "w", typeWikisFlag, "set search type is wikis")
	gh.PersistentFlags().BoolVarP(&typeTopicsFlag, "topics", "t", typeTopicsFlag, "set search type is topics")
	gh.PersistentFlags().BoolVarP(&typeMarketplaceFlag, "marketplace", "m", typeMarketplaceFlag, "set search type is marketplace")

	gh.MarkFlagsMutuallyExclusive("repo", "commits", "issues", "pr", "discussions", "packages", "marketplace", "topics", "wikis", "users")

	gh.PersistentFlags().BoolVarP(&statusOpen, "open", "o", typeWikisFlag, "search issues,pr status is open")
	gh.PersistentFlags().BoolVarP(&statusClosed, "closed", "f", typeUsersFlag, "search issues,pr status is closed")

	gh.PersistentFlags().StringVarP(&languages, "lang", "l", "", "search programming languages. eg: go,java,rust")

	return gh
}

func getType() string {
	var searchType string
	switch {
	case typeRepoFlag:
		searchType = "repositories"
	case typeCommitsFlag:
		searchType = "commits"
	case typeIssuesFlag:
		searchType = "issues"
	case typePrFlag:
		searchType = "pullrequests"
	case typeDiscussionsFlag:
		searchType = "discussions"
	case typePackagesFlag:
		searchType = "packages"
	case typeMarketplaceFlag:
		searchType = "marketplace"
	case typeTopicsFlag:
		searchType = "topics"
	case typeWikisFlag:
		searchType = "wikis"
	case typeUsersFlag:
		searchType = "users"
	}

	return searchType

}
