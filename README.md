# Nolus Core
<div align="center">
<br /><p align="center"><img alt="nolus-core" src="docs/nolus-core-logo.svg" width="100"/></p><br />


[![Go Report Card](https://goreportcard.com/badge/github.com/Nolus-Protocol/nolus-core)](https://goreportcard.com/report/github.com/Nolus-Protocol/nolus-core)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/Nolus-Protocol/nolus-core/blob/main/LICENSE)
[![Lint](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/lint.yaml/badge.svg?branch=main)](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/lint.yaml)
[![Test](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/test.yaml)
[![Test Fuzz](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/test-fuzz.yaml/badge.svg?branch=main)](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/test-fuzz.yaml)
[![Test cosmos-sdk](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/test-cosmos.yaml/badge.svg?branch=main)](https://github.com/Nolus-Protocol/nolus-core/actions/workflows/test-cosmos.yaml)

**Nolus** is a blockchain built using Cosmos SDK and Tendermint.
</div>

## Prerequisites

Install [golang](https://golang.org/), [tomlq](https://tomlq.readthedocs.io/en/latest/installation.html) and [jq](https://stedolan.github.io/jq/).

## Get started

### Build, configure and run a single-node locally deployed Nolus chain

#### Build

  ```sh
  make install
  ```

#### Init, set up the DEX parameters and run

First generate the mnemonic you will use for Hermes:

```sh
nolusd keys mnemonic
```

Then recover osmosis key and use the [UI](https://faucet.osmosis.zone/#/) to get some uosmo:

```sh
osmosisd keys add hermes_key --recover
```

Init and start:

```sh
./scripts/init-local-network.sh --reserve-tokens <reserve_account_init_tokens> --hermes-mnemonic <the_mnemonic_generated_by_the_previous_steps> --dex-network-addr-rpc <dex_network_addr_host_part_rpc> --dex-network-addr-grpc <dex_network_addr_host_part_grpc>
```

Set up the DEX parameter: [Set up the DEX parameters manually](#set-up-the-dex-parameters-manually)

The `make install` command will compile and locally install nolusd on your machine. `init-local-network.sh` generates a node setup, including setting the dex parameter (run `init-local-network.sh --help` for more configuration options). For more details check the [scripts README](./scripts/README.md)

*Notes:

* Make sure "nolus-money-market" repo is checked out as a sibling to this repo.

* The Osmosis binary is required: [Follow the steps](https://gitlab-nomo.credissimo.net/nomo/wiki/-/blob/main/hermes.md#download-the-latest-osmosis-binary).

* Before running the "./scripts/init-local-network.sh" again, make sure the nolusd and hermes processes are killed.

* The "hermes" and "nolusd" logs are stored in ~/hermes and ~/.nolus respectively.

### Run already configured single-node

```sh
nolusd start --home "networks/nolus/local-validator-1"
```

## Install, configure and run a local Hermes relayer manually

Follow the steps [here](https://gitlab-nomo.credissimo.net/nomo/wiki/-/blob/main/hermes.md#install-and-configure-hermes). Write down the connection and channel identifiers at Nolus and Osmosis for further usage.

## Set up the DEX parameters manually

The goal is to let smart contracts know the details of the connectivity to Osmosis. Herebelow is a sample request.
This should be done via sudo gov proposal:

```sh
nolusd tx gov submit-proposal sudo-contract nolus1wn625s4jcmvk0szpl85rj5azkfc6suyvf75q6vrddscjdphtve8s5gg42f '{"setup_dex": {"connection_id": "connection-0", "transfer_channel": {"local_endpoint": "channel-0", "remote_endpoint": "channel-1499"}}}' --title "Set up the DEX parameter" --description "Thе proposal aims to set the DEX parameters in the Leaser contract" --deposit 10000000unls --fees 900unls --gas auto --gas-adjustment 1.1 --from wallet
```

*Notes:
"channel-1499" should be replaced, so you can get the actual channel ID of the remote endpoint with:

```sh
nolusd q ibc channel connections connection-0 --output json | jq '.channels[0].counterparty.channel_id' | tr -d '"'
```

Check the transaction has passed:

```sh
nolusd q wasm contract-state smart nolus1wn625s4jcmvk0szpl85rj5azkfc6suyvf75q6vrddscjdphtve8s5gg42f '{"config":{}}'
```

## Build statically linked binary

By default, `make build` generates a dynamically linked binary. In case someone would like to reproduce the way the binary is built in the pipeline then the command to achieve it locally is:

```sh
docker run --rm -it -v "$(pwd)":/code public.ecr.aws/nolus/builder:<replace_with_the latest_tag> make build -C /code
```

## Upgrade wasmvm

* Update the Go modules
* Update the wasmvm version in the builder Dockerfile at .github/images/builder.Dockerfile
* Increment the IMAGE_TAG and use the same version in the build-binary step in .github/workflows/build.yaml

## Run a full node with docker

https://github.com/Nolus-Protocol/nolus-networks
