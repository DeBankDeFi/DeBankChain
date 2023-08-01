// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DebankMintBurnManagerMetaData contains all meta data concerning the DebankMintBurnManager contract.
var DebankMintBurnManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1TxId\",\"type\":\"bytes32\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1ChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1TokenId\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l1TxId\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l1ChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l1TokenId\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b5060016080819052600060a081905260c082905281610f0c61004a833960006108c80152600061089f015260006108760152610f0c6000f3fe6080604052600436106100645760003560e01c80633092afd5116100435780633092afd5146100be57806354fd4d50146100de578063983b2d561461010957600080fd5b8062f714ce1461006957806313af40351461007e5780631e458bee1461009e575b600080fd5b61007c610077366004610c3c565b610129565b005b34801561008a57600080fd5b5061007c610099366004610c68565b6102ef565b3480156100aa57600080fd5b5061007c6100b9366004610c8a565b6104d4565b3480156100ca57600080fd5b5061007c6100d9366004610c68565b6106b1565b3480156100ea57600080fd5b506100f361086f565b6040516101009190610ced565b60405180910390f35b34801561011557600080fd5b5061007c610124366004610c68565b610912565b600034116101be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f6e6f20746f6b656e2077697468647261772077697468207472616e736163746960448201527f6f6e00000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6040516000908190819034908281818185825af1925050503d8060008114610202576040519150601f19603f3d011682016040523d82523d6000602084013e610207565b606091505b5050905080610298576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f7472616e7366657220696e2077697468647261772066756e6374696f6e20666160448201527f696c65642e00000000000000000000000000000000000000000000000000000060648201526084016101b5565b6040805185815273ffffffffffffffffffffffffffffffffffffffff8516602082015233917f56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b291015b60405180910390a250505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f446562616e6b4d696e744275726e4d616e616765723a2066756e6374696f6e2060448201527f63616e206f6e6c792062652063616c6c656420627920746865206f776e65720060648201526084016101b5565b73ffffffffffffffffffffffffffffffffffffffff8116610439576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f446562616e6b4d696e744275726e4d616e616765723a206f776e65722063616e60448201527f206e6f74206265207a65726f206164647265737300000000000000000000000060648201526084016101b5565b6000546040805173ffffffffffffffffffffffffffffffffffffffff928316815291831660208301527fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c910160405180910390a1600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3360009081526001602052604090205460ff16610573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f446562616e6b4d696e744275726e4d616e616765723a2063616c6c657220646f60448201527f6573206e6f74206861766520746865204d696e74657220726f6c65000000000060648201526084016101b5565b60008373ffffffffffffffffffffffffffffffffffffffff168360405160006040518083038185875af1925050503d80600081146105cd576040519150601f19603f3d011682016040523d82523d6000602084013e6105d2565b606091505b5050905080610663576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f7472616e7366657220696e206d696e742066756e6374696f6e206661696c656460448201527f2e0000000000000000000000000000000000000000000000000000000000000060648201526084016101b5565b604080518481526020810184905273ffffffffffffffffffffffffffffffffffffffff8616917f3dec94b8abc8f801eaade1616d3aadd3114b556a284267905e0a053b2df3989291016102e1565b60005473ffffffffffffffffffffffffffffffffffffffff163314610758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f446562616e6b4d696e744275726e4d616e616765723a2066756e6374696f6e2060448201527f63616e206f6e6c792062652063616c6c656420627920746865206f776e65720060648201526084016101b5565b73ffffffffffffffffffffffffffffffffffffffff81166107fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f446562616e6b4d696e744275726e4d616e616765723a206d696e74657220636160448201527f6e206e6f74206265207a65726f2061646472657373000000000000000000000060648201526084016101b5565b73ffffffffffffffffffffffffffffffffffffffff811660008181526001602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb666929190a250565b606061089a7f0000000000000000000000000000000000000000000000000000000000000000610ad6565b6108c37f0000000000000000000000000000000000000000000000000000000000000000610ad6565b6108ec7f0000000000000000000000000000000000000000000000000000000000000000610ad6565b6040516020016108fe93929190610d3e565b604051602081830303815290604052905090565b60005473ffffffffffffffffffffffffffffffffffffffff1633146109b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f446562616e6b4d696e744275726e4d616e616765723a2066756e6374696f6e2060448201527f63616e206f6e6c792062652063616c6c656420627920746865206f776e65720060648201526084016101b5565b73ffffffffffffffffffffffffffffffffffffffff8116610a5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f446562616e6b4d696e744275726e4d616e616765723a206d696e74657220636160448201527f6e206e6f74206265207a65726f2061646472657373000000000000000000000060648201526084016101b5565b73ffffffffffffffffffffffffffffffffffffffff8116600081815260016020819052604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909217909155517f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f69190a250565b606081600003610b1957505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610b435780610b2d81610de3565b9150610b3c9050600a83610e4a565b9150610b1d565b60008167ffffffffffffffff811115610b5e57610b5e610e5e565b6040519080825280601f01601f191660200182016040528015610b88576020820181803683370190505b5090505b8415610c0b57610b9d600183610e8d565b9150610baa600a86610ea4565b610bb5906030610eb8565b60f81b818381518110610bca57610bca610ed0565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610c04600a86610e4a565b9450610b8c565b949350505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c3757600080fd5b919050565b60008060408385031215610c4f57600080fd5b82359150610c5f60208401610c13565b90509250929050565b600060208284031215610c7a57600080fd5b610c8382610c13565b9392505050565b600080600060608486031215610c9f57600080fd5b610ca884610c13565b95602085013595506040909401359392505050565b60005b83811015610cd8578181015183820152602001610cc0565b83811115610ce7576000848401525b50505050565b6020815260008251806020840152610d0c816040850160208701610cbd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60008451610d50818460208901610cbd565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610d8c816001850160208a01610cbd565b60019201918201528351610da7816002840160208801610cbd565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e1457610e14610db4565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610e5957610e59610e1b565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610e9f57610e9f610db4565b500390565b600082610eb357610eb3610e1b565b500690565b60008219821115610ecb57610ecb610db4565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// DebankMintBurnManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DebankMintBurnManagerMetaData.ABI instead.
var DebankMintBurnManagerABI = DebankMintBurnManagerMetaData.ABI

// DebankMintBurnManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DebankMintBurnManagerMetaData.Bin instead.
var DebankMintBurnManagerBin = DebankMintBurnManagerMetaData.Bin

// DeployDebankMintBurnManager deploys a new Ethereum contract, binding an instance of DebankMintBurnManager to it.
func DeployDebankMintBurnManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebankMintBurnManager, error) {
	parsed, err := DebankMintBurnManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DebankMintBurnManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebankMintBurnManager{DebankMintBurnManagerCaller: DebankMintBurnManagerCaller{contract: contract}, DebankMintBurnManagerTransactor: DebankMintBurnManagerTransactor{contract: contract}, DebankMintBurnManagerFilterer: DebankMintBurnManagerFilterer{contract: contract}}, nil
}

// DebankMintBurnManager is an auto generated Go binding around an Ethereum contract.
type DebankMintBurnManager struct {
	DebankMintBurnManagerCaller     // Read-only binding to the contract
	DebankMintBurnManagerTransactor // Write-only binding to the contract
	DebankMintBurnManagerFilterer   // Log filterer for contract events
}

// DebankMintBurnManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebankMintBurnManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebankMintBurnManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebankMintBurnManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebankMintBurnManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebankMintBurnManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebankMintBurnManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebankMintBurnManagerSession struct {
	Contract     *DebankMintBurnManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DebankMintBurnManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebankMintBurnManagerCallerSession struct {
	Contract *DebankMintBurnManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DebankMintBurnManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebankMintBurnManagerTransactorSession struct {
	Contract     *DebankMintBurnManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DebankMintBurnManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebankMintBurnManagerRaw struct {
	Contract *DebankMintBurnManager // Generic contract binding to access the raw methods on
}

// DebankMintBurnManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebankMintBurnManagerCallerRaw struct {
	Contract *DebankMintBurnManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DebankMintBurnManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebankMintBurnManagerTransactorRaw struct {
	Contract *DebankMintBurnManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebankMintBurnManager creates a new instance of DebankMintBurnManager, bound to a specific deployed contract.
func NewDebankMintBurnManager(address common.Address, backend bind.ContractBackend) (*DebankMintBurnManager, error) {
	contract, err := bindDebankMintBurnManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManager{DebankMintBurnManagerCaller: DebankMintBurnManagerCaller{contract: contract}, DebankMintBurnManagerTransactor: DebankMintBurnManagerTransactor{contract: contract}, DebankMintBurnManagerFilterer: DebankMintBurnManagerFilterer{contract: contract}}, nil
}

// NewDebankMintBurnManagerCaller creates a new read-only instance of DebankMintBurnManager, bound to a specific deployed contract.
func NewDebankMintBurnManagerCaller(address common.Address, caller bind.ContractCaller) (*DebankMintBurnManagerCaller, error) {
	contract, err := bindDebankMintBurnManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerCaller{contract: contract}, nil
}

// NewDebankMintBurnManagerTransactor creates a new write-only instance of DebankMintBurnManager, bound to a specific deployed contract.
func NewDebankMintBurnManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DebankMintBurnManagerTransactor, error) {
	contract, err := bindDebankMintBurnManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerTransactor{contract: contract}, nil
}

// NewDebankMintBurnManagerFilterer creates a new log filterer instance of DebankMintBurnManager, bound to a specific deployed contract.
func NewDebankMintBurnManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DebankMintBurnManagerFilterer, error) {
	contract, err := bindDebankMintBurnManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerFilterer{contract: contract}, nil
}

// bindDebankMintBurnManager binds a generic wrapper to an already deployed contract.
func bindDebankMintBurnManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebankMintBurnManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebankMintBurnManager *DebankMintBurnManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebankMintBurnManager.Contract.DebankMintBurnManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebankMintBurnManager *DebankMintBurnManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.DebankMintBurnManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebankMintBurnManager *DebankMintBurnManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.DebankMintBurnManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebankMintBurnManager *DebankMintBurnManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DebankMintBurnManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DebankMintBurnManager *DebankMintBurnManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DebankMintBurnManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DebankMintBurnManager *DebankMintBurnManagerSession) Version() (string, error) {
	return _DebankMintBurnManager.Contract.Version(&_DebankMintBurnManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DebankMintBurnManager *DebankMintBurnManagerCallerSession) Version() (string, error) {
	return _DebankMintBurnManager.Contract.Version(&_DebankMintBurnManager.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.AddMinter(&_DebankMintBurnManager.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.AddMinter(&_DebankMintBurnManager.TransactOpts, account)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address to, uint256 value, bytes32 l1TxId) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int, l1TxId [32]byte) (*types.Transaction, error) {
	return _DebankMintBurnManager.contract.Transact(opts, "mint", to, value, l1TxId)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address to, uint256 value, bytes32 l1TxId) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerSession) Mint(to common.Address, value *big.Int, l1TxId [32]byte) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.Mint(&_DebankMintBurnManager.TransactOpts, to, value, l1TxId)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address to, uint256 value, bytes32 l1TxId) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorSession) Mint(to common.Address, value *big.Int, l1TxId [32]byte) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.Mint(&_DebankMintBurnManager.TransactOpts, to, value, l1TxId)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address account) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactor) RemoveMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.contract.Transact(opts, "removeMinter", account)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address account) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerSession) RemoveMinter(account common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.RemoveMinter(&_DebankMintBurnManager.TransactOpts, account)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address account) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorSession) RemoveMinter(account common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.RemoveMinter(&_DebankMintBurnManager.TransactOpts, account)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.SetOwner(&_DebankMintBurnManager.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.SetOwner(&_DebankMintBurnManager.TransactOpts, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 l1ChainId, address l1TokenId) payable returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactor) Withdraw(opts *bind.TransactOpts, l1ChainId *big.Int, l1TokenId common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.contract.Transact(opts, "withdraw", l1ChainId, l1TokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 l1ChainId, address l1TokenId) payable returns()
func (_DebankMintBurnManager *DebankMintBurnManagerSession) Withdraw(l1ChainId *big.Int, l1TokenId common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.Withdraw(&_DebankMintBurnManager.TransactOpts, l1ChainId, l1TokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 l1ChainId, address l1TokenId) payable returns()
func (_DebankMintBurnManager *DebankMintBurnManagerTransactorSession) Withdraw(l1ChainId *big.Int, l1TokenId common.Address) (*types.Transaction, error) {
	return _DebankMintBurnManager.Contract.Withdraw(&_DebankMintBurnManager.TransactOpts, l1ChainId, l1TokenId)
}

// DebankMintBurnManagerMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerMintIterator struct {
	Event *DebankMintBurnManagerMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DebankMintBurnManagerMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankMintBurnManagerMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DebankMintBurnManagerMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DebankMintBurnManagerMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankMintBurnManagerMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankMintBurnManagerMint represents a Mint event raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerMint struct {
	To     common.Address
	Value  *big.Int
	L1TxId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x3dec94b8abc8f801eaade1616d3aadd3114b556a284267905e0a053b2df39892.
//
// Solidity: event Mint(address indexed to, uint256 value, bytes32 l1TxId)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*DebankMintBurnManagerMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerMintIterator{contract: _DebankMintBurnManager.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x3dec94b8abc8f801eaade1616d3aadd3114b556a284267905e0a053b2df39892.
//
// Solidity: event Mint(address indexed to, uint256 value, bytes32 l1TxId)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DebankMintBurnManagerMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankMintBurnManagerMint)
				if err := _DebankMintBurnManager.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x3dec94b8abc8f801eaade1616d3aadd3114b556a284267905e0a053b2df39892.
//
// Solidity: event Mint(address indexed to, uint256 value, bytes32 l1TxId)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) ParseMint(log types.Log) (*DebankMintBurnManagerMint, error) {
	event := new(DebankMintBurnManagerMint)
	if err := _DebankMintBurnManager.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebankMintBurnManagerMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerMinterAddedIterator struct {
	Event *DebankMintBurnManagerMinterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DebankMintBurnManagerMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankMintBurnManagerMinterAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DebankMintBurnManagerMinterAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DebankMintBurnManagerMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankMintBurnManagerMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankMintBurnManagerMinterAdded represents a MinterAdded event raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*DebankMintBurnManagerMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerMinterAddedIterator{contract: _DebankMintBurnManager.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *DebankMintBurnManagerMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankMintBurnManagerMinterAdded)
				if err := _DebankMintBurnManager.contract.UnpackLog(event, "MinterAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) ParseMinterAdded(log types.Log) (*DebankMintBurnManagerMinterAdded, error) {
	event := new(DebankMintBurnManagerMinterAdded)
	if err := _DebankMintBurnManager.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebankMintBurnManagerMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerMinterRemovedIterator struct {
	Event *DebankMintBurnManagerMinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DebankMintBurnManagerMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankMintBurnManagerMinterRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DebankMintBurnManagerMinterRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DebankMintBurnManagerMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankMintBurnManagerMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankMintBurnManagerMinterRemoved represents a MinterRemoved event raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*DebankMintBurnManagerMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerMinterRemovedIterator{contract: _DebankMintBurnManager.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *DebankMintBurnManagerMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankMintBurnManagerMinterRemoved)
				if err := _DebankMintBurnManager.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) ParseMinterRemoved(log types.Log) (*DebankMintBurnManagerMinterRemoved, error) {
	event := new(DebankMintBurnManagerMinterRemoved)
	if err := _DebankMintBurnManager.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebankMintBurnManagerOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerOwnerChangedIterator struct {
	Event *DebankMintBurnManagerOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DebankMintBurnManagerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankMintBurnManagerOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DebankMintBurnManagerOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DebankMintBurnManagerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankMintBurnManagerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankMintBurnManagerOwnerChanged represents a OwnerChanged event raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*DebankMintBurnManagerOwnerChangedIterator, error) {

	logs, sub, err := _DebankMintBurnManager.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerOwnerChangedIterator{contract: _DebankMintBurnManager.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *DebankMintBurnManagerOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _DebankMintBurnManager.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankMintBurnManagerOwnerChanged)
				if err := _DebankMintBurnManager.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) ParseOwnerChanged(log types.Log) (*DebankMintBurnManagerOwnerChanged, error) {
	event := new(DebankMintBurnManagerOwnerChanged)
	if err := _DebankMintBurnManager.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DebankMintBurnManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerWithdrawIterator struct {
	Event *DebankMintBurnManagerWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DebankMintBurnManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DebankMintBurnManagerWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DebankMintBurnManagerWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DebankMintBurnManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DebankMintBurnManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DebankMintBurnManagerWithdraw represents a Withdraw event raised by the DebankMintBurnManager contract.
type DebankMintBurnManagerWithdraw struct {
	From      common.Address
	L1ChainId *big.Int
	L1TokenId common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed from, uint256 l1ChainId, address l1TokenId)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address) (*DebankMintBurnManagerWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.FilterLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return &DebankMintBurnManagerWithdrawIterator{contract: _DebankMintBurnManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed from, uint256 l1ChainId, address l1TokenId)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DebankMintBurnManagerWithdraw, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DebankMintBurnManager.contract.WatchLogs(opts, "Withdraw", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DebankMintBurnManagerWithdraw)
				if err := _DebankMintBurnManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed from, uint256 l1ChainId, address l1TokenId)
func (_DebankMintBurnManager *DebankMintBurnManagerFilterer) ParseWithdraw(log types.Log) (*DebankMintBurnManagerWithdraw, error) {
	event := new(DebankMintBurnManagerWithdraw)
	if err := _DebankMintBurnManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
