# Open

open url in browser.

![image](https://user-images.githubusercontent.com/65269574/213114721-8305b43e-2e8f-4a96-ad96-22a35ed2bc1a.gif)

```bash
$ open alias ls | fzf --preview 'open alias info {}' --bind 'enter:execute(open alias run {})'
```

## Installation

```shell
go install github.com/fzdwx/open@latest 
```

## about auth

1. use gh cli: `gh auth login`
2. or `export GH_TOKEN=xxx`

## use case

available commands:

```shell
## Commands
alias [subcommand]                  | Manage custom aliases                                     
completion                          | Generate the autocompletion script for the specified shell
gh [search keyword] | [subcommand]  | open github repo in browser                                  
help [command]                      | Help about any command                                    
history                             | Show open history                                         
log                                 | Show open log                                             
url                                 | open the specified url
```

### open url

open the specified url

```shell
$ open https://www.bilibili.com/
$ open url https://www.bilibili.com/
$ open /root/
```

### open alias

User can add alias for url.

```shell
## add
$ open alias add https://fzdwx.github.io/ --name blog

## delete
$ open alias remove blog

## list
$ open alias list
```

### open gh

open github repo in browser

```shell
## Examples
$ open gh        -> open https://github.com
$ open gh .      -> open current dir(like `open gh repo`) in github
$ open gh fzdwx  -> open https://github.com/search?q=fzdwx
$ open gh -l java sky -> open https://github.com/search?q=sky&l=java
$ open gh fzdwx -u -> https://github.com/search?q=fzdwx&type=users

## Commands
profile p | open your github profile in browser. eg: https://github.com/fzdwx   
repo .    | open github repository in browser. eg: https://github.com/fzdwx/open

## Flags
-f, --closed      | search issues,pr status is closed             
-c, --commits     | set search type is commits                    
-d, --debug       | show log in console                           
-s, --discussions | set search type is discussions                
-h, --help        | help for gh                                   
-i, --issues      | set search type is issues                     
-l, --lang string | search programming languages. eg: go,java,rust
-m, --marketplace | set search type is marketplace                
-o, --open        | search issues,pr status is open               
-g, --packages    | set search type is packages                   
-p, --pr          | set search type is pull requests              
-r, --repo        | set search type is repositories               
-t, --topics      | set search type is topics                     
-u, --users       | set search type is users                      
-w, --wikis       | set search type is wikis
```

## config (env)

```
GH_TOKEN=xxx
OPEN_PREVIEW=bat(fallback to cat)
``` 

## RodeMap

- [x] history
    - [x] history record
    - [x] history view
- [x] custom alias
    - [x] add alias
    - [x] delete alias
    - [x] list alias

