# urlzap

Your own URL manager, statically generated ⚡️

- Keep your (shortned or not) URLs with you
- No need to run a server (although supported)
- Can be used with Github Pages

## How does it work?

You might be asking yourself: how is this done without a server? Well, the answer lies on
`<meta http-equiv="refresh" />`. It basically works as HTTP 301 (Redirect) status code,
but done in the client-side. [There is a bit more explanation on w3 website](https://www.w3.org/TR/WCAG20-TECHS/H76.html).

`urlzap` creates `index.html` files which make use of meta refresh tags. It is not perfect
as a HTTP 301, but is quite close. It is already used by other static web site generators,
such as Hugo.

## Install

### MacOS

Use `brew` to install it

```
brew tap brunoluiz/tap
brew install urlzap
```

### Linux and Windows

[Check the releases section](https://github.com/brunoluiz/urlzap/releases) for more information details 

### go get

Install using `GO111MODULES=off go get github.com/brunoluiz/urlzap/cmd/urlzap` to get the latest version. This will place it in your `$GOPATH`, enabling it to be used anywhere in the system.

**⚠️ Reminder**: the command above download the contents of master, which might not be the stable version. [Check the releases](https://github.com/brunoluiz/urlzap/releases) and get a specific tag for stable versions.

## Usage:

First set-up your URLs in `./config.yml`. Each key in the map will map to `/{key}` routes,
redirecting to `{value}`. In the example below, `https://yourwebsite/google` will 
redirect to Google and `https://yourwebsite/tools/github` to Github.

```yaml
urls:
  google: https://google.com
  tools:
    github: https://github.com
path: './output' # default is './'
```

To generate the static files, call `urlzap generate`.

## Usage with Github Pages:

1. Enable Github Pages and set-up the branch where your static HTML files will be located.
More details at [Github Pages guide](https://pages.github.com/)
1. Update your `config.yml` with new URLs
1. Commit and push to `main`
1. Checkout to your Github Pages branch (usually `gh-pages`) and merge `main` into it
1. Run `urlzap generate`
1. Commit and push the results

If you want to "automate" this process, try the below bash script, after you've updated
your `config.yml`

```sh
#!/bin/bash

# adds, commit and push your changes
git add config.yml
git commit -m 'chore: update config.yml'
git push -u origin main

# make gh-pages branch to be the same as main
git checkout gh-pages
git reset --hard origin/main

# generate files
urlzap generate

# add, commit and push generated files
git add --all
git commit -m 'chore: update HTML files'
git push -u origin gh-pages --force
```

## To-do:

1. Create a landing page
1. Add Github Actions guide
1. Add a better explanation about the tool
