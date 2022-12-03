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
open htts://github.com
open https://google.com

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

# read url from clipboard
# 1. copy url. eg: https://github.com/search?q=fzdwx
# 2. this function does not do any verification, url and file path verification
# run `open`
open 

# open the specified url
open url https://www.bilibili.com/
```

## config (env)

```
GH_TOKEN=xxx
OPEN_LOG_FILE=$tmp$/fzdwx_open.log
``` 

## RodeMap

- [ ] read url from stdin. eg: `echo https://github.com/search?q=fzdwx | open`
- [ ] history
    - [x] history recode
    - [ ] history view
- [ ] custom alias

