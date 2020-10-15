# urlzap

Your own URL manager, statically generated ⚡️

- Keep your (shortned or not) URLs with you
- No need to run a server (although supported)
- Can be used with Github Page

## Usage:

First set-up your URLs in `./config.yml`. Each key in the map will map to `/{key}` routes,
redirecting to `{value}`. In the example below, `https://yourwebsite/google` will 
redirect to Google and `https://yourwebsite/tools/github` to Github.

```
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

If you don't have any cache configured (eg: Cloudflare), it shouldn't a minute to get
the updates.
