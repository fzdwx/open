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
open gh -s fzdwx

# open your github profile
open gh profile
open gh p
```

## config (env)

```
GH_TOKEN=xxx
OPEN_LOG_FILE=$tmp$/fzdwx_open.log
``` 

## RodeMap

- [ ] history
- [ ] custom alias

