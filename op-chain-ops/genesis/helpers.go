package genesis

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// codeNamespace represents the namespace of implementations of predeploys
	codeNamespace = common.HexToAddress("0xc0D3C0d3C0d3C0D3c0d3C0d3c0D3C0d3c0d30000")
	// l2PredeployNamespace represents the namespace of L2 predeploys
	l2PredeployNamespace = common.HexToAddress("0x4200000000000000000000000000000000000000")
	// l1PredeployNamespace represents the namespace of L1 predeploys
	l1PredeployNamespace = common.HexToAddress("0x6900000000000000000000000000000000000000")
	// bigL2PredeployNamespace represents the predeploy namespace as a big.Int
	BigL2PredeployNamespace = new(big.Int).SetBytes(l2PredeployNamespace.Bytes())
	// bigL1PredeployNamespace represents the predeploy namespace as a big.Int
	bigL1PredeployNamespace = new(big.Int).SetBytes(l1PredeployNamespace.Bytes())
	// bigCodeNamespace represents the predeploy namespace as a big.Int
	bigCodeNameSpace = new(big.Int).SetBytes(codeNamespace.Bytes())
	// implementationSlot represents the EIP 1967 implementation storage slot
	ImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
	// implementationSlot represents the EIP 1967 admin storage slot
	AdminSlot = common.HexToHash("0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103")
)

// DevAccounts represent the standard hardhat development accounts.
// These are funded if the deploy config has funding development
// accounts enabled.
var DevAccounts = []common.Address{
	common.HexToAddress("0x9FB6B7a85F17adf3aF29e69fC41d6C3845D92dC1"),
	common.HexToAddress("0xA24066A475f5e8e64E39Dff20bb2b4Aa36769c83"),
	common.HexToAddress("0xfF346dE548f655930A9876B63ce9216164322bf2"),
}

// The devBalance is the amount of wei that a dev account is funded with.
var devBalance = hexutil.MustDecodeBig("0x200000000000000000000000000000000000000000000000000000000000000")

// AddressToCodeNamespace takes a predeploy address and computes
// the implementation address that the implementation should be deployed at
func AddressToCodeNamespace(addr common.Address) (common.Address, error) {
	if !IsL1DevPredeploy(addr) && !IsL2DevPredeploy(addr) {
		return common.Address{}, fmt.Errorf("cannot handle non predeploy: %s", addr)
	}
	bigAddress := new(big.Int).SetBytes(addr[18:])
	num := new(big.Int).Or(bigCodeNameSpace, bigAddress)
	return common.BigToAddress(num), nil
}

func IsL1DevPredeploy(addr common.Address) bool {
	return bytes.Equal(addr[0:2], []byte{0x69, 0x00})
}

func IsL2DevPredeploy(addr common.Address) bool {
	return bytes.Equal(addr[0:2], []byte{0x42, 0x00})
}

// GetBlockFromTag will resolve a Block given an rpc block tag
func GetBlockFromTag(chain ethereum.ChainReader, tag *rpc.BlockNumberOrHash) (*types.Block, error) {
	if hash, ok := tag.Hash(); ok {
		block, err := chain.BlockByHash(context.Background(), hash)
		if err != nil {
			return nil, err
		}
		return block, nil
	} else if num, ok := tag.Number(); ok {
		blockNumber := new(big.Int).SetInt64(num.Int64())
		block, err := chain.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			return nil, err
		}
		return block, nil
	} else {
		return nil, fmt.Errorf("invalid block tag: %v", tag)
	}
}

// uint642Big creates a new *big.Int from a uint64.
func uint642Big(in uint64) *big.Int {
	return new(big.Int).SetUint64(in)
}

func newHexBig(in uint64) *hexutil.Big {
	b := new(big.Int).SetUint64(in)
	hb := hexutil.Big(*b)
	return &hb
}
