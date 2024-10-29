# Signchain Go Client

A Golang client library for interacting with Signchain's secure transaction and wallet management services. This library allows developers to integrate blockchain transaction signing, wallet management, and self-hosted vault services within backend applications.

[![Go Reference](https://pkg.go.dev/badge/github.com/grexie/signchain-go.svg)](https://pkg.go.dev/github.com/grexie/signchain-go)  
[Signchain Documentation](https://signchain.net/docs) | [API Reference](https://signchain.net/docs/api-reference)

---

## Overview

The Signchain Go client provides backend functionality for securely managing Ethereum-compatible wallets and transactions. Designed for server-side use, it allows you to interact with your Signchain vault, create and manage wallets, and sign transactions with ease.

## Installation

Install the package with `go get`:

```bash
go get github.com/grexie/signchain-go/v2
```

## Usage

1. **Initialize Client**  
   Configure and instantiate the Signchain client.

   ```go
   package main

   import (
       "context"
       "log"
       "github.com/grexie/signchain-go/v2"
   )

   func main() {
       options := signchain.NewClientOptions().
           SetAPIKey("YOUR_API_KEY").
           SetVaultID("YOUR_VAULT_ID")

       signchainClient, err := signchain.NewClient(options)
       if err != nil {
           log.Fatalf("Failed to create Signchain client: %v", err)
       }

       // Example usage
       vaultStatus, err := signchainClient.VaultStatus(context.Background())
       if err != nil {
           log.Fatalf("Error getting vault status: %v", err)
       }

       log.Printf("Vault online: %v", vaultStatus.Online())
   }
   ```

2. **Create a Wallet**  
   Use the client to create and manage wallets securely.

   ```go
   wallet, err := signchainClient.CreateWallet(context.Background(), signchain.NewCreateWalletOptions().SetName("My New Wallet"))
   if err != nil {
       log.Fatalf("Error creating wallet: %v", err)
   }
   log.Printf("Wallet created with address: %s", wallet.Address().Hex())
   ```

3. **Sign a Transaction**  
   Sign transactions with options for blockchain compatibility.

   ```go
   signOptions := signchain.NewSignOptions().
       SetChain(client.ChainEthereum).
       SetContract("0x...").
       SetSender("0x...").
       SetArgs([]interface{}{"0x..."}).
       SetABI(map[string]interface{}{}) // Specify the ABI

   result, err := signchainClient.Sign(context.Background(), signOptions)
   if err != nil {
       log.Fatalf("Error signing transaction: %v", err)
   }
   log.Printf("Transaction signed with hash: %s", result.SubmissionHash())
   ```

## API Methods

- `VaultStatus(ctx context.Context) (VaultStatus, error)`: Retrieve the current status of your Signchain vault.
- `CreateWallet(ctx context.Context, options CreateWalletOptions) (Wallet, error)`: Create a new wallet in the vault.
- `GetWallet(ctx context.Context, address common.Address) (Wallet, error)`: Retrieve details of a specific wallet.
- `ListWallets(ctx context.Context, options ListWalletsOptions) (ListWalletsResult, error)`: List all wallets associated with the vault.
- `UpdateWallet(ctx context.Context, address common.Address, options UpdateWalletOptions) (Wallet, error)`: Update the details of a specific wallet.
- `ExpireWallet(ctx context.Context, address common.Address) (Wallet, error)`: Mark a wallet as expired, removing its private key.
- `UnexpireWallet(ctx context.Context, address common.Address) (Wallet, error)`: Reactivate an expired wallet and restore access.

## Authentication & Security

Signchain clients require an API Key and Vault ID to authenticate and access the API. These should be stored securely and never exposed client-side.

## Documentation

For more detailed instructions, see the [Signchain Documentation](https://signchain.net/docs) and [API Reference](https://signchain.net/docs/api-reference).

## License

MIT License.
