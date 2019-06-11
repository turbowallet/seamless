# seamless

## Deploying contract

There is an example contract in the [contracts/](contracts/) folder which this example uses. To deploy it:

* First install [web3](https://github.com/gochain-io/web3).
* Set your private key that has enough GO for gas to deploy with (0.1 GO is plenty for a bunch of deploys) `export WEB3_PRIVATE_KEY=0xAAA`

```sh
cd contracts
web3 contract build increment.sol
web3 contract deploy Incrementor.bin
```

## Running

```sh
make run
```

Then open http://localhost:8080/public

## The JWT Cloud Function

You can use this example cloud function to generate JWT tokens needed to use the TurboWallet API.

```sh
gcloud functions deploy CreateJWT --runtime go111 --trigger-http --set-env-vars SECRET=[YOUR_SECRET_HERE]
```
