# Apple Passkey Demo with Hanko Authentication API

Apple announced at the [WWDC21](https://developer.apple.com/videos/play/wwdc2021/10106/) that WebAuthn credentials will be available as “Passkeys” in the iCloud Keychain, as well as the availability of system-wide WebAuthn APIs on iOS 15, iPadOS 15, and macOS Monterey.

In their WWDC announcement video, Apple demonstrates the creation and seamless synchronization of Passkeys across devices. They even show that WebAuthn works with iOS Apps using the same Passkey.

This example application delivers the "Adopt WebAuthn to your server" part of their video.

It creates a webserver on port 3000 with a basic website and authentication using Passkeys (aka WebAuthn credentials).

You can find this demo in action here: https://apple-passkey.demo.hanko.io/

## Prerequisites

- Go 1.15
- Git
- a (free) account at Hanko's FIDO2/WebAuthn service ([get it here!](https://console.hanko.io/api/oidc?is_register=true))
- the API URL of your Hanko Authentication API
- an API Key ID/secret pair for your Hanko Authentication API

Please refer to the [Getting Started Guide](https://docs.hanko.io/gettingstarted) on how to obtain these values from the Hanko Console.

You need to configure your relying party as http://localhost. If you intend to use any other hostname or IP address you need to configure and use SSL (https)! WebAuthn is designed to work only in a secure context or on localhost.

## Running

1. Clone this repository and change into the folder:

```
git clone https://github.com/teamhanko/apple-wwdc21-webauthn-example

cd apple-wwdc21-webauthn-example
```

2. Configure the API URL, API Key ID and API secret by copying the `config/config.template.yaml` file to `config/config.yaml` and fill in the values:

   - Enter your API URL: `apiUrl: "YOUR_API_URL"`
   - Enter your API Key ID: `apiKeyId: "YOUR_API_KEY_ID"`
   - Enter your API secret: `apiSecret: "YOUR_API_SECRET"`

 and run:

```
go run .
```

The application will start on http://localhost:3000 and is good to go. Use Safari on **macOS Monterey** and **iOS 15** to try out the new Passkey syncing feature. Of course WebAuthn also works on Windows 10 and/or using Chrome or Firefox - just without the Keychain sync.

For more information on the Hanko Authentication API please see the [Hanko Docs](https://docs.hanko.io).