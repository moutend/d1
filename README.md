# D1

An alternative to Cloudflare Wrangler D1.

## Install

```console
go install github.com/moutend/d1/cmd/d1@latest
```

## Usage

Make sure to set the following environment variables:

- `CLOUDFLARE_API_TOKEN`: An API token created from the [Cloudflare dashboard](https://dash.cloudflare.com/profile/api-tokens/). Navigate to My Profile > API Tokens.
- `CLOUDFLARE_ACCOUNT_ID`: Your account ID displayed on the Cloudflare dashboard.
- `CLOUDFLARE_D1_LOCATION`: The location hint, such as "apac".

## License

MIT
