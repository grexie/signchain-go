package client

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
)

type Wallet interface {
	ID() string
	Account() string
	Vault() string
	Name() string
	Address() common.Address
	Created() time.Time
	Updated() time.Time
	Expires() *time.Time
}

type wallet struct {
	ID_      string         `json:"id"`
	Account_ string         `json:"account"`
	Vault_   string         `json:"vault"`
	Name_    string         `json:"name"`
	Address_ common.Address `json:"address"`
	Created_ time.Time      `json:"created"`
	Updated_ time.Time      `json:"updated"`
	Expires_ *time.Time     `json:"expires,omitempty"`
}

// Account implements Wallet.
func (w *wallet) Account() string {
	return w.Account_
}

// Address implements Wallet.
func (w *wallet) Address() common.Address {
	return w.Address_
}

// Created implements Wallet.
func (w *wallet) Created() time.Time {
	return w.Created_
}

// Expires implements Wallet.
func (w *wallet) Expires() *time.Time {
	return w.Expires_
}

// ID implements Wallet.
func (w *wallet) ID() string {
	return w.ID_
}

// Name implements Wallet.
func (w *wallet) Name() string {
	return w.Name_
}

// Updated implements Wallet.
func (w *wallet) Updated() time.Time {
	return w.Updated_
}

// Vault implements Wallet.
func (w *wallet) Vault() string {
	return w.Vault_
}

var _ Wallet = &wallet{}

type CreateWalletOptions interface {
	Name() string
	SetName(name string) CreateWalletOptions
}

type createWalletOptions struct {
	Name_ string `json:"name"`
}

// Name implements CreateWalletOptions.
func (c *createWalletOptions) Name() string {
	return c.Name_
}

// SetName implements CreateWalletOptions.
func (c *createWalletOptions) SetName(name string) CreateWalletOptions {
	c.Name_ = name
	return c
}

var _ CreateWalletOptions = &createWalletOptions{}

func NewCreateWalletOptions() CreateWalletOptions {
	return &createWalletOptions{}
}

type ListWalletsOptions interface {
	Offset() *int64
	SetOffset(offset int64) ListWalletsOptions

	Count() *int64
	SetCount(count int64) ListWalletsOptions
}

type listWalletsOptions struct {
	Offset_ *int64 `url:"offset,omitempty"`
	Count_  *int64 `url:"count,omitempty"`
}

// Count implements ListWalletOptions.
func (l *listWalletsOptions) Count() *int64 {
	return l.Count_
}

// Offset implements ListWalletOptions.
func (l *listWalletsOptions) Offset() *int64 {
	return l.Offset_
}

// SetCount implements ListWalletOptions.
func (l *listWalletsOptions) SetCount(count int64) ListWalletsOptions {
	l.Count_ = &count
	return l
}

// SetOffset implements ListWalletOptions.
func (l *listWalletsOptions) SetOffset(offset int64) ListWalletsOptions {
	l.Offset_ = &offset
	return l
}

var _ ListWalletsOptions = &listWalletsOptions{}

func NewListWalletsOptions() ListWalletsOptions {
	return &listWalletsOptions{}
}

type ListWalletsResult interface {
	Count() int64
	Page() []Wallet
}

type listWalletsResult struct {
	Count_ int64     `json:"count"`
	Page_  []*wallet `json:"page"`
}

// Count implements ListWalletsResult.
func (l *listWalletsResult) Count() int64 {
	return l.Count_
}

// Page implements ListWalletsResult.
func (l *listWalletsResult) Page() []Wallet {
	out := make([]Wallet, len(l.Page_))
	for i, w := range l.Page_ {
		out[i] = w
	}
	return out
}

var _ ListWalletsResult = &listWalletsResult{}

type UpdateWalletOptions interface {
	Name() string
	SetName(name string) UpdateWalletOptions
}

type updateWalletOptions struct {
	Name_ string `json:"name"`
}

// Name implements CreateWalletOptions.
func (c *updateWalletOptions) Name() string {
	return c.Name_
}

// SetName implements CreateWalletOptions.
func (c *updateWalletOptions) SetName(name string) UpdateWalletOptions {
	c.Name_ = name
	return c
}

var _ UpdateWalletOptions = &updateWalletOptions{}

func NewUpdateWalletOptions() UpdateWalletOptions {
	return &updateWalletOptions{}
}

func (c *client) CreateWallet(ctx context.Context, options CreateWalletOptions) (Wallet, error) {
	var res apiResponse[wallet]

	if err := c.Request(ctx, "POST", fmt.Sprintf("/api/v1/vaults/%s/wallets", c.VaultID_), options, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}

func (c *client) GetWallet(ctx context.Context, address common.Address) (Wallet, error) {
	var res apiResponse[wallet]

	if err := c.Request(ctx, "GET", fmt.Sprintf("/api/v1/vaults/%s/wallets/%s", c.VaultID_, address), nil, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}

func (c *client) ListWallets(ctx context.Context, options ListWalletsOptions) (ListWalletsResult, error) {
	var res apiResponse[listWalletsResult]

	if search, err := query.Values(options); err != nil {
		return nil, err
	} else if err := c.Request(ctx, "GET", fmt.Sprintf("/api/v1/vaults/%s/wallets?%s", c.VaultID_, search.Encode()), nil, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}

func (c *client) UpdateWallet(ctx context.Context, address common.Address, options UpdateWalletOptions) (Wallet, error) {
	var res apiResponse[wallet]

	if err := c.Request(ctx, "PUT", fmt.Sprintf("/api/v1/vaults/%s/wallets/%s", c.VaultID_, address), options, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}

func (c *client) ExpireWallet(ctx context.Context, address common.Address) (Wallet, error) {
	var res apiResponse[wallet]

	if err := c.Request(ctx, "DELETE", fmt.Sprintf("/api/v1/vaults/%s/wallets/%s", c.VaultID_, address), nil, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}

func (c *client) UnexpireWallet(ctx context.Context, address common.Address) (Wallet, error) {
	var res apiResponse[wallet]

	if err := c.Request(ctx, "POST", fmt.Sprintf("/api/v1/vaults/%s/wallets/%s", c.VaultID_, address), nil, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}
