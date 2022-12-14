# Open

open url in browser.

```shell
go install github.com/fzdwx/open@latest 
```

## about auth

1. use gh cli: `gh auth login`
2. or `export GH_TOKEN=xxx`

## use case

```shell
# open url
open https://github.com
open https://google.com

# open the specified url
open url https://www.bilibili.com/

# open dir
open .
open /home/fzdwx

# open github.com
open gh

# open https://github.com/search?q=fzdwx
open gh fzdwx

# open https://github.com/search?q=golang
open gh golang

# open https://github.com/search?q=sky&l=java
open gh -l java sky

# open https://github.com/search?q=fzdwx
open gh -l java sky

# open your github profile
open gh profile
open gh p

# open current project git remote url  in browser. https://github.com/fzdwx/open
open gh repo

# open https://github.com/fzdwx/open
open gh repo fzdwx/open
# open # https://github.com/{username}/open`,
open gh repo open

# preview logfile
open log

```

## config (env)

```
GH_TOKEN=xxx
OPEN_PREVIEW=bat(fallback to cat)
``` 

## RodeMap

- [ ] read url from stdin. eg: `echo https://github.com/search?q=fzdwx | open`
- [ ] history
    - [x] history record
    - [ ] history view
- [ ] custom alias

