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

# read url from clipboard
# 1. copy url. eg: https://github.com/search?q=fzdwx
# 2. this function does not do any verification, url and file path verification
# run `open`
open 
```

## config (env)

```
GH_TOKEN=xxx
OPEN_LOG_FILE=$tmp$/fzdwx_open.log
``` 

## RodeMap

- [ ] read url from stdin. eg: `https://github.com/search?q=fzdwx | open`
- [ ] history
    - [x] history recode
    - [ ] history view
- [ ] custom alias

