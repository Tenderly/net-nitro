// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gas_dimensionsgen

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

// ArbSysMetaData contains all meta data concerning the ArbSys contract.
var ArbSysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"arbBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArbSysABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbSysMetaData.ABI instead.
var ArbSysABI = ArbSysMetaData.ABI

// ArbSys is an auto generated Go binding around an Ethereum contract.
type ArbSys struct {
	ArbSysCaller     // Read-only binding to the contract
	ArbSysTransactor // Write-only binding to the contract
	ArbSysFilterer   // Log filterer for contract events
}

// ArbSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSysSession struct {
	Contract     *ArbSys           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbSysCallerSession struct {
	Contract *ArbSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbSysTransactorSession struct {
	Contract     *ArbSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbSysRaw struct {
	Contract *ArbSys // Generic contract binding to access the raw methods on
}

// ArbSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbSysCallerRaw struct {
	Contract *ArbSysCaller // Generic read-only contract binding to access the raw methods on
}

// ArbSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbSysTransactorRaw struct {
	Contract *ArbSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbSys creates a new instance of ArbSys, bound to a specific deployed contract.
func NewArbSys(address common.Address, backend bind.ContractBackend) (*ArbSys, error) {
	contract, err := bindArbSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbSys{ArbSysCaller: ArbSysCaller{contract: contract}, ArbSysTransactor: ArbSysTransactor{contract: contract}, ArbSysFilterer: ArbSysFilterer{contract: contract}}, nil
}

// NewArbSysCaller creates a new read-only instance of ArbSys, bound to a specific deployed contract.
func NewArbSysCaller(address common.Address, caller bind.ContractCaller) (*ArbSysCaller, error) {
	contract, err := bindArbSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysCaller{contract: contract}, nil
}

// NewArbSysTransactor creates a new write-only instance of ArbSys, bound to a specific deployed contract.
func NewArbSysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbSysTransactor, error) {
	contract, err := bindArbSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysTransactor{contract: contract}, nil
}

// NewArbSysFilterer creates a new log filterer instance of ArbSys, bound to a specific deployed contract.
func NewArbSysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbSysFilterer, error) {
	contract, err := bindArbSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbSysFilterer{contract: contract}, nil
}

// bindArbSys binds a generic wrapper to an already deployed contract.
func bindArbSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbSysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbSys *ArbSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.ArbSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbSys *ArbSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbSys *ArbSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbSys *ArbSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbSys *ArbSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbSys *ArbSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transact(opts, method, params...)
}

// ArbBlockNumber is a free data retrieval call binding the contract method 0xa3b1b31d.
//
// Solidity: function arbBlockNumber() view returns(uint256)
func (_ArbSys *ArbSysCaller) ArbBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "arbBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbBlockNumber is a free data retrieval call binding the contract method 0xa3b1b31d.
//
// Solidity: function arbBlockNumber() view returns(uint256)
func (_ArbSys *ArbSysSession) ArbBlockNumber() (*big.Int, error) {
	return _ArbSys.Contract.ArbBlockNumber(&_ArbSys.CallOpts)
}

// ArbBlockNumber is a free data retrieval call binding the contract method 0xa3b1b31d.
//
// Solidity: function arbBlockNumber() view returns(uint256)
func (_ArbSys *ArbSysCallerSession) ArbBlockNumber() (*big.Int, error) {
	return _ArbSys.Contract.ArbBlockNumber(&_ArbSys.CallOpts)
}

// ArbWasmMetaData contains all meta data concerning the ArbWasm contract.
var ArbWasmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"activateProgram\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ArbWasmABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbWasmMetaData.ABI instead.
var ArbWasmABI = ArbWasmMetaData.ABI

// ArbWasm is an auto generated Go binding around an Ethereum contract.
type ArbWasm struct {
	ArbWasmCaller     // Read-only binding to the contract
	ArbWasmTransactor // Write-only binding to the contract
	ArbWasmFilterer   // Log filterer for contract events
}

// ArbWasmCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbWasmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbWasmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbWasmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbWasmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbWasmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbWasmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbWasmSession struct {
	Contract     *ArbWasm          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbWasmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbWasmCallerSession struct {
	Contract *ArbWasmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ArbWasmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbWasmTransactorSession struct {
	Contract     *ArbWasmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ArbWasmRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbWasmRaw struct {
	Contract *ArbWasm // Generic contract binding to access the raw methods on
}

// ArbWasmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbWasmCallerRaw struct {
	Contract *ArbWasmCaller // Generic read-only contract binding to access the raw methods on
}

// ArbWasmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbWasmTransactorRaw struct {
	Contract *ArbWasmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbWasm creates a new instance of ArbWasm, bound to a specific deployed contract.
func NewArbWasm(address common.Address, backend bind.ContractBackend) (*ArbWasm, error) {
	contract, err := bindArbWasm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbWasm{ArbWasmCaller: ArbWasmCaller{contract: contract}, ArbWasmTransactor: ArbWasmTransactor{contract: contract}, ArbWasmFilterer: ArbWasmFilterer{contract: contract}}, nil
}

// NewArbWasmCaller creates a new read-only instance of ArbWasm, bound to a specific deployed contract.
func NewArbWasmCaller(address common.Address, caller bind.ContractCaller) (*ArbWasmCaller, error) {
	contract, err := bindArbWasm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbWasmCaller{contract: contract}, nil
}

// NewArbWasmTransactor creates a new write-only instance of ArbWasm, bound to a specific deployed contract.
func NewArbWasmTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbWasmTransactor, error) {
	contract, err := bindArbWasm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbWasmTransactor{contract: contract}, nil
}

// NewArbWasmFilterer creates a new log filterer instance of ArbWasm, bound to a specific deployed contract.
func NewArbWasmFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbWasmFilterer, error) {
	contract, err := bindArbWasm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbWasmFilterer{contract: contract}, nil
}

// bindArbWasm binds a generic wrapper to an already deployed contract.
func bindArbWasm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbWasmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbWasm *ArbWasmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbWasm.Contract.ArbWasmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbWasm *ArbWasmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbWasm.Contract.ArbWasmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbWasm *ArbWasmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbWasm.Contract.ArbWasmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbWasm *ArbWasmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbWasm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbWasm *ArbWasmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbWasm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbWasm *ArbWasmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbWasm.Contract.contract.Transact(opts, method, params...)
}

// ActivateProgram is a paid mutator transaction binding the contract method 0x58c780c2.
//
// Solidity: function activateProgram(address program) payable returns()
func (_ArbWasm *ArbWasmTransactor) ActivateProgram(opts *bind.TransactOpts, program common.Address) (*types.Transaction, error) {
	return _ArbWasm.contract.Transact(opts, "activateProgram", program)
}

// ActivateProgram is a paid mutator transaction binding the contract method 0x58c780c2.
//
// Solidity: function activateProgram(address program) payable returns()
func (_ArbWasm *ArbWasmSession) ActivateProgram(program common.Address) (*types.Transaction, error) {
	return _ArbWasm.Contract.ActivateProgram(&_ArbWasm.TransactOpts, program)
}

// ActivateProgram is a paid mutator transaction binding the contract method 0x58c780c2.
//
// Solidity: function activateProgram(address program) payable returns()
func (_ArbWasm *ArbWasmTransactorSession) ActivateProgram(program common.Address) (*types.Transaction, error) {
	return _ArbWasm.Contract.ActivateProgram(&_ArbWasm.TransactOpts, program)
}

// BalanceMetaData contains all meta data concerning the Balance contract.
var BalanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"callBalanceCold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callBalanceWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061026a8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806315b4da6814610043578063668beee31461004d5780638381f58a14610057575b5f5ffd5b61004b610075565b005b6100556100ad565b005b61005f61015d565b60405161006c919061017a565b60405180910390f35b5f4790505f63deadbeef73ffffffffffffffffffffffffffffffffffffffff1631905080826100a491906101c0565b5f819055505050565b5f63deadbeef90505f8173ffffffffffffffffffffffffffffffffffffffff165f6040516100da90610220565b5f6040518083038185875af1925050503d805f8114610114576040519150601f19603f3d011682016040523d82523d5f602084013e610119565b606091505b50509050801561012b5760025f819055505b5f8273ffffffffffffffffffffffffffffffffffffffff16319050805f5461015391906101c0565b5f81905550505050565b5f5481565b5f819050919050565b61017481610162565b82525050565b5f60208201905061018d5f83018461016b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101ca82610162565b91506101d583610162565b92508282019050808211156101ed576101ec610193565b5b92915050565b5f81905092915050565b50565b5f61020b5f836101f3565b9150610216826101fd565b5f82019050919050565b5f61022a82610200565b915081905091905056fea2646970667358221220935a0b2636247be6e3977fb5f1f5a28ffeeb2c56438fbad3a399792efa3120fb64736f6c634300081e0033",
}

// BalanceABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceMetaData.ABI instead.
var BalanceABI = BalanceMetaData.ABI

// BalanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BalanceMetaData.Bin instead.
var BalanceBin = BalanceMetaData.Bin

// DeployBalance deploys a new Ethereum contract, binding an instance of Balance to it.
func DeployBalance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Balance, error) {
	parsed, err := BalanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BalanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Balance{BalanceCaller: BalanceCaller{contract: contract}, BalanceTransactor: BalanceTransactor{contract: contract}, BalanceFilterer: BalanceFilterer{contract: contract}}, nil
}

// Balance is an auto generated Go binding around an Ethereum contract.
type Balance struct {
	BalanceCaller     // Read-only binding to the contract
	BalanceTransactor // Write-only binding to the contract
	BalanceFilterer   // Log filterer for contract events
}

// BalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceSession struct {
	Contract     *Balance          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceCallerSession struct {
	Contract *BalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceTransactorSession struct {
	Contract     *BalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceRaw struct {
	Contract *Balance // Generic contract binding to access the raw methods on
}

// BalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceCallerRaw struct {
	Contract *BalanceCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceTransactorRaw struct {
	Contract *BalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalance creates a new instance of Balance, bound to a specific deployed contract.
func NewBalance(address common.Address, backend bind.ContractBackend) (*Balance, error) {
	contract, err := bindBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balance{BalanceCaller: BalanceCaller{contract: contract}, BalanceTransactor: BalanceTransactor{contract: contract}, BalanceFilterer: BalanceFilterer{contract: contract}}, nil
}

// NewBalanceCaller creates a new read-only instance of Balance, bound to a specific deployed contract.
func NewBalanceCaller(address common.Address, caller bind.ContractCaller) (*BalanceCaller, error) {
	contract, err := bindBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCaller{contract: contract}, nil
}

// NewBalanceTransactor creates a new write-only instance of Balance, bound to a specific deployed contract.
func NewBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceTransactor, error) {
	contract, err := bindBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceTransactor{contract: contract}, nil
}

// NewBalanceFilterer creates a new log filterer instance of Balance, bound to a specific deployed contract.
func NewBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceFilterer, error) {
	contract, err := bindBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceFilterer{contract: contract}, nil
}

// bindBalance binds a generic wrapper to an already deployed contract.
func bindBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balance *BalanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balance.Contract.BalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balance *BalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.Contract.BalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balance *BalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balance.Contract.BalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balance *BalanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balance *BalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balance *BalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balance.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Balance *BalanceCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Balance *BalanceSession) Number() (*big.Int, error) {
	return _Balance.Contract.Number(&_Balance.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Balance *BalanceCallerSession) Number() (*big.Int, error) {
	return _Balance.Contract.Number(&_Balance.CallOpts)
}

// CallBalanceCold is a paid mutator transaction binding the contract method 0x15b4da68.
//
// Solidity: function callBalanceCold() returns()
func (_Balance *BalanceTransactor) CallBalanceCold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.contract.Transact(opts, "callBalanceCold")
}

// CallBalanceCold is a paid mutator transaction binding the contract method 0x15b4da68.
//
// Solidity: function callBalanceCold() returns()
func (_Balance *BalanceSession) CallBalanceCold() (*types.Transaction, error) {
	return _Balance.Contract.CallBalanceCold(&_Balance.TransactOpts)
}

// CallBalanceCold is a paid mutator transaction binding the contract method 0x15b4da68.
//
// Solidity: function callBalanceCold() returns()
func (_Balance *BalanceTransactorSession) CallBalanceCold() (*types.Transaction, error) {
	return _Balance.Contract.CallBalanceCold(&_Balance.TransactOpts)
}

// CallBalanceWarm is a paid mutator transaction binding the contract method 0x668beee3.
//
// Solidity: function callBalanceWarm() returns()
func (_Balance *BalanceTransactor) CallBalanceWarm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.contract.Transact(opts, "callBalanceWarm")
}

// CallBalanceWarm is a paid mutator transaction binding the contract method 0x668beee3.
//
// Solidity: function callBalanceWarm() returns()
func (_Balance *BalanceSession) CallBalanceWarm() (*types.Transaction, error) {
	return _Balance.Contract.CallBalanceWarm(&_Balance.TransactOpts)
}

// CallBalanceWarm is a paid mutator transaction binding the contract method 0x668beee3.
//
// Solidity: function callBalanceWarm() returns()
func (_Balance *BalanceTransactorSession) CallBalanceWarm() (*types.Transaction, error) {
	return _Balance.Contract.CallBalanceWarm(&_Balance.TransactOpts)
}

// CallCodeeMetaData contains all meta data concerning the CallCodee contract.
var CallCodeeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061012f8061001c5f395ff3fe6080604052600436106028575f3560e01c80632e52d6061460325780633fb5c1cb14605757602e565b36602e57005b5f5ffd5b348015603c575f5ffd5b506043606f565b604051604e91906093565b60405180910390f35b606d60048036038101906069919060d3565b6074565b005b5f5481565b805f8190555050565b5f819050919050565b608d81607d565b82525050565b5f60208201905060a45f8301846086565b92915050565b5f5ffd5b60b581607d565b811460be575f5ffd5b50565b5f8135905060cd8160ae565b92915050565b5f6020828403121560e55760e460aa565b5b5f60f08482850160c1565b9150509291505056fea26469706673582212205dbee44ee2d77363a78ecbb12f168a5b64c1ce06b40abc3749e50ab2f0397c4664736f6c634300081e0033",
}

// CallCodeeABI is the input ABI used to generate the binding from.
// Deprecated: Use CallCodeeMetaData.ABI instead.
var CallCodeeABI = CallCodeeMetaData.ABI

// CallCodeeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallCodeeMetaData.Bin instead.
var CallCodeeBin = CallCodeeMetaData.Bin

// DeployCallCodee deploys a new Ethereum contract, binding an instance of CallCodee to it.
func DeployCallCodee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallCodee, error) {
	parsed, err := CallCodeeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallCodeeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallCodee{CallCodeeCaller: CallCodeeCaller{contract: contract}, CallCodeeTransactor: CallCodeeTransactor{contract: contract}, CallCodeeFilterer: CallCodeeFilterer{contract: contract}}, nil
}

// CallCodee is an auto generated Go binding around an Ethereum contract.
type CallCodee struct {
	CallCodeeCaller     // Read-only binding to the contract
	CallCodeeTransactor // Write-only binding to the contract
	CallCodeeFilterer   // Log filterer for contract events
}

// CallCodeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallCodeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallCodeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallCodeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallCodeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallCodeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallCodeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallCodeeSession struct {
	Contract     *CallCodee        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallCodeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallCodeeCallerSession struct {
	Contract *CallCodeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CallCodeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallCodeeTransactorSession struct {
	Contract     *CallCodeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CallCodeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallCodeeRaw struct {
	Contract *CallCodee // Generic contract binding to access the raw methods on
}

// CallCodeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallCodeeCallerRaw struct {
	Contract *CallCodeeCaller // Generic read-only contract binding to access the raw methods on
}

// CallCodeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallCodeeTransactorRaw struct {
	Contract *CallCodeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallCodee creates a new instance of CallCodee, bound to a specific deployed contract.
func NewCallCodee(address common.Address, backend bind.ContractBackend) (*CallCodee, error) {
	contract, err := bindCallCodee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallCodee{CallCodeeCaller: CallCodeeCaller{contract: contract}, CallCodeeTransactor: CallCodeeTransactor{contract: contract}, CallCodeeFilterer: CallCodeeFilterer{contract: contract}}, nil
}

// NewCallCodeeCaller creates a new read-only instance of CallCodee, bound to a specific deployed contract.
func NewCallCodeeCaller(address common.Address, caller bind.ContractCaller) (*CallCodeeCaller, error) {
	contract, err := bindCallCodee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallCodeeCaller{contract: contract}, nil
}

// NewCallCodeeTransactor creates a new write-only instance of CallCodee, bound to a specific deployed contract.
func NewCallCodeeTransactor(address common.Address, transactor bind.ContractTransactor) (*CallCodeeTransactor, error) {
	contract, err := bindCallCodee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallCodeeTransactor{contract: contract}, nil
}

// NewCallCodeeFilterer creates a new log filterer instance of CallCodee, bound to a specific deployed contract.
func NewCallCodeeFilterer(address common.Address, filterer bind.ContractFilterer) (*CallCodeeFilterer, error) {
	contract, err := bindCallCodee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallCodeeFilterer{contract: contract}, nil
}

// bindCallCodee binds a generic wrapper to an already deployed contract.
func bindCallCodee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CallCodeeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallCodee *CallCodeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallCodee.Contract.CallCodeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallCodee *CallCodeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallCodee.Contract.CallCodeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallCodee *CallCodeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallCodee.Contract.CallCodeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallCodee *CallCodeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallCodee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallCodee *CallCodeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallCodee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallCodee *CallCodeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallCodee.Contract.contract.Transact(opts, method, params...)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_CallCodee *CallCodeeCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CallCodee.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_CallCodee *CallCodeeSession) N() (*big.Int, error) {
	return _CallCodee.Contract.N(&_CallCodee.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_CallCodee *CallCodeeCallerSession) N() (*big.Int, error) {
	return _CallCodee.Contract.N(&_CallCodee.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _n) payable returns()
func (_CallCodee *CallCodeeTransactor) SetNumber(opts *bind.TransactOpts, _n *big.Int) (*types.Transaction, error) {
	return _CallCodee.contract.Transact(opts, "setNumber", _n)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _n) payable returns()
func (_CallCodee *CallCodeeSession) SetNumber(_n *big.Int) (*types.Transaction, error) {
	return _CallCodee.Contract.SetNumber(&_CallCodee.TransactOpts, _n)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _n) payable returns()
func (_CallCodee *CallCodeeTransactorSession) SetNumber(_n *big.Int) (*types.Transaction, error) {
	return _CallCodee.Contract.SetNumber(&_CallCodee.TransactOpts, _n)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CallCodee *CallCodeeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallCodee.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CallCodee *CallCodeeSession) Receive() (*types.Transaction, error) {
	return _CallCodee.Contract.Receive(&_CallCodee.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CallCodee *CallCodeeTransactorSession) Receive() (*types.Transaction, error) {
	return _CallCodee.Contract.Receive(&_CallCodee.TransactOpts)
}

// CallCoderMetaData contains all meta data concerning the CallCoder contract.
var CallCoderMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldNoTransferMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldNoTransferMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldPayableMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldPayableMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmNoTransferMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmNoTransferMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmPayableMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmPayableMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506106478061001c5f395ff3fe608060405260043610610099575f3560e01c8063692ef31d116100685780638b5c70501161004d5780638b5c705014610196578063b3bbaaa3146101be578063c6c35afc146101e6576100a0565b8063692ef31d146101465780637bde0d8b1461016e576100a0565b80630153b48a146100a457806306d668b9146100cc5780632e52d606146100f457806361f8f2b81461011e576100a0565b366100a057005b5f5ffd5b3480156100af575f5ffd5b506100ca60048036038101906100c5919061054d565b61020e565b005b3480156100d7575f5ffd5b506100f260048036038101906100ed919061054d565b61021d565b005b3480156100ff575f5ffd5b5061010861026e565b6040516101159190610590565b60405180910390f35b348015610129575f5ffd5b50610144600480360381019061013f919061054d565b610273565b005b348015610151575f5ffd5b5061016c6004803603810190610167919061054d565b6102c2565b005b348015610179575f5ffd5b50610194600480360381019061018f919061054d565b61031a565b005b3480156101a1575f5ffd5b506101bc60048036038101906101b7919061054d565b610369565b005b3480156101c9575f5ffd5b506101e460048036038101906101df919061054d565b610376565b005b3480156101f1575f5ffd5b5061020c6004803603810190610207919061054d565b610385565b005b61021a8161abcd610392565b50565b5f8173ffffffffffffffffffffffffffffffffffffffff163190506102448261abcd610392565b8173ffffffffffffffffffffffffffffffffffffffff1631810361026a5760035f819055505b5050565b5f5481565b5f8173ffffffffffffffffffffffffffffffffffffffff16319050610298825f610392565b8173ffffffffffffffffffffffffffffffffffffffff163181036102be5760035f819055505b5050565b60025f819055505f8173ffffffffffffffffffffffffffffffffffffffff163190506102f08261abcd610442565b8173ffffffffffffffffffffffffffffffffffffffff163181036103165760035f819055505b5050565b5f8173ffffffffffffffffffffffffffffffffffffffff1631905061033f825f610442565b8173ffffffffffffffffffffffffffffffffffffffff163181036103655760035f819055505b5050565b610373815f610442565b50565b6103828161abcd610442565b50565b61038f815f610392565b50565b60015f819055505f633fb5c1cb60e01b6113376040516024016103b591906105f8565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f815190505f602059016020818460208701888a5af2915050801561043b5760025f819055505b5050505050565b60015f819055505f633fb5c1cb60e01b61133760405160240161046591906105f8565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f815190505f5f6020828460208701888a5af2905080156104e75760025f819055505b505050505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61051c826104f3565b9050919050565b61052c81610512565b8114610536575f5ffd5b50565b5f8135905061054781610523565b92915050565b5f60208284031215610562576105616104ef565b5b5f61056f84828501610539565b91505092915050565b5f819050919050565b61058a81610578565b82525050565b5f6020820190506105a35f830184610581565b92915050565b5f819050919050565b5f61ffff82169050919050565b5f819050919050565b5f6105e26105dd6105d8846105a9565b6105bf565b6105b2565b9050919050565b6105f2816105c8565b82525050565b5f60208201905061060b5f8301846105e9565b9291505056fea2646970667358221220b2895762d56c83ea2b767266e6055bf48e21ed73b0e8ddf8f72378493837511664736f6c634300081e0033",
}

// CallCoderABI is the input ABI used to generate the binding from.
// Deprecated: Use CallCoderMetaData.ABI instead.
var CallCoderABI = CallCoderMetaData.ABI

// CallCoderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallCoderMetaData.Bin instead.
var CallCoderBin = CallCoderMetaData.Bin

// DeployCallCoder deploys a new Ethereum contract, binding an instance of CallCoder to it.
func DeployCallCoder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallCoder, error) {
	parsed, err := CallCoderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallCoderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallCoder{CallCoderCaller: CallCoderCaller{contract: contract}, CallCoderTransactor: CallCoderTransactor{contract: contract}, CallCoderFilterer: CallCoderFilterer{contract: contract}}, nil
}

// CallCoder is an auto generated Go binding around an Ethereum contract.
type CallCoder struct {
	CallCoderCaller     // Read-only binding to the contract
	CallCoderTransactor // Write-only binding to the contract
	CallCoderFilterer   // Log filterer for contract events
}

// CallCoderCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallCoderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallCoderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallCoderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallCoderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallCoderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallCoderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallCoderSession struct {
	Contract     *CallCoder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallCoderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallCoderCallerSession struct {
	Contract *CallCoderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CallCoderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallCoderTransactorSession struct {
	Contract     *CallCoderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CallCoderRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallCoderRaw struct {
	Contract *CallCoder // Generic contract binding to access the raw methods on
}

// CallCoderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallCoderCallerRaw struct {
	Contract *CallCoderCaller // Generic read-only contract binding to access the raw methods on
}

// CallCoderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallCoderTransactorRaw struct {
	Contract *CallCoderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallCoder creates a new instance of CallCoder, bound to a specific deployed contract.
func NewCallCoder(address common.Address, backend bind.ContractBackend) (*CallCoder, error) {
	contract, err := bindCallCoder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallCoder{CallCoderCaller: CallCoderCaller{contract: contract}, CallCoderTransactor: CallCoderTransactor{contract: contract}, CallCoderFilterer: CallCoderFilterer{contract: contract}}, nil
}

// NewCallCoderCaller creates a new read-only instance of CallCoder, bound to a specific deployed contract.
func NewCallCoderCaller(address common.Address, caller bind.ContractCaller) (*CallCoderCaller, error) {
	contract, err := bindCallCoder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallCoderCaller{contract: contract}, nil
}

// NewCallCoderTransactor creates a new write-only instance of CallCoder, bound to a specific deployed contract.
func NewCallCoderTransactor(address common.Address, transactor bind.ContractTransactor) (*CallCoderTransactor, error) {
	contract, err := bindCallCoder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallCoderTransactor{contract: contract}, nil
}

// NewCallCoderFilterer creates a new log filterer instance of CallCoder, bound to a specific deployed contract.
func NewCallCoderFilterer(address common.Address, filterer bind.ContractFilterer) (*CallCoderFilterer, error) {
	contract, err := bindCallCoder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallCoderFilterer{contract: contract}, nil
}

// bindCallCoder binds a generic wrapper to an already deployed contract.
func bindCallCoder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CallCoderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallCoder *CallCoderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallCoder.Contract.CallCoderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallCoder *CallCoderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallCoder.Contract.CallCoderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallCoder *CallCoderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallCoder.Contract.CallCoderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallCoder *CallCoderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallCoder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallCoder *CallCoderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallCoder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallCoder *CallCoderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallCoder.Contract.contract.Transact(opts, method, params...)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_CallCoder *CallCoderCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CallCoder.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_CallCoder *CallCoderSession) N() (*big.Int, error) {
	return _CallCoder.Contract.N(&_CallCoder.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_CallCoder *CallCoderCallerSession) N() (*big.Int, error) {
	return _CallCoder.Contract.N(&_CallCoder.CallOpts)
}

// ColdNoTransferMemExpansion is a paid mutator transaction binding the contract method 0xc6c35afc.
//
// Solidity: function coldNoTransferMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactor) ColdNoTransferMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "coldNoTransferMemExpansion", target)
}

// ColdNoTransferMemExpansion is a paid mutator transaction binding the contract method 0xc6c35afc.
//
// Solidity: function coldNoTransferMemExpansion(address target) returns()
func (_CallCoder *CallCoderSession) ColdNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdNoTransferMemExpansion(&_CallCoder.TransactOpts, target)
}

// ColdNoTransferMemExpansion is a paid mutator transaction binding the contract method 0xc6c35afc.
//
// Solidity: function coldNoTransferMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactorSession) ColdNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdNoTransferMemExpansion(&_CallCoder.TransactOpts, target)
}

// ColdNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x8b5c7050.
//
// Solidity: function coldNoTransferMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactor) ColdNoTransferMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "coldNoTransferMemUnchanged", target)
}

// ColdNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x8b5c7050.
//
// Solidity: function coldNoTransferMemUnchanged(address target) returns()
func (_CallCoder *CallCoderSession) ColdNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdNoTransferMemUnchanged(&_CallCoder.TransactOpts, target)
}

// ColdNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x8b5c7050.
//
// Solidity: function coldNoTransferMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactorSession) ColdNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdNoTransferMemUnchanged(&_CallCoder.TransactOpts, target)
}

// ColdPayableMemExpansion is a paid mutator transaction binding the contract method 0x0153b48a.
//
// Solidity: function coldPayableMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactor) ColdPayableMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "coldPayableMemExpansion", target)
}

// ColdPayableMemExpansion is a paid mutator transaction binding the contract method 0x0153b48a.
//
// Solidity: function coldPayableMemExpansion(address target) returns()
func (_CallCoder *CallCoderSession) ColdPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdPayableMemExpansion(&_CallCoder.TransactOpts, target)
}

// ColdPayableMemExpansion is a paid mutator transaction binding the contract method 0x0153b48a.
//
// Solidity: function coldPayableMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactorSession) ColdPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdPayableMemExpansion(&_CallCoder.TransactOpts, target)
}

// ColdPayableMemUnchanged is a paid mutator transaction binding the contract method 0xb3bbaaa3.
//
// Solidity: function coldPayableMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactor) ColdPayableMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "coldPayableMemUnchanged", target)
}

// ColdPayableMemUnchanged is a paid mutator transaction binding the contract method 0xb3bbaaa3.
//
// Solidity: function coldPayableMemUnchanged(address target) returns()
func (_CallCoder *CallCoderSession) ColdPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdPayableMemUnchanged(&_CallCoder.TransactOpts, target)
}

// ColdPayableMemUnchanged is a paid mutator transaction binding the contract method 0xb3bbaaa3.
//
// Solidity: function coldPayableMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactorSession) ColdPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.ColdPayableMemUnchanged(&_CallCoder.TransactOpts, target)
}

// WarmNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x61f8f2b8.
//
// Solidity: function warmNoTransferMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactor) WarmNoTransferMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "warmNoTransferMemExpansion", target)
}

// WarmNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x61f8f2b8.
//
// Solidity: function warmNoTransferMemExpansion(address target) returns()
func (_CallCoder *CallCoderSession) WarmNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmNoTransferMemExpansion(&_CallCoder.TransactOpts, target)
}

// WarmNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x61f8f2b8.
//
// Solidity: function warmNoTransferMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactorSession) WarmNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmNoTransferMemExpansion(&_CallCoder.TransactOpts, target)
}

// WarmNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x7bde0d8b.
//
// Solidity: function warmNoTransferMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactor) WarmNoTransferMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "warmNoTransferMemUnchanged", target)
}

// WarmNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x7bde0d8b.
//
// Solidity: function warmNoTransferMemUnchanged(address target) returns()
func (_CallCoder *CallCoderSession) WarmNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmNoTransferMemUnchanged(&_CallCoder.TransactOpts, target)
}

// WarmNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x7bde0d8b.
//
// Solidity: function warmNoTransferMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactorSession) WarmNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmNoTransferMemUnchanged(&_CallCoder.TransactOpts, target)
}

// WarmPayableMemExpansion is a paid mutator transaction binding the contract method 0x06d668b9.
//
// Solidity: function warmPayableMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactor) WarmPayableMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "warmPayableMemExpansion", target)
}

// WarmPayableMemExpansion is a paid mutator transaction binding the contract method 0x06d668b9.
//
// Solidity: function warmPayableMemExpansion(address target) returns()
func (_CallCoder *CallCoderSession) WarmPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmPayableMemExpansion(&_CallCoder.TransactOpts, target)
}

// WarmPayableMemExpansion is a paid mutator transaction binding the contract method 0x06d668b9.
//
// Solidity: function warmPayableMemExpansion(address target) returns()
func (_CallCoder *CallCoderTransactorSession) WarmPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmPayableMemExpansion(&_CallCoder.TransactOpts, target)
}

// WarmPayableMemUnchanged is a paid mutator transaction binding the contract method 0x692ef31d.
//
// Solidity: function warmPayableMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactor) WarmPayableMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _CallCoder.contract.Transact(opts, "warmPayableMemUnchanged", target)
}

// WarmPayableMemUnchanged is a paid mutator transaction binding the contract method 0x692ef31d.
//
// Solidity: function warmPayableMemUnchanged(address target) returns()
func (_CallCoder *CallCoderSession) WarmPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmPayableMemUnchanged(&_CallCoder.TransactOpts, target)
}

// WarmPayableMemUnchanged is a paid mutator transaction binding the contract method 0x692ef31d.
//
// Solidity: function warmPayableMemUnchanged(address target) returns()
func (_CallCoder *CallCoderTransactorSession) WarmPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _CallCoder.Contract.WarmPayableMemUnchanged(&_CallCoder.TransactOpts, target)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CallCoder *CallCoderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallCoder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CallCoder *CallCoderSession) Receive() (*types.Transaction, error) {
	return _CallCoder.Contract.Receive(&_CallCoder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CallCoder *CallCoderTransactorSession) Receive() (*types.Transaction, error) {
	return _CallCoder.Contract.Receive(&_CallCoder.TransactOpts)
}

// CalleeMetaData contains all meta data concerning the Callee contract.
var CalleeMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061012f8061001c5f395ff3fe6080604052600436106028575f3560e01c80633fb5c1cb1460325780638381f58a14604a57602e565b36602e57005b5f5ffd5b604860048036038101906044919060af565b606f565b005b3480156054575f5ffd5b50605b6078565b6040516066919060e2565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6091816081565b8114609a575f5ffd5b50565b5f8135905060a981608a565b92915050565b5f6020828403121560c15760c0607d565b5b5f60cc84828501609d565b91505092915050565b60dc816081565b82525050565b5f60208201905060f35f83018460d5565b9291505056fea2646970667358221220e2f64d7c87ec6a507f07f4a040b187ef61e3e66d21d246d3f6da3692827bc85564736f6c634300081e0033",
}

// CalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use CalleeMetaData.ABI instead.
var CalleeABI = CalleeMetaData.ABI

// CalleeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CalleeMetaData.Bin instead.
var CalleeBin = CalleeMetaData.Bin

// DeployCallee deploys a new Ethereum contract, binding an instance of Callee to it.
func DeployCallee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Callee, error) {
	parsed, err := CalleeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CalleeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Callee{CalleeCaller: CalleeCaller{contract: contract}, CalleeTransactor: CalleeTransactor{contract: contract}, CalleeFilterer: CalleeFilterer{contract: contract}}, nil
}

// Callee is an auto generated Go binding around an Ethereum contract.
type Callee struct {
	CalleeCaller     // Read-only binding to the contract
	CalleeTransactor // Write-only binding to the contract
	CalleeFilterer   // Log filterer for contract events
}

// CalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalleeSession struct {
	Contract     *Callee           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalleeCallerSession struct {
	Contract *CalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalleeTransactorSession struct {
	Contract     *CalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalleeRaw struct {
	Contract *Callee // Generic contract binding to access the raw methods on
}

// CalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalleeCallerRaw struct {
	Contract *CalleeCaller // Generic read-only contract binding to access the raw methods on
}

// CalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalleeTransactorRaw struct {
	Contract *CalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallee creates a new instance of Callee, bound to a specific deployed contract.
func NewCallee(address common.Address, backend bind.ContractBackend) (*Callee, error) {
	contract, err := bindCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Callee{CalleeCaller: CalleeCaller{contract: contract}, CalleeTransactor: CalleeTransactor{contract: contract}, CalleeFilterer: CalleeFilterer{contract: contract}}, nil
}

// NewCalleeCaller creates a new read-only instance of Callee, bound to a specific deployed contract.
func NewCalleeCaller(address common.Address, caller bind.ContractCaller) (*CalleeCaller, error) {
	contract, err := bindCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalleeCaller{contract: contract}, nil
}

// NewCalleeTransactor creates a new write-only instance of Callee, bound to a specific deployed contract.
func NewCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*CalleeTransactor, error) {
	contract, err := bindCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalleeTransactor{contract: contract}, nil
}

// NewCalleeFilterer creates a new log filterer instance of Callee, bound to a specific deployed contract.
func NewCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*CalleeFilterer, error) {
	contract, err := bindCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalleeFilterer{contract: contract}, nil
}

// bindCallee binds a generic wrapper to an already deployed contract.
func bindCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CalleeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callee *CalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callee.Contract.CalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callee *CalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.Contract.CalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callee *CalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callee.Contract.CalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callee *CalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callee *CalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callee *CalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callee.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Callee *CalleeCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Callee.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Callee *CalleeSession) Number() (*big.Int, error) {
	return _Callee.Contract.Number(&_Callee.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Callee *CalleeCallerSession) Number() (*big.Int, error) {
	return _Callee.Contract.Number(&_Callee.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) payable returns()
func (_Callee *CalleeTransactor) SetNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Callee.contract.Transact(opts, "setNumber", _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) payable returns()
func (_Callee *CalleeSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _Callee.Contract.SetNumber(&_Callee.TransactOpts, _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) payable returns()
func (_Callee *CalleeTransactorSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _Callee.Contract.SetNumber(&_Callee.TransactOpts, _number)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Callee *CalleeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Callee *CalleeSession) Receive() (*types.Transaction, error) {
	return _Callee.Contract.Receive(&_Callee.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Callee *CalleeTransactorSession) Receive() (*types.Transaction, error) {
	return _Callee.Contract.Receive(&_Callee.TransactOpts)
}

// CallerMetaData contains all meta data concerning the Caller contract.
var CallerMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldNoTransferMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldNoTransferMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldPayableMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"coldPayableMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmNoTransferMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmNoTransferMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmPayableMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"warmPayableMemUnchanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506106fd8061001c5f395ff3fe608060405260043610610099575f3560e01c8063692ef31d116100685780638b5c70501161004d5780638b5c705014610196578063b3bbaaa3146101be578063c6c35afc146101e6576100a0565b8063692ef31d146101465780637bde0d8b1461016e576100a0565b80630153b48a146100a457806306d668b9146100cc5780632e52d606146100f457806361f8f2b81461011e576100a0565b366100a057005b5f5ffd5b3480156100af575f5ffd5b506100ca60048036038101906100c5919061059b565b61020e565b005b3480156100d7575f5ffd5b506100f260048036038101906100ed919061059b565b61021d565b005b3480156100ff575f5ffd5b5061010861026e565b60405161011591906105de565b60405180910390f35b348015610129575f5ffd5b50610144600480360381019061013f919061059b565b610273565b005b348015610151575f5ffd5b5061016c6004803603810190610167919061059b565b6102c2565b005b348015610179575f5ffd5b50610194600480360381019061018f919061059b565b61031a565b005b3480156101a1575f5ffd5b506101bc60048036038101906101b7919061059b565b610369565b005b3480156101c9575f5ffd5b506101e460048036038101906101df919061059b565b610376565b005b3480156101f1575f5ffd5b5061020c6004803603810190610207919061059b565b610385565b005b61021a8161abcd610392565b50565b5f8173ffffffffffffffffffffffffffffffffffffffff163190506102448261abcd610392565b8173ffffffffffffffffffffffffffffffffffffffff1631810361026a5760035f819055505b5050565b5f5481565b5f8173ffffffffffffffffffffffffffffffffffffffff16319050610298825f610392565b8173ffffffffffffffffffffffffffffffffffffffff163181036102be5760035f819055505b5050565b60025f819055505f8173ffffffffffffffffffffffffffffffffffffffff163190506102f08261abcd610442565b8173ffffffffffffffffffffffffffffffffffffffff163181036103165760035f819055505b5050565b5f8173ffffffffffffffffffffffffffffffffffffffff1631905061033f825f610442565b8173ffffffffffffffffffffffffffffffffffffffff163181036103655760035f819055505b5050565b610373815f610442565b50565b6103828161abcd610442565b50565b61038f815f610392565b50565b60015f819055505f633fb5c1cb60e01b6113376040516024016103b59190610646565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f815190505f602059016020818460208701888a5af1915050801561043b5760025f819055505b5050505050565b60015f819055505f8273ffffffffffffffffffffffffffffffffffffffff1682633fb5c1cb60e01b61133760405160240161047d9190610646565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516104e791906106b1565b5f6040518083038185875af1925050503d805f8114610521576040519150601f19603f3d011682016040523d82523d5f602084013e610526565b606091505b5050905080156105385760025f819055505b505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61056a82610541565b9050919050565b61057a81610560565b8114610584575f5ffd5b50565b5f8135905061059581610571565b92915050565b5f602082840312156105b0576105af61053d565b5b5f6105bd84828501610587565b91505092915050565b5f819050919050565b6105d8816105c6565b82525050565b5f6020820190506105f15f8301846105cf565b92915050565b5f819050919050565b5f61ffff82169050919050565b5f819050919050565b5f61063061062b610626846105f7565b61060d565b610600565b9050919050565b61064081610616565b82525050565b5f6020820190506106595f830184610637565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61068b8261065f565b6106958185610669565b93506106a5818560208601610673565b80840191505092915050565b5f6106bc8284610681565b91508190509291505056fea2646970667358221220ee1d52ae3e3b532c0659845506527913e24bf6e79e2a7e3be0e7280fc09ae6fa64736f6c634300081e0033",
}

// CallerABI is the input ABI used to generate the binding from.
// Deprecated: Use CallerMetaData.ABI instead.
var CallerABI = CallerMetaData.ABI

// CallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallerMetaData.Bin instead.
var CallerBin = CallerMetaData.Bin

// DeployCaller deploys a new Ethereum contract, binding an instance of Caller to it.
func DeployCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Caller, error) {
	parsed, err := CallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// Caller is an auto generated Go binding around an Ethereum contract.
type Caller struct {
	CallerCaller     // Read-only binding to the contract
	CallerTransactor // Write-only binding to the contract
	CallerFilterer   // Log filterer for contract events
}

// CallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallerSession struct {
	Contract     *Caller           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallerCallerSession struct {
	Contract *CallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallerTransactorSession struct {
	Contract     *CallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallerRaw struct {
	Contract *Caller // Generic contract binding to access the raw methods on
}

// CallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerCallerRaw struct {
	Contract *CallerCaller // Generic read-only contract binding to access the raw methods on
}

// CallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallerTransactorRaw struct {
	Contract *CallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaller creates a new instance of Caller, bound to a specific deployed contract.
func NewCaller(address common.Address, backend bind.ContractBackend) (*Caller, error) {
	contract, err := bindCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// NewCallerCaller creates a new read-only instance of Caller, bound to a specific deployed contract.
func NewCallerCaller(address common.Address, caller bind.ContractCaller) (*CallerCaller, error) {
	contract, err := bindCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallerCaller{contract: contract}, nil
}

// NewCallerTransactor creates a new write-only instance of Caller, bound to a specific deployed contract.
func NewCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CallerTransactor, error) {
	contract, err := bindCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallerTransactor{contract: contract}, nil
}

// NewCallerFilterer creates a new log filterer instance of Caller, bound to a specific deployed contract.
func NewCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CallerFilterer, error) {
	contract, err := bindCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallerFilterer{contract: contract}, nil
}

// bindCaller binds a generic wrapper to an already deployed contract.
func bindCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.CallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transact(opts, method, params...)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Caller *CallerCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Caller *CallerSession) N() (*big.Int, error) {
	return _Caller.Contract.N(&_Caller.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Caller *CallerCallerSession) N() (*big.Int, error) {
	return _Caller.Contract.N(&_Caller.CallOpts)
}

// ColdNoTransferMemExpansion is a paid mutator transaction binding the contract method 0xc6c35afc.
//
// Solidity: function coldNoTransferMemExpansion(address target) returns()
func (_Caller *CallerTransactor) ColdNoTransferMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "coldNoTransferMemExpansion", target)
}

// ColdNoTransferMemExpansion is a paid mutator transaction binding the contract method 0xc6c35afc.
//
// Solidity: function coldNoTransferMemExpansion(address target) returns()
func (_Caller *CallerSession) ColdNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdNoTransferMemExpansion(&_Caller.TransactOpts, target)
}

// ColdNoTransferMemExpansion is a paid mutator transaction binding the contract method 0xc6c35afc.
//
// Solidity: function coldNoTransferMemExpansion(address target) returns()
func (_Caller *CallerTransactorSession) ColdNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdNoTransferMemExpansion(&_Caller.TransactOpts, target)
}

// ColdNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x8b5c7050.
//
// Solidity: function coldNoTransferMemUnchanged(address target) returns()
func (_Caller *CallerTransactor) ColdNoTransferMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "coldNoTransferMemUnchanged", target)
}

// ColdNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x8b5c7050.
//
// Solidity: function coldNoTransferMemUnchanged(address target) returns()
func (_Caller *CallerSession) ColdNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdNoTransferMemUnchanged(&_Caller.TransactOpts, target)
}

// ColdNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x8b5c7050.
//
// Solidity: function coldNoTransferMemUnchanged(address target) returns()
func (_Caller *CallerTransactorSession) ColdNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdNoTransferMemUnchanged(&_Caller.TransactOpts, target)
}

// ColdPayableMemExpansion is a paid mutator transaction binding the contract method 0x0153b48a.
//
// Solidity: function coldPayableMemExpansion(address target) returns()
func (_Caller *CallerTransactor) ColdPayableMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "coldPayableMemExpansion", target)
}

// ColdPayableMemExpansion is a paid mutator transaction binding the contract method 0x0153b48a.
//
// Solidity: function coldPayableMemExpansion(address target) returns()
func (_Caller *CallerSession) ColdPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdPayableMemExpansion(&_Caller.TransactOpts, target)
}

// ColdPayableMemExpansion is a paid mutator transaction binding the contract method 0x0153b48a.
//
// Solidity: function coldPayableMemExpansion(address target) returns()
func (_Caller *CallerTransactorSession) ColdPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdPayableMemExpansion(&_Caller.TransactOpts, target)
}

// ColdPayableMemUnchanged is a paid mutator transaction binding the contract method 0xb3bbaaa3.
//
// Solidity: function coldPayableMemUnchanged(address target) returns()
func (_Caller *CallerTransactor) ColdPayableMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "coldPayableMemUnchanged", target)
}

// ColdPayableMemUnchanged is a paid mutator transaction binding the contract method 0xb3bbaaa3.
//
// Solidity: function coldPayableMemUnchanged(address target) returns()
func (_Caller *CallerSession) ColdPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdPayableMemUnchanged(&_Caller.TransactOpts, target)
}

// ColdPayableMemUnchanged is a paid mutator transaction binding the contract method 0xb3bbaaa3.
//
// Solidity: function coldPayableMemUnchanged(address target) returns()
func (_Caller *CallerTransactorSession) ColdPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.ColdPayableMemUnchanged(&_Caller.TransactOpts, target)
}

// WarmNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x61f8f2b8.
//
// Solidity: function warmNoTransferMemExpansion(address target) returns()
func (_Caller *CallerTransactor) WarmNoTransferMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "warmNoTransferMemExpansion", target)
}

// WarmNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x61f8f2b8.
//
// Solidity: function warmNoTransferMemExpansion(address target) returns()
func (_Caller *CallerSession) WarmNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmNoTransferMemExpansion(&_Caller.TransactOpts, target)
}

// WarmNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x61f8f2b8.
//
// Solidity: function warmNoTransferMemExpansion(address target) returns()
func (_Caller *CallerTransactorSession) WarmNoTransferMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmNoTransferMemExpansion(&_Caller.TransactOpts, target)
}

// WarmNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x7bde0d8b.
//
// Solidity: function warmNoTransferMemUnchanged(address target) returns()
func (_Caller *CallerTransactor) WarmNoTransferMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "warmNoTransferMemUnchanged", target)
}

// WarmNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x7bde0d8b.
//
// Solidity: function warmNoTransferMemUnchanged(address target) returns()
func (_Caller *CallerSession) WarmNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmNoTransferMemUnchanged(&_Caller.TransactOpts, target)
}

// WarmNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0x7bde0d8b.
//
// Solidity: function warmNoTransferMemUnchanged(address target) returns()
func (_Caller *CallerTransactorSession) WarmNoTransferMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmNoTransferMemUnchanged(&_Caller.TransactOpts, target)
}

// WarmPayableMemExpansion is a paid mutator transaction binding the contract method 0x06d668b9.
//
// Solidity: function warmPayableMemExpansion(address target) returns()
func (_Caller *CallerTransactor) WarmPayableMemExpansion(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "warmPayableMemExpansion", target)
}

// WarmPayableMemExpansion is a paid mutator transaction binding the contract method 0x06d668b9.
//
// Solidity: function warmPayableMemExpansion(address target) returns()
func (_Caller *CallerSession) WarmPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmPayableMemExpansion(&_Caller.TransactOpts, target)
}

// WarmPayableMemExpansion is a paid mutator transaction binding the contract method 0x06d668b9.
//
// Solidity: function warmPayableMemExpansion(address target) returns()
func (_Caller *CallerTransactorSession) WarmPayableMemExpansion(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmPayableMemExpansion(&_Caller.TransactOpts, target)
}

// WarmPayableMemUnchanged is a paid mutator transaction binding the contract method 0x692ef31d.
//
// Solidity: function warmPayableMemUnchanged(address target) returns()
func (_Caller *CallerTransactor) WarmPayableMemUnchanged(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "warmPayableMemUnchanged", target)
}

// WarmPayableMemUnchanged is a paid mutator transaction binding the contract method 0x692ef31d.
//
// Solidity: function warmPayableMemUnchanged(address target) returns()
func (_Caller *CallerSession) WarmPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmPayableMemUnchanged(&_Caller.TransactOpts, target)
}

// WarmPayableMemUnchanged is a paid mutator transaction binding the contract method 0x692ef31d.
//
// Solidity: function warmPayableMemUnchanged(address target) returns()
func (_Caller *CallerTransactorSession) WarmPayableMemUnchanged(target common.Address) (*types.Transaction, error) {
	return _Caller.Contract.WarmPayableMemUnchanged(&_Caller.TransactOpts, target)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Caller *CallerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Caller *CallerSession) Receive() (*types.Transaction, error) {
	return _Caller.Contract.Receive(&_Caller.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Caller *CallerTransactorSession) Receive() (*types.Transaction, error) {
	return _Caller.Contract.Receive(&_Caller.TransactOpts)
}

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noSpecials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102048061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80633fb5c1cb1461004e578063822ec8611461006a5780638381f58a14610074578063d09de08a14610092575b5f5ffd5b6100686004803603810190610063919061011b565b61009c565b005b6100726100a5565b005b61007c6100b1565b6040516100899190610155565b60405180910390f35b61009a6100b6565b005b805f8190555050565b61696961133701602081f35b5f5481565b5f5f5490506001816100c8919061019b565b90505f5f54905080826100db919061019b565b5f819055505050565b5f5ffd5b5f819050919050565b6100fa816100e8565b8114610104575f5ffd5b50565b5f81359050610115816100f1565b92915050565b5f602082840312156101305761012f6100e4565b5b5f61013d84828501610107565b91505092915050565b61014f816100e8565b82525050565b5f6020820190506101685f830184610146565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101a5826100e8565b91506101b0836100e8565b92508282019050808211156101c8576101c761016e565b5b9291505056fea2646970667358221220285973522b8ad7991c66c263fbd08eae79056517be22fa4eef779c111e61db3764736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Counter *CounterCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Counter *CounterSession) Number() (*big.Int, error) {
	return _Counter.Contract.Number(&_Counter.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_Counter *CounterCallerSession) Number() (*big.Int, error) {
	return _Counter.Contract.Number(&_Counter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// NoSpecials is a paid mutator transaction binding the contract method 0x822ec861.
//
// Solidity: function noSpecials() returns()
func (_Counter *CounterTransactor) NoSpecials(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "noSpecials")
}

// NoSpecials is a paid mutator transaction binding the contract method 0x822ec861.
//
// Solidity: function noSpecials() returns()
func (_Counter *CounterSession) NoSpecials() (*types.Transaction, error) {
	return _Counter.Contract.NoSpecials(&_Counter.TransactOpts)
}

// NoSpecials is a paid mutator transaction binding the contract method 0x822ec861.
//
// Solidity: function noSpecials() returns()
func (_Counter *CounterTransactorSession) NoSpecials() (*types.Transaction, error) {
	return _Counter.Contract.NoSpecials(&_Counter.TransactOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_Counter *CounterTransactor) SetNumber(opts *bind.TransactOpts, newNumber *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "setNumber", newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_Counter *CounterSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.SetNumber(&_Counter.TransactOpts, newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_Counter *CounterTransactorSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.SetNumber(&_Counter.TransactOpts, newNumber)
}

// CounterArrayMetaData contains all meta data concerning the CounterArray contract.
var CounterArrayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"counters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSlotKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slotKey\",\"type\":\"bytes32\"}],\"name\":\"refundFromCalldata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newCounters\",\"type\":\"uint256[]\"}],\"name\":\"setCounters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50601467ffffffffffffffff81111561002b5761002a6100da565b5b6040519080825280602002602001820160405280156100595781602001602082028036833780820191505090505b505f908051906020019061006e929190610074565b50610107565b828054828255905f5260205f209081019282156100ae579160200282015b828111156100ad578251825591602001919060010190610092565b5b5090506100bb91906100bf565b5090565b5b808211156100d6575f815f9055506001016100c0565b5090565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61064d806101145f395ff3fe608060405234801561000f575f5ffd5b5060043610610070575f3560e01c8063df4f5a721161004e578063df4f5a72146100ca578063fba7ffa3146100d4578063ffaab6381461010457610070565b80635db6c8f11461007457806373fd483614610090578063d540a097146100c0575b5f5ffd5b61008e60048036038101906100899190610302565b610120565b005b6100aa60048036038101906100a59190610360565b610126565b6040516100b7919061039a565b60405180910390f35b6100c861013d565b005b6100d261014c565b005b6100ee60048036038101906100e99190610360565b610202565b6040516100fb91906103c2565b60405180910390f35b61011e6004803603810190610119919061052b565b610221565b005b5f815550565b5f5f5f602081208481019250505080915050919050565b5f6020812080545f8255505050565b5f602081205f8155600181015f81556002820190505f81556003820190505f81556004820190505f81556005820190505f81556006820190505f81556007820190505f81556008820190505f81556009820190505f8155600a820190505f8155600b820190505f8155600c820190505f8155600d820190505f8155600e820190505f8155600f820190505f81556010820190505f81556011820190505f81556012820190505f81556013820190505f8155505050565b5f8181548110610210575f80fd5b905f5260205f20015f915090505481565b6014815114610265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025c906105cc565b60405180910390fd5b5f5f90505b60148110156102ba57818181518110610286576102856105ea565b5b60200260200101515f82815481106102a1576102a06105ea565b5b905f5260205f200181905550808060010191505061026a565b5050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6102e1816102cf565b81146102eb575f5ffd5b50565b5f813590506102fc816102d8565b92915050565b5f60208284031215610317576103166102c7565b5b5f610324848285016102ee565b91505092915050565b5f819050919050565b61033f8161032d565b8114610349575f5ffd5b50565b5f8135905061035a81610336565b92915050565b5f60208284031215610375576103746102c7565b5b5f6103828482850161034c565b91505092915050565b610394816102cf565b82525050565b5f6020820190506103ad5f83018461038b565b92915050565b6103bc8161032d565b82525050565b5f6020820190506103d55f8301846103b3565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610425826103df565b810181811067ffffffffffffffff82111715610444576104436103ef565b5b80604052505050565b5f6104566102be565b9050610462828261041c565b919050565b5f67ffffffffffffffff821115610481576104806103ef565b5b602082029050602081019050919050565b5f5ffd5b5f6104a86104a384610467565b61044d565b905080838252602082019050602084028301858111156104cb576104ca610492565b5b835b818110156104f457806104e0888261034c565b8452602084019350506020810190506104cd565b5050509392505050565b5f82601f830112610512576105116103db565b5b8135610522848260208601610496565b91505092915050565b5f602082840312156105405761053f6102c7565b5b5f82013567ffffffffffffffff81111561055d5761055c6102cb565b5b610569848285016104fe565b91505092915050565b5f82825260208201905092915050565b7f4172726179206d75737420626520323020656c656d656e7473206c6f6e6700005f82015250565b5f6105b6601e83610572565b91506105c182610582565b602082019050919050565b5f6020820190508181035f8301526105e3816105aa565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea26469706673582212206df5a2bc7463d52a9615f8e77dff16a3796c44ce3ecaeee54ede65114960539864736f6c634300081e0033",
}

// CounterArrayABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterArrayMetaData.ABI instead.
var CounterArrayABI = CounterArrayMetaData.ABI

// CounterArrayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterArrayMetaData.Bin instead.
var CounterArrayBin = CounterArrayMetaData.Bin

// DeployCounterArray deploys a new Ethereum contract, binding an instance of CounterArray to it.
func DeployCounterArray(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CounterArray, error) {
	parsed, err := CounterArrayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterArray{CounterArrayCaller: CounterArrayCaller{contract: contract}, CounterArrayTransactor: CounterArrayTransactor{contract: contract}, CounterArrayFilterer: CounterArrayFilterer{contract: contract}}, nil
}

// CounterArray is an auto generated Go binding around an Ethereum contract.
type CounterArray struct {
	CounterArrayCaller     // Read-only binding to the contract
	CounterArrayTransactor // Write-only binding to the contract
	CounterArrayFilterer   // Log filterer for contract events
}

// CounterArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterArrayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterArrayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterArraySession struct {
	Contract     *CounterArray     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterArrayCallerSession struct {
	Contract *CounterArrayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CounterArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterArrayTransactorSession struct {
	Contract     *CounterArrayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CounterArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterArrayRaw struct {
	Contract *CounterArray // Generic contract binding to access the raw methods on
}

// CounterArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterArrayCallerRaw struct {
	Contract *CounterArrayCaller // Generic read-only contract binding to access the raw methods on
}

// CounterArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterArrayTransactorRaw struct {
	Contract *CounterArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterArray creates a new instance of CounterArray, bound to a specific deployed contract.
func NewCounterArray(address common.Address, backend bind.ContractBackend) (*CounterArray, error) {
	contract, err := bindCounterArray(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterArray{CounterArrayCaller: CounterArrayCaller{contract: contract}, CounterArrayTransactor: CounterArrayTransactor{contract: contract}, CounterArrayFilterer: CounterArrayFilterer{contract: contract}}, nil
}

// NewCounterArrayCaller creates a new read-only instance of CounterArray, bound to a specific deployed contract.
func NewCounterArrayCaller(address common.Address, caller bind.ContractCaller) (*CounterArrayCaller, error) {
	contract, err := bindCounterArray(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterArrayCaller{contract: contract}, nil
}

// NewCounterArrayTransactor creates a new write-only instance of CounterArray, bound to a specific deployed contract.
func NewCounterArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterArrayTransactor, error) {
	contract, err := bindCounterArray(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterArrayTransactor{contract: contract}, nil
}

// NewCounterArrayFilterer creates a new log filterer instance of CounterArray, bound to a specific deployed contract.
func NewCounterArrayFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterArrayFilterer, error) {
	contract, err := bindCounterArray(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterArrayFilterer{contract: contract}, nil
}

// bindCounterArray binds a generic wrapper to an already deployed contract.
func bindCounterArray(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterArrayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterArray *CounterArrayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterArray.Contract.CounterArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterArray *CounterArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterArray.Contract.CounterArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterArray *CounterArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterArray.Contract.CounterArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterArray *CounterArrayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterArray.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterArray *CounterArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterArray.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterArray *CounterArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterArray.Contract.contract.Transact(opts, method, params...)
}

// Counters is a free data retrieval call binding the contract method 0xfba7ffa3.
//
// Solidity: function counters(uint256 ) view returns(uint256)
func (_CounterArray *CounterArrayCaller) Counters(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CounterArray.contract.Call(opts, &out, "counters", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counters is a free data retrieval call binding the contract method 0xfba7ffa3.
//
// Solidity: function counters(uint256 ) view returns(uint256)
func (_CounterArray *CounterArraySession) Counters(arg0 *big.Int) (*big.Int, error) {
	return _CounterArray.Contract.Counters(&_CounterArray.CallOpts, arg0)
}

// Counters is a free data retrieval call binding the contract method 0xfba7ffa3.
//
// Solidity: function counters(uint256 ) view returns(uint256)
func (_CounterArray *CounterArrayCallerSession) Counters(arg0 *big.Int) (*big.Int, error) {
	return _CounterArray.Contract.Counters(&_CounterArray.CallOpts, arg0)
}

// GetSlotKey is a free data retrieval call binding the contract method 0x73fd4836.
//
// Solidity: function getSlotKey(uint256 index) pure returns(bytes32)
func (_CounterArray *CounterArrayCaller) GetSlotKey(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CounterArray.contract.Call(opts, &out, "getSlotKey", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSlotKey is a free data retrieval call binding the contract method 0x73fd4836.
//
// Solidity: function getSlotKey(uint256 index) pure returns(bytes32)
func (_CounterArray *CounterArraySession) GetSlotKey(index *big.Int) ([32]byte, error) {
	return _CounterArray.Contract.GetSlotKey(&_CounterArray.CallOpts, index)
}

// GetSlotKey is a free data retrieval call binding the contract method 0x73fd4836.
//
// Solidity: function getSlotKey(uint256 index) pure returns(bytes32)
func (_CounterArray *CounterArrayCallerSession) GetSlotKey(index *big.Int) ([32]byte, error) {
	return _CounterArray.Contract.GetSlotKey(&_CounterArray.CallOpts, index)
}

// Refund1 is a paid mutator transaction binding the contract method 0xd540a097.
//
// Solidity: function refund1() returns()
func (_CounterArray *CounterArrayTransactor) Refund1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterArray.contract.Transact(opts, "refund1")
}

// Refund1 is a paid mutator transaction binding the contract method 0xd540a097.
//
// Solidity: function refund1() returns()
func (_CounterArray *CounterArraySession) Refund1() (*types.Transaction, error) {
	return _CounterArray.Contract.Refund1(&_CounterArray.TransactOpts)
}

// Refund1 is a paid mutator transaction binding the contract method 0xd540a097.
//
// Solidity: function refund1() returns()
func (_CounterArray *CounterArrayTransactorSession) Refund1() (*types.Transaction, error) {
	return _CounterArray.Contract.Refund1(&_CounterArray.TransactOpts)
}

// RefundFromCalldata is a paid mutator transaction binding the contract method 0x5db6c8f1.
//
// Solidity: function refundFromCalldata(bytes32 slotKey) returns()
func (_CounterArray *CounterArrayTransactor) RefundFromCalldata(opts *bind.TransactOpts, slotKey [32]byte) (*types.Transaction, error) {
	return _CounterArray.contract.Transact(opts, "refundFromCalldata", slotKey)
}

// RefundFromCalldata is a paid mutator transaction binding the contract method 0x5db6c8f1.
//
// Solidity: function refundFromCalldata(bytes32 slotKey) returns()
func (_CounterArray *CounterArraySession) RefundFromCalldata(slotKey [32]byte) (*types.Transaction, error) {
	return _CounterArray.Contract.RefundFromCalldata(&_CounterArray.TransactOpts, slotKey)
}

// RefundFromCalldata is a paid mutator transaction binding the contract method 0x5db6c8f1.
//
// Solidity: function refundFromCalldata(bytes32 slotKey) returns()
func (_CounterArray *CounterArrayTransactorSession) RefundFromCalldata(slotKey [32]byte) (*types.Transaction, error) {
	return _CounterArray.Contract.RefundFromCalldata(&_CounterArray.TransactOpts, slotKey)
}

// Refunder is a paid mutator transaction binding the contract method 0xdf4f5a72.
//
// Solidity: function refunder() returns()
func (_CounterArray *CounterArrayTransactor) Refunder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterArray.contract.Transact(opts, "refunder")
}

// Refunder is a paid mutator transaction binding the contract method 0xdf4f5a72.
//
// Solidity: function refunder() returns()
func (_CounterArray *CounterArraySession) Refunder() (*types.Transaction, error) {
	return _CounterArray.Contract.Refunder(&_CounterArray.TransactOpts)
}

// Refunder is a paid mutator transaction binding the contract method 0xdf4f5a72.
//
// Solidity: function refunder() returns()
func (_CounterArray *CounterArrayTransactorSession) Refunder() (*types.Transaction, error) {
	return _CounterArray.Contract.Refunder(&_CounterArray.TransactOpts)
}

// SetCounters is a paid mutator transaction binding the contract method 0xffaab638.
//
// Solidity: function setCounters(uint256[] newCounters) returns()
func (_CounterArray *CounterArrayTransactor) SetCounters(opts *bind.TransactOpts, newCounters []*big.Int) (*types.Transaction, error) {
	return _CounterArray.contract.Transact(opts, "setCounters", newCounters)
}

// SetCounters is a paid mutator transaction binding the contract method 0xffaab638.
//
// Solidity: function setCounters(uint256[] newCounters) returns()
func (_CounterArray *CounterArraySession) SetCounters(newCounters []*big.Int) (*types.Transaction, error) {
	return _CounterArray.Contract.SetCounters(&_CounterArray.TransactOpts, newCounters)
}

// SetCounters is a paid mutator transaction binding the contract method 0xffaab638.
//
// Solidity: function setCounters(uint256[] newCounters) returns()
func (_CounterArray *CounterArrayTransactorSession) SetCounters(newCounters []*big.Int) (*types.Transaction, error) {
	return _CounterArray.Contract.SetCounters(&_CounterArray.TransactOpts, newCounters)
}

// CreateeMetaData contains all meta data concerning the Createe contract.
var CreateeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526040516101473803806101478339818101604052810190602391906060565b805f81905550506086565b5f5ffd5b5f819050919050565b6042816032565b8114604b575f5ffd5b50565b5f81519050605a81603b565b92915050565b5f602082840312156072576071602e565b5b5f607d84828501604e565b91505092915050565b60b5806100925f395ff3fe608060405260043610601e575f3560e01c80630c55699c146028576024565b36602457005b5f5ffd5b3480156032575f5ffd5b506039604d565b604051604491906068565b60405180910390f35b5f5481565b5f819050919050565b6062816052565b82525050565b5f60208201905060795f830184605b565b9291505056fea2646970667358221220ef636c07c618bb2f0b1cd146afa834e9bebeacc852d18c65db3542c405a6e25064736f6c634300081e0033",
}

// CreateeABI is the input ABI used to generate the binding from.
// Deprecated: Use CreateeMetaData.ABI instead.
var CreateeABI = CreateeMetaData.ABI

// CreateeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreateeMetaData.Bin instead.
var CreateeBin = CreateeMetaData.Bin

// DeployCreatee deploys a new Ethereum contract, binding an instance of Createe to it.
func DeployCreatee(auth *bind.TransactOpts, backend bind.ContractBackend, _x *big.Int) (common.Address, *types.Transaction, *Createe, error) {
	parsed, err := CreateeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreateeBin), backend, _x)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Createe{CreateeCaller: CreateeCaller{contract: contract}, CreateeTransactor: CreateeTransactor{contract: contract}, CreateeFilterer: CreateeFilterer{contract: contract}}, nil
}

// Createe is an auto generated Go binding around an Ethereum contract.
type Createe struct {
	CreateeCaller     // Read-only binding to the contract
	CreateeTransactor // Write-only binding to the contract
	CreateeFilterer   // Log filterer for contract events
}

// CreateeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreateeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreateeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreateeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreateeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreateeSession struct {
	Contract     *Createe          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreateeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreateeCallerSession struct {
	Contract *CreateeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CreateeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreateeTransactorSession struct {
	Contract     *CreateeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CreateeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreateeRaw struct {
	Contract *Createe // Generic contract binding to access the raw methods on
}

// CreateeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreateeCallerRaw struct {
	Contract *CreateeCaller // Generic read-only contract binding to access the raw methods on
}

// CreateeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreateeTransactorRaw struct {
	Contract *CreateeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreatee creates a new instance of Createe, bound to a specific deployed contract.
func NewCreatee(address common.Address, backend bind.ContractBackend) (*Createe, error) {
	contract, err := bindCreatee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Createe{CreateeCaller: CreateeCaller{contract: contract}, CreateeTransactor: CreateeTransactor{contract: contract}, CreateeFilterer: CreateeFilterer{contract: contract}}, nil
}

// NewCreateeCaller creates a new read-only instance of Createe, bound to a specific deployed contract.
func NewCreateeCaller(address common.Address, caller bind.ContractCaller) (*CreateeCaller, error) {
	contract, err := bindCreatee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreateeCaller{contract: contract}, nil
}

// NewCreateeTransactor creates a new write-only instance of Createe, bound to a specific deployed contract.
func NewCreateeTransactor(address common.Address, transactor bind.ContractTransactor) (*CreateeTransactor, error) {
	contract, err := bindCreatee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreateeTransactor{contract: contract}, nil
}

// NewCreateeFilterer creates a new log filterer instance of Createe, bound to a specific deployed contract.
func NewCreateeFilterer(address common.Address, filterer bind.ContractFilterer) (*CreateeFilterer, error) {
	contract, err := bindCreatee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreateeFilterer{contract: contract}, nil
}

// bindCreatee binds a generic wrapper to an already deployed contract.
func bindCreatee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreateeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Createe *CreateeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Createe.Contract.CreateeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Createe *CreateeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Createe.Contract.CreateeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Createe *CreateeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Createe.Contract.CreateeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Createe *CreateeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Createe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Createe *CreateeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Createe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Createe *CreateeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Createe.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Createe *CreateeCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Createe.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Createe *CreateeSession) X() (*big.Int, error) {
	return _Createe.Contract.X(&_Createe.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Createe *CreateeCallerSession) X() (*big.Int, error) {
	return _Createe.Contract.X(&_Createe.CallOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Createe *CreateeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Createe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Createe *CreateeSession) Receive() (*types.Transaction, error) {
	return _Createe.Contract.Receive(&_Createe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Createe *CreateeTransactorSession) Receive() (*types.Transaction, error) {
	return _Createe.Contract.Receive(&_Createe.TransactOpts)
}

// CreatorMetaData contains all meta data concerning the Creator contract.
var CreatorMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"createNoTransferMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createNoTransferMemUnchanged\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPayableMemExpansion\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPayableMemUnchanged\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103848061001c5f395ff3fe608060405260043610610042575f3560e01c806312dc999b1461004d578063141ca64b14610057578063aeb6cdec1461006d578063c8c6a5661461009757610049565b3661004957005b5f5ffd5b6100556100b5565b005b348015610062575f5ffd5b5061006b6100c3565b005b348015610078575f5ffd5b506100816100ce565b60405161008e91906101bd565b60405180910390f35b61009f6100df565b6040516100ac91906101bd565b60405180910390f35b6100c1620694206100f3565b565b6100cc5f6100f3565b565b5f6100da5f604261012f565b905090565b5f6100ee62069420604261012f565b905090565b5f6040518060200161010490610171565b6020820181038252601f19601f82011660405250905059602082820301806020840185f05050505050565b5f5f838360405161013f90610171565b61014991906101ee565b6040518091039082f0905080158015610164573d5f5f3e3d5ffd5b5090508091505092915050565b6101478061020883390190565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101a78261017e565b9050919050565b6101b78161019d565b82525050565b5f6020820190506101d05f8301846101ae565b92915050565b5f819050919050565b6101e8816101d6565b82525050565b5f6020820190506102015f8301846101df565b9291505056fe60806040526040516101473803806101478339818101604052810190602391906060565b805f81905550506086565b5f5ffd5b5f819050919050565b6042816032565b8114604b575f5ffd5b50565b5f81519050605a81603b565b92915050565b5f602082840312156072576071602e565b5b5f607d84828501604e565b91505092915050565b60b5806100925f395ff3fe608060405260043610601e575f3560e01c80630c55699c146028576024565b36602457005b5f5ffd5b3480156032575f5ffd5b506039604d565b604051604491906068565b60405180910390f35b5f5481565b5f819050919050565b6062816052565b82525050565b5f60208201905060795f830184605b565b9291505056fea2646970667358221220ef636c07c618bb2f0b1cd146afa834e9bebeacc852d18c65db3542c405a6e25064736f6c634300081e0033a2646970667358221220a3068786883effe165d7e5cb78f8e2863bfc2d3a0e7dcc87f978e90945dedc6364736f6c634300081e0033",
}

// CreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatorMetaData.ABI instead.
var CreatorABI = CreatorMetaData.ABI

// CreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreatorMetaData.Bin instead.
var CreatorBin = CreatorMetaData.Bin

// DeployCreator deploys a new Ethereum contract, binding an instance of Creator to it.
func DeployCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Creator, error) {
	parsed, err := CreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Creator{CreatorCaller: CreatorCaller{contract: contract}, CreatorTransactor: CreatorTransactor{contract: contract}, CreatorFilterer: CreatorFilterer{contract: contract}}, nil
}

// Creator is an auto generated Go binding around an Ethereum contract.
type Creator struct {
	CreatorCaller     // Read-only binding to the contract
	CreatorTransactor // Write-only binding to the contract
	CreatorFilterer   // Log filterer for contract events
}

// CreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorSession struct {
	Contract     *Creator          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorCallerSession struct {
	Contract *CreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorTransactorSession struct {
	Contract     *CreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorRaw struct {
	Contract *Creator // Generic contract binding to access the raw methods on
}

// CreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorCallerRaw struct {
	Contract *CreatorCaller // Generic read-only contract binding to access the raw methods on
}

// CreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorTransactorRaw struct {
	Contract *CreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreator creates a new instance of Creator, bound to a specific deployed contract.
func NewCreator(address common.Address, backend bind.ContractBackend) (*Creator, error) {
	contract, err := bindCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Creator{CreatorCaller: CreatorCaller{contract: contract}, CreatorTransactor: CreatorTransactor{contract: contract}, CreatorFilterer: CreatorFilterer{contract: contract}}, nil
}

// NewCreatorCaller creates a new read-only instance of Creator, bound to a specific deployed contract.
func NewCreatorCaller(address common.Address, caller bind.ContractCaller) (*CreatorCaller, error) {
	contract, err := bindCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCaller{contract: contract}, nil
}

// NewCreatorTransactor creates a new write-only instance of Creator, bound to a specific deployed contract.
func NewCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreatorTransactor, error) {
	contract, err := bindCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTransactor{contract: contract}, nil
}

// NewCreatorFilterer creates a new log filterer instance of Creator, bound to a specific deployed contract.
func NewCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreatorFilterer, error) {
	contract, err := bindCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorFilterer{contract: contract}, nil
}

// bindCreator binds a generic wrapper to an already deployed contract.
func bindCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creator *CreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creator.Contract.CreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creator *CreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.Contract.CreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creator *CreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creator.Contract.CreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creator *CreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creator *CreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creator *CreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creator.Contract.contract.Transact(opts, method, params...)
}

// CreateNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x141ca64b.
//
// Solidity: function createNoTransferMemExpansion() returns()
func (_Creator *CreatorTransactor) CreateNoTransferMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "createNoTransferMemExpansion")
}

// CreateNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x141ca64b.
//
// Solidity: function createNoTransferMemExpansion() returns()
func (_Creator *CreatorSession) CreateNoTransferMemExpansion() (*types.Transaction, error) {
	return _Creator.Contract.CreateNoTransferMemExpansion(&_Creator.TransactOpts)
}

// CreateNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x141ca64b.
//
// Solidity: function createNoTransferMemExpansion() returns()
func (_Creator *CreatorTransactorSession) CreateNoTransferMemExpansion() (*types.Transaction, error) {
	return _Creator.Contract.CreateNoTransferMemExpansion(&_Creator.TransactOpts)
}

// CreateNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0xaeb6cdec.
//
// Solidity: function createNoTransferMemUnchanged() returns(address)
func (_Creator *CreatorTransactor) CreateNoTransferMemUnchanged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "createNoTransferMemUnchanged")
}

// CreateNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0xaeb6cdec.
//
// Solidity: function createNoTransferMemUnchanged() returns(address)
func (_Creator *CreatorSession) CreateNoTransferMemUnchanged() (*types.Transaction, error) {
	return _Creator.Contract.CreateNoTransferMemUnchanged(&_Creator.TransactOpts)
}

// CreateNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0xaeb6cdec.
//
// Solidity: function createNoTransferMemUnchanged() returns(address)
func (_Creator *CreatorTransactorSession) CreateNoTransferMemUnchanged() (*types.Transaction, error) {
	return _Creator.Contract.CreateNoTransferMemUnchanged(&_Creator.TransactOpts)
}

// CreatePayableMemExpansion is a paid mutator transaction binding the contract method 0x12dc999b.
//
// Solidity: function createPayableMemExpansion() payable returns()
func (_Creator *CreatorTransactor) CreatePayableMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "createPayableMemExpansion")
}

// CreatePayableMemExpansion is a paid mutator transaction binding the contract method 0x12dc999b.
//
// Solidity: function createPayableMemExpansion() payable returns()
func (_Creator *CreatorSession) CreatePayableMemExpansion() (*types.Transaction, error) {
	return _Creator.Contract.CreatePayableMemExpansion(&_Creator.TransactOpts)
}

// CreatePayableMemExpansion is a paid mutator transaction binding the contract method 0x12dc999b.
//
// Solidity: function createPayableMemExpansion() payable returns()
func (_Creator *CreatorTransactorSession) CreatePayableMemExpansion() (*types.Transaction, error) {
	return _Creator.Contract.CreatePayableMemExpansion(&_Creator.TransactOpts)
}

// CreatePayableMemUnchanged is a paid mutator transaction binding the contract method 0xc8c6a566.
//
// Solidity: function createPayableMemUnchanged() payable returns(address)
func (_Creator *CreatorTransactor) CreatePayableMemUnchanged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "createPayableMemUnchanged")
}

// CreatePayableMemUnchanged is a paid mutator transaction binding the contract method 0xc8c6a566.
//
// Solidity: function createPayableMemUnchanged() payable returns(address)
func (_Creator *CreatorSession) CreatePayableMemUnchanged() (*types.Transaction, error) {
	return _Creator.Contract.CreatePayableMemUnchanged(&_Creator.TransactOpts)
}

// CreatePayableMemUnchanged is a paid mutator transaction binding the contract method 0xc8c6a566.
//
// Solidity: function createPayableMemUnchanged() payable returns(address)
func (_Creator *CreatorTransactorSession) CreatePayableMemUnchanged() (*types.Transaction, error) {
	return _Creator.Contract.CreatePayableMemUnchanged(&_Creator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Creator *CreatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Creator *CreatorSession) Receive() (*types.Transaction, error) {
	return _Creator.Contract.Receive(&_Creator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Creator *CreatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Creator.Contract.Receive(&_Creator.TransactOpts)
}

// CreatorTwoMetaData contains all meta data concerning the CreatorTwo contract.
var CreatorTwoMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"createTwoNoTransferMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"createTwoNoTransferMemUnchanged\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"createTwoPayableMemExpansion\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"createTwoPayableMemUnchanged\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104418061001c5f395ff3fe608060405260043610610042575f3560e01c80634d6aa27d1461004d578063924053b814610075578063ae7f300b146100a5578063d2f40a00146100c157610049565b3661004957005b5f5ffd5b348015610058575f5ffd5b50610073600480360381019061006e9190610210565b6100fd565b005b61008f600480360381019061008a9190610210565b61010a565b60405161009c919061027a565b60405180910390f35b6100bf60048036038101906100ba9190610210565b610121565b005b3480156100cc575f5ffd5b506100e760048036038101906100e29190610210565b610131565b6040516100f4919061027a565b60405180910390f35b6101075f82610145565b50565b5f61011a62069420604284610183565b9050919050565b61012e6204206982610145565b50565b5f61013e5f604284610183565b9050919050565b5f60405180602001610156906101cc565b6020820181038252601f19601f8201166040525090505960208282030183816020850187f5505050505050565b5f5f84839085604051610195906101cc565b61019f91906102ab565b82906040518091039083f590509050801580156101be573d5f5f3e3d5ffd5b509050809150509392505050565b610147806102c583390190565b5f5ffd5b5f819050919050565b6101ef816101dd565b81146101f9575f5ffd5b50565b5f8135905061020a816101e6565b92915050565b5f60208284031215610225576102246101d9565b5b5f610232848285016101fc565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102648261023b565b9050919050565b6102748161025a565b82525050565b5f60208201905061028d5f83018461026b565b92915050565b5f819050919050565b6102a581610293565b82525050565b5f6020820190506102be5f83018461029c565b9291505056fe60806040526040516101473803806101478339818101604052810190602391906060565b805f81905550506086565b5f5ffd5b5f819050919050565b6042816032565b8114604b575f5ffd5b50565b5f81519050605a81603b565b92915050565b5f602082840312156072576071602e565b5b5f607d84828501604e565b91505092915050565b60b5806100925f395ff3fe608060405260043610601e575f3560e01c80630c55699c146028576024565b36602457005b5f5ffd5b3480156032575f5ffd5b506039604d565b604051604491906068565b60405180910390f35b5f5481565b5f819050919050565b6062816052565b82525050565b5f60208201905060795f830184605b565b9291505056fea2646970667358221220ef636c07c618bb2f0b1cd146afa834e9bebeacc852d18c65db3542c405a6e25064736f6c634300081e0033a26469706673582212200808159d83d80f474065391dedea0a209ae3a1b2e18ebdb3db5a72649750c80a64736f6c634300081e0033",
}

// CreatorTwoABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatorTwoMetaData.ABI instead.
var CreatorTwoABI = CreatorTwoMetaData.ABI

// CreatorTwoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreatorTwoMetaData.Bin instead.
var CreatorTwoBin = CreatorTwoMetaData.Bin

// DeployCreatorTwo deploys a new Ethereum contract, binding an instance of CreatorTwo to it.
func DeployCreatorTwo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CreatorTwo, error) {
	parsed, err := CreatorTwoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreatorTwoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreatorTwo{CreatorTwoCaller: CreatorTwoCaller{contract: contract}, CreatorTwoTransactor: CreatorTwoTransactor{contract: contract}, CreatorTwoFilterer: CreatorTwoFilterer{contract: contract}}, nil
}

// CreatorTwo is an auto generated Go binding around an Ethereum contract.
type CreatorTwo struct {
	CreatorTwoCaller     // Read-only binding to the contract
	CreatorTwoTransactor // Write-only binding to the contract
	CreatorTwoFilterer   // Log filterer for contract events
}

// CreatorTwoCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorTwoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTwoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorTwoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTwoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorTwoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTwoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorTwoSession struct {
	Contract     *CreatorTwo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreatorTwoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorTwoCallerSession struct {
	Contract *CreatorTwoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CreatorTwoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorTwoTransactorSession struct {
	Contract     *CreatorTwoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CreatorTwoRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorTwoRaw struct {
	Contract *CreatorTwo // Generic contract binding to access the raw methods on
}

// CreatorTwoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorTwoCallerRaw struct {
	Contract *CreatorTwoCaller // Generic read-only contract binding to access the raw methods on
}

// CreatorTwoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorTwoTransactorRaw struct {
	Contract *CreatorTwoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreatorTwo creates a new instance of CreatorTwo, bound to a specific deployed contract.
func NewCreatorTwo(address common.Address, backend bind.ContractBackend) (*CreatorTwo, error) {
	contract, err := bindCreatorTwo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreatorTwo{CreatorTwoCaller: CreatorTwoCaller{contract: contract}, CreatorTwoTransactor: CreatorTwoTransactor{contract: contract}, CreatorTwoFilterer: CreatorTwoFilterer{contract: contract}}, nil
}

// NewCreatorTwoCaller creates a new read-only instance of CreatorTwo, bound to a specific deployed contract.
func NewCreatorTwoCaller(address common.Address, caller bind.ContractCaller) (*CreatorTwoCaller, error) {
	contract, err := bindCreatorTwo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTwoCaller{contract: contract}, nil
}

// NewCreatorTwoTransactor creates a new write-only instance of CreatorTwo, bound to a specific deployed contract.
func NewCreatorTwoTransactor(address common.Address, transactor bind.ContractTransactor) (*CreatorTwoTransactor, error) {
	contract, err := bindCreatorTwo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTwoTransactor{contract: contract}, nil
}

// NewCreatorTwoFilterer creates a new log filterer instance of CreatorTwo, bound to a specific deployed contract.
func NewCreatorTwoFilterer(address common.Address, filterer bind.ContractFilterer) (*CreatorTwoFilterer, error) {
	contract, err := bindCreatorTwo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorTwoFilterer{contract: contract}, nil
}

// bindCreatorTwo binds a generic wrapper to an already deployed contract.
func bindCreatorTwo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreatorTwoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorTwo *CreatorTwoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorTwo.Contract.CreatorTwoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorTwo *CreatorTwoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreatorTwoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorTwo *CreatorTwoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreatorTwoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorTwo *CreatorTwoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorTwo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorTwo *CreatorTwoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorTwo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorTwo *CreatorTwoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorTwo.Contract.contract.Transact(opts, method, params...)
}

// CreateTwoNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x4d6aa27d.
//
// Solidity: function createTwoNoTransferMemExpansion(bytes32 _salt) returns()
func (_CreatorTwo *CreatorTwoTransactor) CreateTwoNoTransferMemExpansion(opts *bind.TransactOpts, _salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.contract.Transact(opts, "createTwoNoTransferMemExpansion", _salt)
}

// CreateTwoNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x4d6aa27d.
//
// Solidity: function createTwoNoTransferMemExpansion(bytes32 _salt) returns()
func (_CreatorTwo *CreatorTwoSession) CreateTwoNoTransferMemExpansion(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoNoTransferMemExpansion(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoNoTransferMemExpansion is a paid mutator transaction binding the contract method 0x4d6aa27d.
//
// Solidity: function createTwoNoTransferMemExpansion(bytes32 _salt) returns()
func (_CreatorTwo *CreatorTwoTransactorSession) CreateTwoNoTransferMemExpansion(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoNoTransferMemExpansion(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0xd2f40a00.
//
// Solidity: function createTwoNoTransferMemUnchanged(bytes32 _salt) returns(address)
func (_CreatorTwo *CreatorTwoTransactor) CreateTwoNoTransferMemUnchanged(opts *bind.TransactOpts, _salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.contract.Transact(opts, "createTwoNoTransferMemUnchanged", _salt)
}

// CreateTwoNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0xd2f40a00.
//
// Solidity: function createTwoNoTransferMemUnchanged(bytes32 _salt) returns(address)
func (_CreatorTwo *CreatorTwoSession) CreateTwoNoTransferMemUnchanged(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoNoTransferMemUnchanged(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoNoTransferMemUnchanged is a paid mutator transaction binding the contract method 0xd2f40a00.
//
// Solidity: function createTwoNoTransferMemUnchanged(bytes32 _salt) returns(address)
func (_CreatorTwo *CreatorTwoTransactorSession) CreateTwoNoTransferMemUnchanged(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoNoTransferMemUnchanged(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoPayableMemExpansion is a paid mutator transaction binding the contract method 0xae7f300b.
//
// Solidity: function createTwoPayableMemExpansion(bytes32 _salt) payable returns()
func (_CreatorTwo *CreatorTwoTransactor) CreateTwoPayableMemExpansion(opts *bind.TransactOpts, _salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.contract.Transact(opts, "createTwoPayableMemExpansion", _salt)
}

// CreateTwoPayableMemExpansion is a paid mutator transaction binding the contract method 0xae7f300b.
//
// Solidity: function createTwoPayableMemExpansion(bytes32 _salt) payable returns()
func (_CreatorTwo *CreatorTwoSession) CreateTwoPayableMemExpansion(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoPayableMemExpansion(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoPayableMemExpansion is a paid mutator transaction binding the contract method 0xae7f300b.
//
// Solidity: function createTwoPayableMemExpansion(bytes32 _salt) payable returns()
func (_CreatorTwo *CreatorTwoTransactorSession) CreateTwoPayableMemExpansion(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoPayableMemExpansion(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoPayableMemUnchanged is a paid mutator transaction binding the contract method 0x924053b8.
//
// Solidity: function createTwoPayableMemUnchanged(bytes32 _salt) payable returns(address)
func (_CreatorTwo *CreatorTwoTransactor) CreateTwoPayableMemUnchanged(opts *bind.TransactOpts, _salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.contract.Transact(opts, "createTwoPayableMemUnchanged", _salt)
}

// CreateTwoPayableMemUnchanged is a paid mutator transaction binding the contract method 0x924053b8.
//
// Solidity: function createTwoPayableMemUnchanged(bytes32 _salt) payable returns(address)
func (_CreatorTwo *CreatorTwoSession) CreateTwoPayableMemUnchanged(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoPayableMemUnchanged(&_CreatorTwo.TransactOpts, _salt)
}

// CreateTwoPayableMemUnchanged is a paid mutator transaction binding the contract method 0x924053b8.
//
// Solidity: function createTwoPayableMemUnchanged(bytes32 _salt) payable returns(address)
func (_CreatorTwo *CreatorTwoTransactorSession) CreateTwoPayableMemUnchanged(_salt [32]byte) (*types.Transaction, error) {
	return _CreatorTwo.Contract.CreateTwoPayableMemUnchanged(&_CreatorTwo.TransactOpts, _salt)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CreatorTwo *CreatorTwoTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorTwo.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CreatorTwo *CreatorTwoSession) Receive() (*types.Transaction, error) {
	return _CreatorTwo.Contract.Receive(&_CreatorTwo.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CreatorTwo *CreatorTwoTransactorSession) Receive() (*types.Transaction, error) {
	return _CreatorTwo.Contract.Receive(&_CreatorTwo.TransactOpts)
}

// DelegateCalleeMetaData contains all meta data concerning the DelegateCallee contract.
var DelegateCalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061018e8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80633fa4f245146100385780635524107714610056575b5f5ffd5b610040610086565b60405161004d91906100b3565b60405180910390f35b610070600480360381019061006b91906100fa565b61008b565b60405161007d919061013f565b60405180910390f35b5f5481565b5f815f8190555060019050919050565b5f819050919050565b6100ad8161009b565b82525050565b5f6020820190506100c65f8301846100a4565b92915050565b5f5ffd5b6100d98161009b565b81146100e3575f5ffd5b50565b5f813590506100f4816100d0565b92915050565b5f6020828403121561010f5761010e6100cc565b5b5f61011c848285016100e6565b91505092915050565b5f8115159050919050565b61013981610125565b82525050565b5f6020820190506101525f830184610130565b9291505056fea26469706673582212207603e2af7665d6d3d74a8db4dad59f45c5d7a6c2db92d03711964de8a529d5e864736f6c634300081e0033",
}

// DelegateCalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegateCalleeMetaData.ABI instead.
var DelegateCalleeABI = DelegateCalleeMetaData.ABI

// DelegateCalleeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegateCalleeMetaData.Bin instead.
var DelegateCalleeBin = DelegateCalleeMetaData.Bin

// DeployDelegateCallee deploys a new Ethereum contract, binding an instance of DelegateCallee to it.
func DeployDelegateCallee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DelegateCallee, error) {
	parsed, err := DelegateCalleeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegateCalleeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelegateCallee{DelegateCalleeCaller: DelegateCalleeCaller{contract: contract}, DelegateCalleeTransactor: DelegateCalleeTransactor{contract: contract}, DelegateCalleeFilterer: DelegateCalleeFilterer{contract: contract}}, nil
}

// DelegateCallee is an auto generated Go binding around an Ethereum contract.
type DelegateCallee struct {
	DelegateCalleeCaller     // Read-only binding to the contract
	DelegateCalleeTransactor // Write-only binding to the contract
	DelegateCalleeFilterer   // Log filterer for contract events
}

// DelegateCalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegateCalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegateCalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegateCalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegateCalleeSession struct {
	Contract     *DelegateCallee   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegateCalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegateCalleeCallerSession struct {
	Contract *DelegateCalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DelegateCalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegateCalleeTransactorSession struct {
	Contract     *DelegateCalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DelegateCalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegateCalleeRaw struct {
	Contract *DelegateCallee // Generic contract binding to access the raw methods on
}

// DelegateCalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegateCalleeCallerRaw struct {
	Contract *DelegateCalleeCaller // Generic read-only contract binding to access the raw methods on
}

// DelegateCalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegateCalleeTransactorRaw struct {
	Contract *DelegateCalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegateCallee creates a new instance of DelegateCallee, bound to a specific deployed contract.
func NewDelegateCallee(address common.Address, backend bind.ContractBackend) (*DelegateCallee, error) {
	contract, err := bindDelegateCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegateCallee{DelegateCalleeCaller: DelegateCalleeCaller{contract: contract}, DelegateCalleeTransactor: DelegateCalleeTransactor{contract: contract}, DelegateCalleeFilterer: DelegateCalleeFilterer{contract: contract}}, nil
}

// NewDelegateCalleeCaller creates a new read-only instance of DelegateCallee, bound to a specific deployed contract.
func NewDelegateCalleeCaller(address common.Address, caller bind.ContractCaller) (*DelegateCalleeCaller, error) {
	contract, err := bindDelegateCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCalleeCaller{contract: contract}, nil
}

// NewDelegateCalleeTransactor creates a new write-only instance of DelegateCallee, bound to a specific deployed contract.
func NewDelegateCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegateCalleeTransactor, error) {
	contract, err := bindDelegateCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCalleeTransactor{contract: contract}, nil
}

// NewDelegateCalleeFilterer creates a new log filterer instance of DelegateCallee, bound to a specific deployed contract.
func NewDelegateCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegateCalleeFilterer, error) {
	contract, err := bindDelegateCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegateCalleeFilterer{contract: contract}, nil
}

// bindDelegateCallee binds a generic wrapper to an already deployed contract.
func bindDelegateCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegateCalleeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateCallee *DelegateCalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegateCallee.Contract.DelegateCalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateCallee *DelegateCalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCallee.Contract.DelegateCalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateCallee *DelegateCalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateCallee.Contract.DelegateCalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateCallee *DelegateCalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegateCallee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateCallee *DelegateCalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCallee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateCallee *DelegateCalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateCallee.Contract.contract.Transact(opts, method, params...)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_DelegateCallee *DelegateCalleeCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegateCallee.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_DelegateCallee *DelegateCalleeSession) Value() (*big.Int, error) {
	return _DelegateCallee.Contract.Value(&_DelegateCallee.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_DelegateCallee *DelegateCalleeCallerSession) Value() (*big.Int, error) {
	return _DelegateCallee.Contract.Value(&_DelegateCallee.CallOpts)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns(bool)
func (_DelegateCallee *DelegateCalleeTransactor) SetValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DelegateCallee.contract.Transact(opts, "setValue", _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns(bool)
func (_DelegateCallee *DelegateCalleeSession) SetValue(_value *big.Int) (*types.Transaction, error) {
	return _DelegateCallee.Contract.SetValue(&_DelegateCallee.TransactOpts, _value)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 _value) returns(bool)
func (_DelegateCallee *DelegateCalleeTransactorSession) SetValue(_value *big.Int) (*types.Transaction, error) {
	return _DelegateCallee.Contract.SetValue(&_DelegateCallee.TransactOpts, _value)
}

// DelegateCallerMetaData contains all meta data concerning the DelegateCaller contract.
var DelegateCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallEmptyCold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallEmptyColdMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallEmptyWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallEmptyWarmMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallNonEmptyCold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallNonEmptyColdMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallNonEmptyWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testDelegateCallNonEmptyWarmMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061096a8061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a1575f3560e01c8063648a6be811610074578063737c29d011610059578063737c29d01461014f578063978c93bc1461016b578063f55b199814610187576100a1565b8063648a6be814610117578063710814a914610133576100a1565b80632194badc146100a55780633fa4f245146100c15780634289f2a8146100df57806348112963146100fb575b5f5ffd5b6100bf60048036038101906100ba9190610780565b6101a3565b005b6100c96101d1565b6040516100d691906107c3565b60405180910390f35b6100f960048036038101906100f49190610780565b6101d6565b005b61011560048036038101906101109190610780565b61024d565b005b610131600480360381019061012c9190610780565b610354565b005b61014d60048036038101906101489190610780565b610411565b005b61016960048036038101906101649190610780565b610501565b005b61018560048036038101906101809190610780565b610556565b005b6101a1600480360381019061019c9190610780565b6105f4565b005b5f595a6040826040848785f49250505080156101c55760035f819055506101cd565b60045f819055505b5050565b5f5481565b5f8173ffffffffffffffffffffffffffffffffffffffff166040516101fa90610809565b5f60405180830381855af49150503d805f8114610232576040519150601f19603f3d011682016040523d82523d5f602084013e610237565b606091505b5050905080156102495760045f819055505b5050565b5f8173ffffffffffffffffffffffffffffffffffffffff16602a604051602401610277919061086b565b6040516020818303038152906040527f55241077000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161030191906108cc565b5f60405180830381855af49150503d805f8114610339576040519150601f19603f3d011682016040523d82523d5f602084013e61033e565b606091505b5050905080156103505760045f819055505b5050565b5f602a604051602401610367919061086b565b6040516020818303038152906040527f55241077000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f815190505f595a60408285602088018985f492505050801561040b5760045f819055505b50505050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f602b60405160240161043f919061091b565b6040516020818303038152906040527f55241077000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f815190505f595a60408285602088018a85f4925050508080156104e157505f84145b156104f25760035f819055506104fa565b60045f819055505b5050505050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f595a6040826040848885f49250505080801561053857505f82145b156105495760035f81905550610551565b60045f819055505b505050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f8273ffffffffffffffffffffffffffffffffffffffff1660405161059590610809565b5f60405180830381855af49150503d805f81146105cd576040519150601f19603f3d011682016040523d82523d5f602084013e6105d2565b606091505b505090508080156105e257505f82115b156105ef5760045f819055505b505050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f8273ffffffffffffffffffffffffffffffffffffffff16602a604051602401610639919061086b565b6040516020818303038152906040527f55241077000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516106c391906108cc565b5f60405180830381855af49150503d805f81146106fb576040519150601f19603f3d011682016040523d82523d5f602084013e610700565b606091505b5050905080801561071057505f82115b1561071d5760045f819055505b505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61074f82610726565b9050919050565b61075f81610745565b8114610769575f5ffd5b50565b5f8135905061077a81610756565b92915050565b5f6020828403121561079557610794610722565b5b5f6107a28482850161076c565b91505092915050565b5f819050919050565b6107bd816107ab565b82525050565b5f6020820190506107d65f8301846107b4565b92915050565b5f81905092915050565b50565b5f6107f45f836107dc565b91506107ff826107e6565b5f82019050919050565b5f610813826107e9565b9150819050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61085561085061084b8461081d565b610832565b610826565b9050919050565b6108658161083b565b82525050565b5f60208201905061087e5f83018461085c565b92915050565b5f81519050919050565b8281835e5f83830152505050565b5f6108a682610884565b6108b081856107dc565b93506108c081856020860161088e565b80840191505092915050565b5f6108d7828461089c565b915081905092915050565b5f819050919050565b5f6109056109006108fb846108e2565b610832565b610826565b9050919050565b610915816108eb565b82525050565b5f60208201905061092e5f83018461090c565b9291505056fea2646970667358221220b75c91a2513affa275c998257dc2b8112c20a42edebc5298aecc7f53c9eb0eda64736f6c634300081e0033",
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

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_DelegateCaller *DelegateCallerCaller) Value(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegateCaller.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_DelegateCaller *DelegateCallerSession) Value() (*big.Int, error) {
	return _DelegateCaller.Contract.Value(&_DelegateCaller.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns(uint256)
func (_DelegateCaller *DelegateCallerCallerSession) Value() (*big.Int, error) {
	return _DelegateCaller.Contract.Value(&_DelegateCaller.CallOpts)
}

// TestDelegateCallEmptyCold is a paid mutator transaction binding the contract method 0x4289f2a8.
//
// Solidity: function testDelegateCallEmptyCold(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallEmptyCold(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallEmptyCold", emptyAddress)
}

// TestDelegateCallEmptyCold is a paid mutator transaction binding the contract method 0x4289f2a8.
//
// Solidity: function testDelegateCallEmptyCold(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallEmptyCold(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyCold(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyCold is a paid mutator transaction binding the contract method 0x4289f2a8.
//
// Solidity: function testDelegateCallEmptyCold(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallEmptyCold(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyCold(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0x2194badc.
//
// Solidity: function testDelegateCallEmptyColdMemExpansion(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallEmptyColdMemExpansion(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallEmptyColdMemExpansion", emptyAddress)
}

// TestDelegateCallEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0x2194badc.
//
// Solidity: function testDelegateCallEmptyColdMemExpansion(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallEmptyColdMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyColdMemExpansion(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0x2194badc.
//
// Solidity: function testDelegateCallEmptyColdMemExpansion(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallEmptyColdMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyColdMemExpansion(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyWarm is a paid mutator transaction binding the contract method 0x978c93bc.
//
// Solidity: function testDelegateCallEmptyWarm(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallEmptyWarm(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallEmptyWarm", emptyAddress)
}

// TestDelegateCallEmptyWarm is a paid mutator transaction binding the contract method 0x978c93bc.
//
// Solidity: function testDelegateCallEmptyWarm(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallEmptyWarm(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyWarm(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyWarm is a paid mutator transaction binding the contract method 0x978c93bc.
//
// Solidity: function testDelegateCallEmptyWarm(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallEmptyWarm(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyWarm(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x737c29d0.
//
// Solidity: function testDelegateCallEmptyWarmMemExpansion(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallEmptyWarmMemExpansion(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallEmptyWarmMemExpansion", emptyAddress)
}

// TestDelegateCallEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x737c29d0.
//
// Solidity: function testDelegateCallEmptyWarmMemExpansion(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallEmptyWarmMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyWarmMemExpansion(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x737c29d0.
//
// Solidity: function testDelegateCallEmptyWarmMemExpansion(address emptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallEmptyWarmMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallEmptyWarmMemExpansion(&_DelegateCaller.TransactOpts, emptyAddress)
}

// TestDelegateCallNonEmptyCold is a paid mutator transaction binding the contract method 0x48112963.
//
// Solidity: function testDelegateCallNonEmptyCold(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallNonEmptyCold(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallNonEmptyCold", nonEmptyAddress)
}

// TestDelegateCallNonEmptyCold is a paid mutator transaction binding the contract method 0x48112963.
//
// Solidity: function testDelegateCallNonEmptyCold(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallNonEmptyCold(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyCold(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyCold is a paid mutator transaction binding the contract method 0x48112963.
//
// Solidity: function testDelegateCallNonEmptyCold(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallNonEmptyCold(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyCold(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0x648a6be8.
//
// Solidity: function testDelegateCallNonEmptyColdMemExpansion(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallNonEmptyColdMemExpansion(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallNonEmptyColdMemExpansion", nonEmptyAddress)
}

// TestDelegateCallNonEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0x648a6be8.
//
// Solidity: function testDelegateCallNonEmptyColdMemExpansion(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallNonEmptyColdMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyColdMemExpansion(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0x648a6be8.
//
// Solidity: function testDelegateCallNonEmptyColdMemExpansion(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallNonEmptyColdMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyColdMemExpansion(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyWarm is a paid mutator transaction binding the contract method 0xf55b1998.
//
// Solidity: function testDelegateCallNonEmptyWarm(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallNonEmptyWarm(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallNonEmptyWarm", nonEmptyAddress)
}

// TestDelegateCallNonEmptyWarm is a paid mutator transaction binding the contract method 0xf55b1998.
//
// Solidity: function testDelegateCallNonEmptyWarm(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallNonEmptyWarm(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyWarm(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyWarm is a paid mutator transaction binding the contract method 0xf55b1998.
//
// Solidity: function testDelegateCallNonEmptyWarm(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallNonEmptyWarm(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyWarm(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x710814a9.
//
// Solidity: function testDelegateCallNonEmptyWarmMemExpansion(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactor) TestDelegateCallNonEmptyWarmMemExpansion(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.contract.Transact(opts, "testDelegateCallNonEmptyWarmMemExpansion", nonEmptyAddress)
}

// TestDelegateCallNonEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x710814a9.
//
// Solidity: function testDelegateCallNonEmptyWarmMemExpansion(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerSession) TestDelegateCallNonEmptyWarmMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyWarmMemExpansion(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// TestDelegateCallNonEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x710814a9.
//
// Solidity: function testDelegateCallNonEmptyWarmMemExpansion(address nonEmptyAddress) returns()
func (_DelegateCaller *DelegateCallerTransactorSession) TestDelegateCallNonEmptyWarmMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _DelegateCaller.Contract.TestDelegateCallNonEmptyWarmMemExpansion(&_DelegateCaller.TransactOpts, nonEmptyAddress)
}

// ExtCodeCopyMetaData contains all meta data concerning the ExtCodeCopy contract.
var ExtCodeCopyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"extCodeCopyColdMemExpansion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extCodeCopyColdNoMemExpansion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extCodeCopyWarmMemExpansion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extCodeCopyWarmNoMemExpansion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060405160199060d7565b604051809103905ff0801580156031573d5f5f3e3d5ffd5b5060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604051607b9060d7565b604051809103905ff0801580156093573d5f5f3e3d5ffd5b5060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060e4565b610220806104ff83390190565b61040e806100f15f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806313a6a5901461004e578063197bc5ce1461006c578063d2f7b04f1461008a578063d74c69d8146100a8575b5f5ffd5b6100566100c6565b6040516100639190610392565b60405180910390f35b610074610161565b6040516100819190610392565b60405180910390f35b610092610220565b60405161009f9190610392565b60405180910390f35b6100b06102e1565b6040516100bd9190610392565b60405180910390f35b5f5f5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050803b91505f8267ffffffffffffffff81111561010c5761010b6103ab565b5b6040519080825280601f01601f19166020018201604052801561013e5781602001600182028036833780820191505090505b50905059835f60018303853c5080805190602001209350835f8190555050505090565b5f5f5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050803b91505f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f8367ffffffffffffffff8111156101cc576101cb6103ab565b5b6040519080825280601f01601f1916602001820160405280156101fe5781602001600182028036833780820191505090505b509050835f60208301843c80805190602001209450845f819055505050505090565b5f5f5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050803b91505f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f8367ffffffffffffffff81111561028b5761028a6103ab565b5b6040519080825280601f01601f1916602001820160405280156102bd5781602001600182028036833780820191505090505b50905059845f60018303853c5080805190602001209450845f819055505050505090565b5f5f5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050803b91505f8267ffffffffffffffff811115610327576103266103ab565b5b6040519080825280601f01601f1916602001820160405280156103595781602001600182028036833780820191505090505b509050825f60208301843c80805190602001209350835f8190555050505090565b5f819050919050565b61038c8161037a565b82525050565b5f6020820190506103a55f830184610383565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffdfea2646970667358221220fda3cc5e049731b8e903a5b02345946c33a33664221dd1907d7e728b51e8382864736f6c634300081e00336080604052348015600e575f5ffd5b506102048061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80633fb5c1cb1461004e578063822ec8611461006a5780638381f58a14610074578063d09de08a14610092575b5f5ffd5b6100686004803603810190610063919061011b565b61009c565b005b6100726100a5565b005b61007c6100b1565b6040516100899190610155565b60405180910390f35b61009a6100b6565b005b805f8190555050565b61696961133701602081f35b5f5481565b5f5f5490506001816100c8919061019b565b90505f5f54905080826100db919061019b565b5f819055505050565b5f5ffd5b5f819050919050565b6100fa816100e8565b8114610104575f5ffd5b50565b5f81359050610115816100f1565b92915050565b5f602082840312156101305761012f6100e4565b5b5f61013d84828501610107565b91505092915050565b61014f816100e8565b82525050565b5f6020820190506101685f830184610146565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101a5826100e8565b91506101b0836100e8565b92508282019050808211156101c8576101c761016e565b5b9291505056fea2646970667358221220285973522b8ad7991c66c263fbd08eae79056517be22fa4eef779c111e61db3764736f6c634300081e0033",
}

// ExtCodeCopyABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtCodeCopyMetaData.ABI instead.
var ExtCodeCopyABI = ExtCodeCopyMetaData.ABI

// ExtCodeCopyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtCodeCopyMetaData.Bin instead.
var ExtCodeCopyBin = ExtCodeCopyMetaData.Bin

// DeployExtCodeCopy deploys a new Ethereum contract, binding an instance of ExtCodeCopy to it.
func DeployExtCodeCopy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtCodeCopy, error) {
	parsed, err := ExtCodeCopyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtCodeCopyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtCodeCopy{ExtCodeCopyCaller: ExtCodeCopyCaller{contract: contract}, ExtCodeCopyTransactor: ExtCodeCopyTransactor{contract: contract}, ExtCodeCopyFilterer: ExtCodeCopyFilterer{contract: contract}}, nil
}

// ExtCodeCopy is an auto generated Go binding around an Ethereum contract.
type ExtCodeCopy struct {
	ExtCodeCopyCaller     // Read-only binding to the contract
	ExtCodeCopyTransactor // Write-only binding to the contract
	ExtCodeCopyFilterer   // Log filterer for contract events
}

// ExtCodeCopyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtCodeCopyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeCopyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtCodeCopyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeCopyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtCodeCopyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeCopySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtCodeCopySession struct {
	Contract     *ExtCodeCopy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtCodeCopyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtCodeCopyCallerSession struct {
	Contract *ExtCodeCopyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExtCodeCopyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtCodeCopyTransactorSession struct {
	Contract     *ExtCodeCopyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExtCodeCopyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtCodeCopyRaw struct {
	Contract *ExtCodeCopy // Generic contract binding to access the raw methods on
}

// ExtCodeCopyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtCodeCopyCallerRaw struct {
	Contract *ExtCodeCopyCaller // Generic read-only contract binding to access the raw methods on
}

// ExtCodeCopyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtCodeCopyTransactorRaw struct {
	Contract *ExtCodeCopyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtCodeCopy creates a new instance of ExtCodeCopy, bound to a specific deployed contract.
func NewExtCodeCopy(address common.Address, backend bind.ContractBackend) (*ExtCodeCopy, error) {
	contract, err := bindExtCodeCopy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtCodeCopy{ExtCodeCopyCaller: ExtCodeCopyCaller{contract: contract}, ExtCodeCopyTransactor: ExtCodeCopyTransactor{contract: contract}, ExtCodeCopyFilterer: ExtCodeCopyFilterer{contract: contract}}, nil
}

// NewExtCodeCopyCaller creates a new read-only instance of ExtCodeCopy, bound to a specific deployed contract.
func NewExtCodeCopyCaller(address common.Address, caller bind.ContractCaller) (*ExtCodeCopyCaller, error) {
	contract, err := bindExtCodeCopy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeCopyCaller{contract: contract}, nil
}

// NewExtCodeCopyTransactor creates a new write-only instance of ExtCodeCopy, bound to a specific deployed contract.
func NewExtCodeCopyTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtCodeCopyTransactor, error) {
	contract, err := bindExtCodeCopy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeCopyTransactor{contract: contract}, nil
}

// NewExtCodeCopyFilterer creates a new log filterer instance of ExtCodeCopy, bound to a specific deployed contract.
func NewExtCodeCopyFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtCodeCopyFilterer, error) {
	contract, err := bindExtCodeCopy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtCodeCopyFilterer{contract: contract}, nil
}

// bindExtCodeCopy binds a generic wrapper to an already deployed contract.
func bindExtCodeCopy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtCodeCopyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeCopy *ExtCodeCopyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeCopy.Contract.ExtCodeCopyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeCopy *ExtCodeCopyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeCopy *ExtCodeCopyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeCopy *ExtCodeCopyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeCopy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeCopy *ExtCodeCopyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeCopy *ExtCodeCopyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.contract.Transact(opts, method, params...)
}

// ExtCodeCopyColdMemExpansion is a paid mutator transaction binding the contract method 0xd2f7b04f.
//
// Solidity: function extCodeCopyColdMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactor) ExtCodeCopyColdMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeCopy.contract.Transact(opts, "extCodeCopyColdMemExpansion")
}

// ExtCodeCopyColdMemExpansion is a paid mutator transaction binding the contract method 0xd2f7b04f.
//
// Solidity: function extCodeCopyColdMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopySession) ExtCodeCopyColdMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyColdMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyColdMemExpansion is a paid mutator transaction binding the contract method 0xd2f7b04f.
//
// Solidity: function extCodeCopyColdMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactorSession) ExtCodeCopyColdMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyColdMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyColdNoMemExpansion is a paid mutator transaction binding the contract method 0x197bc5ce.
//
// Solidity: function extCodeCopyColdNoMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactor) ExtCodeCopyColdNoMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeCopy.contract.Transact(opts, "extCodeCopyColdNoMemExpansion")
}

// ExtCodeCopyColdNoMemExpansion is a paid mutator transaction binding the contract method 0x197bc5ce.
//
// Solidity: function extCodeCopyColdNoMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopySession) ExtCodeCopyColdNoMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyColdNoMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyColdNoMemExpansion is a paid mutator transaction binding the contract method 0x197bc5ce.
//
// Solidity: function extCodeCopyColdNoMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactorSession) ExtCodeCopyColdNoMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyColdNoMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyWarmMemExpansion is a paid mutator transaction binding the contract method 0x13a6a590.
//
// Solidity: function extCodeCopyWarmMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactor) ExtCodeCopyWarmMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeCopy.contract.Transact(opts, "extCodeCopyWarmMemExpansion")
}

// ExtCodeCopyWarmMemExpansion is a paid mutator transaction binding the contract method 0x13a6a590.
//
// Solidity: function extCodeCopyWarmMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopySession) ExtCodeCopyWarmMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyWarmMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyWarmMemExpansion is a paid mutator transaction binding the contract method 0x13a6a590.
//
// Solidity: function extCodeCopyWarmMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactorSession) ExtCodeCopyWarmMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyWarmMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyWarmNoMemExpansion is a paid mutator transaction binding the contract method 0xd74c69d8.
//
// Solidity: function extCodeCopyWarmNoMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactor) ExtCodeCopyWarmNoMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeCopy.contract.Transact(opts, "extCodeCopyWarmNoMemExpansion")
}

// ExtCodeCopyWarmNoMemExpansion is a paid mutator transaction binding the contract method 0xd74c69d8.
//
// Solidity: function extCodeCopyWarmNoMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopySession) ExtCodeCopyWarmNoMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyWarmNoMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeCopyWarmNoMemExpansion is a paid mutator transaction binding the contract method 0xd74c69d8.
//
// Solidity: function extCodeCopyWarmNoMemExpansion() returns(bytes32 codeHash)
func (_ExtCodeCopy *ExtCodeCopyTransactorSession) ExtCodeCopyWarmNoMemExpansion() (*types.Transaction, error) {
	return _ExtCodeCopy.Contract.ExtCodeCopyWarmNoMemExpansion(&_ExtCodeCopy.TransactOpts)
}

// ExtCodeHashMetaData contains all meta data concerning the ExtCodeHash contract.
var ExtCodeHashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getExtCodeHashCold\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExtCodeHashWarm\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506040516019906074565b604051809103905ff0801580156031573d5f5f3e3d5ffd5b505f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506081565b610142806103f783390190565b6103698061008e5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80635018668e14610038578063ad39eb6214610056575b5f5ffd5b610040610074565b60405161004d919061024a565b60405180910390f35b61005e6100d1565b60405161006b919061024a565b60405180910390f35b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163f905060015f60146101000a81548160ff0219169083151502179055508091505090565b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff165f633fb5c1cb60e01b61133760405160240161012691906102b2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610190919061031d565b5f6040518083038185875af1925050503d805f81146101ca576040519150601f19603f3d011682016040523d82523d5f602084013e6101cf565b606091505b5050905080156101f45760015f60146101000a81548160ff0219169083151502179055505b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163f91505090565b5f819050919050565b61024481610232565b82525050565b5f60208201905061025d5f83018461023b565b92915050565b5f819050919050565b5f61ffff82169050919050565b5f819050919050565b5f61029c61029761029284610263565b610279565b61026c565b9050919050565b6102ac81610282565b82525050565b5f6020820190506102c55f8301846102a3565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6102f7826102cb565b61030181856102d5565b93506103118185602086016102df565b80840191505092915050565b5f61032882846102ed565b91508190509291505056fea264697066735822122076e51c72c412dc0a7b6bf34b5466ba2ee81176a2375ec6c4f3cd785267c7060164736f6c634300081e00336080604052348015600e575f5ffd5b506101268061001c5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80633fb5c1cb1460345780638381f58a14604c575b5f5ffd5b604a60048036038101906046919060a6565b6066565b005b6052606f565b604051605d919060d9565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6088816078565b81146091575f5ffd5b50565b5f8135905060a0816081565b92915050565b5f6020828403121560b85760b76074565b5b5f60c3848285016094565b91505092915050565b60d3816078565b82525050565b5f60208201905060ea5f83018460cc565b9291505056fea2646970667358221220d83aee5466ac18b3aee1511c0addac3e7fdc1e3e519b8bbbfe7ddaff748b3dcd64736f6c634300081e0033",
}

// ExtCodeHashABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtCodeHashMetaData.ABI instead.
var ExtCodeHashABI = ExtCodeHashMetaData.ABI

// ExtCodeHashBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtCodeHashMetaData.Bin instead.
var ExtCodeHashBin = ExtCodeHashMetaData.Bin

// DeployExtCodeHash deploys a new Ethereum contract, binding an instance of ExtCodeHash to it.
func DeployExtCodeHash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtCodeHash, error) {
	parsed, err := ExtCodeHashMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtCodeHashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtCodeHash{ExtCodeHashCaller: ExtCodeHashCaller{contract: contract}, ExtCodeHashTransactor: ExtCodeHashTransactor{contract: contract}, ExtCodeHashFilterer: ExtCodeHashFilterer{contract: contract}}, nil
}

// ExtCodeHash is an auto generated Go binding around an Ethereum contract.
type ExtCodeHash struct {
	ExtCodeHashCaller     // Read-only binding to the contract
	ExtCodeHashTransactor // Write-only binding to the contract
	ExtCodeHashFilterer   // Log filterer for contract events
}

// ExtCodeHashCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtCodeHashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeHashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtCodeHashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeHashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtCodeHashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeHashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtCodeHashSession struct {
	Contract     *ExtCodeHash      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtCodeHashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtCodeHashCallerSession struct {
	Contract *ExtCodeHashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExtCodeHashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtCodeHashTransactorSession struct {
	Contract     *ExtCodeHashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExtCodeHashRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtCodeHashRaw struct {
	Contract *ExtCodeHash // Generic contract binding to access the raw methods on
}

// ExtCodeHashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtCodeHashCallerRaw struct {
	Contract *ExtCodeHashCaller // Generic read-only contract binding to access the raw methods on
}

// ExtCodeHashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtCodeHashTransactorRaw struct {
	Contract *ExtCodeHashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtCodeHash creates a new instance of ExtCodeHash, bound to a specific deployed contract.
func NewExtCodeHash(address common.Address, backend bind.ContractBackend) (*ExtCodeHash, error) {
	contract, err := bindExtCodeHash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHash{ExtCodeHashCaller: ExtCodeHashCaller{contract: contract}, ExtCodeHashTransactor: ExtCodeHashTransactor{contract: contract}, ExtCodeHashFilterer: ExtCodeHashFilterer{contract: contract}}, nil
}

// NewExtCodeHashCaller creates a new read-only instance of ExtCodeHash, bound to a specific deployed contract.
func NewExtCodeHashCaller(address common.Address, caller bind.ContractCaller) (*ExtCodeHashCaller, error) {
	contract, err := bindExtCodeHash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashCaller{contract: contract}, nil
}

// NewExtCodeHashTransactor creates a new write-only instance of ExtCodeHash, bound to a specific deployed contract.
func NewExtCodeHashTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtCodeHashTransactor, error) {
	contract, err := bindExtCodeHash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashTransactor{contract: contract}, nil
}

// NewExtCodeHashFilterer creates a new log filterer instance of ExtCodeHash, bound to a specific deployed contract.
func NewExtCodeHashFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtCodeHashFilterer, error) {
	contract, err := bindExtCodeHash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashFilterer{contract: contract}, nil
}

// bindExtCodeHash binds a generic wrapper to an already deployed contract.
func bindExtCodeHash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtCodeHashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeHash *ExtCodeHashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeHash.Contract.ExtCodeHashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeHash *ExtCodeHashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeHash.Contract.ExtCodeHashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeHash *ExtCodeHashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeHash.Contract.ExtCodeHashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeHash *ExtCodeHashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeHash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeHash *ExtCodeHashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeHash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeHash *ExtCodeHashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeHash.Contract.contract.Transact(opts, method, params...)
}

// GetExtCodeHashCold is a paid mutator transaction binding the contract method 0x5018668e.
//
// Solidity: function getExtCodeHashCold() returns(bytes32)
func (_ExtCodeHash *ExtCodeHashTransactor) GetExtCodeHashCold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeHash.contract.Transact(opts, "getExtCodeHashCold")
}

// GetExtCodeHashCold is a paid mutator transaction binding the contract method 0x5018668e.
//
// Solidity: function getExtCodeHashCold() returns(bytes32)
func (_ExtCodeHash *ExtCodeHashSession) GetExtCodeHashCold() (*types.Transaction, error) {
	return _ExtCodeHash.Contract.GetExtCodeHashCold(&_ExtCodeHash.TransactOpts)
}

// GetExtCodeHashCold is a paid mutator transaction binding the contract method 0x5018668e.
//
// Solidity: function getExtCodeHashCold() returns(bytes32)
func (_ExtCodeHash *ExtCodeHashTransactorSession) GetExtCodeHashCold() (*types.Transaction, error) {
	return _ExtCodeHash.Contract.GetExtCodeHashCold(&_ExtCodeHash.TransactOpts)
}

// GetExtCodeHashWarm is a paid mutator transaction binding the contract method 0xad39eb62.
//
// Solidity: function getExtCodeHashWarm() returns(bytes32)
func (_ExtCodeHash *ExtCodeHashTransactor) GetExtCodeHashWarm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeHash.contract.Transact(opts, "getExtCodeHashWarm")
}

// GetExtCodeHashWarm is a paid mutator transaction binding the contract method 0xad39eb62.
//
// Solidity: function getExtCodeHashWarm() returns(bytes32)
func (_ExtCodeHash *ExtCodeHashSession) GetExtCodeHashWarm() (*types.Transaction, error) {
	return _ExtCodeHash.Contract.GetExtCodeHashWarm(&_ExtCodeHash.TransactOpts)
}

// GetExtCodeHashWarm is a paid mutator transaction binding the contract method 0xad39eb62.
//
// Solidity: function getExtCodeHashWarm() returns(bytes32)
func (_ExtCodeHash *ExtCodeHashTransactorSession) GetExtCodeHashWarm() (*types.Transaction, error) {
	return _ExtCodeHash.Contract.GetExtCodeHashWarm(&_ExtCodeHash.TransactOpts)
}

// ExtCodeHashCalleeMetaData contains all meta data concerning the ExtCodeHashCallee contract.
var ExtCodeHashCalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101268061001c5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80633fb5c1cb1460345780638381f58a14604c575b5f5ffd5b604a60048036038101906046919060a6565b6066565b005b6052606f565b604051605d919060d9565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6088816078565b81146091575f5ffd5b50565b5f8135905060a0816081565b92915050565b5f6020828403121560b85760b76074565b5b5f60c3848285016094565b91505092915050565b60d3816078565b82525050565b5f60208201905060ea5f83018460cc565b9291505056fea2646970667358221220d83aee5466ac18b3aee1511c0addac3e7fdc1e3e519b8bbbfe7ddaff748b3dcd64736f6c634300081e0033",
}

// ExtCodeHashCalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtCodeHashCalleeMetaData.ABI instead.
var ExtCodeHashCalleeABI = ExtCodeHashCalleeMetaData.ABI

// ExtCodeHashCalleeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtCodeHashCalleeMetaData.Bin instead.
var ExtCodeHashCalleeBin = ExtCodeHashCalleeMetaData.Bin

// DeployExtCodeHashCallee deploys a new Ethereum contract, binding an instance of ExtCodeHashCallee to it.
func DeployExtCodeHashCallee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtCodeHashCallee, error) {
	parsed, err := ExtCodeHashCalleeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtCodeHashCalleeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtCodeHashCallee{ExtCodeHashCalleeCaller: ExtCodeHashCalleeCaller{contract: contract}, ExtCodeHashCalleeTransactor: ExtCodeHashCalleeTransactor{contract: contract}, ExtCodeHashCalleeFilterer: ExtCodeHashCalleeFilterer{contract: contract}}, nil
}

// ExtCodeHashCallee is an auto generated Go binding around an Ethereum contract.
type ExtCodeHashCallee struct {
	ExtCodeHashCalleeCaller     // Read-only binding to the contract
	ExtCodeHashCalleeTransactor // Write-only binding to the contract
	ExtCodeHashCalleeFilterer   // Log filterer for contract events
}

// ExtCodeHashCalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtCodeHashCalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeHashCalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtCodeHashCalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeHashCalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtCodeHashCalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeHashCalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtCodeHashCalleeSession struct {
	Contract     *ExtCodeHashCallee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExtCodeHashCalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtCodeHashCalleeCallerSession struct {
	Contract *ExtCodeHashCalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExtCodeHashCalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtCodeHashCalleeTransactorSession struct {
	Contract     *ExtCodeHashCalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExtCodeHashCalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtCodeHashCalleeRaw struct {
	Contract *ExtCodeHashCallee // Generic contract binding to access the raw methods on
}

// ExtCodeHashCalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtCodeHashCalleeCallerRaw struct {
	Contract *ExtCodeHashCalleeCaller // Generic read-only contract binding to access the raw methods on
}

// ExtCodeHashCalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtCodeHashCalleeTransactorRaw struct {
	Contract *ExtCodeHashCalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtCodeHashCallee creates a new instance of ExtCodeHashCallee, bound to a specific deployed contract.
func NewExtCodeHashCallee(address common.Address, backend bind.ContractBackend) (*ExtCodeHashCallee, error) {
	contract, err := bindExtCodeHashCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashCallee{ExtCodeHashCalleeCaller: ExtCodeHashCalleeCaller{contract: contract}, ExtCodeHashCalleeTransactor: ExtCodeHashCalleeTransactor{contract: contract}, ExtCodeHashCalleeFilterer: ExtCodeHashCalleeFilterer{contract: contract}}, nil
}

// NewExtCodeHashCalleeCaller creates a new read-only instance of ExtCodeHashCallee, bound to a specific deployed contract.
func NewExtCodeHashCalleeCaller(address common.Address, caller bind.ContractCaller) (*ExtCodeHashCalleeCaller, error) {
	contract, err := bindExtCodeHashCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashCalleeCaller{contract: contract}, nil
}

// NewExtCodeHashCalleeTransactor creates a new write-only instance of ExtCodeHashCallee, bound to a specific deployed contract.
func NewExtCodeHashCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtCodeHashCalleeTransactor, error) {
	contract, err := bindExtCodeHashCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashCalleeTransactor{contract: contract}, nil
}

// NewExtCodeHashCalleeFilterer creates a new log filterer instance of ExtCodeHashCallee, bound to a specific deployed contract.
func NewExtCodeHashCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtCodeHashCalleeFilterer, error) {
	contract, err := bindExtCodeHashCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtCodeHashCalleeFilterer{contract: contract}, nil
}

// bindExtCodeHashCallee binds a generic wrapper to an already deployed contract.
func bindExtCodeHashCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtCodeHashCalleeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeHashCallee *ExtCodeHashCalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeHashCallee.Contract.ExtCodeHashCalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeHashCallee *ExtCodeHashCalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeHashCallee.Contract.ExtCodeHashCalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeHashCallee *ExtCodeHashCalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeHashCallee.Contract.ExtCodeHashCalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeHashCallee *ExtCodeHashCalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeHashCallee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeHashCallee *ExtCodeHashCalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeHashCallee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeHashCallee *ExtCodeHashCalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeHashCallee.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_ExtCodeHashCallee *ExtCodeHashCalleeCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExtCodeHashCallee.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_ExtCodeHashCallee *ExtCodeHashCalleeSession) Number() (*big.Int, error) {
	return _ExtCodeHashCallee.Contract.Number(&_ExtCodeHashCallee.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_ExtCodeHashCallee *ExtCodeHashCalleeCallerSession) Number() (*big.Int, error) {
	return _ExtCodeHashCallee.Contract.Number(&_ExtCodeHashCallee.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_ExtCodeHashCallee *ExtCodeHashCalleeTransactor) SetNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _ExtCodeHashCallee.contract.Transact(opts, "setNumber", _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_ExtCodeHashCallee *ExtCodeHashCalleeSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _ExtCodeHashCallee.Contract.SetNumber(&_ExtCodeHashCallee.TransactOpts, _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_ExtCodeHashCallee *ExtCodeHashCalleeTransactorSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _ExtCodeHashCallee.Contract.SetNumber(&_ExtCodeHashCallee.TransactOpts, _number)
}

// ExtCodeSizeMetaData contains all meta data concerning the ExtCodeSize contract.
var ExtCodeSizeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getExtCodeSizeCold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExtCodeSizeWarm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506040516019906074565b604051809103905ff0801580156031573d5f5f3e3d5ffd5b505f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506081565b610142806103f783390190565b6103698061008e5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806312a44f9d146100385780633251411e14610056575b5f5ffd5b610040610074565b60405161004d919061024a565b60405180910390f35b61005e6100d1565b60405161006b919061024a565b60405180910390f35b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163b905060015f60146101000a81548160ff0219169083151502179055508091505090565b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff165f633fb5c1cb60e01b61133760405160240161012691906102b2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610190919061031d565b5f6040518083038185875af1925050503d805f81146101ca576040519150601f19603f3d011682016040523d82523d5f602084013e6101cf565b606091505b5050905080156101f45760015f60146101000a81548160ff0219169083151502179055505b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163b91505090565b5f819050919050565b61024481610232565b82525050565b5f60208201905061025d5f83018461023b565b92915050565b5f819050919050565b5f61ffff82169050919050565b5f819050919050565b5f61029c61029761029284610263565b610279565b61026c565b9050919050565b6102ac81610282565b82525050565b5f6020820190506102c55f8301846102a3565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6102f7826102cb565b61030181856102d5565b93506103118185602086016102df565b80840191505092915050565b5f61032882846102ed565b91508190509291505056fea2646970667358221220d965e103a6d892d3020822aa191547290dbfe0195b8f9e85625ac5a8f862702c64736f6c634300081e00336080604052348015600e575f5ffd5b506101268061001c5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80633fb5c1cb1460345780638381f58a14604c575b5f5ffd5b604a60048036038101906046919060a6565b6066565b005b6052606f565b604051605d919060d9565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6088816078565b81146091575f5ffd5b50565b5f8135905060a0816081565b92915050565b5f6020828403121560b85760b76074565b5b5f60c3848285016094565b91505092915050565b60d3816078565b82525050565b5f60208201905060ea5f83018460cc565b9291505056fea26469706673582212206c0243e43dd42e87267d856c9893e85a739991abd00a6d9702bb071805a5302064736f6c634300081e0033",
}

// ExtCodeSizeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtCodeSizeMetaData.ABI instead.
var ExtCodeSizeABI = ExtCodeSizeMetaData.ABI

// ExtCodeSizeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtCodeSizeMetaData.Bin instead.
var ExtCodeSizeBin = ExtCodeSizeMetaData.Bin

// DeployExtCodeSize deploys a new Ethereum contract, binding an instance of ExtCodeSize to it.
func DeployExtCodeSize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtCodeSize, error) {
	parsed, err := ExtCodeSizeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtCodeSizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtCodeSize{ExtCodeSizeCaller: ExtCodeSizeCaller{contract: contract}, ExtCodeSizeTransactor: ExtCodeSizeTransactor{contract: contract}, ExtCodeSizeFilterer: ExtCodeSizeFilterer{contract: contract}}, nil
}

// ExtCodeSize is an auto generated Go binding around an Ethereum contract.
type ExtCodeSize struct {
	ExtCodeSizeCaller     // Read-only binding to the contract
	ExtCodeSizeTransactor // Write-only binding to the contract
	ExtCodeSizeFilterer   // Log filterer for contract events
}

// ExtCodeSizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtCodeSizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeSizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtCodeSizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeSizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtCodeSizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeSizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtCodeSizeSession struct {
	Contract     *ExtCodeSize      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtCodeSizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtCodeSizeCallerSession struct {
	Contract *ExtCodeSizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExtCodeSizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtCodeSizeTransactorSession struct {
	Contract     *ExtCodeSizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExtCodeSizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtCodeSizeRaw struct {
	Contract *ExtCodeSize // Generic contract binding to access the raw methods on
}

// ExtCodeSizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtCodeSizeCallerRaw struct {
	Contract *ExtCodeSizeCaller // Generic read-only contract binding to access the raw methods on
}

// ExtCodeSizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtCodeSizeTransactorRaw struct {
	Contract *ExtCodeSizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtCodeSize creates a new instance of ExtCodeSize, bound to a specific deployed contract.
func NewExtCodeSize(address common.Address, backend bind.ContractBackend) (*ExtCodeSize, error) {
	contract, err := bindExtCodeSize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSize{ExtCodeSizeCaller: ExtCodeSizeCaller{contract: contract}, ExtCodeSizeTransactor: ExtCodeSizeTransactor{contract: contract}, ExtCodeSizeFilterer: ExtCodeSizeFilterer{contract: contract}}, nil
}

// NewExtCodeSizeCaller creates a new read-only instance of ExtCodeSize, bound to a specific deployed contract.
func NewExtCodeSizeCaller(address common.Address, caller bind.ContractCaller) (*ExtCodeSizeCaller, error) {
	contract, err := bindExtCodeSize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeCaller{contract: contract}, nil
}

// NewExtCodeSizeTransactor creates a new write-only instance of ExtCodeSize, bound to a specific deployed contract.
func NewExtCodeSizeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtCodeSizeTransactor, error) {
	contract, err := bindExtCodeSize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeTransactor{contract: contract}, nil
}

// NewExtCodeSizeFilterer creates a new log filterer instance of ExtCodeSize, bound to a specific deployed contract.
func NewExtCodeSizeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtCodeSizeFilterer, error) {
	contract, err := bindExtCodeSize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeFilterer{contract: contract}, nil
}

// bindExtCodeSize binds a generic wrapper to an already deployed contract.
func bindExtCodeSize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtCodeSizeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeSize *ExtCodeSizeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeSize.Contract.ExtCodeSizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeSize *ExtCodeSizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeSize.Contract.ExtCodeSizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeSize *ExtCodeSizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeSize.Contract.ExtCodeSizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeSize *ExtCodeSizeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeSize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeSize *ExtCodeSizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeSize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeSize *ExtCodeSizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeSize.Contract.contract.Transact(opts, method, params...)
}

// GetExtCodeSizeCold is a paid mutator transaction binding the contract method 0x12a44f9d.
//
// Solidity: function getExtCodeSizeCold() returns(uint256)
func (_ExtCodeSize *ExtCodeSizeTransactor) GetExtCodeSizeCold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeSize.contract.Transact(opts, "getExtCodeSizeCold")
}

// GetExtCodeSizeCold is a paid mutator transaction binding the contract method 0x12a44f9d.
//
// Solidity: function getExtCodeSizeCold() returns(uint256)
func (_ExtCodeSize *ExtCodeSizeSession) GetExtCodeSizeCold() (*types.Transaction, error) {
	return _ExtCodeSize.Contract.GetExtCodeSizeCold(&_ExtCodeSize.TransactOpts)
}

// GetExtCodeSizeCold is a paid mutator transaction binding the contract method 0x12a44f9d.
//
// Solidity: function getExtCodeSizeCold() returns(uint256)
func (_ExtCodeSize *ExtCodeSizeTransactorSession) GetExtCodeSizeCold() (*types.Transaction, error) {
	return _ExtCodeSize.Contract.GetExtCodeSizeCold(&_ExtCodeSize.TransactOpts)
}

// GetExtCodeSizeWarm is a paid mutator transaction binding the contract method 0x3251411e.
//
// Solidity: function getExtCodeSizeWarm() returns(uint256)
func (_ExtCodeSize *ExtCodeSizeTransactor) GetExtCodeSizeWarm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeSize.contract.Transact(opts, "getExtCodeSizeWarm")
}

// GetExtCodeSizeWarm is a paid mutator transaction binding the contract method 0x3251411e.
//
// Solidity: function getExtCodeSizeWarm() returns(uint256)
func (_ExtCodeSize *ExtCodeSizeSession) GetExtCodeSizeWarm() (*types.Transaction, error) {
	return _ExtCodeSize.Contract.GetExtCodeSizeWarm(&_ExtCodeSize.TransactOpts)
}

// GetExtCodeSizeWarm is a paid mutator transaction binding the contract method 0x3251411e.
//
// Solidity: function getExtCodeSizeWarm() returns(uint256)
func (_ExtCodeSize *ExtCodeSizeTransactorSession) GetExtCodeSizeWarm() (*types.Transaction, error) {
	return _ExtCodeSize.Contract.GetExtCodeSizeWarm(&_ExtCodeSize.TransactOpts)
}

// ExtCodeSizeCalleeMetaData contains all meta data concerning the ExtCodeSizeCallee contract.
var ExtCodeSizeCalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101268061001c5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80633fb5c1cb1460345780638381f58a14604c575b5f5ffd5b604a60048036038101906046919060a6565b6066565b005b6052606f565b604051605d919060d9565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6088816078565b81146091575f5ffd5b50565b5f8135905060a0816081565b92915050565b5f6020828403121560b85760b76074565b5b5f60c3848285016094565b91505092915050565b60d3816078565b82525050565b5f60208201905060ea5f83018460cc565b9291505056fea26469706673582212206c0243e43dd42e87267d856c9893e85a739991abd00a6d9702bb071805a5302064736f6c634300081e0033",
}

// ExtCodeSizeCalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtCodeSizeCalleeMetaData.ABI instead.
var ExtCodeSizeCalleeABI = ExtCodeSizeCalleeMetaData.ABI

// ExtCodeSizeCalleeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtCodeSizeCalleeMetaData.Bin instead.
var ExtCodeSizeCalleeBin = ExtCodeSizeCalleeMetaData.Bin

// DeployExtCodeSizeCallee deploys a new Ethereum contract, binding an instance of ExtCodeSizeCallee to it.
func DeployExtCodeSizeCallee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtCodeSizeCallee, error) {
	parsed, err := ExtCodeSizeCalleeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtCodeSizeCalleeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtCodeSizeCallee{ExtCodeSizeCalleeCaller: ExtCodeSizeCalleeCaller{contract: contract}, ExtCodeSizeCalleeTransactor: ExtCodeSizeCalleeTransactor{contract: contract}, ExtCodeSizeCalleeFilterer: ExtCodeSizeCalleeFilterer{contract: contract}}, nil
}

// ExtCodeSizeCallee is an auto generated Go binding around an Ethereum contract.
type ExtCodeSizeCallee struct {
	ExtCodeSizeCalleeCaller     // Read-only binding to the contract
	ExtCodeSizeCalleeTransactor // Write-only binding to the contract
	ExtCodeSizeCalleeFilterer   // Log filterer for contract events
}

// ExtCodeSizeCalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtCodeSizeCalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeSizeCalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtCodeSizeCalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeSizeCalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtCodeSizeCalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtCodeSizeCalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtCodeSizeCalleeSession struct {
	Contract     *ExtCodeSizeCallee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExtCodeSizeCalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtCodeSizeCalleeCallerSession struct {
	Contract *ExtCodeSizeCalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExtCodeSizeCalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtCodeSizeCalleeTransactorSession struct {
	Contract     *ExtCodeSizeCalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExtCodeSizeCalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtCodeSizeCalleeRaw struct {
	Contract *ExtCodeSizeCallee // Generic contract binding to access the raw methods on
}

// ExtCodeSizeCalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtCodeSizeCalleeCallerRaw struct {
	Contract *ExtCodeSizeCalleeCaller // Generic read-only contract binding to access the raw methods on
}

// ExtCodeSizeCalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtCodeSizeCalleeTransactorRaw struct {
	Contract *ExtCodeSizeCalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtCodeSizeCallee creates a new instance of ExtCodeSizeCallee, bound to a specific deployed contract.
func NewExtCodeSizeCallee(address common.Address, backend bind.ContractBackend) (*ExtCodeSizeCallee, error) {
	contract, err := bindExtCodeSizeCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeCallee{ExtCodeSizeCalleeCaller: ExtCodeSizeCalleeCaller{contract: contract}, ExtCodeSizeCalleeTransactor: ExtCodeSizeCalleeTransactor{contract: contract}, ExtCodeSizeCalleeFilterer: ExtCodeSizeCalleeFilterer{contract: contract}}, nil
}

// NewExtCodeSizeCalleeCaller creates a new read-only instance of ExtCodeSizeCallee, bound to a specific deployed contract.
func NewExtCodeSizeCalleeCaller(address common.Address, caller bind.ContractCaller) (*ExtCodeSizeCalleeCaller, error) {
	contract, err := bindExtCodeSizeCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeCalleeCaller{contract: contract}, nil
}

// NewExtCodeSizeCalleeTransactor creates a new write-only instance of ExtCodeSizeCallee, bound to a specific deployed contract.
func NewExtCodeSizeCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtCodeSizeCalleeTransactor, error) {
	contract, err := bindExtCodeSizeCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeCalleeTransactor{contract: contract}, nil
}

// NewExtCodeSizeCalleeFilterer creates a new log filterer instance of ExtCodeSizeCallee, bound to a specific deployed contract.
func NewExtCodeSizeCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtCodeSizeCalleeFilterer, error) {
	contract, err := bindExtCodeSizeCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtCodeSizeCalleeFilterer{contract: contract}, nil
}

// bindExtCodeSizeCallee binds a generic wrapper to an already deployed contract.
func bindExtCodeSizeCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtCodeSizeCalleeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeSizeCallee.Contract.ExtCodeSizeCalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.Contract.ExtCodeSizeCalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.Contract.ExtCodeSizeCalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtCodeSizeCallee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExtCodeSizeCallee.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeSession) Number() (*big.Int, error) {
	return _ExtCodeSizeCallee.Contract.Number(&_ExtCodeSizeCallee.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeCallerSession) Number() (*big.Int, error) {
	return _ExtCodeSizeCallee.Contract.Number(&_ExtCodeSizeCallee.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeTransactor) SetNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.contract.Transact(opts, "setNumber", _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.Contract.SetNumber(&_ExtCodeSizeCallee.TransactOpts, _number)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 _number) returns()
func (_ExtCodeSizeCallee *ExtCodeSizeCalleeTransactorSession) SetNumber(_number *big.Int) (*types.Transaction, error) {
	return _ExtCodeSizeCallee.Contract.SetNumber(&_ExtCodeSizeCallee.TransactOpts, _number)
}

// InvalidMetaData contains all meta data concerning the Invalid contract.
var InvalidMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"invalid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertInTryCatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertInTryCatchWithMemoryExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertNoMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertWithMemoryExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertWithMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102798061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c80639fe55433116100595780639fe55433146100b1578063bde9e40e146100bb578063fe557173146100c5578063ffbb0bf9146100cf5761007b565b8063185c38a41461007f5780632e52d606146100895780633599773c146100a7575b5f5ffd5b6100876100d9565b005b61009161011e565b60405161009e919061022a565b60405180910390f35b6100af610123565b005b6100b9610193565b005b6100c3610197565b005b6100cd610199565b005b6100d7610209565b005b5f6040518060400160405280600481526020017f414243440000000000000000000000000000000000000000000000000000000081525090505f815190508060208301fd5b5f5481565b3073ffffffffffffffffffffffffffffffffffffffff1663ffbb0bf96040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610168575f5ffd5b505af1925050508015610179575060015b6101895760015f81905550610191565b60035f819055505b565b5f5ffd5bfe5b3073ffffffffffffffffffffffffffffffffffffffff1663185c38a46040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156101de575f5ffd5b505af19250505080156101ef575060015b6101ff5760015f81905550610207565b60035f819055505b565b59602081018082fd5b5f819050919050565b61022481610212565b82525050565b5f60208201905061023d5f83018461021b565b9291505056fea2646970667358221220a40e6e8d48f919c4a998c8d66be1f5a73ab00d7658345a5b176a70faf4edc5c664736f6c634300081e0033",
}

// InvalidABI is the input ABI used to generate the binding from.
// Deprecated: Use InvalidMetaData.ABI instead.
var InvalidABI = InvalidMetaData.ABI

// InvalidBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InvalidMetaData.Bin instead.
var InvalidBin = InvalidMetaData.Bin

// DeployInvalid deploys a new Ethereum contract, binding an instance of Invalid to it.
func DeployInvalid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Invalid, error) {
	parsed, err := InvalidMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InvalidBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Invalid{InvalidCaller: InvalidCaller{contract: contract}, InvalidTransactor: InvalidTransactor{contract: contract}, InvalidFilterer: InvalidFilterer{contract: contract}}, nil
}

// Invalid is an auto generated Go binding around an Ethereum contract.
type Invalid struct {
	InvalidCaller     // Read-only binding to the contract
	InvalidTransactor // Write-only binding to the contract
	InvalidFilterer   // Log filterer for contract events
}

// InvalidCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvalidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvalidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvalidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvalidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvalidSession struct {
	Contract     *Invalid          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvalidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvalidCallerSession struct {
	Contract *InvalidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// InvalidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvalidTransactorSession struct {
	Contract     *InvalidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InvalidRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvalidRaw struct {
	Contract *Invalid // Generic contract binding to access the raw methods on
}

// InvalidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvalidCallerRaw struct {
	Contract *InvalidCaller // Generic read-only contract binding to access the raw methods on
}

// InvalidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvalidTransactorRaw struct {
	Contract *InvalidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvalid creates a new instance of Invalid, bound to a specific deployed contract.
func NewInvalid(address common.Address, backend bind.ContractBackend) (*Invalid, error) {
	contract, err := bindInvalid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Invalid{InvalidCaller: InvalidCaller{contract: contract}, InvalidTransactor: InvalidTransactor{contract: contract}, InvalidFilterer: InvalidFilterer{contract: contract}}, nil
}

// NewInvalidCaller creates a new read-only instance of Invalid, bound to a specific deployed contract.
func NewInvalidCaller(address common.Address, caller bind.ContractCaller) (*InvalidCaller, error) {
	contract, err := bindInvalid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvalidCaller{contract: contract}, nil
}

// NewInvalidTransactor creates a new write-only instance of Invalid, bound to a specific deployed contract.
func NewInvalidTransactor(address common.Address, transactor bind.ContractTransactor) (*InvalidTransactor, error) {
	contract, err := bindInvalid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvalidTransactor{contract: contract}, nil
}

// NewInvalidFilterer creates a new log filterer instance of Invalid, bound to a specific deployed contract.
func NewInvalidFilterer(address common.Address, filterer bind.ContractFilterer) (*InvalidFilterer, error) {
	contract, err := bindInvalid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvalidFilterer{contract: contract}, nil
}

// bindInvalid binds a generic wrapper to an already deployed contract.
func bindInvalid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InvalidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Invalid *InvalidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Invalid.Contract.InvalidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Invalid *InvalidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.Contract.InvalidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Invalid *InvalidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Invalid.Contract.InvalidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Invalid *InvalidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Invalid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Invalid *InvalidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Invalid *InvalidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Invalid.Contract.contract.Transact(opts, method, params...)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Invalid *InvalidCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Invalid.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Invalid *InvalidSession) N() (*big.Int, error) {
	return _Invalid.Contract.N(&_Invalid.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Invalid *InvalidCallerSession) N() (*big.Int, error) {
	return _Invalid.Contract.N(&_Invalid.CallOpts)
}

// Invalid is a paid mutator transaction binding the contract method 0xbde9e40e.
//
// Solidity: function invalid() returns()
func (_Invalid *InvalidTransactor) Invalid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.contract.Transact(opts, "invalid")
}

// Invalid is a paid mutator transaction binding the contract method 0xbde9e40e.
//
// Solidity: function invalid() returns()
func (_Invalid *InvalidSession) Invalid() (*types.Transaction, error) {
	return _Invalid.Contract.Invalid(&_Invalid.TransactOpts)
}

// Invalid is a paid mutator transaction binding the contract method 0xbde9e40e.
//
// Solidity: function invalid() returns()
func (_Invalid *InvalidTransactorSession) Invalid() (*types.Transaction, error) {
	return _Invalid.Contract.Invalid(&_Invalid.TransactOpts)
}

// RevertInTryCatch is a paid mutator transaction binding the contract method 0xfe557173.
//
// Solidity: function revertInTryCatch() returns()
func (_Invalid *InvalidTransactor) RevertInTryCatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.contract.Transact(opts, "revertInTryCatch")
}

// RevertInTryCatch is a paid mutator transaction binding the contract method 0xfe557173.
//
// Solidity: function revertInTryCatch() returns()
func (_Invalid *InvalidSession) RevertInTryCatch() (*types.Transaction, error) {
	return _Invalid.Contract.RevertInTryCatch(&_Invalid.TransactOpts)
}

// RevertInTryCatch is a paid mutator transaction binding the contract method 0xfe557173.
//
// Solidity: function revertInTryCatch() returns()
func (_Invalid *InvalidTransactorSession) RevertInTryCatch() (*types.Transaction, error) {
	return _Invalid.Contract.RevertInTryCatch(&_Invalid.TransactOpts)
}

// RevertInTryCatchWithMemoryExpansion is a paid mutator transaction binding the contract method 0x3599773c.
//
// Solidity: function revertInTryCatchWithMemoryExpansion() returns()
func (_Invalid *InvalidTransactor) RevertInTryCatchWithMemoryExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.contract.Transact(opts, "revertInTryCatchWithMemoryExpansion")
}

// RevertInTryCatchWithMemoryExpansion is a paid mutator transaction binding the contract method 0x3599773c.
//
// Solidity: function revertInTryCatchWithMemoryExpansion() returns()
func (_Invalid *InvalidSession) RevertInTryCatchWithMemoryExpansion() (*types.Transaction, error) {
	return _Invalid.Contract.RevertInTryCatchWithMemoryExpansion(&_Invalid.TransactOpts)
}

// RevertInTryCatchWithMemoryExpansion is a paid mutator transaction binding the contract method 0x3599773c.
//
// Solidity: function revertInTryCatchWithMemoryExpansion() returns()
func (_Invalid *InvalidTransactorSession) RevertInTryCatchWithMemoryExpansion() (*types.Transaction, error) {
	return _Invalid.Contract.RevertInTryCatchWithMemoryExpansion(&_Invalid.TransactOpts)
}

// RevertNoMessage is a paid mutator transaction binding the contract method 0x9fe55433.
//
// Solidity: function revertNoMessage() returns()
func (_Invalid *InvalidTransactor) RevertNoMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.contract.Transact(opts, "revertNoMessage")
}

// RevertNoMessage is a paid mutator transaction binding the contract method 0x9fe55433.
//
// Solidity: function revertNoMessage() returns()
func (_Invalid *InvalidSession) RevertNoMessage() (*types.Transaction, error) {
	return _Invalid.Contract.RevertNoMessage(&_Invalid.TransactOpts)
}

// RevertNoMessage is a paid mutator transaction binding the contract method 0x9fe55433.
//
// Solidity: function revertNoMessage() returns()
func (_Invalid *InvalidTransactorSession) RevertNoMessage() (*types.Transaction, error) {
	return _Invalid.Contract.RevertNoMessage(&_Invalid.TransactOpts)
}

// RevertWithMemoryExpansion is a paid mutator transaction binding the contract method 0xffbb0bf9.
//
// Solidity: function revertWithMemoryExpansion() returns()
func (_Invalid *InvalidTransactor) RevertWithMemoryExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.contract.Transact(opts, "revertWithMemoryExpansion")
}

// RevertWithMemoryExpansion is a paid mutator transaction binding the contract method 0xffbb0bf9.
//
// Solidity: function revertWithMemoryExpansion() returns()
func (_Invalid *InvalidSession) RevertWithMemoryExpansion() (*types.Transaction, error) {
	return _Invalid.Contract.RevertWithMemoryExpansion(&_Invalid.TransactOpts)
}

// RevertWithMemoryExpansion is a paid mutator transaction binding the contract method 0xffbb0bf9.
//
// Solidity: function revertWithMemoryExpansion() returns()
func (_Invalid *InvalidTransactorSession) RevertWithMemoryExpansion() (*types.Transaction, error) {
	return _Invalid.Contract.RevertWithMemoryExpansion(&_Invalid.TransactOpts)
}

// RevertWithMessage is a paid mutator transaction binding the contract method 0x185c38a4.
//
// Solidity: function revertWithMessage() returns()
func (_Invalid *InvalidTransactor) RevertWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Invalid.contract.Transact(opts, "revertWithMessage")
}

// RevertWithMessage is a paid mutator transaction binding the contract method 0x185c38a4.
//
// Solidity: function revertWithMessage() returns()
func (_Invalid *InvalidSession) RevertWithMessage() (*types.Transaction, error) {
	return _Invalid.Contract.RevertWithMessage(&_Invalid.TransactOpts)
}

// RevertWithMessage is a paid mutator transaction binding the contract method 0x185c38a4.
//
// Solidity: function revertWithMessage() returns()
func (_Invalid *InvalidTransactorSession) RevertWithMessage() (*types.Transaction, error) {
	return _Invalid.Contract.RevertWithMessage(&_Invalid.TransactOpts)
}

// LogEmitterMetaData contains all meta data concerning the LogEmitter contract.
var LogEmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"emitFourTopics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitFourTopicsExtraData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitFourTopicsExtraDataAndMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitOneTopicEmptyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitOneTopicNonEmptyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitOneTopicNonEmptyDataAndMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitThreeTopics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitThreeTopicsExtraData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitThreeTopicsExtraDataAndMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitTwoTopics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitTwoTopicsExtraData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitTwoTopicsExtraDataAndMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitZeroTopicEmptyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitZeroTopicNonEmptyData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitZeroTopicNonEmptyDataAndMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keccakHash\",\"type\":\"bytes32\"}],\"name\":\"LogFourTopics\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keccakHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"LogFourTopicsExtraData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogOneTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"LogOneTopicExtraData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addy\",\"type\":\"address\"}],\"name\":\"LogThreeTopics\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"LogThreeTopicsExtraData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"LogTwoTopics\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"LogTwoTopicsExtraData\",\"type\":\"event\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610ad78061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c806389191c3711610095578063d3c5744411610064578063d3c574441461018e578063f275834014610198578063fa6b885a146101a2578063fe81bed4146101ac576100fe565b806389191c3714610166578063939d4eaa14610170578063ab51095a1461017a578063ba95312614610184576100fe565b806367c93ee1116100d157806367c93ee11461012a5780636e9f9187146101345780637463ebe61461013e5780638381f58a14610148576100fe565b806301db8e6b1461010257806329823b111461010c5780633e59a3c3146101165780634b3be30a14610120575b5f5ffd5b61010a6101b6565b005b6101146101f2565b005b61011e610210565b005b6101286102a5565b005b6101326102cd565b005b61013c610313565b005b610146610357565b005b6101506103ad565b60405161015d91906106ae565b60405180910390f35b61016e6103b2565b005b61017861043c565b005b6101826104f7565b005b61018c610535565b005b610196610580565b005b6101a06105af565b005b6101aa610612565b005b6101b461065a565b005b5f6113375f1b90505f6113385f1b90505f6113395f1b905059818385604084a3505f5f8154809291906101e8906106f4565b9190505550505050565b59604081a0505f5f815480929190610209906106f4565b9190505550565b63feedface73ffffffffffffffffffffffffffffffffffffffff1663b00bb00c7f9111be8bfd08a48f95eadcdb6694214a59e1cc21ca0577e1d07856e916d82bef6040516020016102609061078f565b604051602081830303815290604052610278906107e5565b604051610285919061085a565b60405180910390a35f5f81548092919061029e906106f4565b9190505550565b5f6113375f1b90505981604082a1505f5f8154809291906102c5906106f4565b919050555050565b5f6113375f1b90505f6113385f1b90505f6113395f1b90505f61133a5f1b90505981838587604085a4505f5f815480929190610308906106f4565b919050555050505050565b7f60aa14871395c0cbe2d544024235952891fce304aba1048e0f8ba021339c6beb60405160405180910390a15f5f815480929190610350906106f4565b9190505550565b61caca7f0d95b68954ecebf6f87262a05e0f5f134d5942feba6e7d0c96a647ec0d0a922563deadcafe60405161038d91906108b2565b60405180910390a25f5f8154809291906103a6906106f4565b9190505550565b5f5481565b6040516020016103c190610915565b6040516020818303038152906040528051906020012063beeffeed73ffffffffffffffffffffffffffffffffffffffff16640fadedb00b7f50a18b741ef541cbecce9633dbddc45a366e942a630d038640c3d0f979c8f8e360405160405180910390a45f5f815480929190610435906106f4565b9190505550565b60405160200161044b90610973565b6040516020818303038152906040528051906020012063feedcafe73ffffffffffffffffffffffffffffffffffffffff16640daff0d1777f020c8d0cc6bc20bd920ce3723c79a0eba3a33018f33a401dddd22a9d78973d2b6040516020016104b2906109d1565b6040516020818303038152906040526104ca906107e5565b6040516104d7919061085a565b60405180910390a45f5f8154809291906104f0906106f4565b9190505550565b61051d60405160200161050990610a2f565b60405160208183030381529060405261068c565b5f5f81548092919061052e906106f4565b9190505550565b5f60405160200161054590610a8d565b60405160208183030381529060405290505f6113375f1b905080825160208401a15f5f815480929190610577906106f4565b91905055505050565b61059760405180602001604052805f81525061068c565b5f5f8154809291906105a8906106f4565b9190505550565b63babecafe73ffffffffffffffffffffffffffffffffffffffff1662b1337b7f70a85a3e24dfc01a8d3f33c838120eb893c800ece385461930e9f3d0cfc20e4260405160405180910390a35f5f81548092919061060b906106f4565b9190505550565b62d00d007f6370fe7ee1b1199c3c4673346f536f2a2a55f2a717f6c68b528a14757050aab860405160405180910390a25f5f815480929190610653906106f4565b9190505550565b5f6113375f1b90505f6113385f1b9050598183604083a2505f5f815480929190610683906106f4565b91905055505050565b805160208201a050565b5f819050919050565b6106a881610696565b82525050565b5f6020820190506106c15f83018461069f565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6106fe82610696565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107305761072f6106c7565b5b600182019050919050565b5f81905092915050565b7f48494a4b4c4d4e4f5000000000000000000000000000000000000000000000005f82015250565b5f61077960098361073b565b915061078482610745565b600982019050919050565b5f6107998261076d565b9150819050919050565b5f81519050919050565b5f819050602082019050919050565b5f819050919050565b5f6107d082516107bc565b80915050919050565b5f82821b905092915050565b5f6107ef826107a3565b826107f9846107ad565b9050610804816107c5565b925060208210156108445761083f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff836020036008026107d9565b831692505b5050919050565b610854816107bc565b82525050565b5f60208201905061086d5f83018461084b565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61089c82610873565b9050919050565b6108ac81610892565b82525050565b5f6020820190506108c55f8301846108a3565b92915050565b7f5152535455565758595a000000000000000000000000000000000000000000005f82015250565b5f6108ff600a8361073b565b915061090a826108cb565b600a82019050919050565b5f61091f826108f3565b9150819050919050565b7f41424344000000000000000000000000000000000000000000000000000000005f82015250565b5f61095d60048361073b565b915061096882610929565b600482019050919050565b5f61097d82610951565b9150819050919050565b7f41424344454647000000000000000000000000000000000000000000000000005f82015250565b5f6109bb60078361073b565b91506109c682610987565b600782019050919050565b5f6109db826109af565b9150819050919050565b7f61626364656667000000000000000000000000000000000000000000000000005f82015250565b5f610a1960078361073b565b9150610a24826109e5565b600782019050919050565b5f610a3982610a0d565b9150819050919050565b7f68696a6b6c6d6e6f7000000000000000000000000000000000000000000000005f82015250565b5f610a7760098361073b565b9150610a8282610a43565b600982019050919050565b5f610a9782610a6b565b915081905091905056fea2646970667358221220401e7d1001d5ff801f111c8d9e771e421c7dffa11d220f1f1fc804edd7b350cb64736f6c634300081e0033",
}

// LogEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use LogEmitterMetaData.ABI instead.
var LogEmitterABI = LogEmitterMetaData.ABI

// LogEmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LogEmitterMetaData.Bin instead.
var LogEmitterBin = LogEmitterMetaData.Bin

// DeployLogEmitter deploys a new Ethereum contract, binding an instance of LogEmitter to it.
func DeployLogEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LogEmitter, error) {
	parsed, err := LogEmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LogEmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LogEmitter{LogEmitterCaller: LogEmitterCaller{contract: contract}, LogEmitterTransactor: LogEmitterTransactor{contract: contract}, LogEmitterFilterer: LogEmitterFilterer{contract: contract}}, nil
}

// LogEmitter is an auto generated Go binding around an Ethereum contract.
type LogEmitter struct {
	LogEmitterCaller     // Read-only binding to the contract
	LogEmitterTransactor // Write-only binding to the contract
	LogEmitterFilterer   // Log filterer for contract events
}

// LogEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type LogEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LogEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LogEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LogEmitterSession struct {
	Contract     *LogEmitter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LogEmitterCallerSession struct {
	Contract *LogEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LogEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LogEmitterTransactorSession struct {
	Contract     *LogEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LogEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type LogEmitterRaw struct {
	Contract *LogEmitter // Generic contract binding to access the raw methods on
}

// LogEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LogEmitterCallerRaw struct {
	Contract *LogEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// LogEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LogEmitterTransactorRaw struct {
	Contract *LogEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogEmitter creates a new instance of LogEmitter, bound to a specific deployed contract.
func NewLogEmitter(address common.Address, backend bind.ContractBackend) (*LogEmitter, error) {
	contract, err := bindLogEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LogEmitter{LogEmitterCaller: LogEmitterCaller{contract: contract}, LogEmitterTransactor: LogEmitterTransactor{contract: contract}, LogEmitterFilterer: LogEmitterFilterer{contract: contract}}, nil
}

// NewLogEmitterCaller creates a new read-only instance of LogEmitter, bound to a specific deployed contract.
func NewLogEmitterCaller(address common.Address, caller bind.ContractCaller) (*LogEmitterCaller, error) {
	contract, err := bindLogEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LogEmitterCaller{contract: contract}, nil
}

// NewLogEmitterTransactor creates a new write-only instance of LogEmitter, bound to a specific deployed contract.
func NewLogEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*LogEmitterTransactor, error) {
	contract, err := bindLogEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LogEmitterTransactor{contract: contract}, nil
}

// NewLogEmitterFilterer creates a new log filterer instance of LogEmitter, bound to a specific deployed contract.
func NewLogEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*LogEmitterFilterer, error) {
	contract, err := bindLogEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LogEmitterFilterer{contract: contract}, nil
}

// bindLogEmitter binds a generic wrapper to an already deployed contract.
func bindLogEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LogEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LogEmitter *LogEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LogEmitter.Contract.LogEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LogEmitter *LogEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.Contract.LogEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LogEmitter *LogEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LogEmitter.Contract.LogEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LogEmitter *LogEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LogEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LogEmitter *LogEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LogEmitter *LogEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LogEmitter.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_LogEmitter *LogEmitterCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LogEmitter.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_LogEmitter *LogEmitterSession) Number() (*big.Int, error) {
	return _LogEmitter.Contract.Number(&_LogEmitter.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_LogEmitter *LogEmitterCallerSession) Number() (*big.Int, error) {
	return _LogEmitter.Contract.Number(&_LogEmitter.CallOpts)
}

// EmitFourTopics is a paid mutator transaction binding the contract method 0x89191c37.
//
// Solidity: function emitFourTopics() returns()
func (_LogEmitter *LogEmitterTransactor) EmitFourTopics(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitFourTopics")
}

// EmitFourTopics is a paid mutator transaction binding the contract method 0x89191c37.
//
// Solidity: function emitFourTopics() returns()
func (_LogEmitter *LogEmitterSession) EmitFourTopics() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitFourTopics(&_LogEmitter.TransactOpts)
}

// EmitFourTopics is a paid mutator transaction binding the contract method 0x89191c37.
//
// Solidity: function emitFourTopics() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitFourTopics() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitFourTopics(&_LogEmitter.TransactOpts)
}

// EmitFourTopicsExtraData is a paid mutator transaction binding the contract method 0x939d4eaa.
//
// Solidity: function emitFourTopicsExtraData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitFourTopicsExtraData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitFourTopicsExtraData")
}

// EmitFourTopicsExtraData is a paid mutator transaction binding the contract method 0x939d4eaa.
//
// Solidity: function emitFourTopicsExtraData() returns()
func (_LogEmitter *LogEmitterSession) EmitFourTopicsExtraData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitFourTopicsExtraData(&_LogEmitter.TransactOpts)
}

// EmitFourTopicsExtraData is a paid mutator transaction binding the contract method 0x939d4eaa.
//
// Solidity: function emitFourTopicsExtraData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitFourTopicsExtraData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitFourTopicsExtraData(&_LogEmitter.TransactOpts)
}

// EmitFourTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0x67c93ee1.
//
// Solidity: function emitFourTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactor) EmitFourTopicsExtraDataAndMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitFourTopicsExtraDataAndMemExpansion")
}

// EmitFourTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0x67c93ee1.
//
// Solidity: function emitFourTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterSession) EmitFourTopicsExtraDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitFourTopicsExtraDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitFourTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0x67c93ee1.
//
// Solidity: function emitFourTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitFourTopicsExtraDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitFourTopicsExtraDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitOneTopicEmptyData is a paid mutator transaction binding the contract method 0x6e9f9187.
//
// Solidity: function emitOneTopicEmptyData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitOneTopicEmptyData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitOneTopicEmptyData")
}

// EmitOneTopicEmptyData is a paid mutator transaction binding the contract method 0x6e9f9187.
//
// Solidity: function emitOneTopicEmptyData() returns()
func (_LogEmitter *LogEmitterSession) EmitOneTopicEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitOneTopicEmptyData(&_LogEmitter.TransactOpts)
}

// EmitOneTopicEmptyData is a paid mutator transaction binding the contract method 0x6e9f9187.
//
// Solidity: function emitOneTopicEmptyData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitOneTopicEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitOneTopicEmptyData(&_LogEmitter.TransactOpts)
}

// EmitOneTopicNonEmptyData is a paid mutator transaction binding the contract method 0xba953126.
//
// Solidity: function emitOneTopicNonEmptyData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitOneTopicNonEmptyData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitOneTopicNonEmptyData")
}

// EmitOneTopicNonEmptyData is a paid mutator transaction binding the contract method 0xba953126.
//
// Solidity: function emitOneTopicNonEmptyData() returns()
func (_LogEmitter *LogEmitterSession) EmitOneTopicNonEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitOneTopicNonEmptyData(&_LogEmitter.TransactOpts)
}

// EmitOneTopicNonEmptyData is a paid mutator transaction binding the contract method 0xba953126.
//
// Solidity: function emitOneTopicNonEmptyData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitOneTopicNonEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitOneTopicNonEmptyData(&_LogEmitter.TransactOpts)
}

// EmitOneTopicNonEmptyDataAndMemExpansion is a paid mutator transaction binding the contract method 0x4b3be30a.
//
// Solidity: function emitOneTopicNonEmptyDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactor) EmitOneTopicNonEmptyDataAndMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitOneTopicNonEmptyDataAndMemExpansion")
}

// EmitOneTopicNonEmptyDataAndMemExpansion is a paid mutator transaction binding the contract method 0x4b3be30a.
//
// Solidity: function emitOneTopicNonEmptyDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterSession) EmitOneTopicNonEmptyDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitOneTopicNonEmptyDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitOneTopicNonEmptyDataAndMemExpansion is a paid mutator transaction binding the contract method 0x4b3be30a.
//
// Solidity: function emitOneTopicNonEmptyDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitOneTopicNonEmptyDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitOneTopicNonEmptyDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitThreeTopics is a paid mutator transaction binding the contract method 0xf2758340.
//
// Solidity: function emitThreeTopics() returns()
func (_LogEmitter *LogEmitterTransactor) EmitThreeTopics(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitThreeTopics")
}

// EmitThreeTopics is a paid mutator transaction binding the contract method 0xf2758340.
//
// Solidity: function emitThreeTopics() returns()
func (_LogEmitter *LogEmitterSession) EmitThreeTopics() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitThreeTopics(&_LogEmitter.TransactOpts)
}

// EmitThreeTopics is a paid mutator transaction binding the contract method 0xf2758340.
//
// Solidity: function emitThreeTopics() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitThreeTopics() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitThreeTopics(&_LogEmitter.TransactOpts)
}

// EmitThreeTopicsExtraData is a paid mutator transaction binding the contract method 0x3e59a3c3.
//
// Solidity: function emitThreeTopicsExtraData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitThreeTopicsExtraData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitThreeTopicsExtraData")
}

// EmitThreeTopicsExtraData is a paid mutator transaction binding the contract method 0x3e59a3c3.
//
// Solidity: function emitThreeTopicsExtraData() returns()
func (_LogEmitter *LogEmitterSession) EmitThreeTopicsExtraData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitThreeTopicsExtraData(&_LogEmitter.TransactOpts)
}

// EmitThreeTopicsExtraData is a paid mutator transaction binding the contract method 0x3e59a3c3.
//
// Solidity: function emitThreeTopicsExtraData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitThreeTopicsExtraData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitThreeTopicsExtraData(&_LogEmitter.TransactOpts)
}

// EmitThreeTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0x01db8e6b.
//
// Solidity: function emitThreeTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactor) EmitThreeTopicsExtraDataAndMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitThreeTopicsExtraDataAndMemExpansion")
}

// EmitThreeTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0x01db8e6b.
//
// Solidity: function emitThreeTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterSession) EmitThreeTopicsExtraDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitThreeTopicsExtraDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitThreeTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0x01db8e6b.
//
// Solidity: function emitThreeTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitThreeTopicsExtraDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitThreeTopicsExtraDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitTwoTopics is a paid mutator transaction binding the contract method 0xfa6b885a.
//
// Solidity: function emitTwoTopics() returns()
func (_LogEmitter *LogEmitterTransactor) EmitTwoTopics(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitTwoTopics")
}

// EmitTwoTopics is a paid mutator transaction binding the contract method 0xfa6b885a.
//
// Solidity: function emitTwoTopics() returns()
func (_LogEmitter *LogEmitterSession) EmitTwoTopics() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitTwoTopics(&_LogEmitter.TransactOpts)
}

// EmitTwoTopics is a paid mutator transaction binding the contract method 0xfa6b885a.
//
// Solidity: function emitTwoTopics() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitTwoTopics() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitTwoTopics(&_LogEmitter.TransactOpts)
}

// EmitTwoTopicsExtraData is a paid mutator transaction binding the contract method 0x7463ebe6.
//
// Solidity: function emitTwoTopicsExtraData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitTwoTopicsExtraData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitTwoTopicsExtraData")
}

// EmitTwoTopicsExtraData is a paid mutator transaction binding the contract method 0x7463ebe6.
//
// Solidity: function emitTwoTopicsExtraData() returns()
func (_LogEmitter *LogEmitterSession) EmitTwoTopicsExtraData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitTwoTopicsExtraData(&_LogEmitter.TransactOpts)
}

// EmitTwoTopicsExtraData is a paid mutator transaction binding the contract method 0x7463ebe6.
//
// Solidity: function emitTwoTopicsExtraData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitTwoTopicsExtraData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitTwoTopicsExtraData(&_LogEmitter.TransactOpts)
}

// EmitTwoTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0xfe81bed4.
//
// Solidity: function emitTwoTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactor) EmitTwoTopicsExtraDataAndMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitTwoTopicsExtraDataAndMemExpansion")
}

// EmitTwoTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0xfe81bed4.
//
// Solidity: function emitTwoTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterSession) EmitTwoTopicsExtraDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitTwoTopicsExtraDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitTwoTopicsExtraDataAndMemExpansion is a paid mutator transaction binding the contract method 0xfe81bed4.
//
// Solidity: function emitTwoTopicsExtraDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitTwoTopicsExtraDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitTwoTopicsExtraDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitZeroTopicEmptyData is a paid mutator transaction binding the contract method 0xd3c57444.
//
// Solidity: function emitZeroTopicEmptyData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitZeroTopicEmptyData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitZeroTopicEmptyData")
}

// EmitZeroTopicEmptyData is a paid mutator transaction binding the contract method 0xd3c57444.
//
// Solidity: function emitZeroTopicEmptyData() returns()
func (_LogEmitter *LogEmitterSession) EmitZeroTopicEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitZeroTopicEmptyData(&_LogEmitter.TransactOpts)
}

// EmitZeroTopicEmptyData is a paid mutator transaction binding the contract method 0xd3c57444.
//
// Solidity: function emitZeroTopicEmptyData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitZeroTopicEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitZeroTopicEmptyData(&_LogEmitter.TransactOpts)
}

// EmitZeroTopicNonEmptyData is a paid mutator transaction binding the contract method 0xab51095a.
//
// Solidity: function emitZeroTopicNonEmptyData() returns()
func (_LogEmitter *LogEmitterTransactor) EmitZeroTopicNonEmptyData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitZeroTopicNonEmptyData")
}

// EmitZeroTopicNonEmptyData is a paid mutator transaction binding the contract method 0xab51095a.
//
// Solidity: function emitZeroTopicNonEmptyData() returns()
func (_LogEmitter *LogEmitterSession) EmitZeroTopicNonEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitZeroTopicNonEmptyData(&_LogEmitter.TransactOpts)
}

// EmitZeroTopicNonEmptyData is a paid mutator transaction binding the contract method 0xab51095a.
//
// Solidity: function emitZeroTopicNonEmptyData() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitZeroTopicNonEmptyData() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitZeroTopicNonEmptyData(&_LogEmitter.TransactOpts)
}

// EmitZeroTopicNonEmptyDataAndMemExpansion is a paid mutator transaction binding the contract method 0x29823b11.
//
// Solidity: function emitZeroTopicNonEmptyDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactor) EmitZeroTopicNonEmptyDataAndMemExpansion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LogEmitter.contract.Transact(opts, "emitZeroTopicNonEmptyDataAndMemExpansion")
}

// EmitZeroTopicNonEmptyDataAndMemExpansion is a paid mutator transaction binding the contract method 0x29823b11.
//
// Solidity: function emitZeroTopicNonEmptyDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterSession) EmitZeroTopicNonEmptyDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitZeroTopicNonEmptyDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// EmitZeroTopicNonEmptyDataAndMemExpansion is a paid mutator transaction binding the contract method 0x29823b11.
//
// Solidity: function emitZeroTopicNonEmptyDataAndMemExpansion() returns()
func (_LogEmitter *LogEmitterTransactorSession) EmitZeroTopicNonEmptyDataAndMemExpansion() (*types.Transaction, error) {
	return _LogEmitter.Contract.EmitZeroTopicNonEmptyDataAndMemExpansion(&_LogEmitter.TransactOpts)
}

// LogEmitterLogFourTopicsIterator is returned from FilterLogFourTopics and is used to iterate over the raw logs and unpacked data for LogFourTopics events raised by the LogEmitter contract.
type LogEmitterLogFourTopicsIterator struct {
	Event *LogEmitterLogFourTopics // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogFourTopicsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogFourTopics)
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
		it.Event = new(LogEmitterLogFourTopics)
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
func (it *LogEmitterLogFourTopicsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogFourTopicsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogFourTopics represents a LogFourTopics event raised by the LogEmitter contract.
type LogEmitterLogFourTopics struct {
	Number     *big.Int
	Addy       common.Address
	KeccakHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogFourTopics is a free log retrieval operation binding the contract event 0x50a18b741ef541cbecce9633dbddc45a366e942a630d038640c3d0f979c8f8e3.
//
// Solidity: event LogFourTopics(uint256 indexed number, address indexed addy, bytes32 indexed keccakHash)
func (_LogEmitter *LogEmitterFilterer) FilterLogFourTopics(opts *bind.FilterOpts, number []*big.Int, addy []common.Address, keccakHash [][32]byte) (*LogEmitterLogFourTopicsIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}
	var keccakHashRule []interface{}
	for _, keccakHashItem := range keccakHash {
		keccakHashRule = append(keccakHashRule, keccakHashItem)
	}

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogFourTopics", numberRule, addyRule, keccakHashRule)
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogFourTopicsIterator{contract: _LogEmitter.contract, event: "LogFourTopics", logs: logs, sub: sub}, nil
}

// WatchLogFourTopics is a free log subscription operation binding the contract event 0x50a18b741ef541cbecce9633dbddc45a366e942a630d038640c3d0f979c8f8e3.
//
// Solidity: event LogFourTopics(uint256 indexed number, address indexed addy, bytes32 indexed keccakHash)
func (_LogEmitter *LogEmitterFilterer) WatchLogFourTopics(opts *bind.WatchOpts, sink chan<- *LogEmitterLogFourTopics, number []*big.Int, addy []common.Address, keccakHash [][32]byte) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}
	var keccakHashRule []interface{}
	for _, keccakHashItem := range keccakHash {
		keccakHashRule = append(keccakHashRule, keccakHashItem)
	}

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogFourTopics", numberRule, addyRule, keccakHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogFourTopics)
				if err := _LogEmitter.contract.UnpackLog(event, "LogFourTopics", log); err != nil {
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

// ParseLogFourTopics is a log parse operation binding the contract event 0x50a18b741ef541cbecce9633dbddc45a366e942a630d038640c3d0f979c8f8e3.
//
// Solidity: event LogFourTopics(uint256 indexed number, address indexed addy, bytes32 indexed keccakHash)
func (_LogEmitter *LogEmitterFilterer) ParseLogFourTopics(log types.Log) (*LogEmitterLogFourTopics, error) {
	event := new(LogEmitterLogFourTopics)
	if err := _LogEmitter.contract.UnpackLog(event, "LogFourTopics", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogFourTopicsExtraDataIterator is returned from FilterLogFourTopicsExtraData and is used to iterate over the raw logs and unpacked data for LogFourTopicsExtraData events raised by the LogEmitter contract.
type LogEmitterLogFourTopicsExtraDataIterator struct {
	Event *LogEmitterLogFourTopicsExtraData // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogFourTopicsExtraDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogFourTopicsExtraData)
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
		it.Event = new(LogEmitterLogFourTopicsExtraData)
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
func (it *LogEmitterLogFourTopicsExtraDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogFourTopicsExtraDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogFourTopicsExtraData represents a LogFourTopicsExtraData event raised by the LogEmitter contract.
type LogEmitterLogFourTopicsExtraData struct {
	Number     *big.Int
	Addy       common.Address
	KeccakHash [32]byte
	Arg3       [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogFourTopicsExtraData is a free log retrieval operation binding the contract event 0x020c8d0cc6bc20bd920ce3723c79a0eba3a33018f33a401dddd22a9d78973d2b.
//
// Solidity: event LogFourTopicsExtraData(uint256 indexed number, address indexed addy, bytes32 indexed keccakHash, bytes32 arg3)
func (_LogEmitter *LogEmitterFilterer) FilterLogFourTopicsExtraData(opts *bind.FilterOpts, number []*big.Int, addy []common.Address, keccakHash [][32]byte) (*LogEmitterLogFourTopicsExtraDataIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}
	var keccakHashRule []interface{}
	for _, keccakHashItem := range keccakHash {
		keccakHashRule = append(keccakHashRule, keccakHashItem)
	}

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogFourTopicsExtraData", numberRule, addyRule, keccakHashRule)
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogFourTopicsExtraDataIterator{contract: _LogEmitter.contract, event: "LogFourTopicsExtraData", logs: logs, sub: sub}, nil
}

// WatchLogFourTopicsExtraData is a free log subscription operation binding the contract event 0x020c8d0cc6bc20bd920ce3723c79a0eba3a33018f33a401dddd22a9d78973d2b.
//
// Solidity: event LogFourTopicsExtraData(uint256 indexed number, address indexed addy, bytes32 indexed keccakHash, bytes32 arg3)
func (_LogEmitter *LogEmitterFilterer) WatchLogFourTopicsExtraData(opts *bind.WatchOpts, sink chan<- *LogEmitterLogFourTopicsExtraData, number []*big.Int, addy []common.Address, keccakHash [][32]byte) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}
	var keccakHashRule []interface{}
	for _, keccakHashItem := range keccakHash {
		keccakHashRule = append(keccakHashRule, keccakHashItem)
	}

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogFourTopicsExtraData", numberRule, addyRule, keccakHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogFourTopicsExtraData)
				if err := _LogEmitter.contract.UnpackLog(event, "LogFourTopicsExtraData", log); err != nil {
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

// ParseLogFourTopicsExtraData is a log parse operation binding the contract event 0x020c8d0cc6bc20bd920ce3723c79a0eba3a33018f33a401dddd22a9d78973d2b.
//
// Solidity: event LogFourTopicsExtraData(uint256 indexed number, address indexed addy, bytes32 indexed keccakHash, bytes32 arg3)
func (_LogEmitter *LogEmitterFilterer) ParseLogFourTopicsExtraData(log types.Log) (*LogEmitterLogFourTopicsExtraData, error) {
	event := new(LogEmitterLogFourTopicsExtraData)
	if err := _LogEmitter.contract.UnpackLog(event, "LogFourTopicsExtraData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogOneTopicIterator is returned from FilterLogOneTopic and is used to iterate over the raw logs and unpacked data for LogOneTopic events raised by the LogEmitter contract.
type LogEmitterLogOneTopicIterator struct {
	Event *LogEmitterLogOneTopic // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogOneTopicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogOneTopic)
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
		it.Event = new(LogEmitterLogOneTopic)
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
func (it *LogEmitterLogOneTopicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogOneTopicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogOneTopic represents a LogOneTopic event raised by the LogEmitter contract.
type LogEmitterLogOneTopic struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogOneTopic is a free log retrieval operation binding the contract event 0x60aa14871395c0cbe2d544024235952891fce304aba1048e0f8ba021339c6beb.
//
// Solidity: event LogOneTopic()
func (_LogEmitter *LogEmitterFilterer) FilterLogOneTopic(opts *bind.FilterOpts) (*LogEmitterLogOneTopicIterator, error) {

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogOneTopic")
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogOneTopicIterator{contract: _LogEmitter.contract, event: "LogOneTopic", logs: logs, sub: sub}, nil
}

// WatchLogOneTopic is a free log subscription operation binding the contract event 0x60aa14871395c0cbe2d544024235952891fce304aba1048e0f8ba021339c6beb.
//
// Solidity: event LogOneTopic()
func (_LogEmitter *LogEmitterFilterer) WatchLogOneTopic(opts *bind.WatchOpts, sink chan<- *LogEmitterLogOneTopic) (event.Subscription, error) {

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogOneTopic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogOneTopic)
				if err := _LogEmitter.contract.UnpackLog(event, "LogOneTopic", log); err != nil {
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

// ParseLogOneTopic is a log parse operation binding the contract event 0x60aa14871395c0cbe2d544024235952891fce304aba1048e0f8ba021339c6beb.
//
// Solidity: event LogOneTopic()
func (_LogEmitter *LogEmitterFilterer) ParseLogOneTopic(log types.Log) (*LogEmitterLogOneTopic, error) {
	event := new(LogEmitterLogOneTopic)
	if err := _LogEmitter.contract.UnpackLog(event, "LogOneTopic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogOneTopicExtraDataIterator is returned from FilterLogOneTopicExtraData and is used to iterate over the raw logs and unpacked data for LogOneTopicExtraData events raised by the LogEmitter contract.
type LogEmitterLogOneTopicExtraDataIterator struct {
	Event *LogEmitterLogOneTopicExtraData // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogOneTopicExtraDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogOneTopicExtraData)
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
		it.Event = new(LogEmitterLogOneTopicExtraData)
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
func (it *LogEmitterLogOneTopicExtraDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogOneTopicExtraDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogOneTopicExtraData represents a LogOneTopicExtraData event raised by the LogEmitter contract.
type LogEmitterLogOneTopicExtraData struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogOneTopicExtraData is a free log retrieval operation binding the contract event 0x6c389f22489bf5f86f2f4f867d3f92f40955f94e531dd760afd7446977ab3c19.
//
// Solidity: event LogOneTopicExtraData(bytes arg0)
func (_LogEmitter *LogEmitterFilterer) FilterLogOneTopicExtraData(opts *bind.FilterOpts) (*LogEmitterLogOneTopicExtraDataIterator, error) {

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogOneTopicExtraData")
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogOneTopicExtraDataIterator{contract: _LogEmitter.contract, event: "LogOneTopicExtraData", logs: logs, sub: sub}, nil
}

// WatchLogOneTopicExtraData is a free log subscription operation binding the contract event 0x6c389f22489bf5f86f2f4f867d3f92f40955f94e531dd760afd7446977ab3c19.
//
// Solidity: event LogOneTopicExtraData(bytes arg0)
func (_LogEmitter *LogEmitterFilterer) WatchLogOneTopicExtraData(opts *bind.WatchOpts, sink chan<- *LogEmitterLogOneTopicExtraData) (event.Subscription, error) {

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogOneTopicExtraData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogOneTopicExtraData)
				if err := _LogEmitter.contract.UnpackLog(event, "LogOneTopicExtraData", log); err != nil {
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

// ParseLogOneTopicExtraData is a log parse operation binding the contract event 0x6c389f22489bf5f86f2f4f867d3f92f40955f94e531dd760afd7446977ab3c19.
//
// Solidity: event LogOneTopicExtraData(bytes arg0)
func (_LogEmitter *LogEmitterFilterer) ParseLogOneTopicExtraData(log types.Log) (*LogEmitterLogOneTopicExtraData, error) {
	event := new(LogEmitterLogOneTopicExtraData)
	if err := _LogEmitter.contract.UnpackLog(event, "LogOneTopicExtraData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogThreeTopicsIterator is returned from FilterLogThreeTopics and is used to iterate over the raw logs and unpacked data for LogThreeTopics events raised by the LogEmitter contract.
type LogEmitterLogThreeTopicsIterator struct {
	Event *LogEmitterLogThreeTopics // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogThreeTopicsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogThreeTopics)
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
		it.Event = new(LogEmitterLogThreeTopics)
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
func (it *LogEmitterLogThreeTopicsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogThreeTopicsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogThreeTopics represents a LogThreeTopics event raised by the LogEmitter contract.
type LogEmitterLogThreeTopics struct {
	Number *big.Int
	Addy   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogThreeTopics is a free log retrieval operation binding the contract event 0x70a85a3e24dfc01a8d3f33c838120eb893c800ece385461930e9f3d0cfc20e42.
//
// Solidity: event LogThreeTopics(uint256 indexed number, address indexed addy)
func (_LogEmitter *LogEmitterFilterer) FilterLogThreeTopics(opts *bind.FilterOpts, number []*big.Int, addy []common.Address) (*LogEmitterLogThreeTopicsIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogThreeTopics", numberRule, addyRule)
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogThreeTopicsIterator{contract: _LogEmitter.contract, event: "LogThreeTopics", logs: logs, sub: sub}, nil
}

// WatchLogThreeTopics is a free log subscription operation binding the contract event 0x70a85a3e24dfc01a8d3f33c838120eb893c800ece385461930e9f3d0cfc20e42.
//
// Solidity: event LogThreeTopics(uint256 indexed number, address indexed addy)
func (_LogEmitter *LogEmitterFilterer) WatchLogThreeTopics(opts *bind.WatchOpts, sink chan<- *LogEmitterLogThreeTopics, number []*big.Int, addy []common.Address) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogThreeTopics", numberRule, addyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogThreeTopics)
				if err := _LogEmitter.contract.UnpackLog(event, "LogThreeTopics", log); err != nil {
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

// ParseLogThreeTopics is a log parse operation binding the contract event 0x70a85a3e24dfc01a8d3f33c838120eb893c800ece385461930e9f3d0cfc20e42.
//
// Solidity: event LogThreeTopics(uint256 indexed number, address indexed addy)
func (_LogEmitter *LogEmitterFilterer) ParseLogThreeTopics(log types.Log) (*LogEmitterLogThreeTopics, error) {
	event := new(LogEmitterLogThreeTopics)
	if err := _LogEmitter.contract.UnpackLog(event, "LogThreeTopics", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogThreeTopicsExtraDataIterator is returned from FilterLogThreeTopicsExtraData and is used to iterate over the raw logs and unpacked data for LogThreeTopicsExtraData events raised by the LogEmitter contract.
type LogEmitterLogThreeTopicsExtraDataIterator struct {
	Event *LogEmitterLogThreeTopicsExtraData // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogThreeTopicsExtraDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogThreeTopicsExtraData)
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
		it.Event = new(LogEmitterLogThreeTopicsExtraData)
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
func (it *LogEmitterLogThreeTopicsExtraDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogThreeTopicsExtraDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogThreeTopicsExtraData represents a LogThreeTopicsExtraData event raised by the LogEmitter contract.
type LogEmitterLogThreeTopicsExtraData struct {
	Number *big.Int
	Addy   common.Address
	Arg2   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogThreeTopicsExtraData is a free log retrieval operation binding the contract event 0x9111be8bfd08a48f95eadcdb6694214a59e1cc21ca0577e1d07856e916d82bef.
//
// Solidity: event LogThreeTopicsExtraData(uint256 indexed number, address indexed addy, bytes32 arg2)
func (_LogEmitter *LogEmitterFilterer) FilterLogThreeTopicsExtraData(opts *bind.FilterOpts, number []*big.Int, addy []common.Address) (*LogEmitterLogThreeTopicsExtraDataIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogThreeTopicsExtraData", numberRule, addyRule)
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogThreeTopicsExtraDataIterator{contract: _LogEmitter.contract, event: "LogThreeTopicsExtraData", logs: logs, sub: sub}, nil
}

// WatchLogThreeTopicsExtraData is a free log subscription operation binding the contract event 0x9111be8bfd08a48f95eadcdb6694214a59e1cc21ca0577e1d07856e916d82bef.
//
// Solidity: event LogThreeTopicsExtraData(uint256 indexed number, address indexed addy, bytes32 arg2)
func (_LogEmitter *LogEmitterFilterer) WatchLogThreeTopicsExtraData(opts *bind.WatchOpts, sink chan<- *LogEmitterLogThreeTopicsExtraData, number []*big.Int, addy []common.Address) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}
	var addyRule []interface{}
	for _, addyItem := range addy {
		addyRule = append(addyRule, addyItem)
	}

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogThreeTopicsExtraData", numberRule, addyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogThreeTopicsExtraData)
				if err := _LogEmitter.contract.UnpackLog(event, "LogThreeTopicsExtraData", log); err != nil {
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

// ParseLogThreeTopicsExtraData is a log parse operation binding the contract event 0x9111be8bfd08a48f95eadcdb6694214a59e1cc21ca0577e1d07856e916d82bef.
//
// Solidity: event LogThreeTopicsExtraData(uint256 indexed number, address indexed addy, bytes32 arg2)
func (_LogEmitter *LogEmitterFilterer) ParseLogThreeTopicsExtraData(log types.Log) (*LogEmitterLogThreeTopicsExtraData, error) {
	event := new(LogEmitterLogThreeTopicsExtraData)
	if err := _LogEmitter.contract.UnpackLog(event, "LogThreeTopicsExtraData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogTwoTopicsIterator is returned from FilterLogTwoTopics and is used to iterate over the raw logs and unpacked data for LogTwoTopics events raised by the LogEmitter contract.
type LogEmitterLogTwoTopicsIterator struct {
	Event *LogEmitterLogTwoTopics // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogTwoTopicsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogTwoTopics)
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
		it.Event = new(LogEmitterLogTwoTopics)
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
func (it *LogEmitterLogTwoTopicsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogTwoTopicsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogTwoTopics represents a LogTwoTopics event raised by the LogEmitter contract.
type LogEmitterLogTwoTopics struct {
	Number *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTwoTopics is a free log retrieval operation binding the contract event 0x6370fe7ee1b1199c3c4673346f536f2a2a55f2a717f6c68b528a14757050aab8.
//
// Solidity: event LogTwoTopics(uint256 indexed number)
func (_LogEmitter *LogEmitterFilterer) FilterLogTwoTopics(opts *bind.FilterOpts, number []*big.Int) (*LogEmitterLogTwoTopicsIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogTwoTopics", numberRule)
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogTwoTopicsIterator{contract: _LogEmitter.contract, event: "LogTwoTopics", logs: logs, sub: sub}, nil
}

// WatchLogTwoTopics is a free log subscription operation binding the contract event 0x6370fe7ee1b1199c3c4673346f536f2a2a55f2a717f6c68b528a14757050aab8.
//
// Solidity: event LogTwoTopics(uint256 indexed number)
func (_LogEmitter *LogEmitterFilterer) WatchLogTwoTopics(opts *bind.WatchOpts, sink chan<- *LogEmitterLogTwoTopics, number []*big.Int) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogTwoTopics", numberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogTwoTopics)
				if err := _LogEmitter.contract.UnpackLog(event, "LogTwoTopics", log); err != nil {
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

// ParseLogTwoTopics is a log parse operation binding the contract event 0x6370fe7ee1b1199c3c4673346f536f2a2a55f2a717f6c68b528a14757050aab8.
//
// Solidity: event LogTwoTopics(uint256 indexed number)
func (_LogEmitter *LogEmitterFilterer) ParseLogTwoTopics(log types.Log) (*LogEmitterLogTwoTopics, error) {
	event := new(LogEmitterLogTwoTopics)
	if err := _LogEmitter.contract.UnpackLog(event, "LogTwoTopics", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LogEmitterLogTwoTopicsExtraDataIterator is returned from FilterLogTwoTopicsExtraData and is used to iterate over the raw logs and unpacked data for LogTwoTopicsExtraData events raised by the LogEmitter contract.
type LogEmitterLogTwoTopicsExtraDataIterator struct {
	Event *LogEmitterLogTwoTopicsExtraData // Event containing the contract specifics and raw log

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
func (it *LogEmitterLogTwoTopicsExtraDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogEmitterLogTwoTopicsExtraData)
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
		it.Event = new(LogEmitterLogTwoTopicsExtraData)
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
func (it *LogEmitterLogTwoTopicsExtraDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogEmitterLogTwoTopicsExtraDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogEmitterLogTwoTopicsExtraData represents a LogTwoTopicsExtraData event raised by the LogEmitter contract.
type LogEmitterLogTwoTopicsExtraData struct {
	Number *big.Int
	Arg1   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTwoTopicsExtraData is a free log retrieval operation binding the contract event 0x0d95b68954ecebf6f87262a05e0f5f134d5942feba6e7d0c96a647ec0d0a9225.
//
// Solidity: event LogTwoTopicsExtraData(uint256 indexed number, address arg1)
func (_LogEmitter *LogEmitterFilterer) FilterLogTwoTopicsExtraData(opts *bind.FilterOpts, number []*big.Int) (*LogEmitterLogTwoTopicsExtraDataIterator, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _LogEmitter.contract.FilterLogs(opts, "LogTwoTopicsExtraData", numberRule)
	if err != nil {
		return nil, err
	}
	return &LogEmitterLogTwoTopicsExtraDataIterator{contract: _LogEmitter.contract, event: "LogTwoTopicsExtraData", logs: logs, sub: sub}, nil
}

// WatchLogTwoTopicsExtraData is a free log subscription operation binding the contract event 0x0d95b68954ecebf6f87262a05e0f5f134d5942feba6e7d0c96a647ec0d0a9225.
//
// Solidity: event LogTwoTopicsExtraData(uint256 indexed number, address arg1)
func (_LogEmitter *LogEmitterFilterer) WatchLogTwoTopicsExtraData(opts *bind.WatchOpts, sink chan<- *LogEmitterLogTwoTopicsExtraData, number []*big.Int) (event.Subscription, error) {

	var numberRule []interface{}
	for _, numberItem := range number {
		numberRule = append(numberRule, numberItem)
	}

	logs, sub, err := _LogEmitter.contract.WatchLogs(opts, "LogTwoTopicsExtraData", numberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogEmitterLogTwoTopicsExtraData)
				if err := _LogEmitter.contract.UnpackLog(event, "LogTwoTopicsExtraData", log); err != nil {
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

// ParseLogTwoTopicsExtraData is a log parse operation binding the contract event 0x0d95b68954ecebf6f87262a05e0f5f134d5942feba6e7d0c96a647ec0d0a9225.
//
// Solidity: event LogTwoTopicsExtraData(uint256 indexed number, address arg1)
func (_LogEmitter *LogEmitterFilterer) ParseLogTwoTopicsExtraData(log types.Log) (*LogEmitterLogTwoTopicsExtraData, error) {
	event := new(LogEmitterLogTwoTopicsExtraData)
	if err := _LogEmitter.contract.UnpackLog(event, "LogTwoTopicsExtraData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NestedCallMetaData contains all meta data concerning the NestedCall contract.
var NestedCallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"callCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"entrypoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"intermediateCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nestedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105248061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610065575f3560e01c80632eb8cdd81161004e5780632eb8cdd8146100a15780634b28f9a2146100bf578063f647dfeb146100dd57610065565b80630b2db306146100695780630ddad78a14610085575b5f5ffd5b610083600480360381019061007e9190610316565b6100fb565b005b61009f600480360381019061009a9190610316565b61021a565b005b6100a96102a7565b6040516100b69190610359565b60405180910390f35b6100c76102ac565b6040516100d49190610359565b60405180910390f35b6100e56102b2565b6040516100f29190610359565b60405180910390f35b63a4b1cafe5f819055505f8173ffffffffffffffffffffffffffffffffffffffff1663d5378f6b60e01b604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161019591906103c4565b5f60405180830381855af49150503d805f81146101cd576040519150601f19603f3d011682016040523d82523d5f602084013e6101d2565b606091505b5050905080610216576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020d90610434565b60405180910390fd5b5050565b660ffc13a1171ab55f8190555060025f8154809291906102399061047f565b91905055503073ffffffffffffffffffffffffffffffffffffffff16630b2db306826040518263ffffffff1660e01b815260040161027791906104d5565b5f604051808303815f87803b15801561028e575f5ffd5b505af11580156102a0573d5f5f3e3d5ffd5b5050505050565b5f5481565b60025481565b60015481565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102e5826102bc565b9050919050565b6102f5816102db565b81146102ff575f5ffd5b50565b5f81359050610310816102ec565b92915050565b5f6020828403121561032b5761032a6102b8565b5b5f61033884828501610302565b91505092915050565b5f819050919050565b61035381610341565b82525050565b5f60208201905061036c5f83018461034a565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61039e82610372565b6103a8818561037c565b93506103b8818560208601610386565b80840191505092915050565b5f6103cf8284610394565b915081905092915050565b5f82825260208201905092915050565b7f44656c656761746563616c6c206661696c6564000000000000000000000000005f82015250565b5f61041e6013836103da565b9150610429826103ea565b602082019050919050565b5f6020820190508181035f83015261044b81610412565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61048982610341565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036104bb576104ba610452565b5b600182019050919050565b6104cf816102db565b82525050565b5f6020820190506104e85f8301846104c6565b9291505056fea26469706673582212208aa9b1ea3d3778e1a2a3a24bf6bef9ca90dc69bf23a5e9a0e1b325a8de72522f64736f6c634300081e0033",
}

// NestedCallABI is the input ABI used to generate the binding from.
// Deprecated: Use NestedCallMetaData.ABI instead.
var NestedCallABI = NestedCallMetaData.ABI

// NestedCallBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NestedCallMetaData.Bin instead.
var NestedCallBin = NestedCallMetaData.Bin

// DeployNestedCall deploys a new Ethereum contract, binding an instance of NestedCall to it.
func DeployNestedCall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NestedCall, error) {
	parsed, err := NestedCallMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NestedCallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NestedCall{NestedCallCaller: NestedCallCaller{contract: contract}, NestedCallTransactor: NestedCallTransactor{contract: contract}, NestedCallFilterer: NestedCallFilterer{contract: contract}}, nil
}

// NestedCall is an auto generated Go binding around an Ethereum contract.
type NestedCall struct {
	NestedCallCaller     // Read-only binding to the contract
	NestedCallTransactor // Write-only binding to the contract
	NestedCallFilterer   // Log filterer for contract events
}

// NestedCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type NestedCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NestedCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NestedCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NestedCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NestedCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NestedCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NestedCallSession struct {
	Contract     *NestedCall       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NestedCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NestedCallCallerSession struct {
	Contract *NestedCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NestedCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NestedCallTransactorSession struct {
	Contract     *NestedCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NestedCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type NestedCallRaw struct {
	Contract *NestedCall // Generic contract binding to access the raw methods on
}

// NestedCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NestedCallCallerRaw struct {
	Contract *NestedCallCaller // Generic read-only contract binding to access the raw methods on
}

// NestedCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NestedCallTransactorRaw struct {
	Contract *NestedCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNestedCall creates a new instance of NestedCall, bound to a specific deployed contract.
func NewNestedCall(address common.Address, backend bind.ContractBackend) (*NestedCall, error) {
	contract, err := bindNestedCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NestedCall{NestedCallCaller: NestedCallCaller{contract: contract}, NestedCallTransactor: NestedCallTransactor{contract: contract}, NestedCallFilterer: NestedCallFilterer{contract: contract}}, nil
}

// NewNestedCallCaller creates a new read-only instance of NestedCall, bound to a specific deployed contract.
func NewNestedCallCaller(address common.Address, caller bind.ContractCaller) (*NestedCallCaller, error) {
	contract, err := bindNestedCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NestedCallCaller{contract: contract}, nil
}

// NewNestedCallTransactor creates a new write-only instance of NestedCall, bound to a specific deployed contract.
func NewNestedCallTransactor(address common.Address, transactor bind.ContractTransactor) (*NestedCallTransactor, error) {
	contract, err := bindNestedCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NestedCallTransactor{contract: contract}, nil
}

// NewNestedCallFilterer creates a new log filterer instance of NestedCall, bound to a specific deployed contract.
func NewNestedCallFilterer(address common.Address, filterer bind.ContractFilterer) (*NestedCallFilterer, error) {
	contract, err := bindNestedCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NestedCallFilterer{contract: contract}, nil
}

// bindNestedCall binds a generic wrapper to an already deployed contract.
func bindNestedCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NestedCallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NestedCall *NestedCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NestedCall.Contract.NestedCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NestedCall *NestedCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NestedCall.Contract.NestedCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NestedCall *NestedCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NestedCall.Contract.NestedCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NestedCall *NestedCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NestedCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NestedCall *NestedCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NestedCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NestedCall *NestedCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NestedCall.Contract.contract.Transact(opts, method, params...)
}

// CallCount is a free data retrieval call binding the contract method 0x4b28f9a2.
//
// Solidity: function callCount() view returns(uint256)
func (_NestedCall *NestedCallCaller) CallCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NestedCall.contract.Call(opts, &out, "callCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallCount is a free data retrieval call binding the contract method 0x4b28f9a2.
//
// Solidity: function callCount() view returns(uint256)
func (_NestedCall *NestedCallSession) CallCount() (*big.Int, error) {
	return _NestedCall.Contract.CallCount(&_NestedCall.CallOpts)
}

// CallCount is a free data retrieval call binding the contract method 0x4b28f9a2.
//
// Solidity: function callCount() view returns(uint256)
func (_NestedCall *NestedCallCallerSession) CallCount() (*big.Int, error) {
	return _NestedCall.Contract.CallCount(&_NestedCall.CallOpts)
}

// MainValue is a free data retrieval call binding the contract method 0x2eb8cdd8.
//
// Solidity: function mainValue() view returns(uint256)
func (_NestedCall *NestedCallCaller) MainValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NestedCall.contract.Call(opts, &out, "mainValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainValue is a free data retrieval call binding the contract method 0x2eb8cdd8.
//
// Solidity: function mainValue() view returns(uint256)
func (_NestedCall *NestedCallSession) MainValue() (*big.Int, error) {
	return _NestedCall.Contract.MainValue(&_NestedCall.CallOpts)
}

// MainValue is a free data retrieval call binding the contract method 0x2eb8cdd8.
//
// Solidity: function mainValue() view returns(uint256)
func (_NestedCall *NestedCallCallerSession) MainValue() (*big.Int, error) {
	return _NestedCall.Contract.MainValue(&_NestedCall.CallOpts)
}

// NestedValue is a free data retrieval call binding the contract method 0xf647dfeb.
//
// Solidity: function nestedValue() view returns(uint256)
func (_NestedCall *NestedCallCaller) NestedValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NestedCall.contract.Call(opts, &out, "nestedValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NestedValue is a free data retrieval call binding the contract method 0xf647dfeb.
//
// Solidity: function nestedValue() view returns(uint256)
func (_NestedCall *NestedCallSession) NestedValue() (*big.Int, error) {
	return _NestedCall.Contract.NestedValue(&_NestedCall.CallOpts)
}

// NestedValue is a free data retrieval call binding the contract method 0xf647dfeb.
//
// Solidity: function nestedValue() view returns(uint256)
func (_NestedCall *NestedCallCallerSession) NestedValue() (*big.Int, error) {
	return _NestedCall.Contract.NestedValue(&_NestedCall.CallOpts)
}

// Entrypoint is a paid mutator transaction binding the contract method 0x0ddad78a.
//
// Solidity: function entrypoint(address target) returns()
func (_NestedCall *NestedCallTransactor) Entrypoint(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _NestedCall.contract.Transact(opts, "entrypoint", target)
}

// Entrypoint is a paid mutator transaction binding the contract method 0x0ddad78a.
//
// Solidity: function entrypoint(address target) returns()
func (_NestedCall *NestedCallSession) Entrypoint(target common.Address) (*types.Transaction, error) {
	return _NestedCall.Contract.Entrypoint(&_NestedCall.TransactOpts, target)
}

// Entrypoint is a paid mutator transaction binding the contract method 0x0ddad78a.
//
// Solidity: function entrypoint(address target) returns()
func (_NestedCall *NestedCallTransactorSession) Entrypoint(target common.Address) (*types.Transaction, error) {
	return _NestedCall.Contract.Entrypoint(&_NestedCall.TransactOpts, target)
}

// IntermediateCall is a paid mutator transaction binding the contract method 0x0b2db306.
//
// Solidity: function intermediateCall(address target) returns()
func (_NestedCall *NestedCallTransactor) IntermediateCall(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _NestedCall.contract.Transact(opts, "intermediateCall", target)
}

// IntermediateCall is a paid mutator transaction binding the contract method 0x0b2db306.
//
// Solidity: function intermediateCall(address target) returns()
func (_NestedCall *NestedCallSession) IntermediateCall(target common.Address) (*types.Transaction, error) {
	return _NestedCall.Contract.IntermediateCall(&_NestedCall.TransactOpts, target)
}

// IntermediateCall is a paid mutator transaction binding the contract method 0x0b2db306.
//
// Solidity: function intermediateCall(address target) returns()
func (_NestedCall *NestedCallTransactorSession) IntermediateCall(target common.Address) (*types.Transaction, error) {
	return _NestedCall.Contract.IntermediateCall(&_NestedCall.TransactOpts, target)
}

// NestedTargetMetaData contains all meta data concerning the NestedTarget contract.
var NestedTargetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"performNestedAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50607080601a5f395ff3fe6080604052348015600e575f5ffd5b50600436106026575f3560e01c8063d5378f6b14602a575b5f5ffd5b60306032565b005b61a4b160015556fea2646970667358221220c42b7bb550dd61a1a85e8f063d166d455dfa2b03fdc34388f85570221c3f209164736f6c634300081e0033",
}

// NestedTargetABI is the input ABI used to generate the binding from.
// Deprecated: Use NestedTargetMetaData.ABI instead.
var NestedTargetABI = NestedTargetMetaData.ABI

// NestedTargetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NestedTargetMetaData.Bin instead.
var NestedTargetBin = NestedTargetMetaData.Bin

// DeployNestedTarget deploys a new Ethereum contract, binding an instance of NestedTarget to it.
func DeployNestedTarget(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NestedTarget, error) {
	parsed, err := NestedTargetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NestedTargetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NestedTarget{NestedTargetCaller: NestedTargetCaller{contract: contract}, NestedTargetTransactor: NestedTargetTransactor{contract: contract}, NestedTargetFilterer: NestedTargetFilterer{contract: contract}}, nil
}

// NestedTarget is an auto generated Go binding around an Ethereum contract.
type NestedTarget struct {
	NestedTargetCaller     // Read-only binding to the contract
	NestedTargetTransactor // Write-only binding to the contract
	NestedTargetFilterer   // Log filterer for contract events
}

// NestedTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type NestedTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NestedTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NestedTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NestedTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NestedTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NestedTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NestedTargetSession struct {
	Contract     *NestedTarget     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NestedTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NestedTargetCallerSession struct {
	Contract *NestedTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NestedTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NestedTargetTransactorSession struct {
	Contract     *NestedTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NestedTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type NestedTargetRaw struct {
	Contract *NestedTarget // Generic contract binding to access the raw methods on
}

// NestedTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NestedTargetCallerRaw struct {
	Contract *NestedTargetCaller // Generic read-only contract binding to access the raw methods on
}

// NestedTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NestedTargetTransactorRaw struct {
	Contract *NestedTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNestedTarget creates a new instance of NestedTarget, bound to a specific deployed contract.
func NewNestedTarget(address common.Address, backend bind.ContractBackend) (*NestedTarget, error) {
	contract, err := bindNestedTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NestedTarget{NestedTargetCaller: NestedTargetCaller{contract: contract}, NestedTargetTransactor: NestedTargetTransactor{contract: contract}, NestedTargetFilterer: NestedTargetFilterer{contract: contract}}, nil
}

// NewNestedTargetCaller creates a new read-only instance of NestedTarget, bound to a specific deployed contract.
func NewNestedTargetCaller(address common.Address, caller bind.ContractCaller) (*NestedTargetCaller, error) {
	contract, err := bindNestedTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NestedTargetCaller{contract: contract}, nil
}

// NewNestedTargetTransactor creates a new write-only instance of NestedTarget, bound to a specific deployed contract.
func NewNestedTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*NestedTargetTransactor, error) {
	contract, err := bindNestedTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NestedTargetTransactor{contract: contract}, nil
}

// NewNestedTargetFilterer creates a new log filterer instance of NestedTarget, bound to a specific deployed contract.
func NewNestedTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*NestedTargetFilterer, error) {
	contract, err := bindNestedTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NestedTargetFilterer{contract: contract}, nil
}

// bindNestedTarget binds a generic wrapper to an already deployed contract.
func bindNestedTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NestedTargetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NestedTarget *NestedTargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NestedTarget.Contract.NestedTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NestedTarget *NestedTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NestedTarget.Contract.NestedTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NestedTarget *NestedTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NestedTarget.Contract.NestedTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NestedTarget *NestedTargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NestedTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NestedTarget *NestedTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NestedTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NestedTarget *NestedTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NestedTarget.Contract.contract.Transact(opts, method, params...)
}

// PerformNestedAction is a paid mutator transaction binding the contract method 0xd5378f6b.
//
// Solidity: function performNestedAction() returns()
func (_NestedTarget *NestedTargetTransactor) PerformNestedAction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NestedTarget.contract.Transact(opts, "performNestedAction")
}

// PerformNestedAction is a paid mutator transaction binding the contract method 0xd5378f6b.
//
// Solidity: function performNestedAction() returns()
func (_NestedTarget *NestedTargetSession) PerformNestedAction() (*types.Transaction, error) {
	return _NestedTarget.Contract.PerformNestedAction(&_NestedTarget.TransactOpts)
}

// PerformNestedAction is a paid mutator transaction binding the contract method 0xd5378f6b.
//
// Solidity: function performNestedAction() returns()
func (_NestedTarget *NestedTargetTransactorSession) PerformNestedAction() (*types.Transaction, error) {
	return _NestedTarget.Contract.PerformNestedAction(&_NestedTarget.TransactOpts)
}

// OutOfGasMetaData contains all meta data concerning the OutOfGas contract.
var OutOfGasMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"callOutOfGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasErrorOccurred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outOfGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061029a8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80630c55699c1461004e57806331fe52e81461006c5780637dc376c1146100765780637e5d21cb14610080575b5f5ffd5b61005661009e565b60405161006391906101a4565b60405180910390f35b6100746100a3565b005b61007e6100c8565b005b61008861017a565b60405161009591906101d7565b60405180910390f35b5f5481565b5b6001156100c6575f5f8154809291906100bc9061021d565b91905055506100a4565b565b5f60015f6101000a81548160ff0219169083151502179055503073ffffffffffffffffffffffffffffffffffffffff166331fe52e8620186a06040518263ffffffff1660e01b81526004015f604051808303815f88803b15801561012a575f5ffd5b5087f19350505050801561013c575060015b61015e576001805f6101000a81548160ff021916908315150217905550610178565b5f60015f6101000a81548160ff0219169083151502179055505b565b60015f9054906101000a900460ff1681565b5f819050919050565b61019e8161018c565b82525050565b5f6020820190506101b75f830184610195565b92915050565b5f8115159050919050565b6101d1816101bd565b82525050565b5f6020820190506101ea5f8301846101c8565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6102278261018c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610259576102586101f0565b5b60018201905091905056fea264697066735822122029bb4f324f43ecd275c403ec33e0efedf96c2bbabcf1a7035eab60031357767164736f6c634300081e0033",
}

// OutOfGasABI is the input ABI used to generate the binding from.
// Deprecated: Use OutOfGasMetaData.ABI instead.
var OutOfGasABI = OutOfGasMetaData.ABI

// OutOfGasBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutOfGasMetaData.Bin instead.
var OutOfGasBin = OutOfGasMetaData.Bin

// DeployOutOfGas deploys a new Ethereum contract, binding an instance of OutOfGas to it.
func DeployOutOfGas(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutOfGas, error) {
	parsed, err := OutOfGasMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutOfGasBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutOfGas{OutOfGasCaller: OutOfGasCaller{contract: contract}, OutOfGasTransactor: OutOfGasTransactor{contract: contract}, OutOfGasFilterer: OutOfGasFilterer{contract: contract}}, nil
}

// OutOfGas is an auto generated Go binding around an Ethereum contract.
type OutOfGas struct {
	OutOfGasCaller     // Read-only binding to the contract
	OutOfGasTransactor // Write-only binding to the contract
	OutOfGasFilterer   // Log filterer for contract events
}

// OutOfGasCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutOfGasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutOfGasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutOfGasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutOfGasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutOfGasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutOfGasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutOfGasSession struct {
	Contract     *OutOfGas         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutOfGasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutOfGasCallerSession struct {
	Contract *OutOfGasCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OutOfGasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutOfGasTransactorSession struct {
	Contract     *OutOfGasTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OutOfGasRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutOfGasRaw struct {
	Contract *OutOfGas // Generic contract binding to access the raw methods on
}

// OutOfGasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutOfGasCallerRaw struct {
	Contract *OutOfGasCaller // Generic read-only contract binding to access the raw methods on
}

// OutOfGasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutOfGasTransactorRaw struct {
	Contract *OutOfGasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutOfGas creates a new instance of OutOfGas, bound to a specific deployed contract.
func NewOutOfGas(address common.Address, backend bind.ContractBackend) (*OutOfGas, error) {
	contract, err := bindOutOfGas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutOfGas{OutOfGasCaller: OutOfGasCaller{contract: contract}, OutOfGasTransactor: OutOfGasTransactor{contract: contract}, OutOfGasFilterer: OutOfGasFilterer{contract: contract}}, nil
}

// NewOutOfGasCaller creates a new read-only instance of OutOfGas, bound to a specific deployed contract.
func NewOutOfGasCaller(address common.Address, caller bind.ContractCaller) (*OutOfGasCaller, error) {
	contract, err := bindOutOfGas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutOfGasCaller{contract: contract}, nil
}

// NewOutOfGasTransactor creates a new write-only instance of OutOfGas, bound to a specific deployed contract.
func NewOutOfGasTransactor(address common.Address, transactor bind.ContractTransactor) (*OutOfGasTransactor, error) {
	contract, err := bindOutOfGas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutOfGasTransactor{contract: contract}, nil
}

// NewOutOfGasFilterer creates a new log filterer instance of OutOfGas, bound to a specific deployed contract.
func NewOutOfGasFilterer(address common.Address, filterer bind.ContractFilterer) (*OutOfGasFilterer, error) {
	contract, err := bindOutOfGas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutOfGasFilterer{contract: contract}, nil
}

// bindOutOfGas binds a generic wrapper to an already deployed contract.
func bindOutOfGas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OutOfGasMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutOfGas *OutOfGasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutOfGas.Contract.OutOfGasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutOfGas *OutOfGasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutOfGas.Contract.OutOfGasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutOfGas *OutOfGasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutOfGas.Contract.OutOfGasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutOfGas *OutOfGasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutOfGas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutOfGas *OutOfGasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutOfGas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutOfGas *OutOfGasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutOfGas.Contract.contract.Transact(opts, method, params...)
}

// GasErrorOccurred is a free data retrieval call binding the contract method 0x7e5d21cb.
//
// Solidity: function gasErrorOccurred() view returns(bool)
func (_OutOfGas *OutOfGasCaller) GasErrorOccurred(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OutOfGas.contract.Call(opts, &out, "gasErrorOccurred")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GasErrorOccurred is a free data retrieval call binding the contract method 0x7e5d21cb.
//
// Solidity: function gasErrorOccurred() view returns(bool)
func (_OutOfGas *OutOfGasSession) GasErrorOccurred() (bool, error) {
	return _OutOfGas.Contract.GasErrorOccurred(&_OutOfGas.CallOpts)
}

// GasErrorOccurred is a free data retrieval call binding the contract method 0x7e5d21cb.
//
// Solidity: function gasErrorOccurred() view returns(bool)
func (_OutOfGas *OutOfGasCallerSession) GasErrorOccurred() (bool, error) {
	return _OutOfGas.Contract.GasErrorOccurred(&_OutOfGas.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_OutOfGas *OutOfGasCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutOfGas.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_OutOfGas *OutOfGasSession) X() (*big.Int, error) {
	return _OutOfGas.Contract.X(&_OutOfGas.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_OutOfGas *OutOfGasCallerSession) X() (*big.Int, error) {
	return _OutOfGas.Contract.X(&_OutOfGas.CallOpts)
}

// CallOutOfGas is a paid mutator transaction binding the contract method 0x7dc376c1.
//
// Solidity: function callOutOfGas() returns()
func (_OutOfGas *OutOfGasTransactor) CallOutOfGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutOfGas.contract.Transact(opts, "callOutOfGas")
}

// CallOutOfGas is a paid mutator transaction binding the contract method 0x7dc376c1.
//
// Solidity: function callOutOfGas() returns()
func (_OutOfGas *OutOfGasSession) CallOutOfGas() (*types.Transaction, error) {
	return _OutOfGas.Contract.CallOutOfGas(&_OutOfGas.TransactOpts)
}

// CallOutOfGas is a paid mutator transaction binding the contract method 0x7dc376c1.
//
// Solidity: function callOutOfGas() returns()
func (_OutOfGas *OutOfGasTransactorSession) CallOutOfGas() (*types.Transaction, error) {
	return _OutOfGas.Contract.CallOutOfGas(&_OutOfGas.TransactOpts)
}

// OutOfGas is a paid mutator transaction binding the contract method 0x31fe52e8.
//
// Solidity: function outOfGas() returns()
func (_OutOfGas *OutOfGasTransactor) OutOfGas(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutOfGas.contract.Transact(opts, "outOfGas")
}

// OutOfGas is a paid mutator transaction binding the contract method 0x31fe52e8.
//
// Solidity: function outOfGas() returns()
func (_OutOfGas *OutOfGasSession) OutOfGas() (*types.Transaction, error) {
	return _OutOfGas.Contract.OutOfGas(&_OutOfGas.TransactOpts)
}

// OutOfGas is a paid mutator transaction binding the contract method 0x31fe52e8.
//
// Solidity: function outOfGas() returns()
func (_OutOfGas *OutOfGasTransactorSession) OutOfGas() (*types.Transaction, error) {
	return _OutOfGas.Contract.OutOfGas(&_OutOfGas.TransactOpts)
}

// PayableCounterMetaData contains all meta data concerning the PayableCounter contract.
var PayableCounterMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101338061001c5f395ff3fe6080604052600436106028575f3560e01c80633fb5c1cb14602b5780638381f58a14604e576029565b5b005b3480156035575f5ffd5b50604c60048036038101906048919060b3565b6073565b005b3480156058575f5ffd5b50605f607c565b604051606a919060e6565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6095816085565b8114609e575f5ffd5b50565b5f8135905060ad81608e565b92915050565b5f6020828403121560c55760c46081565b5b5f60d08482850160a1565b91505092915050565b60e0816085565b82525050565b5f60208201905060f75f83018460d9565b9291505056fea26469706673582212207be8e90a394087dc6904b55c5d169902279ec7597786236ee4413868cc3dd25e64736f6c634300081e0033",
}

// PayableCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use PayableCounterMetaData.ABI instead.
var PayableCounterABI = PayableCounterMetaData.ABI

// PayableCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PayableCounterMetaData.Bin instead.
var PayableCounterBin = PayableCounterMetaData.Bin

// DeployPayableCounter deploys a new Ethereum contract, binding an instance of PayableCounter to it.
func DeployPayableCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PayableCounter, error) {
	parsed, err := PayableCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PayableCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PayableCounter{PayableCounterCaller: PayableCounterCaller{contract: contract}, PayableCounterTransactor: PayableCounterTransactor{contract: contract}, PayableCounterFilterer: PayableCounterFilterer{contract: contract}}, nil
}

// PayableCounter is an auto generated Go binding around an Ethereum contract.
type PayableCounter struct {
	PayableCounterCaller     // Read-only binding to the contract
	PayableCounterTransactor // Write-only binding to the contract
	PayableCounterFilterer   // Log filterer for contract events
}

// PayableCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayableCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayableCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayableCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayableCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayableCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayableCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayableCounterSession struct {
	Contract     *PayableCounter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayableCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayableCounterCallerSession struct {
	Contract *PayableCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PayableCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayableCounterTransactorSession struct {
	Contract     *PayableCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PayableCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayableCounterRaw struct {
	Contract *PayableCounter // Generic contract binding to access the raw methods on
}

// PayableCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayableCounterCallerRaw struct {
	Contract *PayableCounterCaller // Generic read-only contract binding to access the raw methods on
}

// PayableCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayableCounterTransactorRaw struct {
	Contract *PayableCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayableCounter creates a new instance of PayableCounter, bound to a specific deployed contract.
func NewPayableCounter(address common.Address, backend bind.ContractBackend) (*PayableCounter, error) {
	contract, err := bindPayableCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayableCounter{PayableCounterCaller: PayableCounterCaller{contract: contract}, PayableCounterTransactor: PayableCounterTransactor{contract: contract}, PayableCounterFilterer: PayableCounterFilterer{contract: contract}}, nil
}

// NewPayableCounterCaller creates a new read-only instance of PayableCounter, bound to a specific deployed contract.
func NewPayableCounterCaller(address common.Address, caller bind.ContractCaller) (*PayableCounterCaller, error) {
	contract, err := bindPayableCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayableCounterCaller{contract: contract}, nil
}

// NewPayableCounterTransactor creates a new write-only instance of PayableCounter, bound to a specific deployed contract.
func NewPayableCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*PayableCounterTransactor, error) {
	contract, err := bindPayableCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayableCounterTransactor{contract: contract}, nil
}

// NewPayableCounterFilterer creates a new log filterer instance of PayableCounter, bound to a specific deployed contract.
func NewPayableCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*PayableCounterFilterer, error) {
	contract, err := bindPayableCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayableCounterFilterer{contract: contract}, nil
}

// bindPayableCounter binds a generic wrapper to an already deployed contract.
func bindPayableCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PayableCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayableCounter *PayableCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayableCounter.Contract.PayableCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayableCounter *PayableCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayableCounter.Contract.PayableCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayableCounter *PayableCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayableCounter.Contract.PayableCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayableCounter *PayableCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayableCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayableCounter *PayableCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayableCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayableCounter *PayableCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayableCounter.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_PayableCounter *PayableCounterCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PayableCounter.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_PayableCounter *PayableCounterSession) Number() (*big.Int, error) {
	return _PayableCounter.Contract.Number(&_PayableCounter.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint256)
func (_PayableCounter *PayableCounterCallerSession) Number() (*big.Int, error) {
	return _PayableCounter.Contract.Number(&_PayableCounter.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_PayableCounter *PayableCounterTransactor) SetNumber(opts *bind.TransactOpts, newNumber *big.Int) (*types.Transaction, error) {
	return _PayableCounter.contract.Transact(opts, "setNumber", newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_PayableCounter *PayableCounterSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _PayableCounter.Contract.SetNumber(&_PayableCounter.TransactOpts, newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(uint256 newNumber) returns()
func (_PayableCounter *PayableCounterTransactorSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _PayableCounter.Contract.SetNumber(&_PayableCounter.TransactOpts, newNumber)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PayableCounter *PayableCounterTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PayableCounter.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PayableCounter *PayableCounterSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PayableCounter.Contract.Fallback(&_PayableCounter.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PayableCounter *PayableCounterTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PayableCounter.Contract.Fallback(&_PayableCounter.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PayableCounter *PayableCounterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayableCounter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PayableCounter *PayableCounterSession) Receive() (*types.Transaction, error) {
	return _PayableCounter.Contract.Receive(&_PayableCounter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PayableCounter *PayableCounterTransactorSession) Receive() (*types.Transaction, error) {
	return _PayableCounter.Contract.Receive(&_PayableCounter.TransactOpts)
}

// PrecompileMetaData contains all meta data concerning the Precompile contract.
var PrecompileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"testActivateProgram\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testArbSysArbBlockNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102ed8061001c5f395ff3fe608060405260043610610033575f3560e01c80632e52d6061461003757806393cefdc414610061578063993d794214610077575b5f5ffd5b348015610042575f5ffd5b5061004b610093565b6040516100589190610198565b60405180910390f35b34801561006c575f5ffd5b50610075610098565b005b610091600480360381019061008c919061020f565b61010d565b005b5f5481565b606473ffffffffffffffffffffffffffffffffffffffff1663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100e2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101069190610264565b5f81905550565b60015f81905550607173ffffffffffffffffffffffffffffffffffffffff166358c780c234836040518363ffffffff1660e01b815260040161014f919061029e565b5f604051808303818588803b158015610166575f5ffd5b505af1158015610178573d5f5f3e3d5ffd5b505050505050565b5f819050919050565b61019281610180565b82525050565b5f6020820190506101ab5f830184610189565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101de826101b5565b9050919050565b6101ee816101d4565b81146101f8575f5ffd5b50565b5f81359050610209816101e5565b92915050565b5f60208284031215610224576102236101b1565b5b5f610231848285016101fb565b91505092915050565b61024381610180565b811461024d575f5ffd5b50565b5f8151905061025e8161023a565b92915050565b5f60208284031215610279576102786101b1565b5b5f61028684828501610250565b91505092915050565b610298816101d4565b82525050565b5f6020820190506102b15f83018461028f565b9291505056fea26469706673582212200e3ea15747cdcd6a8f74499310a73f5eb8627dfb01a7e75bf3067bcd8f487ce564736f6c634300081e0033",
}

// PrecompileABI is the input ABI used to generate the binding from.
// Deprecated: Use PrecompileMetaData.ABI instead.
var PrecompileABI = PrecompileMetaData.ABI

// PrecompileBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrecompileMetaData.Bin instead.
var PrecompileBin = PrecompileMetaData.Bin

// DeployPrecompile deploys a new Ethereum contract, binding an instance of Precompile to it.
func DeployPrecompile(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Precompile, error) {
	parsed, err := PrecompileMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrecompileBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Precompile{PrecompileCaller: PrecompileCaller{contract: contract}, PrecompileTransactor: PrecompileTransactor{contract: contract}, PrecompileFilterer: PrecompileFilterer{contract: contract}}, nil
}

// Precompile is an auto generated Go binding around an Ethereum contract.
type Precompile struct {
	PrecompileCaller     // Read-only binding to the contract
	PrecompileTransactor // Write-only binding to the contract
	PrecompileFilterer   // Log filterer for contract events
}

// PrecompileCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompileSession struct {
	Contract     *Precompile       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrecompileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompileCallerSession struct {
	Contract *PrecompileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PrecompileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompileTransactorSession struct {
	Contract     *PrecompileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PrecompileRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompileRaw struct {
	Contract *Precompile // Generic contract binding to access the raw methods on
}

// PrecompileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompileCallerRaw struct {
	Contract *PrecompileCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompileTransactorRaw struct {
	Contract *PrecompileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompile creates a new instance of Precompile, bound to a specific deployed contract.
func NewPrecompile(address common.Address, backend bind.ContractBackend) (*Precompile, error) {
	contract, err := bindPrecompile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Precompile{PrecompileCaller: PrecompileCaller{contract: contract}, PrecompileTransactor: PrecompileTransactor{contract: contract}, PrecompileFilterer: PrecompileFilterer{contract: contract}}, nil
}

// NewPrecompileCaller creates a new read-only instance of Precompile, bound to a specific deployed contract.
func NewPrecompileCaller(address common.Address, caller bind.ContractCaller) (*PrecompileCaller, error) {
	contract, err := bindPrecompile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileCaller{contract: contract}, nil
}

// NewPrecompileTransactor creates a new write-only instance of Precompile, bound to a specific deployed contract.
func NewPrecompileTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompileTransactor, error) {
	contract, err := bindPrecompile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileTransactor{contract: contract}, nil
}

// NewPrecompileFilterer creates a new log filterer instance of Precompile, bound to a specific deployed contract.
func NewPrecompileFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompileFilterer, error) {
	contract, err := bindPrecompile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompileFilterer{contract: contract}, nil
}

// bindPrecompile binds a generic wrapper to an already deployed contract.
func bindPrecompile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrecompileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompile *PrecompileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Precompile.Contract.PrecompileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompile *PrecompileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompile.Contract.PrecompileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompile *PrecompileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompile.Contract.PrecompileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompile *PrecompileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Precompile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompile *PrecompileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompile *PrecompileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompile.Contract.contract.Transact(opts, method, params...)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Precompile *PrecompileCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Precompile *PrecompileSession) N() (*big.Int, error) {
	return _Precompile.Contract.N(&_Precompile.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_Precompile *PrecompileCallerSession) N() (*big.Int, error) {
	return _Precompile.Contract.N(&_Precompile.CallOpts)
}

// TestActivateProgram is a paid mutator transaction binding the contract method 0x993d7942.
//
// Solidity: function testActivateProgram(address program) payable returns()
func (_Precompile *PrecompileTransactor) TestActivateProgram(opts *bind.TransactOpts, program common.Address) (*types.Transaction, error) {
	return _Precompile.contract.Transact(opts, "testActivateProgram", program)
}

// TestActivateProgram is a paid mutator transaction binding the contract method 0x993d7942.
//
// Solidity: function testActivateProgram(address program) payable returns()
func (_Precompile *PrecompileSession) TestActivateProgram(program common.Address) (*types.Transaction, error) {
	return _Precompile.Contract.TestActivateProgram(&_Precompile.TransactOpts, program)
}

// TestActivateProgram is a paid mutator transaction binding the contract method 0x993d7942.
//
// Solidity: function testActivateProgram(address program) payable returns()
func (_Precompile *PrecompileTransactorSession) TestActivateProgram(program common.Address) (*types.Transaction, error) {
	return _Precompile.Contract.TestActivateProgram(&_Precompile.TransactOpts, program)
}

// TestArbSysArbBlockNumber is a paid mutator transaction binding the contract method 0x93cefdc4.
//
// Solidity: function testArbSysArbBlockNumber() returns()
func (_Precompile *PrecompileTransactor) TestArbSysArbBlockNumber(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompile.contract.Transact(opts, "testArbSysArbBlockNumber")
}

// TestArbSysArbBlockNumber is a paid mutator transaction binding the contract method 0x93cefdc4.
//
// Solidity: function testArbSysArbBlockNumber() returns()
func (_Precompile *PrecompileSession) TestArbSysArbBlockNumber() (*types.Transaction, error) {
	return _Precompile.Contract.TestArbSysArbBlockNumber(&_Precompile.TransactOpts)
}

// TestArbSysArbBlockNumber is a paid mutator transaction binding the contract method 0x93cefdc4.
//
// Solidity: function testArbSysArbBlockNumber() returns()
func (_Precompile *PrecompileTransactorSession) TestArbSysArbBlockNumber() (*types.Transaction, error) {
	return _Precompile.Contract.TestArbSysArbBlockNumber(&_Precompile.TransactOpts)
}

// SelfDestructorMetaData contains all meta data concerning the SelfDestructor contract.
var SelfDestructorMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"selfDestruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"warmEmptySelfDestructor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"warmSelfDestructor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061031a8061001c5f395ff3fe608060405260043610610037575f3560e01c80633f5a0bdd1461003a57806388b6388c14610062578063e52f97671461008a57610038565b5b005b348015610045575f5ffd5b50610060600480360381019061005b9190610214565b6100b2565b005b34801561006d575f5ffd5b5061008860048036038101906100839190610214565b6100cb565b005b348015610095575f5ffd5b506100b060048036038101906100ab9190610214565b61013e565b005b8073ffffffffffffffffffffffffffffffffffffffff16ff5b5f8173ffffffffffffffffffffffffffffffffffffffff166040516100ef9061026c565b5f604051808303815f865af19150503d805f8114610128576040519150601f19603f3d011682016040523d82523d5f602084013e61012d565b606091505b5050905061013a826100b2565b5050565b5f8190508073ffffffffffffffffffffffffffffffffffffffff16633fb5c1cb60016040518263ffffffff1660e01b815260040161017c91906102cb565b5f604051808303815f87803b158015610193575f5ffd5b505af11580156101a5573d5f5f3e3d5ffd5b505050506101b2826100b2565b5050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101e3826101ba565b9050919050565b6101f3816101d9565b81146101fd575f5ffd5b50565b5f8135905061020e816101ea565b92915050565b5f60208284031215610229576102286101b6565b5b5f61023684828501610200565b91505092915050565b5f81905092915050565b50565b5f6102575f8361023f565b915061026282610249565b5f82019050919050565b5f6102768261024c565b9150819050919050565b5f819050919050565b5f819050919050565b5f819050919050565b5f6102b56102b06102ab84610280565b610292565b610289565b9050919050565b6102c58161029b565b82525050565b5f6020820190506102de5f8301846102bc565b9291505056fea264697066735822122021b98924cbc0a43f73707b6c190b555f7c2f76409fa98735daef6404192f8bfe64736f6c634300081e0033",
}

// SelfDestructorABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfDestructorMetaData.ABI instead.
var SelfDestructorABI = SelfDestructorMetaData.ABI

// SelfDestructorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SelfDestructorMetaData.Bin instead.
var SelfDestructorBin = SelfDestructorMetaData.Bin

// DeploySelfDestructor deploys a new Ethereum contract, binding an instance of SelfDestructor to it.
func DeploySelfDestructor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SelfDestructor, error) {
	parsed, err := SelfDestructorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SelfDestructorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SelfDestructor{SelfDestructorCaller: SelfDestructorCaller{contract: contract}, SelfDestructorTransactor: SelfDestructorTransactor{contract: contract}, SelfDestructorFilterer: SelfDestructorFilterer{contract: contract}}, nil
}

// SelfDestructor is an auto generated Go binding around an Ethereum contract.
type SelfDestructor struct {
	SelfDestructorCaller     // Read-only binding to the contract
	SelfDestructorTransactor // Write-only binding to the contract
	SelfDestructorFilterer   // Log filterer for contract events
}

// SelfDestructorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelfDestructorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfDestructorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelfDestructorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfDestructorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelfDestructorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfDestructorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelfDestructorSession struct {
	Contract     *SelfDestructor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelfDestructorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelfDestructorCallerSession struct {
	Contract *SelfDestructorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SelfDestructorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelfDestructorTransactorSession struct {
	Contract     *SelfDestructorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SelfDestructorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelfDestructorRaw struct {
	Contract *SelfDestructor // Generic contract binding to access the raw methods on
}

// SelfDestructorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelfDestructorCallerRaw struct {
	Contract *SelfDestructorCaller // Generic read-only contract binding to access the raw methods on
}

// SelfDestructorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelfDestructorTransactorRaw struct {
	Contract *SelfDestructorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelfDestructor creates a new instance of SelfDestructor, bound to a specific deployed contract.
func NewSelfDestructor(address common.Address, backend bind.ContractBackend) (*SelfDestructor, error) {
	contract, err := bindSelfDestructor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfDestructor{SelfDestructorCaller: SelfDestructorCaller{contract: contract}, SelfDestructorTransactor: SelfDestructorTransactor{contract: contract}, SelfDestructorFilterer: SelfDestructorFilterer{contract: contract}}, nil
}

// NewSelfDestructorCaller creates a new read-only instance of SelfDestructor, bound to a specific deployed contract.
func NewSelfDestructorCaller(address common.Address, caller bind.ContractCaller) (*SelfDestructorCaller, error) {
	contract, err := bindSelfDestructor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfDestructorCaller{contract: contract}, nil
}

// NewSelfDestructorTransactor creates a new write-only instance of SelfDestructor, bound to a specific deployed contract.
func NewSelfDestructorTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfDestructorTransactor, error) {
	contract, err := bindSelfDestructor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfDestructorTransactor{contract: contract}, nil
}

// NewSelfDestructorFilterer creates a new log filterer instance of SelfDestructor, bound to a specific deployed contract.
func NewSelfDestructorFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfDestructorFilterer, error) {
	contract, err := bindSelfDestructor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfDestructorFilterer{contract: contract}, nil
}

// bindSelfDestructor binds a generic wrapper to an already deployed contract.
func bindSelfDestructor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SelfDestructorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfDestructor *SelfDestructorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfDestructor.Contract.SelfDestructorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfDestructor *SelfDestructorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfDestructor.Contract.SelfDestructorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfDestructor *SelfDestructorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfDestructor.Contract.SelfDestructorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfDestructor *SelfDestructorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfDestructor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfDestructor *SelfDestructorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfDestructor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfDestructor *SelfDestructorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfDestructor.Contract.contract.Transact(opts, method, params...)
}

// SelfDestruct is a paid mutator transaction binding the contract method 0x3f5a0bdd.
//
// Solidity: function selfDestruct(address who) returns()
func (_SelfDestructor *SelfDestructorTransactor) SelfDestruct(opts *bind.TransactOpts, who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.contract.Transact(opts, "selfDestruct", who)
}

// SelfDestruct is a paid mutator transaction binding the contract method 0x3f5a0bdd.
//
// Solidity: function selfDestruct(address who) returns()
func (_SelfDestructor *SelfDestructorSession) SelfDestruct(who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.Contract.SelfDestruct(&_SelfDestructor.TransactOpts, who)
}

// SelfDestruct is a paid mutator transaction binding the contract method 0x3f5a0bdd.
//
// Solidity: function selfDestruct(address who) returns()
func (_SelfDestructor *SelfDestructorTransactorSession) SelfDestruct(who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.Contract.SelfDestruct(&_SelfDestructor.TransactOpts, who)
}

// WarmEmptySelfDestructor is a paid mutator transaction binding the contract method 0x88b6388c.
//
// Solidity: function warmEmptySelfDestructor(address who) returns()
func (_SelfDestructor *SelfDestructorTransactor) WarmEmptySelfDestructor(opts *bind.TransactOpts, who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.contract.Transact(opts, "warmEmptySelfDestructor", who)
}

// WarmEmptySelfDestructor is a paid mutator transaction binding the contract method 0x88b6388c.
//
// Solidity: function warmEmptySelfDestructor(address who) returns()
func (_SelfDestructor *SelfDestructorSession) WarmEmptySelfDestructor(who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.Contract.WarmEmptySelfDestructor(&_SelfDestructor.TransactOpts, who)
}

// WarmEmptySelfDestructor is a paid mutator transaction binding the contract method 0x88b6388c.
//
// Solidity: function warmEmptySelfDestructor(address who) returns()
func (_SelfDestructor *SelfDestructorTransactorSession) WarmEmptySelfDestructor(who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.Contract.WarmEmptySelfDestructor(&_SelfDestructor.TransactOpts, who)
}

// WarmSelfDestructor is a paid mutator transaction binding the contract method 0xe52f9767.
//
// Solidity: function warmSelfDestructor(address who) returns()
func (_SelfDestructor *SelfDestructorTransactor) WarmSelfDestructor(opts *bind.TransactOpts, who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.contract.Transact(opts, "warmSelfDestructor", who)
}

// WarmSelfDestructor is a paid mutator transaction binding the contract method 0xe52f9767.
//
// Solidity: function warmSelfDestructor(address who) returns()
func (_SelfDestructor *SelfDestructorSession) WarmSelfDestructor(who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.Contract.WarmSelfDestructor(&_SelfDestructor.TransactOpts, who)
}

// WarmSelfDestructor is a paid mutator transaction binding the contract method 0xe52f9767.
//
// Solidity: function warmSelfDestructor(address who) returns()
func (_SelfDestructor *SelfDestructorTransactorSession) WarmSelfDestructor(who common.Address) (*types.Transaction, error) {
	return _SelfDestructor.Contract.WarmSelfDestructor(&_SelfDestructor.TransactOpts, who)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SelfDestructor *SelfDestructorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SelfDestructor.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SelfDestructor *SelfDestructorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SelfDestructor.Contract.Fallback(&_SelfDestructor.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SelfDestructor *SelfDestructorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SelfDestructor.Contract.Fallback(&_SelfDestructor.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SelfDestructor *SelfDestructorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfDestructor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SelfDestructor *SelfDestructorSession) Receive() (*types.Transaction, error) {
	return _SelfDestructor.Contract.Receive(&_SelfDestructor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SelfDestructor *SelfDestructorTransactorSession) Receive() (*types.Transaction, error) {
	return _SelfDestructor.Contract.Receive(&_SelfDestructor.TransactOpts)
}

// SloadMetaData contains all meta data concerning the Sload contract.
var SloadMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coldSload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"warmSload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060035f8190555061011c806100235f395ff3fe6080604052348015600e575f5ffd5b5060043610603a575f3560e01c80630dbe671f14603e578063931c4f1f146058578063b4185e06146072575b5f5ffd5b6044608c565b604051604f919060cf565b60405180910390f35b605e6091565b6040516069919060cf565b60405180910390f35b607860a5565b6040516083919060cf565b60405180910390f35b5f5481565b5f5f5f54905060055f819055508091505090565b5f60045f819055505f5f5490508091505090565b5f819050919050565b60c98160b9565b82525050565b5f60208201905060e05f83018460c2565b9291505056fea2646970667358221220c4cf0ec2e66c9ae6e99edfb2a1807952b68493f085400c6261dd516eddd8dbd364736f6c634300081e0033",
}

// SloadABI is the input ABI used to generate the binding from.
// Deprecated: Use SloadMetaData.ABI instead.
var SloadABI = SloadMetaData.ABI

// SloadBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SloadMetaData.Bin instead.
var SloadBin = SloadMetaData.Bin

// DeploySload deploys a new Ethereum contract, binding an instance of Sload to it.
func DeploySload(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sload, error) {
	parsed, err := SloadMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SloadBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sload{SloadCaller: SloadCaller{contract: contract}, SloadTransactor: SloadTransactor{contract: contract}, SloadFilterer: SloadFilterer{contract: contract}}, nil
}

// Sload is an auto generated Go binding around an Ethereum contract.
type Sload struct {
	SloadCaller     // Read-only binding to the contract
	SloadTransactor // Write-only binding to the contract
	SloadFilterer   // Log filterer for contract events
}

// SloadCaller is an auto generated read-only Go binding around an Ethereum contract.
type SloadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SloadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SloadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SloadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SloadSession struct {
	Contract     *Sload            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SloadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SloadCallerSession struct {
	Contract *SloadCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SloadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SloadTransactorSession struct {
	Contract     *SloadTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SloadRaw is an auto generated low-level Go binding around an Ethereum contract.
type SloadRaw struct {
	Contract *Sload // Generic contract binding to access the raw methods on
}

// SloadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SloadCallerRaw struct {
	Contract *SloadCaller // Generic read-only contract binding to access the raw methods on
}

// SloadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SloadTransactorRaw struct {
	Contract *SloadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSload creates a new instance of Sload, bound to a specific deployed contract.
func NewSload(address common.Address, backend bind.ContractBackend) (*Sload, error) {
	contract, err := bindSload(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sload{SloadCaller: SloadCaller{contract: contract}, SloadTransactor: SloadTransactor{contract: contract}, SloadFilterer: SloadFilterer{contract: contract}}, nil
}

// NewSloadCaller creates a new read-only instance of Sload, bound to a specific deployed contract.
func NewSloadCaller(address common.Address, caller bind.ContractCaller) (*SloadCaller, error) {
	contract, err := bindSload(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SloadCaller{contract: contract}, nil
}

// NewSloadTransactor creates a new write-only instance of Sload, bound to a specific deployed contract.
func NewSloadTransactor(address common.Address, transactor bind.ContractTransactor) (*SloadTransactor, error) {
	contract, err := bindSload(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SloadTransactor{contract: contract}, nil
}

// NewSloadFilterer creates a new log filterer instance of Sload, bound to a specific deployed contract.
func NewSloadFilterer(address common.Address, filterer bind.ContractFilterer) (*SloadFilterer, error) {
	contract, err := bindSload(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SloadFilterer{contract: contract}, nil
}

// bindSload binds a generic wrapper to an already deployed contract.
func bindSload(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SloadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sload *SloadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sload.Contract.SloadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sload *SloadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.Contract.SloadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sload *SloadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sload.Contract.SloadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sload *SloadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sload.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sload *SloadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sload *SloadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sload.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Sload *SloadCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sload.contract.Call(opts, &out, "a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Sload *SloadSession) A() (*big.Int, error) {
	return _Sload.Contract.A(&_Sload.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(uint256)
func (_Sload *SloadCallerSession) A() (*big.Int, error) {
	return _Sload.Contract.A(&_Sload.CallOpts)
}

// ColdSload is a paid mutator transaction binding the contract method 0x931c4f1f.
//
// Solidity: function coldSload() returns(uint256)
func (_Sload *SloadTransactor) ColdSload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.contract.Transact(opts, "coldSload")
}

// ColdSload is a paid mutator transaction binding the contract method 0x931c4f1f.
//
// Solidity: function coldSload() returns(uint256)
func (_Sload *SloadSession) ColdSload() (*types.Transaction, error) {
	return _Sload.Contract.ColdSload(&_Sload.TransactOpts)
}

// ColdSload is a paid mutator transaction binding the contract method 0x931c4f1f.
//
// Solidity: function coldSload() returns(uint256)
func (_Sload *SloadTransactorSession) ColdSload() (*types.Transaction, error) {
	return _Sload.Contract.ColdSload(&_Sload.TransactOpts)
}

// WarmSload is a paid mutator transaction binding the contract method 0xb4185e06.
//
// Solidity: function warmSload() returns(uint256)
func (_Sload *SloadTransactor) WarmSload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sload.contract.Transact(opts, "warmSload")
}

// WarmSload is a paid mutator transaction binding the contract method 0xb4185e06.
//
// Solidity: function warmSload() returns(uint256)
func (_Sload *SloadSession) WarmSload() (*types.Transaction, error) {
	return _Sload.Contract.WarmSload(&_Sload.TransactOpts)
}

// WarmSload is a paid mutator transaction binding the contract method 0xb4185e06.
//
// Solidity: function warmSload() returns(uint256)
func (_Sload *SloadTransactorSession) WarmSload() (*types.Transaction, error) {
	return _Sload.Contract.WarmSload(&_Sload.TransactOpts)
}

// SstoreMetaData contains all meta data concerning the Sstore contract.
var SstoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"nonZeroFromStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreColdNonZeroToDifferentNonZeroValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreColdNonZeroToSameNonZeroValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreColdNonZeroValueToZero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreColdZeroToNonZero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreColdZeroToZero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreMultipleWarmNonZeroToNonZeroToNonZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreMultipleWarmNonZeroToNonZeroToSameNonZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreMultipleWarmNonZeroToZeroToNonZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreMultipleWarmNonZeroToZeroToSameNonZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreMultipleWarmZeroToNonZeroBackToZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreMultipleWarmZeroToNonZeroToNonZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreWarmNonZeroToDifferentNonZeroValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreWarmNonZeroToSameNonZeroValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreWarmNonZeroValueToZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreWarmZeroToNonZeroValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sstoreWarmZeroToZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeroFromStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060026001819055506104bf806100245f395ff3fe608060405234801561000f575f5ffd5b5060043610610134575f3560e01c806381b2333e116100b0578063bc83c5141161007f578063c9286bec11610064578063c9286bec146102aa578063e4c89164146102c8578063fe54f262146102d257610134565b8063bc83c51414610282578063bd7c534f146102a057610134565b806381b2333e1461021e5780638ae19b231461023c5780638bd99ea814610246578063ac69ed4f1461026457610134565b80635b4a53111161010757806366ca3ecf116100ec57806366ca3ecf146101d857806372f7e165146101f65780637ab978051461021457610134565b80635b4a53111461019c5780635d0c43c2146101ba57610134565b80631633aeab146101385780631c5a081014610142578063501da24e1461016057806357625b121461017e575b5f5ffd5b6101406102f0565b005b61014a6102fa565b6040516101579190610470565b60405180910390f35b61016861031c565b6040516101759190610470565b60405180910390f35b610186610332565b6040516101939190610470565b60405180910390f35b6101a4610346565b6040516101b19190610470565b60405180910390f35b6101c2610368565b6040516101cf9190610470565b60405180910390f35b6101e061036e565b6040516101ed9190610470565b60405180910390f35b6101fe610391565b60405161020b9190610470565b60405180910390f35b61021c6103af565b005b6102266103b8565b6040516102339190610470565b60405180910390f35b6102446103db565b005b61024e6103e5565b60405161025b9190610470565b60405180910390f35b61026c6103f8565b6040516102799190610470565b60405180910390f35b61028a610417565b6040516102979190610470565b60405180910390f35b6102a861041c565b005b6102b2610425565b6040516102bf9190610470565b60405180910390f35b6102d061043b565b005b6102da610443565b6040516102e79190610470565b60405180910390f35b6003600181905550565b5f5f60015490505f600181905550600154905060026001819055508091505090565b5f5f600154905060036001819055508091505090565b5f5f5f54905060015f819055508091505090565b5f5f60015490505f600181905550600154905060046001819055508091505090565b60015481565b5f5f60015490506003600181905550600154905060026001819055508091505090565b5f5f5f54905060015f819055505f5490505f5f819055508091505090565b5f600181905550565b5f5f60015490506003600181905550600154905060046001819055508091505090565b6002600181905550565b5f5f5f5490505f5f819055508091505090565b5f5f5f54905060015f819055505f54905060025f819055508091505090565b5f5481565b60015f81905550565b5f5f600154905060026001819055508091505090565b5f5f81905550565b5f5f60015490505f6001819055508091505090565b5f819050919050565b61046a81610458565b82525050565b5f6020820190506104835f830184610461565b9291505056fea2646970667358221220dc419819e13695cd7cd927804d729e5dad1909af7e1c1b15f8566ad94bf49ad264736f6c634300081e0033",
}

// SstoreABI is the input ABI used to generate the binding from.
// Deprecated: Use SstoreMetaData.ABI instead.
var SstoreABI = SstoreMetaData.ABI

// SstoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SstoreMetaData.Bin instead.
var SstoreBin = SstoreMetaData.Bin

// DeploySstore deploys a new Ethereum contract, binding an instance of Sstore to it.
func DeploySstore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sstore, error) {
	parsed, err := SstoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SstoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sstore{SstoreCaller: SstoreCaller{contract: contract}, SstoreTransactor: SstoreTransactor{contract: contract}, SstoreFilterer: SstoreFilterer{contract: contract}}, nil
}

// Sstore is an auto generated Go binding around an Ethereum contract.
type Sstore struct {
	SstoreCaller     // Read-only binding to the contract
	SstoreTransactor // Write-only binding to the contract
	SstoreFilterer   // Log filterer for contract events
}

// SstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type SstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SstoreSession struct {
	Contract     *Sstore           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SstoreCallerSession struct {
	Contract *SstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SstoreTransactorSession struct {
	Contract     *SstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type SstoreRaw struct {
	Contract *Sstore // Generic contract binding to access the raw methods on
}

// SstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SstoreCallerRaw struct {
	Contract *SstoreCaller // Generic read-only contract binding to access the raw methods on
}

// SstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SstoreTransactorRaw struct {
	Contract *SstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSstore creates a new instance of Sstore, bound to a specific deployed contract.
func NewSstore(address common.Address, backend bind.ContractBackend) (*Sstore, error) {
	contract, err := bindSstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sstore{SstoreCaller: SstoreCaller{contract: contract}, SstoreTransactor: SstoreTransactor{contract: contract}, SstoreFilterer: SstoreFilterer{contract: contract}}, nil
}

// NewSstoreCaller creates a new read-only instance of Sstore, bound to a specific deployed contract.
func NewSstoreCaller(address common.Address, caller bind.ContractCaller) (*SstoreCaller, error) {
	contract, err := bindSstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SstoreCaller{contract: contract}, nil
}

// NewSstoreTransactor creates a new write-only instance of Sstore, bound to a specific deployed contract.
func NewSstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*SstoreTransactor, error) {
	contract, err := bindSstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SstoreTransactor{contract: contract}, nil
}

// NewSstoreFilterer creates a new log filterer instance of Sstore, bound to a specific deployed contract.
func NewSstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*SstoreFilterer, error) {
	contract, err := bindSstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SstoreFilterer{contract: contract}, nil
}

// bindSstore binds a generic wrapper to an already deployed contract.
func bindSstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SstoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sstore *SstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sstore.Contract.SstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sstore *SstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.Contract.SstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sstore *SstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sstore.Contract.SstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sstore *SstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sstore *SstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sstore *SstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sstore.Contract.contract.Transact(opts, method, params...)
}

// NonZeroFromStart is a free data retrieval call binding the contract method 0x5d0c43c2.
//
// Solidity: function nonZeroFromStart() view returns(uint256)
func (_Sstore *SstoreCaller) NonZeroFromStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sstore.contract.Call(opts, &out, "nonZeroFromStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonZeroFromStart is a free data retrieval call binding the contract method 0x5d0c43c2.
//
// Solidity: function nonZeroFromStart() view returns(uint256)
func (_Sstore *SstoreSession) NonZeroFromStart() (*big.Int, error) {
	return _Sstore.Contract.NonZeroFromStart(&_Sstore.CallOpts)
}

// NonZeroFromStart is a free data retrieval call binding the contract method 0x5d0c43c2.
//
// Solidity: function nonZeroFromStart() view returns(uint256)
func (_Sstore *SstoreCallerSession) NonZeroFromStart() (*big.Int, error) {
	return _Sstore.Contract.NonZeroFromStart(&_Sstore.CallOpts)
}

// ZeroFromStart is a free data retrieval call binding the contract method 0xbc83c514.
//
// Solidity: function zeroFromStart() view returns(uint256)
func (_Sstore *SstoreCaller) ZeroFromStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sstore.contract.Call(opts, &out, "zeroFromStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZeroFromStart is a free data retrieval call binding the contract method 0xbc83c514.
//
// Solidity: function zeroFromStart() view returns(uint256)
func (_Sstore *SstoreSession) ZeroFromStart() (*big.Int, error) {
	return _Sstore.Contract.ZeroFromStart(&_Sstore.CallOpts)
}

// ZeroFromStart is a free data retrieval call binding the contract method 0xbc83c514.
//
// Solidity: function zeroFromStart() view returns(uint256)
func (_Sstore *SstoreCallerSession) ZeroFromStart() (*big.Int, error) {
	return _Sstore.Contract.ZeroFromStart(&_Sstore.CallOpts)
}

// SstoreColdNonZeroToDifferentNonZeroValue is a paid mutator transaction binding the contract method 0x1633aeab.
//
// Solidity: function sstoreColdNonZeroToDifferentNonZeroValue() returns()
func (_Sstore *SstoreTransactor) SstoreColdNonZeroToDifferentNonZeroValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreColdNonZeroToDifferentNonZeroValue")
}

// SstoreColdNonZeroToDifferentNonZeroValue is a paid mutator transaction binding the contract method 0x1633aeab.
//
// Solidity: function sstoreColdNonZeroToDifferentNonZeroValue() returns()
func (_Sstore *SstoreSession) SstoreColdNonZeroToDifferentNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdNonZeroToDifferentNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreColdNonZeroToDifferentNonZeroValue is a paid mutator transaction binding the contract method 0x1633aeab.
//
// Solidity: function sstoreColdNonZeroToDifferentNonZeroValue() returns()
func (_Sstore *SstoreTransactorSession) SstoreColdNonZeroToDifferentNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdNonZeroToDifferentNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreColdNonZeroToSameNonZeroValue is a paid mutator transaction binding the contract method 0x8ae19b23.
//
// Solidity: function sstoreColdNonZeroToSameNonZeroValue() returns()
func (_Sstore *SstoreTransactor) SstoreColdNonZeroToSameNonZeroValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreColdNonZeroToSameNonZeroValue")
}

// SstoreColdNonZeroToSameNonZeroValue is a paid mutator transaction binding the contract method 0x8ae19b23.
//
// Solidity: function sstoreColdNonZeroToSameNonZeroValue() returns()
func (_Sstore *SstoreSession) SstoreColdNonZeroToSameNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdNonZeroToSameNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreColdNonZeroToSameNonZeroValue is a paid mutator transaction binding the contract method 0x8ae19b23.
//
// Solidity: function sstoreColdNonZeroToSameNonZeroValue() returns()
func (_Sstore *SstoreTransactorSession) SstoreColdNonZeroToSameNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdNonZeroToSameNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreColdNonZeroValueToZero is a paid mutator transaction binding the contract method 0x7ab97805.
//
// Solidity: function sstoreColdNonZeroValueToZero() returns()
func (_Sstore *SstoreTransactor) SstoreColdNonZeroValueToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreColdNonZeroValueToZero")
}

// SstoreColdNonZeroValueToZero is a paid mutator transaction binding the contract method 0x7ab97805.
//
// Solidity: function sstoreColdNonZeroValueToZero() returns()
func (_Sstore *SstoreSession) SstoreColdNonZeroValueToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdNonZeroValueToZero(&_Sstore.TransactOpts)
}

// SstoreColdNonZeroValueToZero is a paid mutator transaction binding the contract method 0x7ab97805.
//
// Solidity: function sstoreColdNonZeroValueToZero() returns()
func (_Sstore *SstoreTransactorSession) SstoreColdNonZeroValueToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdNonZeroValueToZero(&_Sstore.TransactOpts)
}

// SstoreColdZeroToNonZero is a paid mutator transaction binding the contract method 0xbd7c534f.
//
// Solidity: function sstoreColdZeroToNonZero() returns()
func (_Sstore *SstoreTransactor) SstoreColdZeroToNonZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreColdZeroToNonZero")
}

// SstoreColdZeroToNonZero is a paid mutator transaction binding the contract method 0xbd7c534f.
//
// Solidity: function sstoreColdZeroToNonZero() returns()
func (_Sstore *SstoreSession) SstoreColdZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreColdZeroToNonZero is a paid mutator transaction binding the contract method 0xbd7c534f.
//
// Solidity: function sstoreColdZeroToNonZero() returns()
func (_Sstore *SstoreTransactorSession) SstoreColdZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreColdZeroToZero is a paid mutator transaction binding the contract method 0xe4c89164.
//
// Solidity: function sstoreColdZeroToZero() returns()
func (_Sstore *SstoreTransactor) SstoreColdZeroToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreColdZeroToZero")
}

// SstoreColdZeroToZero is a paid mutator transaction binding the contract method 0xe4c89164.
//
// Solidity: function sstoreColdZeroToZero() returns()
func (_Sstore *SstoreSession) SstoreColdZeroToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdZeroToZero(&_Sstore.TransactOpts)
}

// SstoreColdZeroToZero is a paid mutator transaction binding the contract method 0xe4c89164.
//
// Solidity: function sstoreColdZeroToZero() returns()
func (_Sstore *SstoreTransactorSession) SstoreColdZeroToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreColdZeroToZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToNonZeroToNonZero is a paid mutator transaction binding the contract method 0x81b2333e.
//
// Solidity: function sstoreMultipleWarmNonZeroToNonZeroToNonZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreMultipleWarmNonZeroToNonZeroToNonZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreMultipleWarmNonZeroToNonZeroToNonZero")
}

// SstoreMultipleWarmNonZeroToNonZeroToNonZero is a paid mutator transaction binding the contract method 0x81b2333e.
//
// Solidity: function sstoreMultipleWarmNonZeroToNonZeroToNonZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreMultipleWarmNonZeroToNonZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToNonZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToNonZeroToNonZero is a paid mutator transaction binding the contract method 0x81b2333e.
//
// Solidity: function sstoreMultipleWarmNonZeroToNonZeroToNonZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreMultipleWarmNonZeroToNonZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToNonZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToNonZeroToSameNonZero is a paid mutator transaction binding the contract method 0x66ca3ecf.
//
// Solidity: function sstoreMultipleWarmNonZeroToNonZeroToSameNonZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreMultipleWarmNonZeroToNonZeroToSameNonZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreMultipleWarmNonZeroToNonZeroToSameNonZero")
}

// SstoreMultipleWarmNonZeroToNonZeroToSameNonZero is a paid mutator transaction binding the contract method 0x66ca3ecf.
//
// Solidity: function sstoreMultipleWarmNonZeroToNonZeroToSameNonZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreMultipleWarmNonZeroToNonZeroToSameNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToNonZeroToSameNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToNonZeroToSameNonZero is a paid mutator transaction binding the contract method 0x66ca3ecf.
//
// Solidity: function sstoreMultipleWarmNonZeroToNonZeroToSameNonZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreMultipleWarmNonZeroToNonZeroToSameNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToNonZeroToSameNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToZeroToNonZero is a paid mutator transaction binding the contract method 0x5b4a5311.
//
// Solidity: function sstoreMultipleWarmNonZeroToZeroToNonZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreMultipleWarmNonZeroToZeroToNonZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreMultipleWarmNonZeroToZeroToNonZero")
}

// SstoreMultipleWarmNonZeroToZeroToNonZero is a paid mutator transaction binding the contract method 0x5b4a5311.
//
// Solidity: function sstoreMultipleWarmNonZeroToZeroToNonZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreMultipleWarmNonZeroToZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToZeroToNonZero is a paid mutator transaction binding the contract method 0x5b4a5311.
//
// Solidity: function sstoreMultipleWarmNonZeroToZeroToNonZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreMultipleWarmNonZeroToZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToZeroToSameNonZero is a paid mutator transaction binding the contract method 0x1c5a0810.
//
// Solidity: function sstoreMultipleWarmNonZeroToZeroToSameNonZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreMultipleWarmNonZeroToZeroToSameNonZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreMultipleWarmNonZeroToZeroToSameNonZero")
}

// SstoreMultipleWarmNonZeroToZeroToSameNonZero is a paid mutator transaction binding the contract method 0x1c5a0810.
//
// Solidity: function sstoreMultipleWarmNonZeroToZeroToSameNonZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreMultipleWarmNonZeroToZeroToSameNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToZeroToSameNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmNonZeroToZeroToSameNonZero is a paid mutator transaction binding the contract method 0x1c5a0810.
//
// Solidity: function sstoreMultipleWarmNonZeroToZeroToSameNonZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreMultipleWarmNonZeroToZeroToSameNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmNonZeroToZeroToSameNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmZeroToNonZeroBackToZero is a paid mutator transaction binding the contract method 0x72f7e165.
//
// Solidity: function sstoreMultipleWarmZeroToNonZeroBackToZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreMultipleWarmZeroToNonZeroBackToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreMultipleWarmZeroToNonZeroBackToZero")
}

// SstoreMultipleWarmZeroToNonZeroBackToZero is a paid mutator transaction binding the contract method 0x72f7e165.
//
// Solidity: function sstoreMultipleWarmZeroToNonZeroBackToZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreMultipleWarmZeroToNonZeroBackToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmZeroToNonZeroBackToZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmZeroToNonZeroBackToZero is a paid mutator transaction binding the contract method 0x72f7e165.
//
// Solidity: function sstoreMultipleWarmZeroToNonZeroBackToZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreMultipleWarmZeroToNonZeroBackToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmZeroToNonZeroBackToZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmZeroToNonZeroToNonZero is a paid mutator transaction binding the contract method 0xac69ed4f.
//
// Solidity: function sstoreMultipleWarmZeroToNonZeroToNonZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreMultipleWarmZeroToNonZeroToNonZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreMultipleWarmZeroToNonZeroToNonZero")
}

// SstoreMultipleWarmZeroToNonZeroToNonZero is a paid mutator transaction binding the contract method 0xac69ed4f.
//
// Solidity: function sstoreMultipleWarmZeroToNonZeroToNonZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreMultipleWarmZeroToNonZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmZeroToNonZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreMultipleWarmZeroToNonZeroToNonZero is a paid mutator transaction binding the contract method 0xac69ed4f.
//
// Solidity: function sstoreMultipleWarmZeroToNonZeroToNonZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreMultipleWarmZeroToNonZeroToNonZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreMultipleWarmZeroToNonZeroToNonZero(&_Sstore.TransactOpts)
}

// SstoreWarmNonZeroToDifferentNonZeroValue is a paid mutator transaction binding the contract method 0x501da24e.
//
// Solidity: function sstoreWarmNonZeroToDifferentNonZeroValue() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreWarmNonZeroToDifferentNonZeroValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreWarmNonZeroToDifferentNonZeroValue")
}

// SstoreWarmNonZeroToDifferentNonZeroValue is a paid mutator transaction binding the contract method 0x501da24e.
//
// Solidity: function sstoreWarmNonZeroToDifferentNonZeroValue() returns(uint256)
func (_Sstore *SstoreSession) SstoreWarmNonZeroToDifferentNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmNonZeroToDifferentNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreWarmNonZeroToDifferentNonZeroValue is a paid mutator transaction binding the contract method 0x501da24e.
//
// Solidity: function sstoreWarmNonZeroToDifferentNonZeroValue() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreWarmNonZeroToDifferentNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmNonZeroToDifferentNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreWarmNonZeroToSameNonZeroValue is a paid mutator transaction binding the contract method 0xc9286bec.
//
// Solidity: function sstoreWarmNonZeroToSameNonZeroValue() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreWarmNonZeroToSameNonZeroValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreWarmNonZeroToSameNonZeroValue")
}

// SstoreWarmNonZeroToSameNonZeroValue is a paid mutator transaction binding the contract method 0xc9286bec.
//
// Solidity: function sstoreWarmNonZeroToSameNonZeroValue() returns(uint256)
func (_Sstore *SstoreSession) SstoreWarmNonZeroToSameNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmNonZeroToSameNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreWarmNonZeroToSameNonZeroValue is a paid mutator transaction binding the contract method 0xc9286bec.
//
// Solidity: function sstoreWarmNonZeroToSameNonZeroValue() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreWarmNonZeroToSameNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmNonZeroToSameNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreWarmNonZeroValueToZero is a paid mutator transaction binding the contract method 0xfe54f262.
//
// Solidity: function sstoreWarmNonZeroValueToZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreWarmNonZeroValueToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreWarmNonZeroValueToZero")
}

// SstoreWarmNonZeroValueToZero is a paid mutator transaction binding the contract method 0xfe54f262.
//
// Solidity: function sstoreWarmNonZeroValueToZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreWarmNonZeroValueToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmNonZeroValueToZero(&_Sstore.TransactOpts)
}

// SstoreWarmNonZeroValueToZero is a paid mutator transaction binding the contract method 0xfe54f262.
//
// Solidity: function sstoreWarmNonZeroValueToZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreWarmNonZeroValueToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmNonZeroValueToZero(&_Sstore.TransactOpts)
}

// SstoreWarmZeroToNonZeroValue is a paid mutator transaction binding the contract method 0x57625b12.
//
// Solidity: function sstoreWarmZeroToNonZeroValue() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreWarmZeroToNonZeroValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreWarmZeroToNonZeroValue")
}

// SstoreWarmZeroToNonZeroValue is a paid mutator transaction binding the contract method 0x57625b12.
//
// Solidity: function sstoreWarmZeroToNonZeroValue() returns(uint256)
func (_Sstore *SstoreSession) SstoreWarmZeroToNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmZeroToNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreWarmZeroToNonZeroValue is a paid mutator transaction binding the contract method 0x57625b12.
//
// Solidity: function sstoreWarmZeroToNonZeroValue() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreWarmZeroToNonZeroValue() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmZeroToNonZeroValue(&_Sstore.TransactOpts)
}

// SstoreWarmZeroToZero is a paid mutator transaction binding the contract method 0x8bd99ea8.
//
// Solidity: function sstoreWarmZeroToZero() returns(uint256)
func (_Sstore *SstoreTransactor) SstoreWarmZeroToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sstore.contract.Transact(opts, "sstoreWarmZeroToZero")
}

// SstoreWarmZeroToZero is a paid mutator transaction binding the contract method 0x8bd99ea8.
//
// Solidity: function sstoreWarmZeroToZero() returns(uint256)
func (_Sstore *SstoreSession) SstoreWarmZeroToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmZeroToZero(&_Sstore.TransactOpts)
}

// SstoreWarmZeroToZero is a paid mutator transaction binding the contract method 0x8bd99ea8.
//
// Solidity: function sstoreWarmZeroToZero() returns(uint256)
func (_Sstore *SstoreTransactorSession) SstoreWarmZeroToZero() (*types.Transaction, error) {
	return _Sstore.Contract.SstoreWarmZeroToZero(&_Sstore.TransactOpts)
}

// StaticCalleeMetaData contains all meta data concerning the StaticCallee contract.
var StaticCalleeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50624142435f8190555060af8060235f395ff3fe6080604052348015600e575f5ffd5b50600436106026575f3560e01c8063f2c9ecd814602a575b5f5ffd5b60306044565b604051603b91906062565b60405180910390f35b5f5f54905090565b5f819050919050565b605c81604c565b82525050565b5f60208201905060735f8301846055565b9291505056fea2646970667358221220ee66ec990cc7d3426ac0a225f397662a55ee4547a56e3279842075ea65bd962364736f6c634300081e0033",
}

// StaticCalleeABI is the input ABI used to generate the binding from.
// Deprecated: Use StaticCalleeMetaData.ABI instead.
var StaticCalleeABI = StaticCalleeMetaData.ABI

// StaticCalleeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StaticCalleeMetaData.Bin instead.
var StaticCalleeBin = StaticCalleeMetaData.Bin

// DeployStaticCallee deploys a new Ethereum contract, binding an instance of StaticCallee to it.
func DeployStaticCallee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StaticCallee, error) {
	parsed, err := StaticCalleeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StaticCalleeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StaticCallee{StaticCalleeCaller: StaticCalleeCaller{contract: contract}, StaticCalleeTransactor: StaticCalleeTransactor{contract: contract}, StaticCalleeFilterer: StaticCalleeFilterer{contract: contract}}, nil
}

// StaticCallee is an auto generated Go binding around an Ethereum contract.
type StaticCallee struct {
	StaticCalleeCaller     // Read-only binding to the contract
	StaticCalleeTransactor // Write-only binding to the contract
	StaticCalleeFilterer   // Log filterer for contract events
}

// StaticCalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaticCalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticCalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaticCalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticCalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaticCalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticCalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaticCalleeSession struct {
	Contract     *StaticCallee     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaticCalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaticCalleeCallerSession struct {
	Contract *StaticCalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StaticCalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaticCalleeTransactorSession struct {
	Contract     *StaticCalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StaticCalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaticCalleeRaw struct {
	Contract *StaticCallee // Generic contract binding to access the raw methods on
}

// StaticCalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaticCalleeCallerRaw struct {
	Contract *StaticCalleeCaller // Generic read-only contract binding to access the raw methods on
}

// StaticCalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaticCalleeTransactorRaw struct {
	Contract *StaticCalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaticCallee creates a new instance of StaticCallee, bound to a specific deployed contract.
func NewStaticCallee(address common.Address, backend bind.ContractBackend) (*StaticCallee, error) {
	contract, err := bindStaticCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaticCallee{StaticCalleeCaller: StaticCalleeCaller{contract: contract}, StaticCalleeTransactor: StaticCalleeTransactor{contract: contract}, StaticCalleeFilterer: StaticCalleeFilterer{contract: contract}}, nil
}

// NewStaticCalleeCaller creates a new read-only instance of StaticCallee, bound to a specific deployed contract.
func NewStaticCalleeCaller(address common.Address, caller bind.ContractCaller) (*StaticCalleeCaller, error) {
	contract, err := bindStaticCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaticCalleeCaller{contract: contract}, nil
}

// NewStaticCalleeTransactor creates a new write-only instance of StaticCallee, bound to a specific deployed contract.
func NewStaticCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*StaticCalleeTransactor, error) {
	contract, err := bindStaticCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaticCalleeTransactor{contract: contract}, nil
}

// NewStaticCalleeFilterer creates a new log filterer instance of StaticCallee, bound to a specific deployed contract.
func NewStaticCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*StaticCalleeFilterer, error) {
	contract, err := bindStaticCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaticCalleeFilterer{contract: contract}, nil
}

// bindStaticCallee binds a generic wrapper to an already deployed contract.
func bindStaticCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaticCalleeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticCallee *StaticCalleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticCallee.Contract.StaticCalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticCallee *StaticCalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticCallee.Contract.StaticCalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticCallee *StaticCalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticCallee.Contract.StaticCalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticCallee *StaticCalleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticCallee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticCallee *StaticCalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticCallee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticCallee *StaticCalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticCallee.Contract.contract.Transact(opts, method, params...)
}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_StaticCallee *StaticCalleeCaller) GetNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaticCallee.contract.Call(opts, &out, "getNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_StaticCallee *StaticCalleeSession) GetNumber() (*big.Int, error) {
	return _StaticCallee.Contract.GetNumber(&_StaticCallee.CallOpts)
}

// GetNumber is a free data retrieval call binding the contract method 0xf2c9ecd8.
//
// Solidity: function getNumber() view returns(uint256)
func (_StaticCallee *StaticCalleeCallerSession) GetNumber() (*big.Int, error) {
	return _StaticCallee.Contract.GetNumber(&_StaticCallee.CallOpts)
}

// StaticCallerMetaData contains all meta data concerning the StaticCaller contract.
var StaticCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallEmptyCold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallEmptyColdMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallEmptyWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallEmptyWarmMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallNonEmptyCold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallNonEmptyColdMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallNonEmptyWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nonEmptyAddress\",\"type\":\"address\"}],\"name\":\"testStaticCallNonEmptyWarmMemExpansion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107168061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a1575f3560e01c8063afa67c5c11610074578063e2c8580511610059578063e2c858051461014f578063e53f24b91461016b578063fd8c5a4c14610187576100a1565b8063afa67c5c14610117578063b7f2182114610133576100a1565b80632e52d606146100a5578063525ef06a146100c35780637f4a0527146100df57806395ed58a3146100fb575b5f5ffd5b6100ad6101a3565b6040516100ba9190610569565b60405180910390f35b6100dd60048036038101906100d891906105e0565b6101a8565b005b6100f960048036038101906100f491906105e0565b6101b5565b005b610115600480360381019061011091906105e0565b6101c2565b005b610131600480360381019061012c91906105e0565b610214565b005b61014d600480360381019061014891906105e0565b610266565b005b610169600480360381019061016491906105e0565b6102b8565b005b610185600480360381019061018091906105e0565b6102c5565b005b6101a1600480360381019061019c91906105e0565b610317565b005b5f5481565b6101b181610324565b5050565b6101be81610324565b5050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f6101e78361043d565b90508080156101f557505f82145b156102065760425f8190555061020f565b6113375f819055505b505050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f61023983610324565b905080801561024757505f82145b156102585760425f81905550610261565b6113375f819055505b505050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f61028b83610324565b905080801561029957505f82145b156102aa5760425f819055506102b3565b6113375f819055505b505050565b6102c18161048f565b5050565b5f8173ffffffffffffffffffffffffffffffffffffffff163190505f6102ea8361048f565b90508080156102f857505f82145b156103095760425f81905550610312565b6113375f819055505b505050565b6103208161043d565b5050565b5f5f602067ffffffffffffffff8111156103415761034061060b565b5b6040519080825280601f01601f1916602001820160405280156103735781602001600182028036833780820191505090505b5090505f6040516024016040516020818303038152906040527ff2c9ecd8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f8151905060208084018260208501885afa9350831561042c57826104209061067a565b5f1c5f81905550610435565b6113375f819055505b505050919050565b5f6060602059015a6020826004848885fa935060208201604052831561046257815192505b5050811561048057806104749061067a565b5f1c5f81905550610489565b6113375f819055505b50919050565b5f5f6040516024016040516020818303038152906040527ff2c9ecd8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f81519050602059015a60208284602087018985fa9450505082156105415760015f8190555061054a565b6113375f819055505b5050919050565b5f819050919050565b61056381610551565b82525050565b5f60208201905061057c5f83018461055a565b92915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105af82610586565b9050919050565b6105bf816105a5565b81146105c9575f5ffd5b50565b5f813590506105da816105b6565b92915050565b5f602082840312156105f5576105f4610582565b5b5f610602848285016105cc565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f81519050919050565b5f819050602082019050919050565b5f819050919050565b5f6106658251610651565b80915050919050565b5f82821b905092915050565b5f61068482610638565b8261068e84610642565b90506106998161065a565b925060208210156106d9576106d47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360080261066e565b831692505b505091905056fea2646970667358221220bd99c6cf32e96ea55bddfd5d5da11846e6179d6e747caf59ef7632bfcd3aec7a64736f6c634300081e0033",
}

// StaticCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use StaticCallerMetaData.ABI instead.
var StaticCallerABI = StaticCallerMetaData.ABI

// StaticCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StaticCallerMetaData.Bin instead.
var StaticCallerBin = StaticCallerMetaData.Bin

// DeployStaticCaller deploys a new Ethereum contract, binding an instance of StaticCaller to it.
func DeployStaticCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StaticCaller, error) {
	parsed, err := StaticCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StaticCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StaticCaller{StaticCallerCaller: StaticCallerCaller{contract: contract}, StaticCallerTransactor: StaticCallerTransactor{contract: contract}, StaticCallerFilterer: StaticCallerFilterer{contract: contract}}, nil
}

// StaticCaller is an auto generated Go binding around an Ethereum contract.
type StaticCaller struct {
	StaticCallerCaller     // Read-only binding to the contract
	StaticCallerTransactor // Write-only binding to the contract
	StaticCallerFilterer   // Log filterer for contract events
}

// StaticCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaticCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaticCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaticCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaticCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaticCallerSession struct {
	Contract     *StaticCaller     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaticCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaticCallerCallerSession struct {
	Contract *StaticCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StaticCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaticCallerTransactorSession struct {
	Contract     *StaticCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StaticCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaticCallerRaw struct {
	Contract *StaticCaller // Generic contract binding to access the raw methods on
}

// StaticCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaticCallerCallerRaw struct {
	Contract *StaticCallerCaller // Generic read-only contract binding to access the raw methods on
}

// StaticCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaticCallerTransactorRaw struct {
	Contract *StaticCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaticCaller creates a new instance of StaticCaller, bound to a specific deployed contract.
func NewStaticCaller(address common.Address, backend bind.ContractBackend) (*StaticCaller, error) {
	contract, err := bindStaticCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaticCaller{StaticCallerCaller: StaticCallerCaller{contract: contract}, StaticCallerTransactor: StaticCallerTransactor{contract: contract}, StaticCallerFilterer: StaticCallerFilterer{contract: contract}}, nil
}

// NewStaticCallerCaller creates a new read-only instance of StaticCaller, bound to a specific deployed contract.
func NewStaticCallerCaller(address common.Address, caller bind.ContractCaller) (*StaticCallerCaller, error) {
	contract, err := bindStaticCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaticCallerCaller{contract: contract}, nil
}

// NewStaticCallerTransactor creates a new write-only instance of StaticCaller, bound to a specific deployed contract.
func NewStaticCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*StaticCallerTransactor, error) {
	contract, err := bindStaticCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaticCallerTransactor{contract: contract}, nil
}

// NewStaticCallerFilterer creates a new log filterer instance of StaticCaller, bound to a specific deployed contract.
func NewStaticCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*StaticCallerFilterer, error) {
	contract, err := bindStaticCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaticCallerFilterer{contract: contract}, nil
}

// bindStaticCaller binds a generic wrapper to an already deployed contract.
func bindStaticCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaticCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticCaller *StaticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticCaller.Contract.StaticCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticCaller *StaticCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticCaller.Contract.StaticCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticCaller *StaticCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticCaller.Contract.StaticCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaticCaller *StaticCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaticCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaticCaller *StaticCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaticCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaticCaller *StaticCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaticCaller.Contract.contract.Transact(opts, method, params...)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_StaticCaller *StaticCallerCaller) N(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaticCaller.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_StaticCaller *StaticCallerSession) N() (*big.Int, error) {
	return _StaticCaller.Contract.N(&_StaticCaller.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(uint256)
func (_StaticCaller *StaticCallerCallerSession) N() (*big.Int, error) {
	return _StaticCaller.Contract.N(&_StaticCaller.CallOpts)
}

// TestStaticCallEmptyCold is a paid mutator transaction binding the contract method 0x525ef06a.
//
// Solidity: function testStaticCallEmptyCold(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallEmptyCold(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallEmptyCold", emptyAddress)
}

// TestStaticCallEmptyCold is a paid mutator transaction binding the contract method 0x525ef06a.
//
// Solidity: function testStaticCallEmptyCold(address emptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallEmptyCold(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyCold(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyCold is a paid mutator transaction binding the contract method 0x525ef06a.
//
// Solidity: function testStaticCallEmptyCold(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallEmptyCold(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyCold(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0xfd8c5a4c.
//
// Solidity: function testStaticCallEmptyColdMemExpansion(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallEmptyColdMemExpansion(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallEmptyColdMemExpansion", emptyAddress)
}

// TestStaticCallEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0xfd8c5a4c.
//
// Solidity: function testStaticCallEmptyColdMemExpansion(address emptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallEmptyColdMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyColdMemExpansion(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0xfd8c5a4c.
//
// Solidity: function testStaticCallEmptyColdMemExpansion(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallEmptyColdMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyColdMemExpansion(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyWarm is a paid mutator transaction binding the contract method 0xafa67c5c.
//
// Solidity: function testStaticCallEmptyWarm(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallEmptyWarm(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallEmptyWarm", emptyAddress)
}

// TestStaticCallEmptyWarm is a paid mutator transaction binding the contract method 0xafa67c5c.
//
// Solidity: function testStaticCallEmptyWarm(address emptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallEmptyWarm(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyWarm(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyWarm is a paid mutator transaction binding the contract method 0xafa67c5c.
//
// Solidity: function testStaticCallEmptyWarm(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallEmptyWarm(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyWarm(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x95ed58a3.
//
// Solidity: function testStaticCallEmptyWarmMemExpansion(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallEmptyWarmMemExpansion(opts *bind.TransactOpts, emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallEmptyWarmMemExpansion", emptyAddress)
}

// TestStaticCallEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x95ed58a3.
//
// Solidity: function testStaticCallEmptyWarmMemExpansion(address emptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallEmptyWarmMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyWarmMemExpansion(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0x95ed58a3.
//
// Solidity: function testStaticCallEmptyWarmMemExpansion(address emptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallEmptyWarmMemExpansion(emptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallEmptyWarmMemExpansion(&_StaticCaller.TransactOpts, emptyAddress)
}

// TestStaticCallNonEmptyCold is a paid mutator transaction binding the contract method 0x7f4a0527.
//
// Solidity: function testStaticCallNonEmptyCold(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallNonEmptyCold(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallNonEmptyCold", nonEmptyAddress)
}

// TestStaticCallNonEmptyCold is a paid mutator transaction binding the contract method 0x7f4a0527.
//
// Solidity: function testStaticCallNonEmptyCold(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallNonEmptyCold(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyCold(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyCold is a paid mutator transaction binding the contract method 0x7f4a0527.
//
// Solidity: function testStaticCallNonEmptyCold(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallNonEmptyCold(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyCold(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0xe2c85805.
//
// Solidity: function testStaticCallNonEmptyColdMemExpansion(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallNonEmptyColdMemExpansion(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallNonEmptyColdMemExpansion", nonEmptyAddress)
}

// TestStaticCallNonEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0xe2c85805.
//
// Solidity: function testStaticCallNonEmptyColdMemExpansion(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallNonEmptyColdMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyColdMemExpansion(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyColdMemExpansion is a paid mutator transaction binding the contract method 0xe2c85805.
//
// Solidity: function testStaticCallNonEmptyColdMemExpansion(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallNonEmptyColdMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyColdMemExpansion(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyWarm is a paid mutator transaction binding the contract method 0xb7f21821.
//
// Solidity: function testStaticCallNonEmptyWarm(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallNonEmptyWarm(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallNonEmptyWarm", nonEmptyAddress)
}

// TestStaticCallNonEmptyWarm is a paid mutator transaction binding the contract method 0xb7f21821.
//
// Solidity: function testStaticCallNonEmptyWarm(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallNonEmptyWarm(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyWarm(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyWarm is a paid mutator transaction binding the contract method 0xb7f21821.
//
// Solidity: function testStaticCallNonEmptyWarm(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallNonEmptyWarm(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyWarm(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0xe53f24b9.
//
// Solidity: function testStaticCallNonEmptyWarmMemExpansion(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactor) TestStaticCallNonEmptyWarmMemExpansion(opts *bind.TransactOpts, nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.contract.Transact(opts, "testStaticCallNonEmptyWarmMemExpansion", nonEmptyAddress)
}

// TestStaticCallNonEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0xe53f24b9.
//
// Solidity: function testStaticCallNonEmptyWarmMemExpansion(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerSession) TestStaticCallNonEmptyWarmMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyWarmMemExpansion(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// TestStaticCallNonEmptyWarmMemExpansion is a paid mutator transaction binding the contract method 0xe53f24b9.
//
// Solidity: function testStaticCallNonEmptyWarmMemExpansion(address nonEmptyAddress) returns()
func (_StaticCaller *StaticCallerTransactorSession) TestStaticCallNonEmptyWarmMemExpansion(nonEmptyAddress common.Address) (*types.Transaction, error) {
	return _StaticCaller.Contract.TestStaticCallNonEmptyWarmMemExpansion(&_StaticCaller.TransactOpts, nonEmptyAddress)
}

// StylusCallerMetaData contains all meta data concerning the StylusCaller contract.
var StylusCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"programAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"callKeccak\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104158061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c80631d00bae41461002d575b5f5ffd5b610047600480360381019061004291906102a0565b610049565b005b5f600160405160200161005c919061033a565b6040516020818303038152906040528260405160200161007d9291906103a6565b60405160208183030381529060405290508273ffffffffffffffffffffffffffffffffffffffff16816040516100b391906103c9565b5f604051808303815f865af19150503d805f81146100ec576040519150601f19603f3d011682016040523d82523d5f602084013e6100f1565b606091505b505050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101338261010a565b9050919050565b61014381610129565b811461014d575f5ffd5b50565b5f8135905061015e8161013a565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101b28261016c565b810181811067ffffffffffffffff821117156101d1576101d061017c565b5b80604052505050565b5f6101e36100f9565b90506101ef82826101a9565b919050565b5f67ffffffffffffffff82111561020e5761020d61017c565b5b6102178261016c565b9050602081019050919050565b828183375f83830152505050565b5f61024461023f846101f4565b6101da565b9050828152602081018484840111156102605761025f610168565b5b61026b848285610224565b509392505050565b5f82601f83011261028757610286610164565b5b8135610297848260208601610232565b91505092915050565b5f5f604083850312156102b6576102b5610102565b5b5f6102c385828601610150565b925050602083013567ffffffffffffffff8111156102e4576102e3610106565b5b6102f085828601610273565b9150509250929050565b5f60ff82169050919050565b5f8160f81b9050919050565b5f61031c82610306565b9050919050565b61033461032f826102fa565b610312565b82525050565b5f6103458284610323565b60018201915081905092915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61038082610354565b61038a818561035e565b935061039a818560208601610368565b80840191505092915050565b5f6103b18285610376565b91506103bd8284610376565b91508190509392505050565b5f6103d48284610376565b91508190509291505056fea26469706673582212209e90f0d4b0beae554aae516753fd6db841cdb574e1c8fda3a04a21d7741aac7a64736f6c634300081e0033",
}

// StylusCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use StylusCallerMetaData.ABI instead.
var StylusCallerABI = StylusCallerMetaData.ABI

// StylusCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StylusCallerMetaData.Bin instead.
var StylusCallerBin = StylusCallerMetaData.Bin

// DeployStylusCaller deploys a new Ethereum contract, binding an instance of StylusCaller to it.
func DeployStylusCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StylusCaller, error) {
	parsed, err := StylusCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StylusCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StylusCaller{StylusCallerCaller: StylusCallerCaller{contract: contract}, StylusCallerTransactor: StylusCallerTransactor{contract: contract}, StylusCallerFilterer: StylusCallerFilterer{contract: contract}}, nil
}

// StylusCaller is an auto generated Go binding around an Ethereum contract.
type StylusCaller struct {
	StylusCallerCaller     // Read-only binding to the contract
	StylusCallerTransactor // Write-only binding to the contract
	StylusCallerFilterer   // Log filterer for contract events
}

// StylusCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StylusCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StylusCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StylusCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StylusCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StylusCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StylusCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StylusCallerSession struct {
	Contract     *StylusCaller     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StylusCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StylusCallerCallerSession struct {
	Contract *StylusCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StylusCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StylusCallerTransactorSession struct {
	Contract     *StylusCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StylusCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StylusCallerRaw struct {
	Contract *StylusCaller // Generic contract binding to access the raw methods on
}

// StylusCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StylusCallerCallerRaw struct {
	Contract *StylusCallerCaller // Generic read-only contract binding to access the raw methods on
}

// StylusCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StylusCallerTransactorRaw struct {
	Contract *StylusCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStylusCaller creates a new instance of StylusCaller, bound to a specific deployed contract.
func NewStylusCaller(address common.Address, backend bind.ContractBackend) (*StylusCaller, error) {
	contract, err := bindStylusCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StylusCaller{StylusCallerCaller: StylusCallerCaller{contract: contract}, StylusCallerTransactor: StylusCallerTransactor{contract: contract}, StylusCallerFilterer: StylusCallerFilterer{contract: contract}}, nil
}

// NewStylusCallerCaller creates a new read-only instance of StylusCaller, bound to a specific deployed contract.
func NewStylusCallerCaller(address common.Address, caller bind.ContractCaller) (*StylusCallerCaller, error) {
	contract, err := bindStylusCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StylusCallerCaller{contract: contract}, nil
}

// NewStylusCallerTransactor creates a new write-only instance of StylusCaller, bound to a specific deployed contract.
func NewStylusCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*StylusCallerTransactor, error) {
	contract, err := bindStylusCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StylusCallerTransactor{contract: contract}, nil
}

// NewStylusCallerFilterer creates a new log filterer instance of StylusCaller, bound to a specific deployed contract.
func NewStylusCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*StylusCallerFilterer, error) {
	contract, err := bindStylusCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StylusCallerFilterer{contract: contract}, nil
}

// bindStylusCaller binds a generic wrapper to an already deployed contract.
func bindStylusCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StylusCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StylusCaller *StylusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StylusCaller.Contract.StylusCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StylusCaller *StylusCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StylusCaller.Contract.StylusCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StylusCaller *StylusCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StylusCaller.Contract.StylusCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StylusCaller *StylusCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StylusCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StylusCaller *StylusCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StylusCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StylusCaller *StylusCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StylusCaller.Contract.contract.Transact(opts, method, params...)
}

// CallKeccak is a paid mutator transaction binding the contract method 0x1d00bae4.
//
// Solidity: function callKeccak(address programAddress, bytes input) returns()
func (_StylusCaller *StylusCallerTransactor) CallKeccak(opts *bind.TransactOpts, programAddress common.Address, input []byte) (*types.Transaction, error) {
	return _StylusCaller.contract.Transact(opts, "callKeccak", programAddress, input)
}

// CallKeccak is a paid mutator transaction binding the contract method 0x1d00bae4.
//
// Solidity: function callKeccak(address programAddress, bytes input) returns()
func (_StylusCaller *StylusCallerSession) CallKeccak(programAddress common.Address, input []byte) (*types.Transaction, error) {
	return _StylusCaller.Contract.CallKeccak(&_StylusCaller.TransactOpts, programAddress, input)
}

// CallKeccak is a paid mutator transaction binding the contract method 0x1d00bae4.
//
// Solidity: function callKeccak(address programAddress, bytes input) returns()
func (_StylusCaller *StylusCallerTransactorSession) CallKeccak(programAddress common.Address, input []byte) (*types.Transaction, error) {
	return _StylusCaller.Contract.CallKeccak(&_StylusCaller.TransactOpts, programAddress, input)
}
