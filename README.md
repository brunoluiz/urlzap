<h1 align="center">
  URLZap
</h1>

<p align="center">
  Your own static URL shortener ⚡️
</p>

Static site generators, published on Github Pages, are quite popular nowadays. But what about a
static URL shortener (to not say generator), which allows you to redirect URLs based on static files?

Usually, developers end-up setting up a server with redirects for this (not statically). That is
where URLZap comes in.  It generates URLs using files and HTML wizardry, allowing users to host
their own URL redirects into Github Pages.

- 🔗 Similar to static website generators, but for URLs
- 🔒 Keep your (shortened or not) URLs with you
- 🌎 Can be used with Github Pages
- ☕️ No need to run a server or set-up HTTP 301 redirects

Example project: [`brunoluiz/_`](https://github.com/brunoluiz/_)

## ☕️ How does it work?

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

These files can be uploaded to Github Pages for example, not requiring any server.  On
[`brunoluiz/_`](https://github.com/brunoluiz/_) you can see an example `config.yml` and checkout
the output in `gh-pages` branch 😉

## 📀 Install

### Linux and Windows

[Check the releases section](https://github.com/brunoluiz/urlzap/releases) for more information details

### MacOS

Use `brew` to install it

```
brew tap brunoluiz/tap
brew install urlzap
```

## ⚙️ Usage

### Generate files locally

Using the previous YAML example:

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

### Automatic deploy using Github Actions

Most likely you will end-up using Github Pages together with this tool. If so, perhaps the best
way to use it and reap its benefits is through Github Actions. Head to
[`brunoluiz/urlzap-github-action`](https://github.com/marketplace/actions/urlzap) for more details
on how to install it, covering generation & deployment.

> ⚠️ You might need to manually enable Github Pages in your repository! More details at
> [Github Pages guide](https://pages.github.com/)

### Manual deploy to Github Pages or similars

If Github Actions are not for you, try the following manual process instead.

1. Enable Github Pages and set-up the branch where your static HTML files will be located.
More details at [Github Pages guide](https://pages.github.com/)
1. Set-up your `config.yml`
1. Commit and push to `main`
1. Checkout to your Github Pages branch (usually `gh-pages`) and run `git reset --hard origin/main`
(this will reset the HEAD to `master`)
1. Run `urlzap generate`
1. Commit and push

The following script follows what is described on the steps above:

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
