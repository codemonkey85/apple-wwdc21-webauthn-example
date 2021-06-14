# Simple Webauthn Example (usernameless)

The example application demonstrates FIDO2/WebAuthn usernameless flow (discoverable credentials) and how to utilize the
Hanko Authentication API in a web application written in Go.

## Prerequisites

- Go 1.15
- an API URL of your Hanko Authentication instance and an API Key ID/API secret pair in order to make authorized calls
  of the Hanko API
    - see [Documentation](https://docs.hanko.io/gettingstarted) on how to get these

## Running

### From Source

Configure the API URL, API Key ID and API secret for making authorized calls to the Hanko Authentication API by editing
the `config/config.yaml` file:

1. Edit your API URL: `apiUrl: "YOUR_API_URL"`
2. Edit your API Key ID: `apiKeyId: "YOUR_API_KEY_ID"`
3. Edit your API secret: `apiSecret: "YOUR_API_SECRET"`

Clone the repository and run:

```
go run .
```

For more information on the Hanko Authentication API see the [Hanko Docs](https://docs.hanko.io). 