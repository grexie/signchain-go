package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

var ErrNotImplemented = errors.New("not implemented")

type Client interface {
	URL() string
	APIKey() string
	VaultID() string

	Sign(ctx context.Context, options SignOptions) (SignResult, error)
	
	CreateWallet(ctx context.Context, options CreateWalletOptions) (Wallet, error)
	GetWallet(ctx context.Context, address common.Address) (Wallet, error)
	ListWallets(ctx context.Context, options ListWalletsOptions) (ListWalletsResult, error)
	UpdateWallet(ctx context.Context, address common.Address, options UpdateWalletOptions) (Wallet, error)
	ExpireWallet(ctx context.Context, address common.Address) (Wallet, error)
	UnexpireWallet(ctx context.Context, address common.Address) (Wallet, error)

	VaultStatus(ctx context.Context) (VaultStatus, error)
}

type client struct {
	URL_     string
	APIKey_  string
	VaultID_ string
}

// APIKey implements Client.
func (c *client) APIKey() string {
	return c.APIKey_
}

// URL implements Client.
func (c *client) URL() string {
	return c.URL_
}

// VaultID implements Client.
func (c *client) VaultID() string {
	return c.VaultID_
}

var _ Client = &client{}

type ClientOptions interface {
	URL() *string
	SetURL(url string) ClientOptions
	UnsetURL() ClientOptions

	APIKey() string
	SetAPIKey(apiKey string) ClientOptions

	VaultID() string
	SetVaultID(vaultId string) ClientOptions
}

type clientOptions struct {
	URL_     *string
	APIKey_  string
	VaultID_ string
}

// APIKey implements ClientOptions.
func (c *clientOptions) APIKey() string {
	return c.APIKey_
}

// SetAPIKey implements ClientOptions.
func (c *clientOptions) SetAPIKey(apiKey string) ClientOptions {
	c.APIKey_ = apiKey
	return c
}

// SetURL implements ClientOptions.
func (c *clientOptions) SetURL(url string) ClientOptions {
	c.URL_ = &url
	return c
}

// SetVaultID implements ClientOptions.
func (c *clientOptions) SetVaultID(vaultId string) ClientOptions {
	c.VaultID_ = vaultId
	return c
}

// URL implements ClientOptions.
func (c *clientOptions) URL() *string {
	return c.URL_
}

// UnsetURL implements ClientOptions.
func (c *clientOptions) UnsetURL() ClientOptions {
	c.URL_ = nil
	return c
}

// VaultID implements ClientOptions.
func (c *clientOptions) VaultID() string {
	return c.VaultID_
}

var _ ClientOptions = &clientOptions{}

func NewClientOptions() ClientOptions {
	return &clientOptions{}
}

func NewClient(options ClientOptions) (Client, error) {
	c := client{}
	if options.URL() == nil {
		c.URL_ = "https://signchain.net"
	} else {
		c.URL_ = *options.URL()
	}
	c.APIKey_ = options.APIKey()
	c.VaultID_ = options.VaultID()

	return &c, nil
}

type apiResponse[T any] struct {
	Success bool `json:"success"`
	Data T `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}

func (c *client) Request(ctx context.Context, method string, path string, body any, response any) (error) {
	var r *http.Request
	if body == nil {
		if req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", c.URL_, path), nil); err != nil {
			return err
		} else {
			r = req
		}
	} else {
		if b, err := json.Marshal(body); err != nil {
			return err
		} else if req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", c.URL_, path), bytes.NewBuffer(b)); err != nil {
			return err
		} else {
			r = req
			r.Header.Set("Content-Type", "application/json")
		}
	}

	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey_))

	if res, err := http.DefaultClient.Do(r); err != nil {
		return err
	} else if res.StatusCode >= 400 {
		var response apiResponse[any]
		if b, err := io.ReadAll(res.Body); err != nil {
			return err
		} else if err := json.Unmarshal(b, &response); err != nil {
			return err
		} else {
			return errors.New(*response.Error)
		}
	} else if res.Header.Get("Content-Type") == "application/json" {
		if b, err := io.ReadAll(res.Body); err != nil {
			return err
		} else if err := json.Unmarshal(b, &response); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return nil
	}
}
