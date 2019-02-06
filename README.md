# seamless

## Running

```sh
make run
```

Then open http://localhost:8080

## The JWT Cloud Function

You can use this example cloud function to generate JWT tokens needed to use the TurboWallet API.

```sh
gcloud functions deploy CreateJWT --runtime go111 --trigger-http --set-env-vars SECRET=[YOUR_SECRET_HERE]
```
