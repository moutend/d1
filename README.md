# D1

An alternative to Cloudflare Wrangler D1 command line tool.

## Prerequisites

Please install the latest version of the [Go toolchain](https://go.dev/doc/install).

## Installation

```console
go install github.com/moutend/d1/cmd/d1@latest
```

## Usage

Make sure to set the following environment variables:

- `CLOUDFLARE_API_TOKEN`: An API token created from the [Cloudflare dashboard](https://dash.cloudflare.com/profile/api-tokens/). Navigate to My Profile > API Tokens.
- `CLOUDFLARE_ACCOUNT_ID`: Your account ID displayed on the Cloudflare dashboard.
- `CLOUDFLARE_D1_LOCATION`: The location hint, such as "apac".

## Motivation

The official Wrangler command occasionally returns the error message "fetch failed." I believe this could be a network issue rather than a problem with the Wrangler implementation. However, since the Wrangler command is under frequent development, there is a possibility that it could be an implementation problem. This is why I implement another d1 command line tool.

## License

MIT
