[![Twitter Follow](https://img.shields.io/twitter/follow/DeBankDeFi?style=social)](https://twitter.com/DeBankDeFi)
[![Discord](https://img.shields.io/discord/984015101017346058?color=%235865F2&label=Discord&logo=discord&logoColor=%23fff)](https://discord.com/invite/KYuj8DE)
# DeBankChain - DeBank layer 2 solution

The DeBankChain is the Layer 2 scaling solution powered by [bedrock version](https://community.optimism.io/docs/developers/bedrock/) of Optimism OP Stack.


## Features

Besides the [differentiators of bedrock](https://community.optimism.io/docs/developers/bedrock/differences/), DeBankChain is the solution that we aim to provide the better optimistic rollup solution.

### Node optimization

- The consensus logic has been modified reducing the L1 data storage gas cost by 100 to 400 times
- An Abstract Accounts(AA)-like account system has been implemented natively, giving users a web2-like experience while maintaining 100% compatibility with existing EVM standards
- The new account system supports transactions signed with second-layer private keys specifically generated, reducing the use of L1 private keys in low-value scenarios and enhancing the security of users' large assets.

### JSONRPC-API enhancement

- EVM Pre Execution
- EVM Transaction Tracing
- Multicall

## Public RPC & APIs

| Parameter | Value |
| -------------- | ------------------- |
| Network Name   | DeBank Testnet      |
| Chain ID       | 2021398             |
| Explorer       | https://explorer.testnet.debank.com |
| RPC Endpoint   | https://rpc.testnet.debank.com |

## Documentation

If you want to build on top of DeBankChain, refer to the [DeBankChain Doc](https://github.com/DeBankDeFi/DeBankChain/tree/main/docs/build).

## Community

To get help from other developers, discuss ideas, and stay up-to-date on what's happening, become a part of our community on Telegram. Join our [official Telegram Channel](https://t.me/DeBankEN).

## Directory Structure

<pre>
~ DeBankChain ~~
├── <a href="./packages">packages</a>
│   └── <a href="./packages/contracts-bedrock">contracts-bedrock</a>: Bedrock smart contracts.
├── <a href="./op-bindings">op-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./op-batcher">op-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./op-e2e">op-e2e</a>: End-to-End testing of all bedrock components in Go
├── <a href="./op-node">op-node</a>: rollup consensus-layer client.
├── <a href="./op-proposer">op-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
└── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade
</pre>

## License

All files within this repository are licensed under the [MIT License](https://github.com/DeBankDeFi/DeBankChain/blob/main/LICENSE) unless stated otherwise.
