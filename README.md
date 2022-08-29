# Open

open url in browser.

```shell
go install github.com/fzdwx/open@latest 
```

## gh

open github

```shell
# open https://github.com/
open gh 

# open your github profile. https://github.com/{username}
# must use gh auth.
open gh  -p

# open https://github.com/fzdwx
open gh  fzdwx

# open https://github.com/fzdwx/open
open gh  fzdwx/open

# open https://github.com/search?q=just
open gh just -s

# open https://github.com/fzdwx?tab=stars&q=qwe
open gh qwe --star
```