# urlzap

Your own static URL generator ⚡️

- Similar to static website generators, but for URLs
- Keep your (shortned or not) URLs with you
- Can be used with Github Pages
- No need to run a server or set-up HTTP 301 redirects

## How does it work?

You might be asking yourself: how is this done without a server? Well, the answer lies on
`<meta http-equiv="refresh" />`. It works as HTTP 301 (Redirect) status code, but it is done
in the client-side. [There is a bit more explanation on w3c website](https://www.w3.org/TR/WCAG20-TECHS/H76.html).

Based on a `config.yml` containing the desired path and URL, `urlzap` will create `index.html`
files which make use of meta refresh tags. It is not perfect as a HTTP 301, but is quite close.
A similar strategy is used by other static website generators, such as Hugo.

An example would be:

```yaml
path: './links' # default is './'
urls:
  google: https://google.com
  tools:
    github: https://github.com
```

Each key in the map will map to `{.path param}/{key}` routes, redirecting to `{value}`.
This would generate the following:

```
- links/
  - google/
    - index.html (contains redirect)
  - tools/
    - github/
      - index.html (contains redirect)
```

These files can be uploaded to Github Pages for example, not requiring any server.

## Install

### Linux and Windows

[Check the releases section](https://github.com/brunoluiz/urlzap/releases) for more information details

### MacOS

Use `brew` to install it

```
brew tap brunoluiz/tap
brew install urlzap
```

## Usage

### Generate files

Using the previous example:

```yaml
path: './links' # default is './'
urls:
  google: https://google.com
  tools:
    github: https://github.com
```

- `urls`: desired URL map, following the `{key}:{redirect URL}` pattern
- `path`: output path

To generate the static files, run `urlzap generate`.

### Usage with Github Pages or similars

1. Enable Github Pages and set-up the branch where your static HTML files will be located.
More details at [Github Pages guide](https://pages.github.com/)
1. Set-up your `config.yml`
1. Commit and push to `main`
1. Checkout to your Github Pages branch (usually `gh-pages`) and run `git reset --hard origin/main`
(this will reset the HEAD to `master`)
1. Run `urlzap generate`
1. Commit and push

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
1. Create a Github Actions
1. Add a better explanation about the tool
