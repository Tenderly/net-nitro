// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yulgen

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

// InvalidJumpMetaData contains all meta data concerning the InvalidJump contract.
var InvalidJumpMetaData = &bind.MetaData{
	ABI: "null",
	Bin: "0x602380600c6000396000f3fe34601e5760003560e01c63c59e9bfd14601757600080fd5b6001600156005b600080fd",
}

// InvalidJumpABI is the input ABI used to generate the binding from.
// Deprecated: Use InvalidJumpMetaData.ABI instead.
var InvalidJumpABI = InvalidJumpMetaData.ABI

// InvalidJumpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InvalidJumpMetaData.Bin instead.
var InvalidJumpBin = InvalidJumpMetaData.Bin

// DeployInvalidJump deploys a new Ethereum contract, binding an instance of InvalidJump to it.
func DeployInvalidJump(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InvalidJump, error) {
	parsed, err := InvalidJumpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InvalidJumpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InvalidJump{InvalidJumpCaller: InvalidJumpCaller{contract: contract}, InvalidJumpTransactor: InvalidJumpTransactor{contract: contract}, InvalidJumpFilterer: InvalidJumpFilterer{contract: contract}}, nil
}

// InvalidJump is an auto generated Go binding around an Ethereum contract.
type InvalidJump struct {
	InvalidJumpCaller     // Read-only binding to the contract
	InvalidJumpTransactor // Write-only binding to the contract
	InvalidJumpFilterer   // Log filterer for contract events
}

// InvalidJumpCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvalidJumpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidJumpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvalidJumpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidJumpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvalidJumpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidJumpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvalidJumpSession struct {
	Contract     *InvalidJump      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvalidJumpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvalidJumpCallerSession struct {
	Contract *InvalidJumpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InvalidJumpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvalidJumpTransactorSession struct {
	Contract     *InvalidJumpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InvalidJumpRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvalidJumpRaw struct {
	Contract *InvalidJump // Generic contract binding to access the raw methods on
}

// InvalidJumpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvalidJumpCallerRaw struct {
	Contract *InvalidJumpCaller // Generic read-only contract binding to access the raw methods on
}

// InvalidJumpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvalidJumpTransactorRaw struct {
	Contract *InvalidJumpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvalidJump creates a new instance of InvalidJump, bound to a specific deployed contract.
func NewInvalidJump(address common.Address, backend bind.ContractBackend) (*InvalidJump, error) {
	contract, err := bindInvalidJump(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InvalidJump{InvalidJumpCaller: InvalidJumpCaller{contract: contract}, InvalidJumpTransactor: InvalidJumpTransactor{contract: contract}, InvalidJumpFilterer: InvalidJumpFilterer{contract: contract}}, nil
}

// NewInvalidJumpCaller creates a new read-only instance of InvalidJump, bound to a specific deployed contract.
func NewInvalidJumpCaller(address common.Address, caller bind.ContractCaller) (*InvalidJumpCaller, error) {
	contract, err := bindInvalidJump(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvalidJumpCaller{contract: contract}, nil
}

// NewInvalidJumpTransactor creates a new write-only instance of InvalidJump, bound to a specific deployed contract.
func NewInvalidJumpTransactor(address common.Address, transactor bind.ContractTransactor) (*InvalidJumpTransactor, error) {
	contract, err := bindInvalidJump(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvalidJumpTransactor{contract: contract}, nil
}

// NewInvalidJumpFilterer creates a new log filterer instance of InvalidJump, bound to a specific deployed contract.
func NewInvalidJumpFilterer(address common.Address, filterer bind.ContractFilterer) (*InvalidJumpFilterer, error) {
	contract, err := bindInvalidJump(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvalidJumpFilterer{contract: contract}, nil
}

// bindInvalidJump binds a generic wrapper to an already deployed contract.
func bindInvalidJump(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InvalidJumpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InvalidJump *InvalidJumpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InvalidJump.Contract.InvalidJumpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InvalidJump *InvalidJumpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InvalidJump.Contract.InvalidJumpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InvalidJump *InvalidJumpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InvalidJump.Contract.InvalidJumpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InvalidJump *InvalidJumpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InvalidJump.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InvalidJump *InvalidJumpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InvalidJump.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InvalidJump *InvalidJumpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InvalidJump.Contract.contract.Transact(opts, method, params...)
}

// Reader4844MetaData contains all meta data concerning the Reader4844 contract.
var Reader4844MetaData = &bind.MetaData{
	ABI: "null",
	Bin: "0x605780600a5f395ff3fe346053575f3560e01c8063e83a2d8214602757631f6d6ef714601f575f80fd5b4a5f5260205ff35b5f5b804990811560415760019160408260051b0152016029565b60409060205f528060205260051b015ff35b5f80fd",
}

// Reader4844ABI is the input ABI used to generate the binding from.
// Deprecated: Use Reader4844MetaData.ABI instead.
var Reader4844ABI = Reader4844MetaData.ABI

// Reader4844Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Reader4844MetaData.Bin instead.
var Reader4844Bin = Reader4844MetaData.Bin

// DeployReader4844 deploys a new Ethereum contract, binding an instance of Reader4844 to it.
func DeployReader4844(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reader4844, error) {
	parsed, err := Reader4844MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Reader4844Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reader4844{Reader4844Caller: Reader4844Caller{contract: contract}, Reader4844Transactor: Reader4844Transactor{contract: contract}, Reader4844Filterer: Reader4844Filterer{contract: contract}}, nil
}

// Reader4844 is an auto generated Go binding around an Ethereum contract.
type Reader4844 struct {
	Reader4844Caller     // Read-only binding to the contract
	Reader4844Transactor // Write-only binding to the contract
	Reader4844Filterer   // Log filterer for contract events
}

// Reader4844Caller is an auto generated read-only Go binding around an Ethereum contract.
type Reader4844Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Reader4844Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Reader4844Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Reader4844Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Reader4844Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Reader4844Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Reader4844Session struct {
	Contract     *Reader4844       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Reader4844CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Reader4844CallerSession struct {
	Contract *Reader4844Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Reader4844TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Reader4844TransactorSession struct {
	Contract     *Reader4844Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Reader4844Raw is an auto generated low-level Go binding around an Ethereum contract.
type Reader4844Raw struct {
	Contract *Reader4844 // Generic contract binding to access the raw methods on
}

// Reader4844CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Reader4844CallerRaw struct {
	Contract *Reader4844Caller // Generic read-only contract binding to access the raw methods on
}

// Reader4844TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Reader4844TransactorRaw struct {
	Contract *Reader4844Transactor // Generic write-only contract binding to access the raw methods on
}

// NewReader4844 creates a new instance of Reader4844, bound to a specific deployed contract.
func NewReader4844(address common.Address, backend bind.ContractBackend) (*Reader4844, error) {
	contract, err := bindReader4844(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reader4844{Reader4844Caller: Reader4844Caller{contract: contract}, Reader4844Transactor: Reader4844Transactor{contract: contract}, Reader4844Filterer: Reader4844Filterer{contract: contract}}, nil
}

// NewReader4844Caller creates a new read-only instance of Reader4844, bound to a specific deployed contract.
func NewReader4844Caller(address common.Address, caller bind.ContractCaller) (*Reader4844Caller, error) {
	contract, err := bindReader4844(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Reader4844Caller{contract: contract}, nil
}

// NewReader4844Transactor creates a new write-only instance of Reader4844, bound to a specific deployed contract.
func NewReader4844Transactor(address common.Address, transactor bind.ContractTransactor) (*Reader4844Transactor, error) {
	contract, err := bindReader4844(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Reader4844Transactor{contract: contract}, nil
}

// NewReader4844Filterer creates a new log filterer instance of Reader4844, bound to a specific deployed contract.
func NewReader4844Filterer(address common.Address, filterer bind.ContractFilterer) (*Reader4844Filterer, error) {
	contract, err := bindReader4844(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Reader4844Filterer{contract: contract}, nil
}

// bindReader4844 binds a generic wrapper to an already deployed contract.
func bindReader4844(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Reader4844MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader4844 *Reader4844Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reader4844.Contract.Reader4844Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader4844 *Reader4844Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader4844.Contract.Reader4844Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader4844 *Reader4844Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader4844.Contract.Reader4844Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader4844 *Reader4844CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reader4844.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader4844 *Reader4844TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader4844.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader4844 *Reader4844TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader4844.Contract.contract.Transact(opts, method, params...)
}
