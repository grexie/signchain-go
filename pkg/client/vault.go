package client

import (
	"context"
	"fmt"
	"time"
)

type VaultStatus interface {
	Timestamp() time.Time
	Online() bool
	VaultKeys() int64
	Wallets() int64
	Version() string
}

type vaultStatus struct {
	Timestamp_ time.Time `json:"timestamp"`
	Online_    bool      `json:"online"`
	VaultKeys_ int64     `json:"vaultKeys"`
	Wallets_   int64     `json:"wallets"`
	Version_   string    `json:"version"`
}

// Online implements VaultStatus.
func (v *vaultStatus) Online() bool {
	return v.Online_
}

// Timestamp implements VaultStatus.
func (v *vaultStatus) Timestamp() time.Time {
	return v.Timestamp_
}

// VaultKeys implements VaultStatus.
func (v *vaultStatus) VaultKeys() int64 {
	return v.VaultKeys_
}

// Version implements VaultStatus.
func (v *vaultStatus) Version() string {
	return v.Version_
}

// Wallets implements VaultStatus.
func (v *vaultStatus) Wallets() int64 {
	return v.Wallets_
}

var _ VaultStatus = &vaultStatus{}

func (c *client) VaultStatus(ctx context.Context) (VaultStatus, error) {
	var res apiResponse[vaultStatus]

	if err := c.Request(ctx, "GET", fmt.Sprintf("/api/v1/vaults/%s", c.VaultID_), nil, &res); err != nil {
		return nil, err
	} else {
		return &res.Data, nil
	}
}