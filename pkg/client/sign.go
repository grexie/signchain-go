package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Chain string

const (
	ChainLocal      Chain = "local"
	ChainEthereum   Chain = "ethereum"
	ChainSepolia    Chain = "sepolia"
	ChainBSC        Chain = "bsc"
	ChainBSCTestnet Chain = "bsc-testnet"
	ChainPolygon    Chain = "polygon"
	ChainAmoy       Chain = "amoy"
	ChainAvalanche  Chain = "avalanche"
	ChainFuji       Chain = "fuji"
)

type SignOptions interface {
	Chain() Chain
	SetChain(chain Chain) SignOptions

	Contract() common.Address
	SetContract(contract common.Address) SignOptions

	Sender() common.Address
	SetSender(sender common.Address) SignOptions

	Uniq() *[]byte
	SetUniq(uniq []byte) SignOptions
	UnsetUniq() SignOptions

	Signer() *common.Address
	SetSigner(signer common.Address) SignOptions
	UnsetSigner() SignOptions

	ABI() map[string]any
	SetABI(abi map[string]any) SignOptions

	Args() []any
	SetArgs(args []any) SignOptions
}

type signOptions struct {
	Chain_    Chain           `json:"chain"`
	Contract_ common.Address  `json:"contract"`
	Sender_   common.Address  `json:"sender"`
	Uniq_     *[]byte         `json:"uniq,omitempty"`
	Signer_   *common.Address `json:"signer,omitempty"`
	ABI_      map[string]any  `json:"abi"`
	Args_     []any           `json:"args"`
}

// ABI implements SignOptions.
func (s *signOptions) ABI() map[string]any {
	return s.ABI_
}

// Args implements SignOptions.
func (s *signOptions) Args() []any {
	return s.Args_
}

// Chain implements SignOptions.
func (s *signOptions) Chain() Chain {
	return s.Chain_
}

// Contract implements SignOptions.
func (s *signOptions) Contract() common.Address {
	return s.Contract_
}

// Sender implements SignOptions.
func (s *signOptions) Sender() common.Address {
	return s.Sender_
}

// SetABI implements SignOptions.
func (s *signOptions) SetABI(abi map[string]any) SignOptions {
	s.ABI_ = abi
	return s
}

// SetArgs implements SignOptions.
func (s *signOptions) SetArgs(args []any) SignOptions {
	s.Args_ = args
	return s
}

// SetChain implements SignOptions.
func (s *signOptions) SetChain(chain Chain) SignOptions {
	s.Chain_ = chain
	return s
}

// SetContract implements SignOptions.
func (s *signOptions) SetContract(contract common.Address) SignOptions {
	s.Contract_ = contract
	return s
}

// SetSender implements SignOptions.
func (s *signOptions) SetSender(sender common.Address) SignOptions {
	s.Sender_ = sender
	return s
}

// SetSigner implements SignOptions.
func (s *signOptions) SetSigner(signer common.Address) SignOptions {
	s.Signer_ = &signer
	return s
}

// SetUniq implements SignOptions.
func (s *signOptions) SetUniq(uniq []byte) SignOptions {
	s.Uniq_ = &uniq
	return s
}

// Signer implements SignOptions.
func (s *signOptions) Signer() *common.Address {
	return s.Signer_
}

// Uniq implements SignOptions.
func (s *signOptions) Uniq() *[]byte {
	return s.Uniq_
}

// UnsetSigner implements SignOptions.
func (s *signOptions) UnsetSigner() SignOptions {
	s.Signer_ = nil
	return s
}

// UnsetUniq implements SignOptions.
func (s *signOptions) UnsetUniq() SignOptions {
	s.Uniq_ = nil
	return s
}

var _ SignOptions = &signOptions{}

func NewSignOptions() SignOptions {
	return &signOptions{}
}

type SignResult interface {
	SubmissionHash() string
	Args() []any
}

type signResult struct {
	SubmissionHash_ string `json:"submissionHash"`
	Args_           []any  `json:"args"`
}

// Args implements SignResult.
func (s *signResult) Args() []any {
	return s.Args_
}

// SubmissionHash implements SignResult.
func (s *signResult) SubmissionHash() string {
	return s.SubmissionHash_
}

var _ SignResult = &signResult{}

func (c *client) Sign(ctx context.Context, options SignOptions) (SignResult, error) {
	var response apiResponse[signResult]

	if err := c.Request(ctx, "POST", fmt.Sprintf("/api/v1/vaults/%s/sign", c.VaultID_), options, &response); err != nil {
		return nil, err
	} else {
		return &response.Data, nil
	}
}
