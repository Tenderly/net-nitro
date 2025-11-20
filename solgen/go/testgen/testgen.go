// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testgen

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
	_ = abi.ConvertType
)

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// DelegateCallerMetaData contains all meta data concerning the DelegateCaller contract.
var DelegateCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_called\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"makeDelegatecall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610298806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e632e17214610030575b600080fd5b6101096004803603604081101561004657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561008357600080fd5b82018360208201111561009557600080fd5b803590602001918460018302840111640100000000831117156100b757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061018d565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610151578082015181840152602081019050610136565b50505050905090810190601f16801561017e5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b602083106101dc57805182526020820191506020810190506020830392506101b9565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461023c576040519150601f19603f3d011682016040523d82523d6000602084013e610241565b606091505b5080925081935050508161025b576040513d6000823e3d81fd5b925092905056fea2646970667358221220d84f4ab0b3696a6c4a0b664b25c4aa624aaadfad1e965e5ceedb6f18f4dd86dc64736f6c63430007060033",
}

// DelegateCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegateCallerMetaData.ABI instead.
var DelegateCallerABI = DelegateCallerMetaData.ABI

// DelegateCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegateCallerMetaData.Bin instead.
var DelegateCallerBin = DelegateCallerMetaData.Bin

// DeployDelegateCaller deploys a new Ethereum contract, binding an instance of DelegateCaller to it.
func DeployDelegateCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DelegateCaller, error) {
	parsed, err := DelegateCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegateCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelegateCaller{DelegateCallerCaller: DelegateCallerCaller{contract: contract}, DelegateCallerTransactor: DelegateCallerTransactor{contract: contract}, DelegateCallerFilterer: DelegateCallerFilterer{contract: contract}}, nil
}

// DelegateCaller is an auto generated Go binding around an Ethereum contract.
type DelegateCaller struct {
	DelegateCallerCaller     // Read-only binding to the contract
	DelegateCallerTransactor // Write-only binding to the contract
	DelegateCallerFilterer   // Log filterer for contract events
}

// DelegateCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegateCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegateCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegateCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegateCallerSession struct {
	Contract     *DelegateCaller   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegateCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegateCallerCallerSession struct {
	Contract *DelegateCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DelegateCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegateCallerTransactorSession struct {
	Contract     *DelegateCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DelegateCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegateCallerRaw struct {
	Contract *DelegateCaller // Generic contract binding to access the raw methods on
}

// DelegateCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegateCallerCallerRaw struct {
	Contract *DelegateCallerCaller // Generic read-only contract binding to access the raw methods on
}

// DelegateCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegateCallerTransactorRaw struct {
	Contract *DelegateCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegateCaller creates a new instance of DelegateCaller, bound to a specific deployed contract.
func NewDelegateCaller(address common.Address, backend bind.ContractBackend) (*DelegateCaller, error) {
	contract, err := bindDelegateCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegateCaller{DelegateCallerCaller: DelegateCallerCaller{contract: contract}, DelegateCallerTransactor: DelegateCallerTransactor{contract: contract}, DelegateCallerFilterer: DelegateCallerFilterer{contract: contract}}, nil
}

// NewDelegateCallerCaller creates a new read-only instance of DelegateCaller, bound to a specific deployed contract.
func NewDelegateCallerCaller(address common.Address, caller bind.ContractCaller) (*DelegateCallerCaller, error) {
	contract, err := bindDelegateCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCallerCaller{contract: contract}, nil
}

// NewDelegateCallerTransactor creates a new write-only instance of DelegateCaller, bound to a specific deployed contract.
func NewDelegateCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegateCallerTransactor, error) {
	contract, err := bindDelegateCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCallerTransactor{contract: contract}, nil
}

// NewDelegateCallerFilterer creates a new log filterer instance of DelegateCaller, bound to a specific deployed contract.
func NewDelegateCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegateCallerFilterer, error) {
	contract, err := bindDelegateCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegateCallerFilterer{contract: contract}, nil
}

// bindDelegateCaller binds a generic wrapper to an already deployed contract.
func bindDelegateCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegateCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateCaller *DelegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegateCaller.Contract.DelegateCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateCaller *DelegateCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCaller.Contract.DelegateCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateCaller *DelegateCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateCaller.Contract.DelegateCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateCaller *DelegateCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegateCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateCaller *DelegateCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateCaller *DelegateCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateCaller.Contract.contract.Transact(opts, method, params...)
}

// MakeDelegatecall is a paid mutator transaction binding the contract method 0xe632e172.
//
// Solidity: function makeDelegatecall(address _called, bytes _calldata) returns(bool success, bytes returnData)
func (_DelegateCaller *DelegateCallerTransactor) MakeDelegatecall(opts *bind.TransactOpts, _called common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "makeDelegatecall", _called, _calldata)
}

// MakeDelegatecall is a paid mutator transaction binding the contract method 0xe632e172.
//
// Solidity: function makeDelegatecall(address _called, bytes _calldata) returns(bool success, bytes returnData)
func (_DelegateCaller *DelegateCallerSession) MakeDelegatecall(_called common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DelegateCaller.Contract.MakeDelegatecall(&_DelegateCaller.TransactOpts, _called, _calldata)
}

// MakeDelegatecall is a paid mutator transaction binding the contract method 0xe632e172.
//
// Solidity: function makeDelegatecall(address _called, bytes _calldata) returns(bool success, bytes returnData)
func (_DelegateCaller *DelegateCallerTransactorSession) MakeDelegatecall(_called common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DelegateCaller.Contract.MakeDelegatecall(&_DelegateCaller.TransactOpts, _called, _calldata)
}

// ERC1155TokenMetaData contains all meta data concerning the ERC1155Token contract.
var ERC1155TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"trickFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280601381526020017f68747470733a2f2f6578616d706c652e636f6d00000000000000000000000000815250620000606301ffc9a760e01b620000a860201b60201c565b6200007181620001b160201b60201c565b6200008963d9b67a2660e01b620000a860201b60201c565b620000a1630e89341c60e01b620000a860201b60201c565b5062000283565b63ffffffff60e01b817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141562000145576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433136353a20696e76616c696420696e746572666163652069640000000081525060200191505060405180910390fd5b6001600080837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b8060039080519060200190620001c9929190620001cd565b5050565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000205576000855562000251565b82601f106200022057805160ff191683800117855562000251565b8280016001018555821562000251579182015b828111156200025057825182559160200191906001019062000233565b5b50905062000260919062000264565b5090565b5b808211156200027f57600081600090555060010162000265565b5090565b612b3880620002936000396000f3fe608060405234801561001057600080fd5b50600436106100ce5760003560e01c80635624b25b1161008c578063b4faba0911610066578063b4faba09146108f4578063dbaabecb146109cf578063e985e9c514610a13578063f242432a14610a8d576100ce565b80635624b25b14610746578063731133e9146107f7578063a22cb465146108a4576100ce565b8062fdd58e146100d357806301ffc9a7146101355780630e89341c146101985780631f7fdffa1461023f5780632eb2c2d6146103825780634e1273f4146105a5575b600080fd5b61011f600480360360408110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b9c565b6040518082815260200191505060405180910390f35b6101806004803603602081101561014b57600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610c7c565b60405180821515815260200191505060405180910390f35b6101c4600480360360208110156101ae57600080fd5b8101908080359060200190929190505050610ce3565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102045780820151818401526020810190506101e9565b50505050905090810190601f1680156102315780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103806004803603608081101561025557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561029257600080fd5b8201836020820111156102a457600080fd5b803590602001918460208302840111640100000000831117156102c657600080fd5b9091929391929390803590602001906401000000008111156102e757600080fd5b8201836020820111156102f957600080fd5b8035906020019184602083028401116401000000008311171561031b57600080fd5b90919293919293908035906020019064010000000081111561033c57600080fd5b82018360208201111561034e57600080fd5b8035906020019184600183028401116401000000008311171561037057600080fd5b9091929391929390505050610d87565b005b6105a3600480360360a081101561039857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156103f557600080fd5b82018360208201111561040757600080fd5b8035906020019184602083028401116401000000008311171561042957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561048957600080fd5b82018360208201111561049b57600080fd5b803590602001918460208302840111640100000000831117156104bd57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561051d57600080fd5b82018360208201111561052f57600080fd5b8035906020019184600183028401116401000000008311171561055157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e62565b005b6106ef600480360360408110156105bb57600080fd5b81019080803590602001906401000000008111156105d857600080fd5b8201836020820111156105ea57600080fd5b8035906020019184602083028401116401000000008311171561060c57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561066c57600080fd5b82018360208201111561067e57600080fd5b803590602001918460208302840111640100000000831117156106a057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506112ed565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610732578082015181840152602081019050610717565b505050509050019250505060405180910390f35b61077c6004803603604081101561075c57600080fd5b8101908080359060200190929190803590602001909291905050506113ff565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107bc5780820151818401526020810190506107a1565b50505050905090810190601f1680156107e95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6108a26004803603608081101561080d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561085e57600080fd5b82018360208201111561087057600080fd5b8035906020019184600183028401116401000000008311171561089257600080fd5b9091929391929390505050611485565b005b6108f2600480360360408110156108ba57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506114dc565b005b6109cd6004803603604081101561090a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561094757600080fd5b82018360208201111561095957600080fd5b8035906020019184600183028401116401000000008311171561097b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611675565b005b610a11600480360360208110156109e557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061169c565b005b610a7560048036036040811015610a2957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116c2565b60405180821515815260200191505060405180910390f35b610b9a600480360360a0811015610aa357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190640100000000811115610b1457600080fd5b820183602082011115610b2657600080fd5b80359060200191846001830284011164010000000083111715610b4857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611756565b005b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180612993602b913960400191505060405180910390fd5b6001600083815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000806000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b606060038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d7b5780601f10610d5057610100808354040283529160200191610d7b565b820191906000526020600020905b815481529060010190602001808311610d5e57829003601f168201915b50505050509050919050565b610e5987878780806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050611acb565b50505050505050565b8151835114610ebc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180612aba6028913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610f42576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806129e76025913960400191505060405180910390fd5b610f4a611ded565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610f905750610f8f85610f8a611ded565b6116c2565b5b610fe5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526032815260200180612a0c6032913960400191505060405180910390fd5b6000610fef611ded565b9050610fff818787878787611df5565b60005b84518110156111d057600085828151811061101957fe5b60200260200101519050600085838151811061103157fe5b602002602001015190506110b8816040518060600160405280602a8152602001612a3e602a91396001600086815260200190815260200160002060008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dfd9092919063ffffffff16565b6001600084815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061116f816001600085815260200190815260200160002060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611eb790919063ffffffff16565b6001600084815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050806001019050611002565b508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611280578082015181840152602081019050611265565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156112c25780820151818401526020810190506112a7565b5050505090500194505050505060405180910390a46112e5818787878787611f3f565b505050505050565b60608151835114611349576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180612a916029913960400191505060405180910390fd5b6000835167ffffffffffffffff8111801561136357600080fd5b506040519080825280602002602001820160405280156113925781602001602082028036833780820191505090505b50905060005b84518110156113f4576113d18582815181106113b057fe5b60200260200101518583815181106113c457fe5b6020026020010151610b9c565b8282815181106113dd57fe5b602002602001018181525050806001019050611398565b508091505092915050565b60606000600583901b67ffffffffffffffff8111801561141e57600080fd5b506040519080825280601f01601f1916602001820160405280156114515781602001600182028036833780820191505090505b50905060005b8381101561147a5780850154806020830260208501015250806001019050611457565b508091505092915050565b6114d585858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506122ce565b5050505050565b8173ffffffffffffffffffffffffffffffffffffffff166114fb611ded565b73ffffffffffffffffffffffffffffffffffffffff161415611568576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180612a686029913960400191505060405180910390fd5b8060026000611575611ded565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611622611ded565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405180821515815260200191505060405180910390a35050565b600080825160208401855af46040518181523d60208201523d6000604083013e60403d0181fd5b807f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d55550565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614156117dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806129e76025913960400191505060405180910390fd5b6117e4611ded565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061182a575061182985611824611ded565b6116c2565b5b61187f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806129be6029913960400191505060405180910390fd5b6000611889611ded565b90506118a981878761189a886124d1565b6118a3886124d1565b87611df5565b611926836040518060600160405280602a8152602001612a3e602a91396001600088815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dfd9092919063ffffffff16565b6001600086815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506119dd836001600087815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611eb790919063ffffffff16565b6001600086815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628787604051808381526020018281526020019250505060405180910390a4611ac3818787878787612542565b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611b51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612ae26021913960400191505060405180910390fd5b8151835114611bab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180612aba6028913960400191505060405180910390fd5b6000611bb5611ded565b9050611bc681600087878787611df5565b60005b8451811015611ccf57611c5b60016000878481518110611be557fe5b6020026020010151815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054858381518110611c4557fe5b6020026020010151611eb790919063ffffffff16565b60016000878481518110611c6b57fe5b6020026020010151815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508080600101915050611bc9565b508473ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015611d80578082015181840152602081019050611d65565b50505050905001838103825284818151815260200191508051906020019060200280838360005b83811015611dc2578082015181840152602081019050611da7565b5050505090500194505050505060405180910390a4611de681600087878787611f3f565b5050505050565b600033905090565b505050505050565b6000838311158290611eaa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611e6f578082015181840152602081019050611e54565b50505050905090810190601f168015611e9c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390509392505050565b600080828401905083811015611f35576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b611f5e8473ffffffffffffffffffffffffffffffffffffffff1661284f565b156122c6578373ffffffffffffffffffffffffffffffffffffffff1663bc197c8187878686866040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015612016578082015181840152602081019050611ffb565b50505050905001848103835286818151815260200191508051906020019060200280838360005b8381101561205857808201518184015260208101905061203d565b50505050905001848103825285818151815260200191508051906020019080838360005b8381101561209757808201518184015260208101905061207c565b50505050905090810190601f1680156120c45780820380516001836020036101000a031916815260200191505b5098505050505050505050602060405180830381600087803b1580156120e957600080fd5b505af192505050801561211d57506040513d602081101561210957600080fd5b810190808051906020019092919050505060015b61222757612129612880565b8061213457506121d6565b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561219b578082015181840152602081019050612180565b50505050905090810190601f1680156121c85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129376034913960400191505060405180910390fd5b63bc197c8160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146122c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061296b6028913960400191505060405180910390fd5b505b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415612354576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612ae26021913960400191505060405180910390fd5b600061235e611ded565b905061237f81600087612370886124d1565b612379886124d1565b87611df5565b6123e2836001600087815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611eb790919063ffffffff16565b6001600086815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628787604051808381526020018281526020019250505060405180910390a46124ca81600087878787612542565b5050505050565b60606000600167ffffffffffffffff811180156124ed57600080fd5b5060405190808252806020026020018201604052801561251c5781602001602082028036833780820191505090505b509050828160008151811061252d57fe5b60200260200101818152505080915050919050565b6125618473ffffffffffffffffffffffffffffffffffffffff1661284f565b15612847578373ffffffffffffffffffffffffffffffffffffffff1663f23a6e6187878686866040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561261a5780820151818401526020810190506125ff565b50505050905090810190601f1680156126475780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561266a57600080fd5b505af192505050801561269e57506040513d602081101561268a57600080fd5b810190808051906020019092919050505060015b6127a8576126aa612880565b806126b55750612757565b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561271c578082015181840152602081019050612701565b50505050905090810190601f1680156127495780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129376034913960400191505060405180910390fd5b63f23a6e6160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614612845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061296b6028913960400191505060405180910390fd5b505b505050505050565b600080823b905060008111915050919050565b6000601f19601f8301169050919050565b60008160e01c9050919050565b600060443d101561289057612933565b60046000803e6128a1600051612873565b6308c379a081146128b25750612933565b60405160043d036004823e80513d602482011167ffffffffffffffff821117156128de57505050612933565b808201805167ffffffffffffffff8111156128fd575050505050612933565b8060208301013d850181111561291857505050505050612933565b61292182612862565b60208401016040528296505050505050505b9056fe455243313135353a207472616e7366657220746f206e6f6e2045524331313535526563656976657220696d706c656d656e746572455243313135353a204552433131353552656365697665722072656a656374656420746f6b656e73455243313135353a2062616c616e636520717565727920666f7220746865207a65726f2061646472657373455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243313135353a207472616e7366657220746f20746865207a65726f2061646472657373455243313135353a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243313135353a20696e73756666696369656e742062616c616e636520666f72207472616e73666572455243313135353a2073657474696e6720617070726f76616c2073746174757320666f722073656c66455243313135353a206163636f756e747320616e6420696473206c656e677468206d69736d61746368455243313135353a2069647320616e6420616d6f756e7473206c656e677468206d69736d61746368455243313135353a206d696e7420746f20746865207a65726f2061646472657373a2646970667358221220f19f248c56577f5fa53609a0dc21bf09993bb055d129245307f9dc2285429f2864736f6c63430007060033",
}

// ERC1155TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155TokenMetaData.ABI instead.
var ERC1155TokenABI = ERC1155TokenMetaData.ABI

// ERC1155TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1155TokenMetaData.Bin instead.
var ERC1155TokenBin = ERC1155TokenMetaData.Bin

// DeployERC1155Token deploys a new Ethereum contract, binding an instance of ERC1155Token to it.
func DeployERC1155Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC1155Token, error) {
	parsed, err := ERC1155TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1155TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1155Token{ERC1155TokenCaller: ERC1155TokenCaller{contract: contract}, ERC1155TokenTransactor: ERC1155TokenTransactor{contract: contract}, ERC1155TokenFilterer: ERC1155TokenFilterer{contract: contract}}, nil
}

// ERC1155Token is an auto generated Go binding around an Ethereum contract.
type ERC1155Token struct {
	ERC1155TokenCaller     // Read-only binding to the contract
	ERC1155TokenTransactor // Write-only binding to the contract
	ERC1155TokenFilterer   // Log filterer for contract events
}

// ERC1155TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155TokenSession struct {
	Contract     *ERC1155Token     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155TokenCallerSession struct {
	Contract *ERC1155TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC1155TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TokenTransactorSession struct {
	Contract     *ERC1155TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC1155TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155TokenRaw struct {
	Contract *ERC1155Token // Generic contract binding to access the raw methods on
}

// ERC1155TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155TokenCallerRaw struct {
	Contract *ERC1155TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TokenTransactorRaw struct {
	Contract *ERC1155TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Token creates a new instance of ERC1155Token, bound to a specific deployed contract.
func NewERC1155Token(address common.Address, backend bind.ContractBackend) (*ERC1155Token, error) {
	contract, err := bindERC1155Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Token{ERC1155TokenCaller: ERC1155TokenCaller{contract: contract}, ERC1155TokenTransactor: ERC1155TokenTransactor{contract: contract}, ERC1155TokenFilterer: ERC1155TokenFilterer{contract: contract}}, nil
}

// NewERC1155TokenCaller creates a new read-only instance of ERC1155Token, bound to a specific deployed contract.
func NewERC1155TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC1155TokenCaller, error) {
	contract, err := bindERC1155Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenCaller{contract: contract}, nil
}

// NewERC1155TokenTransactor creates a new write-only instance of ERC1155Token, bound to a specific deployed contract.
func NewERC1155TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155TokenTransactor, error) {
	contract, err := bindERC1155Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenTransactor{contract: contract}, nil
}

// NewERC1155TokenFilterer creates a new log filterer instance of ERC1155Token, bound to a specific deployed contract.
func NewERC1155TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155TokenFilterer, error) {
	contract, err := bindERC1155Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenFilterer{contract: contract}, nil
}

// bindERC1155Token binds a generic wrapper to an already deployed contract.
func bindERC1155Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Token *ERC1155TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Token.Contract.ERC1155TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Token *ERC1155TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Token.Contract.ERC1155TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Token *ERC1155TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Token.Contract.ERC1155TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Token *ERC1155TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Token *ERC1155TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Token *ERC1155TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155Token *ERC1155TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155Token.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155Token *ERC1155TokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155Token.Contract.BalanceOf(&_ERC1155Token.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155Token *ERC1155TokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155Token.Contract.BalanceOf(&_ERC1155Token.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155Token *ERC1155TokenCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC1155Token.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155Token *ERC1155TokenSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155Token.Contract.BalanceOfBatch(&_ERC1155Token.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155Token *ERC1155TokenCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155Token.Contract.BalanceOfBatch(&_ERC1155Token.CallOpts, accounts, ids)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_ERC1155Token *ERC1155TokenCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ERC1155Token.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_ERC1155Token *ERC1155TokenSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _ERC1155Token.Contract.GetStorageAt(&_ERC1155Token.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_ERC1155Token *ERC1155TokenCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _ERC1155Token.Contract.GetStorageAt(&_ERC1155Token.CallOpts, offset, length)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155Token *ERC1155TokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155Token.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155Token *ERC1155TokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155Token.Contract.IsApprovedForAll(&_ERC1155Token.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155Token *ERC1155TokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155Token.Contract.IsApprovedForAll(&_ERC1155Token.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Token *ERC1155TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155Token.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Token *ERC1155TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Token.Contract.SupportsInterface(&_ERC1155Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155Token *ERC1155TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155Token.Contract.SupportsInterface(&_ERC1155Token.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155Token *ERC1155TokenCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155Token.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155Token *ERC1155TokenSession) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155Token.Contract.Uri(&_ERC1155Token.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155Token *ERC1155TokenCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155Token.Contract.Uri(&_ERC1155Token.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "mint", to, id, value, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155Token *ERC1155TokenSession) Mint(to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.Mint(&_ERC1155Token.TransactOpts, to, id, value, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) Mint(to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.Mint(&_ERC1155Token.TransactOpts, to, id, value, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "mintBatch", to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155Token *ERC1155TokenSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.MintBatch(&_ERC1155Token.TransactOpts, to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.MintBatch(&_ERC1155Token.TransactOpts, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155Token *ERC1155TokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SafeBatchTransferFrom(&_ERC1155Token.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SafeBatchTransferFrom(&_ERC1155Token.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155Token *ERC1155TokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SafeTransferFrom(&_ERC1155Token.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SafeTransferFrom(&_ERC1155Token.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155Token *ERC1155TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155Token *ERC1155TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SetApprovalForAll(&_ERC1155Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SetApprovalForAll(&_ERC1155Token.TransactOpts, operator, approved)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_ERC1155Token *ERC1155TokenTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_ERC1155Token *ERC1155TokenSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SimulateAndRevert(&_ERC1155Token.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _ERC1155Token.Contract.SimulateAndRevert(&_ERC1155Token.TransactOpts, targetContract, calldataPayload)
}

// TrickFallbackHandler is a paid mutator transaction binding the contract method 0xdbaabecb.
//
// Solidity: function trickFallbackHandler(address fallbackHandler) returns()
func (_ERC1155Token *ERC1155TokenTransactor) TrickFallbackHandler(opts *bind.TransactOpts, fallbackHandler common.Address) (*types.Transaction, error) {
	return _ERC1155Token.contract.Transact(opts, "trickFallbackHandler", fallbackHandler)
}

// TrickFallbackHandler is a paid mutator transaction binding the contract method 0xdbaabecb.
//
// Solidity: function trickFallbackHandler(address fallbackHandler) returns()
func (_ERC1155Token *ERC1155TokenSession) TrickFallbackHandler(fallbackHandler common.Address) (*types.Transaction, error) {
	return _ERC1155Token.Contract.TrickFallbackHandler(&_ERC1155Token.TransactOpts, fallbackHandler)
}

// TrickFallbackHandler is a paid mutator transaction binding the contract method 0xdbaabecb.
//
// Solidity: function trickFallbackHandler(address fallbackHandler) returns()
func (_ERC1155Token *ERC1155TokenTransactorSession) TrickFallbackHandler(fallbackHandler common.Address) (*types.Transaction, error) {
	return _ERC1155Token.Contract.TrickFallbackHandler(&_ERC1155Token.TransactOpts, fallbackHandler)
}

// ERC1155TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155Token contract.
type ERC1155TokenApprovalForAllIterator struct {
	Event *ERC1155TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC1155TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TokenApprovalForAll)
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
		it.Event = new(ERC1155TokenApprovalForAll)
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
func (it *ERC1155TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TokenApprovalForAll represents a ApprovalForAll event raised by the ERC1155Token contract.
type ERC1155TokenApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155Token *ERC1155TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155TokenApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155Token.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenApprovalForAllIterator{contract: _ERC1155Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155Token *ERC1155TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155TokenApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155Token.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TokenApprovalForAll)
				if err := _ERC1155Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155Token *ERC1155TokenFilterer) ParseApprovalForAll(log types.Log) (*ERC1155TokenApprovalForAll, error) {
	event := new(ERC1155TokenApprovalForAll)
	if err := _ERC1155Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155Token contract.
type ERC1155TokenTransferBatchIterator struct {
	Event *ERC1155TokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *ERC1155TokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TokenTransferBatch)
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
		it.Event = new(ERC1155TokenTransferBatch)
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
func (it *ERC1155TokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TokenTransferBatch represents a TransferBatch event raised by the ERC1155Token contract.
type ERC1155TokenTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155Token *ERC1155TokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TokenTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Token.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenTransferBatchIterator{contract: _ERC1155Token.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155Token *ERC1155TokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155TokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Token.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TokenTransferBatch)
				if err := _ERC1155Token.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155Token *ERC1155TokenFilterer) ParseTransferBatch(log types.Log) (*ERC1155TokenTransferBatch, error) {
	event := new(ERC1155TokenTransferBatch)
	if err := _ERC1155Token.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155Token contract.
type ERC1155TokenTransferSingleIterator struct {
	Event *ERC1155TokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *ERC1155TokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TokenTransferSingle)
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
		it.Event = new(ERC1155TokenTransferSingle)
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
func (it *ERC1155TokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TokenTransferSingle represents a TransferSingle event raised by the ERC1155Token contract.
type ERC1155TokenTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155Token *ERC1155TokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TokenTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Token.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenTransferSingleIterator{contract: _ERC1155Token.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155Token *ERC1155TokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155TokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155Token.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TokenTransferSingle)
				if err := _ERC1155Token.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155Token *ERC1155TokenFilterer) ParseTransferSingle(log types.Log) (*ERC1155TokenTransferSingle, error) {
	event := new(ERC1155TokenTransferSingle)
	if err := _ERC1155Token.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155Token contract.
type ERC1155TokenURIIterator struct {
	Event *ERC1155TokenURI // Event containing the contract specifics and raw log

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
func (it *ERC1155TokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TokenURI)
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
		it.Event = new(ERC1155TokenURI)
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
func (it *ERC1155TokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TokenURI represents a URI event raised by the ERC1155Token contract.
type ERC1155TokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155Token *ERC1155TokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155TokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155Token.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TokenURIIterator{contract: _ERC1155Token.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155Token *ERC1155TokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155TokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155Token.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TokenURI)
				if err := _ERC1155Token.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155Token *ERC1155TokenFilterer) ParseURI(log types.Log) (*ERC1155TokenURI, error) {
	event := new(ERC1155TokenURI)
	if err := _ERC1155Token.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenMetaData contains all meta data concerning the ERC20Token contract.
var ERC20TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600981526020017f54657374546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f545400000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200009692919062000359565b508060049080519060200190620000af92919062000359565b506012600560006101000a81548160ff021916908360ff1602179055505050620000e73366038d7ea4c68000620000ed60201b60201c565b6200040f565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000191576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b620001a560008383620002cb60201b60201c565b620001c181600254620002d060201b620009a01790919060201c565b6002819055506200021f816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620002d060201b620009a01790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b6000808284019050838110156200034f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620003915760008555620003dd565b82601f10620003ac57805160ff1916838001178555620003dd565b82800160010185558215620003dd579182015b82811115620003dc578251825591602001919060010190620003bf565b5b509050620003ec9190620003f0565b5090565b5b808211156200040b576000816000905550600101620003f1565b5090565b6110de806200041f6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461025857806370a08231146102bc57806395d89b4114610314578063a457c2d714610397578063a9059cbb146103fb578063dd62ed3e1461045f576100a9565b806306fdde03146100ae578063095ea7b31461013157806318160ddd1461019557806323b872dd146101b3578063313ce56714610237575b600080fd5b6100b66104d7565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f65780820151818401526020810190506100db565b50505050905090810190601f1680156101235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017d6004803603604081101561014757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610579565b60405180821515815260200191505060405180910390f35b61019d610597565b6040518082815260200191505060405180910390f35b61021f600480360360608110156101c957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105a1565b60405180821515815260200191505060405180910390f35b61023f61067a565b604051808260ff16815260200191505060405180910390f35b6102a46004803603604081101561026e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610691565b60405180821515815260200191505060405180910390f35b6102fe600480360360208110156102d257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610744565b6040518082815260200191505060405180910390f35b61031c61078c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561035c578082015181840152602081019050610341565b50505050905090810190601f1680156103895780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103e3600480360360408110156103ad57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061082e565b60405180821515815260200191505060405180910390f35b6104476004803603604081101561041157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108fb565b60405180821515815260200191505060405180910390f35b6104c16004803603604081101561047557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610919565b6040518082815260200191505060405180910390f35b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561056f5780601f106105445761010080835404028352916020019161056f565b820191906000526020600020905b81548152906001019060200180831161055257829003601f168201915b5050505050905090565b600061058d610586610a28565b8484610a30565b6001905092915050565b6000600254905090565b60006105ae848484610c27565b61066f846105ba610a28565b61066a8560405180606001604052806028815260200161101360289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610620610a28565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ee89092919063ffffffff16565b610a30565b600190509392505050565b6000600560009054906101000a900460ff16905090565b600061073a61069e610a28565b8461073585600160006106af610a28565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109a090919063ffffffff16565b610a30565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108245780601f106107f957610100808354040283529160200191610824565b820191906000526020600020905b81548152906001019060200180831161080757829003601f168201915b5050505050905090565b60006108f161083b610a28565b846108ec856040518060600160405280602581526020016110846025913960016000610865610a28565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ee89092919063ffffffff16565b610a30565b6001905092915050565b600061090f610908610a28565b8484610c27565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600080828401905083811015610a1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ab6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806110606024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610fcb6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610cad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061103b6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180610fa86023913960400191505060405180910390fd5b610d3e838383610fa2565b610da981604051806060016040528060268152602001610fed602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ee89092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610e3c816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109a090919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290610f95576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f5a578082015181840152602081019050610f3f565b50505050905090810190601f168015610f875780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390509392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212204470df7321b34c7aff6cd7c89b9c17463284c03ea82b644023944d765a34206964736f6c63430007060033",
}

// ERC20TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenMetaData.ABI instead.
var ERC20TokenABI = ERC20TokenMetaData.ABI

// ERC20TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenMetaData.Bin instead.
var ERC20TokenBin = ERC20TokenMetaData.Bin

// DeployERC20Token deploys a new Ethereum contract, binding an instance of ERC20Token to it.
func DeployERC20Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Token, error) {
	parsed, err := ERC20TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Token{ERC20TokenCaller: ERC20TokenCaller{contract: contract}, ERC20TokenTransactor: ERC20TokenTransactor{contract: contract}, ERC20TokenFilterer: ERC20TokenFilterer{contract: contract}}, nil
}

// ERC20Token is an auto generated Go binding around an Ethereum contract.
type ERC20Token struct {
	ERC20TokenCaller     // Read-only binding to the contract
	ERC20TokenTransactor // Write-only binding to the contract
	ERC20TokenFilterer   // Log filterer for contract events
}

// ERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenSession struct {
	Contract     *ERC20Token       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenCallerSession struct {
	Contract *ERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenTransactorSession struct {
	Contract     *ERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenRaw struct {
	Contract *ERC20Token // Generic contract binding to access the raw methods on
}

// ERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenCallerRaw struct {
	Contract *ERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenTransactorRaw struct {
	Contract *ERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Token creates a new instance of ERC20Token, bound to a specific deployed contract.
func NewERC20Token(address common.Address, backend bind.ContractBackend) (*ERC20Token, error) {
	contract, err := bindERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Token{ERC20TokenCaller: ERC20TokenCaller{contract: contract}, ERC20TokenTransactor: ERC20TokenTransactor{contract: contract}, ERC20TokenFilterer: ERC20TokenFilterer{contract: contract}}, nil
}

// NewERC20TokenCaller creates a new read-only instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenCaller, error) {
	contract, err := bindERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenCaller{contract: contract}, nil
}

// NewERC20TokenTransactor creates a new write-only instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenTransactor, error) {
	contract, err := bindERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenTransactor{contract: contract}, nil
}

// NewERC20TokenFilterer creates a new log filterer instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenFilterer, error) {
	contract, err := bindERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenFilterer{contract: contract}, nil
}

// bindERC20Token binds a generic wrapper to an already deployed contract.
func bindERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Token *ERC20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Token.Contract.ERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Token *ERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Token.Contract.ERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Token *ERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Token.Contract.ERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Token *ERC20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Token *ERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Token *ERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Token *ERC20TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Token *ERC20TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.Allowance(&_ERC20Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.Allowance(&_ERC20Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Token *ERC20TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Token *ERC20TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.BalanceOf(&_ERC20Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.BalanceOf(&_ERC20Token.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Token *ERC20TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Token *ERC20TokenSession) Decimals() (uint8, error) {
	return _ERC20Token.Contract.Decimals(&_ERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Token *ERC20TokenCallerSession) Decimals() (uint8, error) {
	return _ERC20Token.Contract.Decimals(&_ERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Token *ERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Token *ERC20TokenSession) Name() (string, error) {
	return _ERC20Token.Contract.Name(&_ERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Token *ERC20TokenCallerSession) Name() (string, error) {
	return _ERC20Token.Contract.Name(&_ERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Token *ERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Token *ERC20TokenSession) Symbol() (string, error) {
	return _ERC20Token.Contract.Symbol(&_ERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Token *ERC20TokenCallerSession) Symbol() (string, error) {
	return _ERC20Token.Contract.Symbol(&_ERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Token *ERC20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Token *ERC20TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC20Token.Contract.TotalSupply(&_ERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Token.Contract.TotalSupply(&_ERC20Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Approve(&_ERC20Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Approve(&_ERC20Token.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Token *ERC20TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.DecreaseAllowance(&_ERC20Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.DecreaseAllowance(&_ERC20Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Token *ERC20TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.IncreaseAllowance(&_ERC20Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.IncreaseAllowance(&_ERC20Token.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Transfer(&_ERC20Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Transfer(&_ERC20Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.TransferFrom(&_ERC20Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.TransferFrom(&_ERC20Token.TransactOpts, sender, recipient, amount)
}

// ERC20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Token contract.
type ERC20TokenApprovalIterator struct {
	Event *ERC20TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenApproval)
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
		it.Event = new(ERC20TokenApproval)
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
func (it *ERC20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenApproval represents a Approval event raised by the ERC20Token contract.
type ERC20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenApprovalIterator{contract: _ERC20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenApproval)
				if err := _ERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) ParseApproval(log types.Log) (*ERC20TokenApproval, error) {
	event := new(ERC20TokenApproval)
	if err := _ERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Token contract.
type ERC20TokenTransferIterator struct {
	Event *ERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenTransfer)
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
		it.Event = new(ERC20TokenTransfer)
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
func (it *ERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenTransfer represents a Transfer event raised by the ERC20Token contract.
type ERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenTransferIterator{contract: _ERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenTransfer)
				if err := _ERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) ParseTransfer(log types.Log) (*ERC20TokenTransfer, error) {
	event := new(ERC20TokenTransfer)
	if err := _ERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TokenMetaData contains all meta data concerning the ERC721Token contract.
var ERC721TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"trickFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600981526020017f54657374546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f5454000000000000000000000000000000000000000000000000000000000000815250620000966301ffc9a760e01b6200011860201b60201c565b8160069080519060200190620000ae92919062000221565b508060079080519060200190620000c792919062000221565b50620000e06380ac58cd60e01b6200011860201b60201c565b620000f8635b5e139f60e01b6200011860201b60201c565b6200011063780e9d6360e01b6200011860201b60201c565b5050620002d7565b63ffffffff60e01b817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415620001b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433136353a20696e76616c696420696e746572666163652069640000000081525060200191505060405180910390fd5b6001600080837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620002595760008555620002a5565b82601f106200027457805160ff1916838001178555620002a5565b82800160010185558215620002a5579182015b82811115620002a457825182559160200191906001019062000287565b5b509050620002b49190620002b8565b5090565b5b80821115620002d3576000816000905550600101620002b9565b5090565b612c4380620002e76000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80635624b25b116100b8578063a22cb4651161007c578063a22cb4651461071b578063b4faba091461076b578063b88d4fde14610846578063c87b56dd1461094b578063dbaabecb146109f2578063e985e9c514610a3657610137565b80635624b25b146104b45780636352211e146105655780636c0360eb146105bd57806370a082311461064057806395d89b411461069857610137565b806323b872dd116100ff57806323b872dd146102e65780632f745c591461035457806340c10f19146103b657806342842e0e146104045780634f6ccce71461047257610137565b806301ffc9a71461013c57806306fdde031461019f578063081812fc14610222578063095ea7b31461027a57806318160ddd146102c8575b600080fd5b6101876004803603602081101561015257600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610ab0565b60405180821515815260200191505060405180910390f35b6101a7610b17565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101e75780820151818401526020810190506101cc565b50505050905090810190601f1680156102145780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61024e6004803603602081101561023857600080fd5b8101908080359060200190929190505050610bb9565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102c66004803603604081101561029057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c54565b005b6102d0610d98565b6040518082815260200191505060405180910390f35b610352600480360360608110156102fc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610da9565b005b6103a06004803603604081101561036a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e1f565b6040518082815260200191505060405180910390f35b610402600480360360408110156103cc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e7a565b005b6104706004803603606081101561041a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e88565b005b61049e6004803603602081101561048857600080fd5b8101908080359060200190929190505050610ea8565b6040518082815260200191505060405180910390f35b6104ea600480360360408110156104ca57600080fd5b810190808035906020019092919080359060200190929190505050610ecb565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561052a57808201518184015260208101905061050f565b50505050905090810190601f1680156105575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6105916004803603602081101561057b57600080fd5b8101908080359060200190929190505050610f51565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105c5610f88565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106055780820151818401526020810190506105ea565b50505050905090810190601f1680156106325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106826004803603602081101561065657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061102a565b6040518082815260200191505060405180910390f35b6106a06110ff565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106e05780820151818401526020810190506106c5565b50505050905090810190601f16801561070d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6107696004803603604081101561073157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506111a1565b005b6108446004803603604081101561078157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156107be57600080fd5b8201836020820111156107d057600080fd5b803590602001918460018302840111640100000000831117156107f257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611357565b005b6109496004803603608081101561085c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156108c357600080fd5b8201836020820111156108d557600080fd5b803590602001918460018302840111640100000000831117156108f757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061137e565b005b6109776004803603602081101561096157600080fd5b81019080803590602001909291905050506113f6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109b757808201518184015260208101905061099c565b50505050905090810190601f1680156109e45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610a3460048036036020811015610a0857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116c7565b005b610a9860048036036040811015610a4c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116ed565b60405180821515815260200191505060405180910390f35b6000806000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050919050565b606060068054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610baf5780601f10610b8457610100808354040283529160200191610baf565b820191906000526020600020905b815481529060010190602001808311610b9257829003601f168201915b5050505050905090565b6000610bc482611781565b610c19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612b38602c913960400191505060405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000610c5f82610f51565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ce6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612bbc6021913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610d0561179e565b73ffffffffffffffffffffffffffffffffffffffff161480610d345750610d3381610d2e61179e565b6116ed565b5b610d89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526038815260200180612a8b6038913960400191505060405180910390fd5b610d9383836117a6565b505050565b6000610da4600261185f565b905090565b610dba610db461179e565b82611874565b610e0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526031815260200180612bdd6031913960400191505060405180910390fd5b610e1a838383611968565b505050565b6000610e7282600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611bab90919063ffffffff16565b905092915050565b610e848282611bc5565b5050565b610ea38383836040518060200160405280600081525061137e565b505050565b600080610ebf836002611db990919063ffffffff16565b50905080915050919050565b60606000600583901b67ffffffffffffffff81118015610eea57600080fd5b506040519080825280601f01601f191660200182016040528015610f1d5781602001600182028036833780820191505090505b50905060005b83811015610f465780850154806020830260208501015250806001019050610f23565b508091505092915050565b6000610f8182604051806060016040528060298152602001612aed602991396002611de59092919063ffffffff16565b9050919050565b606060098054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110205780601f10610ff557610100808354040283529160200191611020565b820191906000526020600020905b81548152906001019060200180831161100357829003601f168201915b5050505050905090565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156110b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612ac3602a913960400191505060405180910390fd5b6110f8600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020611e04565b9050919050565b606060078054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111975780601f1061116c57610100808354040283529160200191611197565b820191906000526020600020905b81548152906001019060200180831161117a57829003601f168201915b5050505050905090565b6111a961179e565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561124a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4552433732313a20617070726f766520746f2063616c6c65720000000000000081525060200191505060405180910390fd5b806005600061125761179e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff1661130461179e565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405180821515815260200191505060405180910390a35050565b600080825160208401855af46040518181523d60208201523d6000604083013e60403d0181fd5b61138f61138961179e565b83611874565b6113e4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526031815260200180612bdd6031913960400191505060405180910390fd5b6113f084848484611e19565b50505050565b606061140182611781565b611456576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180612b8d602f913960400191505060405180910390fd5b6000600860008481526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114ff5780601f106114d4576101008083540402835291602001916114ff565b820191906000526020600020905b8154815290600101906020018083116114e257829003601f168201915b505050505090506000611510610f88565b90506000815114156115265781925050506116c2565b6000825111156115f75780826040516020018083805190602001908083835b602083106115685780518252602082019150602081019050602083039250611545565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083106115b95780518252602082019150602081019050602083039250611596565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052925050506116c2565b8061160185611e8b565b6040516020018083805190602001908083835b602083106116375780518252602082019150602081019050602083039250611614565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083106116885780518252602082019150602081019050602083039250611665565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052925050505b919050565b807f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d55550565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000611797826002611fd290919063ffffffff16565b9050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff1661181983610f51565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600061186d82600001611fec565b9050919050565b600061187f82611781565b6118d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612a5f602c913960400191505060405180910390fd5b60006118df83610f51565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061194e57508373ffffffffffffffffffffffffffffffffffffffff1661193684610bb9565b73ffffffffffffffffffffffffffffffffffffffff16145b8061195f575061195e81856116ed565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff1661198882610f51565b73ffffffffffffffffffffffffffffffffffffffff16146119f4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180612b646029913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612a156024913960400191505060405180910390fd5b611a85838383611ffd565b611a906000826117a6565b611ae181600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061200290919063ffffffff16565b50611b3381600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061201c90919063ffffffff16565b50611b4a818360026120369092919063ffffffff16565b50808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b6000611bba836000018361206b565b60001c905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c68576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4552433732313a206d696e7420746f20746865207a65726f206164647265737381525060200191505060405180910390fd5b611c7181611781565b15611ce4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000081525060200191505060405180910390fd5b611cf060008383611ffd565b611d4181600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061201c90919063ffffffff16565b50611d58818360026120369092919063ffffffff16565b50808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b600080600080611dcc86600001866120ee565b915091508160001c8160001c9350935050509250929050565b6000611df8846000018460001b84612187565b60001c90509392505050565b6000611e128260000161227d565b9050919050565b611e24848484611968565b611e308484848461228e565b611e85576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806129e36032913960400191505060405180910390fd5b50505050565b60606000821415611ed3576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611fcd565b600082905060005b60008214611efd578080600101915050600a8281611ef557fe5b049150611edb565b60008167ffffffffffffffff81118015611f1657600080fd5b506040519080825280601f01601f191660200182016040528015611f495781602001600182028036833780820191505090505b50905060006001830390508593505b60008414611fc557600a8481611f6a57fe5b0660300160f81b82828060019003935081518110611f8457fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8481611fbd57fe5b049350611f58565b819450505050505b919050565b6000611fe4836000018360001b6124a7565b905092915050565b600081600001805490509050919050565b505050565b6000612014836000018360001b6124ca565b905092915050565b600061202e836000018360001b6125b2565b905092915050565b6000612062846000018460001b8473ffffffffffffffffffffffffffffffffffffffff1660001b612622565b90509392505050565b6000818360000180549050116120cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806129c16022913960400191505060405180910390fd5b8260000182815481106120db57fe5b9060005260206000200154905092915050565b60008082846000018054905011612150576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612b166022913960400191505060405180910390fd5b600084600001848154811061216157fe5b906000526020600020906002020190508060000154816001015492509250509250929050565b6000808460010160008581526020019081526020016000205490506000811415839061224e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156122135780820151818401526020810190506121f8565b50505050905090810190601f1680156122405780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5084600001600182038154811061226157fe5b9060005260206000209060020201600101549150509392505050565b600081600001805490509050919050565b60006122af8473ffffffffffffffffffffffffffffffffffffffff166126fe565b6122bc576001905061249f565b600061242663150b7a0260e01b6122d161179e565b888787604051602401808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561235557808201518184015260208101905061233a565b50505050905090810190601f1680156123825780820380516001836020036101000a031916815260200191505b5095505050505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518060600160405280603281526020016129e3603291398773ffffffffffffffffffffffffffffffffffffffff166127119092919063ffffffff16565b9050600081806020019051602081101561243f57600080fd5b8101908080519060200190929190505050905063150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614925050505b949350505050565b600080836001016000848152602001908152602001600020541415905092915050565b600080836001016000848152602001908152602001600020549050600081146125a6576000600182039050600060018660000180549050039050600086600001828154811061251557fe5b906000526020600020015490508087600001848154811061253257fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061256a57fe5b600190038181906000526020600020016000905590558660010160008781526020019081526020016000206000905560019450505050506125ac565b60009150505b92915050565b60006125be8383612729565b61261757826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061261c565b600090505b92915050565b60008084600101600085815260200190815260200160002054905060008114156126c9578460000160405180604001604052808681526020018581525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010155505084600001805490508560010160008681526020019081526020016000208190555060019150506126f7565b828560000160018303815481106126dc57fe5b90600052602060002090600202016001018190555060009150505b9392505050565b600080823b905060008111915050919050565b6060612720848460008561274c565b90509392505050565b600080836001016000848152602001908152602001600020541415905092915050565b6060824710156127a7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612a396026913960400191505060405180910390fd5b6127b0856126fe565b612822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b60208310612871578051825260208201915060208101905060208303925061284e565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146128d3576040519150601f19603f3d011682016040523d82523d6000602084013e6128d8565b606091505b50915091506128e88282866128f4565b92505050949350505050565b60608315612904578290506129b9565b6000835111156129175782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561297e578082015181840152602081019050612963565b50505050905090810190601f1680156129ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b939250505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64734552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c4552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e456e756d657261626c654d61703a20696e646578206f7574206f6620626f756e64734552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a264697066735822122015610f81d530ee4ac0046812900a60b887c6c2fe73de8ba61792211fe001467964736f6c63430007060033",
}

// ERC721TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721TokenMetaData.ABI instead.
var ERC721TokenABI = ERC721TokenMetaData.ABI

// ERC721TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721TokenMetaData.Bin instead.
var ERC721TokenBin = ERC721TokenMetaData.Bin

// DeployERC721Token deploys a new Ethereum contract, binding an instance of ERC721Token to it.
func DeployERC721Token(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Token, error) {
	parsed, err := ERC721TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Token{ERC721TokenCaller: ERC721TokenCaller{contract: contract}, ERC721TokenTransactor: ERC721TokenTransactor{contract: contract}, ERC721TokenFilterer: ERC721TokenFilterer{contract: contract}}, nil
}

// ERC721Token is an auto generated Go binding around an Ethereum contract.
type ERC721Token struct {
	ERC721TokenCaller     // Read-only binding to the contract
	ERC721TokenTransactor // Write-only binding to the contract
	ERC721TokenFilterer   // Log filterer for contract events
}

// ERC721TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721TokenSession struct {
	Contract     *ERC721Token      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721TokenCallerSession struct {
	Contract *ERC721TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC721TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TokenTransactorSession struct {
	Contract     *ERC721TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC721TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721TokenRaw struct {
	Contract *ERC721Token // Generic contract binding to access the raw methods on
}

// ERC721TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721TokenCallerRaw struct {
	Contract *ERC721TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TokenTransactorRaw struct {
	Contract *ERC721TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Token creates a new instance of ERC721Token, bound to a specific deployed contract.
func NewERC721Token(address common.Address, backend bind.ContractBackend) (*ERC721Token, error) {
	contract, err := bindERC721Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Token{ERC721TokenCaller: ERC721TokenCaller{contract: contract}, ERC721TokenTransactor: ERC721TokenTransactor{contract: contract}, ERC721TokenFilterer: ERC721TokenFilterer{contract: contract}}, nil
}

// NewERC721TokenCaller creates a new read-only instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC721TokenCaller, error) {
	contract, err := bindERC721Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenCaller{contract: contract}, nil
}

// NewERC721TokenTransactor creates a new write-only instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721TokenTransactor, error) {
	contract, err := bindERC721Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenTransactor{contract: contract}, nil
}

// NewERC721TokenFilterer creates a new log filterer instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721TokenFilterer, error) {
	contract, err := bindERC721Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenFilterer{contract: contract}, nil
}

// bindERC721Token binds a generic wrapper to an already deployed contract.
func bindERC721Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Token *ERC721TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Token.Contract.ERC721TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Token *ERC721TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Token.Contract.ERC721TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Token *ERC721TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Token.Contract.ERC721TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Token *ERC721TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Token *ERC721TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Token *ERC721TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Token *ERC721TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Token *ERC721TokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Token.Contract.BalanceOf(&_ERC721Token.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Token.Contract.BalanceOf(&_ERC721Token.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ERC721Token *ERC721TokenCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ERC721Token *ERC721TokenSession) BaseURI() (string, error) {
	return _ERC721Token.Contract.BaseURI(&_ERC721Token.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_ERC721Token *ERC721TokenCallerSession) BaseURI() (string, error) {
	return _ERC721Token.Contract.BaseURI(&_ERC721Token.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Token *ERC721TokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Token *ERC721TokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.GetApproved(&_ERC721Token.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Token *ERC721TokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.GetApproved(&_ERC721Token.CallOpts, tokenId)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_ERC721Token *ERC721TokenCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_ERC721Token *ERC721TokenSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _ERC721Token.Contract.GetStorageAt(&_ERC721Token.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_ERC721Token *ERC721TokenCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _ERC721Token.Contract.GetStorageAt(&_ERC721Token.CallOpts, offset, length)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Token *ERC721TokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Token *ERC721TokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Token.Contract.IsApprovedForAll(&_ERC721Token.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Token.Contract.IsApprovedForAll(&_ERC721Token.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Token *ERC721TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Token *ERC721TokenSession) Name() (string, error) {
	return _ERC721Token.Contract.Name(&_ERC721Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Token *ERC721TokenCallerSession) Name() (string, error) {
	return _ERC721Token.Contract.Name(&_ERC721Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Token *ERC721TokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Token *ERC721TokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.OwnerOf(&_ERC721Token.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Token *ERC721TokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.OwnerOf(&_ERC721Token.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Token *ERC721TokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Token *ERC721TokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Token.Contract.SupportsInterface(&_ERC721Token.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Token.Contract.SupportsInterface(&_ERC721Token.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Token *ERC721TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Token *ERC721TokenSession) Symbol() (string, error) {
	return _ERC721Token.Contract.Symbol(&_ERC721Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Token *ERC721TokenCallerSession) Symbol() (string, error) {
	return _ERC721Token.Contract.Symbol(&_ERC721Token.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Token *ERC721TokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenByIndex(&_ERC721Token.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenByIndex(&_ERC721Token.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Token *ERC721TokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenOfOwnerByIndex(&_ERC721Token.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenOfOwnerByIndex(&_ERC721Token.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Token *ERC721TokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Token *ERC721TokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Token.Contract.TokenURI(&_ERC721Token.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Token *ERC721TokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Token.Contract.TokenURI(&_ERC721Token.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Token *ERC721TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC721Token.Contract.TotalSupply(&_ERC721Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Token.Contract.TotalSupply(&_ERC721Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Approve(&_ERC721Token.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Approve(&_ERC721Token.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 token) returns()
func (_ERC721Token *ERC721TokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, token *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "mint", to, token)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 token) returns()
func (_ERC721Token *ERC721TokenSession) Mint(to common.Address, token *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Mint(&_ERC721Token.TransactOpts, to, token)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 token) returns()
func (_ERC721Token *ERC721TokenTransactorSession) Mint(to common.Address, token *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Mint(&_ERC721Token.TransactOpts, to, token)
}

// SafeTransferFrom42842e0e is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenTransactor) SafeTransferFrom42842e0e(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom42842e0e is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenSession) SafeTransferFrom42842e0e(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFrom42842e0e(&_ERC721Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom42842e0e is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SafeTransferFrom42842e0e(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFrom42842e0e(&_ERC721Token.TransactOpts, from, to, tokenId)
}

// SafeTransferFromb88d4fde is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Token *ERC721TokenTransactor) SafeTransferFromb88d4fde(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFromb88d4fde is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Token *ERC721TokenSession) SafeTransferFromb88d4fde(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFromb88d4fde(&_ERC721Token.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFromb88d4fde is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SafeTransferFromb88d4fde(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFromb88d4fde(&_ERC721Token.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Token *ERC721TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Token *ERC721TokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Token.Contract.SetApprovalForAll(&_ERC721Token.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Token.Contract.SetApprovalForAll(&_ERC721Token.TransactOpts, operator, approved)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_ERC721Token *ERC721TokenTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_ERC721Token *ERC721TokenSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SimulateAndRevert(&_ERC721Token.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SimulateAndRevert(&_ERC721Token.TransactOpts, targetContract, calldataPayload)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.TransferFrom(&_ERC721Token.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Token *ERC721TokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.TransferFrom(&_ERC721Token.TransactOpts, from, to, tokenId)
}

// TrickFallbackHandler is a paid mutator transaction binding the contract method 0xdbaabecb.
//
// Solidity: function trickFallbackHandler(address fallbackHandler) returns()
func (_ERC721Token *ERC721TokenTransactor) TrickFallbackHandler(opts *bind.TransactOpts, fallbackHandler common.Address) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "trickFallbackHandler", fallbackHandler)
}

// TrickFallbackHandler is a paid mutator transaction binding the contract method 0xdbaabecb.
//
// Solidity: function trickFallbackHandler(address fallbackHandler) returns()
func (_ERC721Token *ERC721TokenSession) TrickFallbackHandler(fallbackHandler common.Address) (*types.Transaction, error) {
	return _ERC721Token.Contract.TrickFallbackHandler(&_ERC721Token.TransactOpts, fallbackHandler)
}

// TrickFallbackHandler is a paid mutator transaction binding the contract method 0xdbaabecb.
//
// Solidity: function trickFallbackHandler(address fallbackHandler) returns()
func (_ERC721Token *ERC721TokenTransactorSession) TrickFallbackHandler(fallbackHandler common.Address) (*types.Transaction, error) {
	return _ERC721Token.Contract.TrickFallbackHandler(&_ERC721Token.TransactOpts, fallbackHandler)
}

// ERC721TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Token contract.
type ERC721TokenApprovalIterator struct {
	Event *ERC721TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC721TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenApproval)
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
		it.Event = new(ERC721TokenApproval)
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
func (it *ERC721TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenApproval represents a Approval event raised by the ERC721Token contract.
type ERC721TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Token *ERC721TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenApprovalIterator{contract: _ERC721Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Token *ERC721TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721TokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenApproval)
				if err := _ERC721Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Token *ERC721TokenFilterer) ParseApproval(log types.Log) (*ERC721TokenApproval, error) {
	event := new(ERC721TokenApproval)
	if err := _ERC721Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Token contract.
type ERC721TokenApprovalForAllIterator struct {
	Event *ERC721TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenApprovalForAll)
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
		it.Event = new(ERC721TokenApprovalForAll)
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
func (it *ERC721TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenApprovalForAll represents a ApprovalForAll event raised by the ERC721Token contract.
type ERC721TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Token *ERC721TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721TokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenApprovalForAllIterator{contract: _ERC721Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Token *ERC721TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721TokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenApprovalForAll)
				if err := _ERC721Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Token *ERC721TokenFilterer) ParseApprovalForAll(log types.Log) (*ERC721TokenApprovalForAll, error) {
	event := new(ERC721TokenApprovalForAll)
	if err := _ERC721Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Token contract.
type ERC721TokenTransferIterator struct {
	Event *ERC721TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenTransfer)
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
		it.Event = new(ERC721TokenTransfer)
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
func (it *ERC721TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenTransfer represents a Transfer event raised by the ERC721Token contract.
type ERC721TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Token *ERC721TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenTransferIterator{contract: _ERC721Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Token *ERC721TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721TokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenTransfer)
				if err := _ERC721Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Token *ERC721TokenFilterer) ParseTransfer(log types.Log) (*ERC721TokenTransfer, error) {
	event := new(ERC721TokenTransfer)
	if err := _ERC721Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISafeMetaData contains all meta data concerning the ISafe contract.
var ISafeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISafeABI is the input ABI used to generate the binding from.
// Deprecated: Use ISafeMetaData.ABI instead.
var ISafeABI = ISafeMetaData.ABI

// ISafe is an auto generated Go binding around an Ethereum contract.
type ISafe struct {
	ISafeCaller     // Read-only binding to the contract
	ISafeTransactor // Write-only binding to the contract
	ISafeFilterer   // Log filterer for contract events
}

// ISafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISafeSession struct {
	Contract     *ISafe            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISafeCallerSession struct {
	Contract *ISafeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISafeTransactorSession struct {
	Contract     *ISafeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISafeRaw struct {
	Contract *ISafe // Generic contract binding to access the raw methods on
}

// ISafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISafeCallerRaw struct {
	Contract *ISafeCaller // Generic read-only contract binding to access the raw methods on
}

// ISafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISafeTransactorRaw struct {
	Contract *ISafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISafe creates a new instance of ISafe, bound to a specific deployed contract.
func NewISafe(address common.Address, backend bind.ContractBackend) (*ISafe, error) {
	contract, err := bindISafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISafe{ISafeCaller: ISafeCaller{contract: contract}, ISafeTransactor: ISafeTransactor{contract: contract}, ISafeFilterer: ISafeFilterer{contract: contract}}, nil
}

// NewISafeCaller creates a new read-only instance of ISafe, bound to a specific deployed contract.
func NewISafeCaller(address common.Address, caller bind.ContractCaller) (*ISafeCaller, error) {
	contract, err := bindISafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISafeCaller{contract: contract}, nil
}

// NewISafeTransactor creates a new write-only instance of ISafe, bound to a specific deployed contract.
func NewISafeTransactor(address common.Address, transactor bind.ContractTransactor) (*ISafeTransactor, error) {
	contract, err := bindISafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISafeTransactor{contract: contract}, nil
}

// NewISafeFilterer creates a new log filterer instance of ISafe, bound to a specific deployed contract.
func NewISafeFilterer(address common.Address, filterer bind.ContractFilterer) (*ISafeFilterer, error) {
	contract, err := bindISafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISafeFilterer{contract: contract}, nil
}

// bindISafe binds a generic wrapper to an already deployed contract.
func bindISafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISafe *ISafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISafe.Contract.ISafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISafe *ISafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISafe.Contract.ISafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISafe *ISafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISafe.Contract.ISafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISafe *ISafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISafe *ISafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISafe *ISafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISafe.Contract.contract.Transact(opts, method, params...)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_ISafe *ISafeTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _ISafe.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_ISafe *ISafeSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _ISafe.Contract.EnableModule(&_ISafe.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_ISafe *ISafeTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _ISafe.Contract.EnableModule(&_ISafe.TransactOpts, module)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_ISafe *ISafeTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ISafe.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_ISafe *ISafeSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ISafe.Contract.ExecTransactionFromModule(&_ISafe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_ISafe *ISafeTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _ISafe.Contract.ExecTransactionFromModule(&_ISafe.TransactOpts, to, value, data, operation)
}

// Test4337ModuleAndHandlerMetaData contains all meta data concerning the Test4337ModuleAndHandler contract.
var Test4337ModuleAndHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entryPointAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ENTRYPOINT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MY_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableMyself\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516109d93803806109d9833981810160405281019061003291906100bb565b8073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250503073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250505061012d565b6000815190506100b581610116565b92915050565b6000602082840312156100cd57600080fd5b60006100db848285016100a6565b91505092915050565b60006100ef826100f6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b61011f816100e4565b811461012a57600080fd5b50565b60805160601c60a05160601c6108796101606000398061017c5280610391525080610117528061023b52506108796000f3fe60806040526004361061004a5760003560e01c80633a756cec1461004f5780633a871cdd1461007a578063a798b2b1146100b7578063ab4ed83e146100ce578063e8eb3cc6146100ea575b600080fd5b34801561005b57600080fd5b50610064610115565b6040516100719190610646565b60405180910390f35b34801561008657600080fd5b506100a1600480360381019061009c919061052b565b610139565b6040516100ae9190610719565b60405180910390f35b3480156100c357600080fd5b506100cc61021d565b005b6100e860048036038101906100e39190610496565b6102aa565b005b3480156100f657600080fd5b506100ff61038f565b60405161010c9190610646565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60008084600001602081019061014f919061046d565b9050600081905060008414610210578073ffffffffffffffffffffffffffffffffffffffff1663468721a77f00000000000000000000000000000000000000000000000000000000000000008660006040518463ffffffff1660e01b81526004016101bc939291906106af565b602060405180830381600087803b1580156101d657600080fd5b505af11580156101ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020e9190610502565b505b6000925050509392505050565b3073ffffffffffffffffffffffffffffffffffffffff1663610b59257f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b81526004016102769190610646565b600060405180830381600087803b15801561029057600080fd5b505af11580156102a4573d6000803e3d6000fd5b50505050565b600033905060008190508073ffffffffffffffffffffffffffffffffffffffff1663468721a78787878760006040518663ffffffff1660e01b81526004016102f6959493929190610661565b602060405180830381600087803b15801561031057600080fd5b505af1158015610324573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103489190610502565b610387576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037e906106f9565b60405180910390fd5b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000813590506103c2816107e7565b92915050565b6000815190506103d7816107fe565b92915050565b6000813590506103ec81610815565b92915050565b60008083601f84011261040457600080fd5b8235905067ffffffffffffffff81111561041d57600080fd5b60208301915083600182028301111561043557600080fd5b9250929050565b6000610160828403121561044f57600080fd5b81905092915050565b6000813590506104678161082c565b92915050565b60006020828403121561047f57600080fd5b600061048d848285016103b3565b91505092915050565b600080600080606085870312156104ac57600080fd5b60006104ba878288016103b3565b94505060206104cb87828801610458565b935050604085013567ffffffffffffffff8111156104e857600080fd5b6104f4878288016103f2565b925092505092959194509250565b60006020828403121561051457600080fd5b6000610522848285016103c8565b91505092915050565b60008060006060848603121561054057600080fd5b600084013567ffffffffffffffff81111561055a57600080fd5b6105668682870161043c565b9350506020610577868287016103dd565b925050604061058886828701610458565b9150509250925092565b61059b81610756565b82525050565b60006105ad8385610734565b93506105ba8385846107c7565b6105c3836107d6565b840190509392505050565b6105d7816107b5565b82525050565b60006105ea600983610745565b91507f7478206661696c656400000000000000000000000000000000000000000000006000830152602082019050919050565b600061062a600083610734565b9150600082019050919050565b6106408161079e565b82525050565b600060208201905061065b6000830184610592565b92915050565b60006080820190506106766000830188610592565b6106836020830187610637565b81810360408301526106968185876105a1565b90506106a560608301846105ce565b9695505050505050565b60006080820190506106c46000830186610592565b6106d16020830185610637565b81810360408301526106e28161061d565b90506106f160608301846105ce565b949350505050565b60006020820190508181036000830152610712816105dd565b9050919050565b600060208201905061072e6000830184610637565b92915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006107618261077e565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006107c0826107a8565b9050919050565b82818337600083830152505050565b6000601f19601f8301169050919050565b6107f081610756565b81146107fb57600080fd5b50565b61080781610768565b811461081257600080fd5b50565b61081e81610774565b811461082957600080fd5b50565b6108358161079e565b811461084057600080fd5b5056fea2646970667358221220060335c4dd303a70e25315c3d37bfa6e78c74996854b5d49b14f44e2f4acf1b264736f6c63430007060033",
}

// Test4337ModuleAndHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use Test4337ModuleAndHandlerMetaData.ABI instead.
var Test4337ModuleAndHandlerABI = Test4337ModuleAndHandlerMetaData.ABI

// Test4337ModuleAndHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Test4337ModuleAndHandlerMetaData.Bin instead.
var Test4337ModuleAndHandlerBin = Test4337ModuleAndHandlerMetaData.Bin

// DeployTest4337ModuleAndHandler deploys a new Ethereum contract, binding an instance of Test4337ModuleAndHandler to it.
func DeployTest4337ModuleAndHandler(auth *bind.TransactOpts, backend bind.ContractBackend, entryPointAddress common.Address) (common.Address, *types.Transaction, *Test4337ModuleAndHandler, error) {
	parsed, err := Test4337ModuleAndHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Test4337ModuleAndHandlerBin), backend, entryPointAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test4337ModuleAndHandler{Test4337ModuleAndHandlerCaller: Test4337ModuleAndHandlerCaller{contract: contract}, Test4337ModuleAndHandlerTransactor: Test4337ModuleAndHandlerTransactor{contract: contract}, Test4337ModuleAndHandlerFilterer: Test4337ModuleAndHandlerFilterer{contract: contract}}, nil
}

// Test4337ModuleAndHandler is an auto generated Go binding around an Ethereum contract.
type Test4337ModuleAndHandler struct {
	Test4337ModuleAndHandlerCaller     // Read-only binding to the contract
	Test4337ModuleAndHandlerTransactor // Write-only binding to the contract
	Test4337ModuleAndHandlerFilterer   // Log filterer for contract events
}

// Test4337ModuleAndHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Test4337ModuleAndHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test4337ModuleAndHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Test4337ModuleAndHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test4337ModuleAndHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Test4337ModuleAndHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test4337ModuleAndHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Test4337ModuleAndHandlerSession struct {
	Contract     *Test4337ModuleAndHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Test4337ModuleAndHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Test4337ModuleAndHandlerCallerSession struct {
	Contract *Test4337ModuleAndHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// Test4337ModuleAndHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Test4337ModuleAndHandlerTransactorSession struct {
	Contract     *Test4337ModuleAndHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// Test4337ModuleAndHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Test4337ModuleAndHandlerRaw struct {
	Contract *Test4337ModuleAndHandler // Generic contract binding to access the raw methods on
}

// Test4337ModuleAndHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Test4337ModuleAndHandlerCallerRaw struct {
	Contract *Test4337ModuleAndHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// Test4337ModuleAndHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Test4337ModuleAndHandlerTransactorRaw struct {
	Contract *Test4337ModuleAndHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest4337ModuleAndHandler creates a new instance of Test4337ModuleAndHandler, bound to a specific deployed contract.
func NewTest4337ModuleAndHandler(address common.Address, backend bind.ContractBackend) (*Test4337ModuleAndHandler, error) {
	contract, err := bindTest4337ModuleAndHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test4337ModuleAndHandler{Test4337ModuleAndHandlerCaller: Test4337ModuleAndHandlerCaller{contract: contract}, Test4337ModuleAndHandlerTransactor: Test4337ModuleAndHandlerTransactor{contract: contract}, Test4337ModuleAndHandlerFilterer: Test4337ModuleAndHandlerFilterer{contract: contract}}, nil
}

// NewTest4337ModuleAndHandlerCaller creates a new read-only instance of Test4337ModuleAndHandler, bound to a specific deployed contract.
func NewTest4337ModuleAndHandlerCaller(address common.Address, caller bind.ContractCaller) (*Test4337ModuleAndHandlerCaller, error) {
	contract, err := bindTest4337ModuleAndHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Test4337ModuleAndHandlerCaller{contract: contract}, nil
}

// NewTest4337ModuleAndHandlerTransactor creates a new write-only instance of Test4337ModuleAndHandler, bound to a specific deployed contract.
func NewTest4337ModuleAndHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*Test4337ModuleAndHandlerTransactor, error) {
	contract, err := bindTest4337ModuleAndHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Test4337ModuleAndHandlerTransactor{contract: contract}, nil
}

// NewTest4337ModuleAndHandlerFilterer creates a new log filterer instance of Test4337ModuleAndHandler, bound to a specific deployed contract.
func NewTest4337ModuleAndHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*Test4337ModuleAndHandlerFilterer, error) {
	contract, err := bindTest4337ModuleAndHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Test4337ModuleAndHandlerFilterer{contract: contract}, nil
}

// bindTest4337ModuleAndHandler binds a generic wrapper to an already deployed contract.
func bindTest4337ModuleAndHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Test4337ModuleAndHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test4337ModuleAndHandler.Contract.Test4337ModuleAndHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.Test4337ModuleAndHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.Test4337ModuleAndHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test4337ModuleAndHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.contract.Transact(opts, method, params...)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerCaller) ENTRYPOINT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Test4337ModuleAndHandler.contract.Call(opts, &out, "ENTRYPOINT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerSession) ENTRYPOINT() (common.Address, error) {
	return _Test4337ModuleAndHandler.Contract.ENTRYPOINT(&_Test4337ModuleAndHandler.CallOpts)
}

// ENTRYPOINT is a free data retrieval call binding the contract method 0xe8eb3cc6.
//
// Solidity: function ENTRYPOINT() view returns(address)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerCallerSession) ENTRYPOINT() (common.Address, error) {
	return _Test4337ModuleAndHandler.Contract.ENTRYPOINT(&_Test4337ModuleAndHandler.CallOpts)
}

// MYADDRESS is a free data retrieval call binding the contract method 0x3a756cec.
//
// Solidity: function MY_ADDRESS() view returns(address)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerCaller) MYADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Test4337ModuleAndHandler.contract.Call(opts, &out, "MY_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MYADDRESS is a free data retrieval call binding the contract method 0x3a756cec.
//
// Solidity: function MY_ADDRESS() view returns(address)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerSession) MYADDRESS() (common.Address, error) {
	return _Test4337ModuleAndHandler.Contract.MYADDRESS(&_Test4337ModuleAndHandler.CallOpts)
}

// MYADDRESS is a free data retrieval call binding the contract method 0x3a756cec.
//
// Solidity: function MY_ADDRESS() view returns(address)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerCallerSession) MYADDRESS() (common.Address, error) {
	return _Test4337ModuleAndHandler.Contract.MYADDRESS(&_Test4337ModuleAndHandler.CallOpts)
}

// EnableMyself is a paid mutator transaction binding the contract method 0xa798b2b1.
//
// Solidity: function enableMyself() returns()
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactor) EnableMyself(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.contract.Transact(opts, "enableMyself")
}

// EnableMyself is a paid mutator transaction binding the contract method 0xa798b2b1.
//
// Solidity: function enableMyself() returns()
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerSession) EnableMyself() (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.EnableMyself(&_Test4337ModuleAndHandler.TransactOpts)
}

// EnableMyself is a paid mutator transaction binding the contract method 0xa798b2b1.
//
// Solidity: function enableMyself() returns()
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactorSession) EnableMyself() (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.EnableMyself(&_Test4337ModuleAndHandler.TransactOpts)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0xab4ed83e.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data) payable returns()
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.contract.Transact(opts, "execTransaction", to, value, data)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0xab4ed83e.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data) payable returns()
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerSession) ExecTransaction(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.ExecTransaction(&_Test4337ModuleAndHandler.TransactOpts, to, value, data)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0xab4ed83e.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data) payable returns()
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.ExecTransaction(&_Test4337ModuleAndHandler.TransactOpts, to, value, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.contract.Transact(opts, "validateUserOp", userOp, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerSession) ValidateUserOp(userOp UserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.ValidateUserOp(&_Test4337ModuleAndHandler.TransactOpts, userOp, arg1, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 missingAccountFunds) returns(uint256 validationData)
func (_Test4337ModuleAndHandler *Test4337ModuleAndHandlerTransactorSession) ValidateUserOp(userOp UserOperation, arg1 [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Test4337ModuleAndHandler.Contract.ValidateUserOp(&_Test4337ModuleAndHandler.TransactOpts, userOp, arg1, missingAccountFunds)
}

// TestHandlerMetaData contains all meta data concerning the TestHandler contract.
var TestHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"dudududu\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610164806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806354955e5914610030575b600080fd5b610038610081565b604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b60008061008c61009c565b610094610126565b915091509091565b6000601460003690501015610119576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f496e76616c69642063616c6c64617461206c656e67746800000000000000000081525060200191505060405180910390fd5b601436033560601c905090565b60003390509056fea2646970667358221220128e278302c66b69b8e8ebb334d77f3e0b96f4b9a78c6a86be6eda5481bf811164736f6c63430007060033",
}

// TestHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use TestHandlerMetaData.ABI instead.
var TestHandlerABI = TestHandlerMetaData.ABI

// TestHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestHandlerMetaData.Bin instead.
var TestHandlerBin = TestHandlerMetaData.Bin

// DeployTestHandler deploys a new Ethereum contract, binding an instance of TestHandler to it.
func DeployTestHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestHandler, error) {
	parsed, err := TestHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestHandler{TestHandlerCaller: TestHandlerCaller{contract: contract}, TestHandlerTransactor: TestHandlerTransactor{contract: contract}, TestHandlerFilterer: TestHandlerFilterer{contract: contract}}, nil
}

// TestHandler is an auto generated Go binding around an Ethereum contract.
type TestHandler struct {
	TestHandlerCaller     // Read-only binding to the contract
	TestHandlerTransactor // Write-only binding to the contract
	TestHandlerFilterer   // Log filterer for contract events
}

// TestHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestHandlerSession struct {
	Contract     *TestHandler      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestHandlerCallerSession struct {
	Contract *TestHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestHandlerTransactorSession struct {
	Contract     *TestHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestHandlerRaw struct {
	Contract *TestHandler // Generic contract binding to access the raw methods on
}

// TestHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestHandlerCallerRaw struct {
	Contract *TestHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// TestHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestHandlerTransactorRaw struct {
	Contract *TestHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestHandler creates a new instance of TestHandler, bound to a specific deployed contract.
func NewTestHandler(address common.Address, backend bind.ContractBackend) (*TestHandler, error) {
	contract, err := bindTestHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestHandler{TestHandlerCaller: TestHandlerCaller{contract: contract}, TestHandlerTransactor: TestHandlerTransactor{contract: contract}, TestHandlerFilterer: TestHandlerFilterer{contract: contract}}, nil
}

// NewTestHandlerCaller creates a new read-only instance of TestHandler, bound to a specific deployed contract.
func NewTestHandlerCaller(address common.Address, caller bind.ContractCaller) (*TestHandlerCaller, error) {
	contract, err := bindTestHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestHandlerCaller{contract: contract}, nil
}

// NewTestHandlerTransactor creates a new write-only instance of TestHandler, bound to a specific deployed contract.
func NewTestHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*TestHandlerTransactor, error) {
	contract, err := bindTestHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestHandlerTransactor{contract: contract}, nil
}

// NewTestHandlerFilterer creates a new log filterer instance of TestHandler, bound to a specific deployed contract.
func NewTestHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*TestHandlerFilterer, error) {
	contract, err := bindTestHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestHandlerFilterer{contract: contract}, nil
}

// bindTestHandler binds a generic wrapper to an already deployed contract.
func bindTestHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestHandler *TestHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestHandler.Contract.TestHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestHandler *TestHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestHandler.Contract.TestHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestHandler *TestHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestHandler.Contract.TestHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestHandler *TestHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestHandler *TestHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestHandler *TestHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestHandler.Contract.contract.Transact(opts, method, params...)
}

// Dudududu is a free data retrieval call binding the contract method 0x54955e59.
//
// Solidity: function dudududu() view returns(address sender, address manager)
func (_TestHandler *TestHandlerCaller) Dudududu(opts *bind.CallOpts) (struct {
	Sender  common.Address
	Manager common.Address
}, error) {
	var out []interface{}
	err := _TestHandler.contract.Call(opts, &out, "dudududu")

	outstruct := new(struct {
		Sender  common.Address
		Manager common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Manager = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Dudududu is a free data retrieval call binding the contract method 0x54955e59.
//
// Solidity: function dudududu() view returns(address sender, address manager)
func (_TestHandler *TestHandlerSession) Dudududu() (struct {
	Sender  common.Address
	Manager common.Address
}, error) {
	return _TestHandler.Contract.Dudududu(&_TestHandler.CallOpts)
}

// Dudududu is a free data retrieval call binding the contract method 0x54955e59.
//
// Solidity: function dudududu() view returns(address sender, address manager)
func (_TestHandler *TestHandlerCallerSession) Dudududu() (struct {
	Sender  common.Address
	Manager common.Address
}, error) {
	return _TestHandler.Contract.Dudududu(&_TestHandler.CallOpts)
}

// TestMarshalLibMetaData contains all meta data concerning the TestMarshalLib contract.
var TestMarshalLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"decode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStatic\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"decodeWithSelector\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStatic\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isStatic\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"encode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isStatic\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"encodeWithSelector\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103cd806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063508dd39f14610051578063b9148704146100d8578063be2b620e1461013c578063c3748ef1146101c9575b600080fd5b61007d6004803603602081101561006757600080fd5b810190808035906020019092919050505061022a565b604051808415158152602001837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018273ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390f35b610126600480360360408110156100ee57600080fd5b81019080803515159060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610245565b6040518082815260200191505060405180910390f35b6101b36004803603606081101561015257600080fd5b8101908080351515906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610259565b6040518082815260200191505060405180910390f35b6101f5600480360360208110156101df57600080fd5b810190808035906020019092919050505061026f565b6040518083151581526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b600080600061023884610284565b9250925092509193909250565b600061025183836102b9565b905092915050565b600061026684848461030c565b90509392505050565b60008061027b8361036f565b91509150915091565b60008060008360f81c15925073ffffffffffffffffffffffffffffffffffffffff841690508360a01c60a81b91509193909250565b6000826102e6577f01000000000000000000000000000000000000000000000000000000000000006102e9565b60005b8273ffffffffffffffffffffffffffffffffffffffff161760001b905092915050565b600060d88360e01c63ffffffff16901b84610347577f010000000000000000000000000000000000000000000000000000000000000061034a565b60005b8373ffffffffffffffffffffffffffffffffffffffff16171760001b90509392505050565b6000808260f81c15915073ffffffffffffffffffffffffffffffffffffffff8316905091509156fea2646970667358221220d4c11780e0aa930ecff4fb19cc638c47a85055cd849f9593ba2c04ec495dd07364736f6c63430007060033",
}

// TestMarshalLibABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMarshalLibMetaData.ABI instead.
var TestMarshalLibABI = TestMarshalLibMetaData.ABI

// TestMarshalLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMarshalLibMetaData.Bin instead.
var TestMarshalLibBin = TestMarshalLibMetaData.Bin

// DeployTestMarshalLib deploys a new Ethereum contract, binding an instance of TestMarshalLib to it.
func DeployTestMarshalLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestMarshalLib, error) {
	parsed, err := TestMarshalLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestMarshalLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestMarshalLib{TestMarshalLibCaller: TestMarshalLibCaller{contract: contract}, TestMarshalLibTransactor: TestMarshalLibTransactor{contract: contract}, TestMarshalLibFilterer: TestMarshalLibFilterer{contract: contract}}, nil
}

// TestMarshalLib is an auto generated Go binding around an Ethereum contract.
type TestMarshalLib struct {
	TestMarshalLibCaller     // Read-only binding to the contract
	TestMarshalLibTransactor // Write-only binding to the contract
	TestMarshalLibFilterer   // Log filterer for contract events
}

// TestMarshalLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestMarshalLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMarshalLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestMarshalLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMarshalLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestMarshalLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMarshalLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestMarshalLibSession struct {
	Contract     *TestMarshalLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestMarshalLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestMarshalLibCallerSession struct {
	Contract *TestMarshalLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TestMarshalLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestMarshalLibTransactorSession struct {
	Contract     *TestMarshalLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TestMarshalLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestMarshalLibRaw struct {
	Contract *TestMarshalLib // Generic contract binding to access the raw methods on
}

// TestMarshalLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestMarshalLibCallerRaw struct {
	Contract *TestMarshalLibCaller // Generic read-only contract binding to access the raw methods on
}

// TestMarshalLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestMarshalLibTransactorRaw struct {
	Contract *TestMarshalLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestMarshalLib creates a new instance of TestMarshalLib, bound to a specific deployed contract.
func NewTestMarshalLib(address common.Address, backend bind.ContractBackend) (*TestMarshalLib, error) {
	contract, err := bindTestMarshalLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestMarshalLib{TestMarshalLibCaller: TestMarshalLibCaller{contract: contract}, TestMarshalLibTransactor: TestMarshalLibTransactor{contract: contract}, TestMarshalLibFilterer: TestMarshalLibFilterer{contract: contract}}, nil
}

// NewTestMarshalLibCaller creates a new read-only instance of TestMarshalLib, bound to a specific deployed contract.
func NewTestMarshalLibCaller(address common.Address, caller bind.ContractCaller) (*TestMarshalLibCaller, error) {
	contract, err := bindTestMarshalLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestMarshalLibCaller{contract: contract}, nil
}

// NewTestMarshalLibTransactor creates a new write-only instance of TestMarshalLib, bound to a specific deployed contract.
func NewTestMarshalLibTransactor(address common.Address, transactor bind.ContractTransactor) (*TestMarshalLibTransactor, error) {
	contract, err := bindTestMarshalLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestMarshalLibTransactor{contract: contract}, nil
}

// NewTestMarshalLibFilterer creates a new log filterer instance of TestMarshalLib, bound to a specific deployed contract.
func NewTestMarshalLibFilterer(address common.Address, filterer bind.ContractFilterer) (*TestMarshalLibFilterer, error) {
	contract, err := bindTestMarshalLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestMarshalLibFilterer{contract: contract}, nil
}

// bindTestMarshalLib binds a generic wrapper to an already deployed contract.
func bindTestMarshalLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestMarshalLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMarshalLib *TestMarshalLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMarshalLib.Contract.TestMarshalLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMarshalLib *TestMarshalLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMarshalLib.Contract.TestMarshalLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMarshalLib *TestMarshalLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMarshalLib.Contract.TestMarshalLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMarshalLib *TestMarshalLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMarshalLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMarshalLib *TestMarshalLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMarshalLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMarshalLib *TestMarshalLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMarshalLib.Contract.contract.Transact(opts, method, params...)
}

// Decode is a free data retrieval call binding the contract method 0xc3748ef1.
//
// Solidity: function decode(bytes32 data) pure returns(bool isStatic, address handler)
func (_TestMarshalLib *TestMarshalLibCaller) Decode(opts *bind.CallOpts, data [32]byte) (struct {
	IsStatic bool
	Handler  common.Address
}, error) {
	var out []interface{}
	err := _TestMarshalLib.contract.Call(opts, &out, "decode", data)

	outstruct := new(struct {
		IsStatic bool
		Handler  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStatic = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Handler = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Decode is a free data retrieval call binding the contract method 0xc3748ef1.
//
// Solidity: function decode(bytes32 data) pure returns(bool isStatic, address handler)
func (_TestMarshalLib *TestMarshalLibSession) Decode(data [32]byte) (struct {
	IsStatic bool
	Handler  common.Address
}, error) {
	return _TestMarshalLib.Contract.Decode(&_TestMarshalLib.CallOpts, data)
}

// Decode is a free data retrieval call binding the contract method 0xc3748ef1.
//
// Solidity: function decode(bytes32 data) pure returns(bool isStatic, address handler)
func (_TestMarshalLib *TestMarshalLibCallerSession) Decode(data [32]byte) (struct {
	IsStatic bool
	Handler  common.Address
}, error) {
	return _TestMarshalLib.Contract.Decode(&_TestMarshalLib.CallOpts, data)
}

// DecodeWithSelector is a free data retrieval call binding the contract method 0x508dd39f.
//
// Solidity: function decodeWithSelector(bytes32 data) pure returns(bool isStatic, bytes4 selector, address handler)
func (_TestMarshalLib *TestMarshalLibCaller) DecodeWithSelector(opts *bind.CallOpts, data [32]byte) (struct {
	IsStatic bool
	Selector [4]byte
	Handler  common.Address
}, error) {
	var out []interface{}
	err := _TestMarshalLib.contract.Call(opts, &out, "decodeWithSelector", data)

	outstruct := new(struct {
		IsStatic bool
		Selector [4]byte
		Handler  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStatic = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Selector = *abi.ConvertType(out[1], new([4]byte)).(*[4]byte)
	outstruct.Handler = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DecodeWithSelector is a free data retrieval call binding the contract method 0x508dd39f.
//
// Solidity: function decodeWithSelector(bytes32 data) pure returns(bool isStatic, bytes4 selector, address handler)
func (_TestMarshalLib *TestMarshalLibSession) DecodeWithSelector(data [32]byte) (struct {
	IsStatic bool
	Selector [4]byte
	Handler  common.Address
}, error) {
	return _TestMarshalLib.Contract.DecodeWithSelector(&_TestMarshalLib.CallOpts, data)
}

// DecodeWithSelector is a free data retrieval call binding the contract method 0x508dd39f.
//
// Solidity: function decodeWithSelector(bytes32 data) pure returns(bool isStatic, bytes4 selector, address handler)
func (_TestMarshalLib *TestMarshalLibCallerSession) DecodeWithSelector(data [32]byte) (struct {
	IsStatic bool
	Selector [4]byte
	Handler  common.Address
}, error) {
	return _TestMarshalLib.Contract.DecodeWithSelector(&_TestMarshalLib.CallOpts, data)
}

// Encode is a free data retrieval call binding the contract method 0xb9148704.
//
// Solidity: function encode(bool isStatic, address handler) pure returns(bytes32 data)
func (_TestMarshalLib *TestMarshalLibCaller) Encode(opts *bind.CallOpts, isStatic bool, handler common.Address) ([32]byte, error) {
	var out []interface{}
	err := _TestMarshalLib.contract.Call(opts, &out, "encode", isStatic, handler)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Encode is a free data retrieval call binding the contract method 0xb9148704.
//
// Solidity: function encode(bool isStatic, address handler) pure returns(bytes32 data)
func (_TestMarshalLib *TestMarshalLibSession) Encode(isStatic bool, handler common.Address) ([32]byte, error) {
	return _TestMarshalLib.Contract.Encode(&_TestMarshalLib.CallOpts, isStatic, handler)
}

// Encode is a free data retrieval call binding the contract method 0xb9148704.
//
// Solidity: function encode(bool isStatic, address handler) pure returns(bytes32 data)
func (_TestMarshalLib *TestMarshalLibCallerSession) Encode(isStatic bool, handler common.Address) ([32]byte, error) {
	return _TestMarshalLib.Contract.Encode(&_TestMarshalLib.CallOpts, isStatic, handler)
}

// EncodeWithSelector is a free data retrieval call binding the contract method 0xbe2b620e.
//
// Solidity: function encodeWithSelector(bool isStatic, bytes4 selector, address handler) pure returns(bytes32 data)
func (_TestMarshalLib *TestMarshalLibCaller) EncodeWithSelector(opts *bind.CallOpts, isStatic bool, selector [4]byte, handler common.Address) ([32]byte, error) {
	var out []interface{}
	err := _TestMarshalLib.contract.Call(opts, &out, "encodeWithSelector", isStatic, selector, handler)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeWithSelector is a free data retrieval call binding the contract method 0xbe2b620e.
//
// Solidity: function encodeWithSelector(bool isStatic, bytes4 selector, address handler) pure returns(bytes32 data)
func (_TestMarshalLib *TestMarshalLibSession) EncodeWithSelector(isStatic bool, selector [4]byte, handler common.Address) ([32]byte, error) {
	return _TestMarshalLib.Contract.EncodeWithSelector(&_TestMarshalLib.CallOpts, isStatic, selector, handler)
}

// EncodeWithSelector is a free data retrieval call binding the contract method 0xbe2b620e.
//
// Solidity: function encodeWithSelector(bool isStatic, bytes4 selector, address handler) pure returns(bytes32 data)
func (_TestMarshalLib *TestMarshalLibCallerSession) EncodeWithSelector(isStatic bool, selector [4]byte, handler common.Address) ([32]byte, error) {
	return _TestMarshalLib.Contract.EncodeWithSelector(&_TestMarshalLib.CallOpts, isStatic, selector, handler)
}

// TestNativeTokenReceiverMetaData contains all meta data concerning the TestNativeTokenReceiver contract.
var TestNativeTokenReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forwardedGas\",\"type\":\"uint256\"}],\"name\":\"BreadReceived\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50609280601d6000396000f3fe60806040523373ffffffffffffffffffffffffffffffffffffffff167f16549311ba52796916987df5401f791fb06b998524a5a8684010010415850bb3345a604051808381526020018281526020019250505060405180910390a200fea264697066735822122035663a4184b682e3d2c1649228db3273b6a2439d885e4203ca9ef996501e7b4c64736f6c63430007060033",
}

// TestNativeTokenReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use TestNativeTokenReceiverMetaData.ABI instead.
var TestNativeTokenReceiverABI = TestNativeTokenReceiverMetaData.ABI

// TestNativeTokenReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestNativeTokenReceiverMetaData.Bin instead.
var TestNativeTokenReceiverBin = TestNativeTokenReceiverMetaData.Bin

// DeployTestNativeTokenReceiver deploys a new Ethereum contract, binding an instance of TestNativeTokenReceiver to it.
func DeployTestNativeTokenReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestNativeTokenReceiver, error) {
	parsed, err := TestNativeTokenReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestNativeTokenReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestNativeTokenReceiver{TestNativeTokenReceiverCaller: TestNativeTokenReceiverCaller{contract: contract}, TestNativeTokenReceiverTransactor: TestNativeTokenReceiverTransactor{contract: contract}, TestNativeTokenReceiverFilterer: TestNativeTokenReceiverFilterer{contract: contract}}, nil
}

// TestNativeTokenReceiver is an auto generated Go binding around an Ethereum contract.
type TestNativeTokenReceiver struct {
	TestNativeTokenReceiverCaller     // Read-only binding to the contract
	TestNativeTokenReceiverTransactor // Write-only binding to the contract
	TestNativeTokenReceiverFilterer   // Log filterer for contract events
}

// TestNativeTokenReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestNativeTokenReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestNativeTokenReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestNativeTokenReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestNativeTokenReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestNativeTokenReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestNativeTokenReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestNativeTokenReceiverSession struct {
	Contract     *TestNativeTokenReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestNativeTokenReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestNativeTokenReceiverCallerSession struct {
	Contract *TestNativeTokenReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TestNativeTokenReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestNativeTokenReceiverTransactorSession struct {
	Contract     *TestNativeTokenReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TestNativeTokenReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestNativeTokenReceiverRaw struct {
	Contract *TestNativeTokenReceiver // Generic contract binding to access the raw methods on
}

// TestNativeTokenReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestNativeTokenReceiverCallerRaw struct {
	Contract *TestNativeTokenReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// TestNativeTokenReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestNativeTokenReceiverTransactorRaw struct {
	Contract *TestNativeTokenReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestNativeTokenReceiver creates a new instance of TestNativeTokenReceiver, bound to a specific deployed contract.
func NewTestNativeTokenReceiver(address common.Address, backend bind.ContractBackend) (*TestNativeTokenReceiver, error) {
	contract, err := bindTestNativeTokenReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestNativeTokenReceiver{TestNativeTokenReceiverCaller: TestNativeTokenReceiverCaller{contract: contract}, TestNativeTokenReceiverTransactor: TestNativeTokenReceiverTransactor{contract: contract}, TestNativeTokenReceiverFilterer: TestNativeTokenReceiverFilterer{contract: contract}}, nil
}

// NewTestNativeTokenReceiverCaller creates a new read-only instance of TestNativeTokenReceiver, bound to a specific deployed contract.
func NewTestNativeTokenReceiverCaller(address common.Address, caller bind.ContractCaller) (*TestNativeTokenReceiverCaller, error) {
	contract, err := bindTestNativeTokenReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestNativeTokenReceiverCaller{contract: contract}, nil
}

// NewTestNativeTokenReceiverTransactor creates a new write-only instance of TestNativeTokenReceiver, bound to a specific deployed contract.
func NewTestNativeTokenReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*TestNativeTokenReceiverTransactor, error) {
	contract, err := bindTestNativeTokenReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestNativeTokenReceiverTransactor{contract: contract}, nil
}

// NewTestNativeTokenReceiverFilterer creates a new log filterer instance of TestNativeTokenReceiver, bound to a specific deployed contract.
func NewTestNativeTokenReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*TestNativeTokenReceiverFilterer, error) {
	contract, err := bindTestNativeTokenReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestNativeTokenReceiverFilterer{contract: contract}, nil
}

// bindTestNativeTokenReceiver binds a generic wrapper to an already deployed contract.
func bindTestNativeTokenReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestNativeTokenReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestNativeTokenReceiver *TestNativeTokenReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestNativeTokenReceiver.Contract.TestNativeTokenReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestNativeTokenReceiver *TestNativeTokenReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.Contract.TestNativeTokenReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestNativeTokenReceiver *TestNativeTokenReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.Contract.TestNativeTokenReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestNativeTokenReceiver *TestNativeTokenReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestNativeTokenReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestNativeTokenReceiver *TestNativeTokenReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestNativeTokenReceiver *TestNativeTokenReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestNativeTokenReceiver *TestNativeTokenReceiverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestNativeTokenReceiver *TestNativeTokenReceiverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.Contract.Fallback(&_TestNativeTokenReceiver.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestNativeTokenReceiver *TestNativeTokenReceiverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestNativeTokenReceiver.Contract.Fallback(&_TestNativeTokenReceiver.TransactOpts, calldata)
}

// TestNativeTokenReceiverBreadReceivedIterator is returned from FilterBreadReceived and is used to iterate over the raw logs and unpacked data for BreadReceived events raised by the TestNativeTokenReceiver contract.
type TestNativeTokenReceiverBreadReceivedIterator struct {
	Event *TestNativeTokenReceiverBreadReceived // Event containing the contract specifics and raw log

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
func (it *TestNativeTokenReceiverBreadReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestNativeTokenReceiverBreadReceived)
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
		it.Event = new(TestNativeTokenReceiverBreadReceived)
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
func (it *TestNativeTokenReceiverBreadReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestNativeTokenReceiverBreadReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestNativeTokenReceiverBreadReceived represents a BreadReceived event raised by the TestNativeTokenReceiver contract.
type TestNativeTokenReceiverBreadReceived struct {
	From         common.Address
	Amount       *big.Int
	ForwardedGas *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBreadReceived is a free log retrieval operation binding the contract event 0x16549311ba52796916987df5401f791fb06b998524a5a8684010010415850bb3.
//
// Solidity: event BreadReceived(address indexed from, uint256 amount, uint256 forwardedGas)
func (_TestNativeTokenReceiver *TestNativeTokenReceiverFilterer) FilterBreadReceived(opts *bind.FilterOpts, from []common.Address) (*TestNativeTokenReceiverBreadReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TestNativeTokenReceiver.contract.FilterLogs(opts, "BreadReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &TestNativeTokenReceiverBreadReceivedIterator{contract: _TestNativeTokenReceiver.contract, event: "BreadReceived", logs: logs, sub: sub}, nil
}

// WatchBreadReceived is a free log subscription operation binding the contract event 0x16549311ba52796916987df5401f791fb06b998524a5a8684010010415850bb3.
//
// Solidity: event BreadReceived(address indexed from, uint256 amount, uint256 forwardedGas)
func (_TestNativeTokenReceiver *TestNativeTokenReceiverFilterer) WatchBreadReceived(opts *bind.WatchOpts, sink chan<- *TestNativeTokenReceiverBreadReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TestNativeTokenReceiver.contract.WatchLogs(opts, "BreadReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestNativeTokenReceiverBreadReceived)
				if err := _TestNativeTokenReceiver.contract.UnpackLog(event, "BreadReceived", log); err != nil {
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

// ParseBreadReceived is a log parse operation binding the contract event 0x16549311ba52796916987df5401f791fb06b998524a5a8684010010415850bb3.
//
// Solidity: event BreadReceived(address indexed from, uint256 amount, uint256 forwardedGas)
func (_TestNativeTokenReceiver *TestNativeTokenReceiverFilterer) ParseBreadReceived(log types.Log) (*TestNativeTokenReceiverBreadReceived, error) {
	event := new(TestNativeTokenReceiverBreadReceived)
	if err := _TestNativeTokenReceiver.contract.UnpackLog(event, "BreadReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSafeSignatureVerifierMetaData contains all meta data concerning the TestSafeSignatureVerifier contract.
var TestSafeSignatureVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISafe\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"typeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"encodeData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"isValidSafeSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610349806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806353f00b1414610030575b600080fd5b61015a600480360360e081101561004657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803590602001906401000000008111156100c157600080fd5b8201836020820111156100d357600080fd5b803590602001918460018302840111640100000000831117156100f557600080fd5b90919293919293908035906020019064010000000081111561011657600080fd5b82018360208201111561012857600080fd5b8035906020019184600183028401116401000000008311171561014a57600080fd5b909192939192939050505061018f565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b60006101e0878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061020b565b805190602001208814156101fd57631626ba7e60e01b90506101fe565b5b9998505050505050505050565b6060601960f81b600160f81b8585856040516020018083815260200182805190602001908083835b602083106102565780518252602082019150602081019050602083039250610233565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012060405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101847effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018381526020018281526020019450505050506040516020818303038152906040529050939250505056fea2646970667358221220fa448552905e4d2c6996c8dfe643dd61baa5601328f0516fa4847ec80ec75c8f64736f6c63430007060033",
}

// TestSafeSignatureVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use TestSafeSignatureVerifierMetaData.ABI instead.
var TestSafeSignatureVerifierABI = TestSafeSignatureVerifierMetaData.ABI

// TestSafeSignatureVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestSafeSignatureVerifierMetaData.Bin instead.
var TestSafeSignatureVerifierBin = TestSafeSignatureVerifierMetaData.Bin

// DeployTestSafeSignatureVerifier deploys a new Ethereum contract, binding an instance of TestSafeSignatureVerifier to it.
func DeployTestSafeSignatureVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestSafeSignatureVerifier, error) {
	parsed, err := TestSafeSignatureVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestSafeSignatureVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestSafeSignatureVerifier{TestSafeSignatureVerifierCaller: TestSafeSignatureVerifierCaller{contract: contract}, TestSafeSignatureVerifierTransactor: TestSafeSignatureVerifierTransactor{contract: contract}, TestSafeSignatureVerifierFilterer: TestSafeSignatureVerifierFilterer{contract: contract}}, nil
}

// TestSafeSignatureVerifier is an auto generated Go binding around an Ethereum contract.
type TestSafeSignatureVerifier struct {
	TestSafeSignatureVerifierCaller     // Read-only binding to the contract
	TestSafeSignatureVerifierTransactor // Write-only binding to the contract
	TestSafeSignatureVerifierFilterer   // Log filterer for contract events
}

// TestSafeSignatureVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestSafeSignatureVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSafeSignatureVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestSafeSignatureVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSafeSignatureVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestSafeSignatureVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSafeSignatureVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSafeSignatureVerifierSession struct {
	Contract     *TestSafeSignatureVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TestSafeSignatureVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestSafeSignatureVerifierCallerSession struct {
	Contract *TestSafeSignatureVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TestSafeSignatureVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestSafeSignatureVerifierTransactorSession struct {
	Contract     *TestSafeSignatureVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TestSafeSignatureVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestSafeSignatureVerifierRaw struct {
	Contract *TestSafeSignatureVerifier // Generic contract binding to access the raw methods on
}

// TestSafeSignatureVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestSafeSignatureVerifierCallerRaw struct {
	Contract *TestSafeSignatureVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// TestSafeSignatureVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestSafeSignatureVerifierTransactorRaw struct {
	Contract *TestSafeSignatureVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestSafeSignatureVerifier creates a new instance of TestSafeSignatureVerifier, bound to a specific deployed contract.
func NewTestSafeSignatureVerifier(address common.Address, backend bind.ContractBackend) (*TestSafeSignatureVerifier, error) {
	contract, err := bindTestSafeSignatureVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestSafeSignatureVerifier{TestSafeSignatureVerifierCaller: TestSafeSignatureVerifierCaller{contract: contract}, TestSafeSignatureVerifierTransactor: TestSafeSignatureVerifierTransactor{contract: contract}, TestSafeSignatureVerifierFilterer: TestSafeSignatureVerifierFilterer{contract: contract}}, nil
}

// NewTestSafeSignatureVerifierCaller creates a new read-only instance of TestSafeSignatureVerifier, bound to a specific deployed contract.
func NewTestSafeSignatureVerifierCaller(address common.Address, caller bind.ContractCaller) (*TestSafeSignatureVerifierCaller, error) {
	contract, err := bindTestSafeSignatureVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestSafeSignatureVerifierCaller{contract: contract}, nil
}

// NewTestSafeSignatureVerifierTransactor creates a new write-only instance of TestSafeSignatureVerifier, bound to a specific deployed contract.
func NewTestSafeSignatureVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*TestSafeSignatureVerifierTransactor, error) {
	contract, err := bindTestSafeSignatureVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestSafeSignatureVerifierTransactor{contract: contract}, nil
}

// NewTestSafeSignatureVerifierFilterer creates a new log filterer instance of TestSafeSignatureVerifier, bound to a specific deployed contract.
func NewTestSafeSignatureVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*TestSafeSignatureVerifierFilterer, error) {
	contract, err := bindTestSafeSignatureVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestSafeSignatureVerifierFilterer{contract: contract}, nil
}

// bindTestSafeSignatureVerifier binds a generic wrapper to an already deployed contract.
func bindTestSafeSignatureVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestSafeSignatureVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSafeSignatureVerifier.Contract.TestSafeSignatureVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSafeSignatureVerifier.Contract.TestSafeSignatureVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSafeSignatureVerifier.Contract.TestSafeSignatureVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSafeSignatureVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSafeSignatureVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSafeSignatureVerifier.Contract.contract.Transact(opts, method, params...)
}

// IsValidSafeSignature is a free data retrieval call binding the contract method 0x53f00b14.
//
// Solidity: function isValidSafeSignature(address , address , bytes32 _hash, bytes32 domainSeparator, bytes32 typeHash, bytes encodeData, bytes ) pure returns(bytes4 magic)
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierCaller) IsValidSafeSignature(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, _hash [32]byte, domainSeparator [32]byte, typeHash [32]byte, encodeData []byte, arg6 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TestSafeSignatureVerifier.contract.Call(opts, &out, "isValidSafeSignature", arg0, arg1, _hash, domainSeparator, typeHash, encodeData, arg6)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSafeSignature is a free data retrieval call binding the contract method 0x53f00b14.
//
// Solidity: function isValidSafeSignature(address , address , bytes32 _hash, bytes32 domainSeparator, bytes32 typeHash, bytes encodeData, bytes ) pure returns(bytes4 magic)
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierSession) IsValidSafeSignature(arg0 common.Address, arg1 common.Address, _hash [32]byte, domainSeparator [32]byte, typeHash [32]byte, encodeData []byte, arg6 []byte) ([4]byte, error) {
	return _TestSafeSignatureVerifier.Contract.IsValidSafeSignature(&_TestSafeSignatureVerifier.CallOpts, arg0, arg1, _hash, domainSeparator, typeHash, encodeData, arg6)
}

// IsValidSafeSignature is a free data retrieval call binding the contract method 0x53f00b14.
//
// Solidity: function isValidSafeSignature(address , address , bytes32 _hash, bytes32 domainSeparator, bytes32 typeHash, bytes encodeData, bytes ) pure returns(bytes4 magic)
func (_TestSafeSignatureVerifier *TestSafeSignatureVerifierCallerSession) IsValidSafeSignature(arg0 common.Address, arg1 common.Address, _hash [32]byte, domainSeparator [32]byte, typeHash [32]byte, encodeData []byte, arg6 []byte) ([4]byte, error) {
	return _TestSafeSignatureVerifier.Contract.IsValidSafeSignature(&_TestSafeSignatureVerifier.CallOpts, arg0, arg1, _hash, domainSeparator, typeHash, encodeData, arg6)
}

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 value) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 value) returns(bool)
func (_Token *TokenSession) Transfer(_to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 value) returns(bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, value)
}
