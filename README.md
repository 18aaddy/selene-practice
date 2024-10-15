![tests](https://github.com/18aaddy/selene-practice/actions/workflows/test.yml/badge.svg)
![linter](https://github.com/18aaddy/selene-practice/actions/workflows/cilint.yml/badge.svg)

# Introduction

Selene is a fast, open source, portable & secure light client for Ethereum written in Golang. We plan to ship Selene as the underlying software behind wallets that use light clients. We derived our inspiration from [Helios](https://github.com/a16z/helios) which is a light client written in Rust. The project is in active maintenance on the [dev](https://github.com/18aaddy/selene-practice/tree/dev) branch. 

# Architecture
![Selene Architecture](https://github.com/user-attachments/assets/db7eb9d7-5bc3-4911-a849-1b2d05239942)
## Data Flow

1. The Consensus Layer provides sync information and verified block headers to the RPC Server.
2. The RPC Server passes block information and verified headers to the Execution Layer.
3. The Execution Layer validates Merkle proofs based on the state root and requested data.
   
## Centralised RPC Server 
This server acts as an intermediary between the Consensus and Execution layers. It handles:

* Providing block headers of previous blocks from checkpoint to the latest block<br>
* Transmitting block gossip of block head<br>
* Passing verified block headers to the Execution Layer<br>

## Execution Layer 
The Execution Layer is responsible for processing transactions and maintaining the current state of the blockchain. It includes:

A `Validate Merkle Proofs` field that:

Takes state root and requested data as input
Outputs a boolean (true/false) indicating the validity of the Merkle proof

## Consensus Layer

The Consensus Layer is responsible for maintaining agreement on the state of the blockchain across the network. It includes:

* Getting weak subjectivity checkpoints
* Logic for determining **current and next sync committees**
* A **Syncing** process that:
   * Uses sync committee makeup to fetch previous block headers
   * Syncs for each sync committee period (~27 hours) up to the latest block
* A **verify bls sig** function that:
   * Takes `blsaggsig` and `blspubkey[]` as input
   * This function verifies a BLS aggregate signature. It accepts the aggregated signature (blsaggsig) and an array of public keys (blspubkey[]), returning a boolean value that indicates whether the signature is 
     valid.

# Installing
Yet to come.

# Usage
Yet to come.

# Testing
In progress.

# Warning
Selene is still experimental software. We hope to ship v0.1 by November 2024.

# Contributing
We openly welcome contributions to selene from the broader ethereum community. For details, refer [CONTRIBUTION GUIDELINES](https://github.com/18aaddy/selene-practice/blob/dev/CONTRIBUTING.md).
