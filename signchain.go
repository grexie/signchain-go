package signchain

import "github.com/grexie/signchain-go/v2/pkg/client"

type Client = client.Client
var NewClient = client.NewClient
type ClientOptions = client.ClientOptions
var NewClientOptions = client.NewClientOptions

type Wallet = client.Wallet

type SignOptions = client.SignOptions
var NewSignOptions = client.NewSignOptions
type SignResult = client.SignResult

type CreateWalletOptions = client.CreateWalletOptions
var NewCreateWalletOptions = client.NewCreateWalletOptions

type ListWalletOptions = client.ListWalletsOptions
var NewListWalletsOptions = client.NewListWalletsOptions
type ListWalletsResult = client.ListWalletsResult

type UpdateWalletOptions = client.UpdateWalletOptions
var NewUpdateWalletOptions = client.NewUpdateWalletOptions

type VaultStatus = client.VaultStatus


