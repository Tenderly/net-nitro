// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package express_lane_auction_legacy_gen

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

// Bid is an auto generated low-level Go binding around an user-defined struct.
type Bid struct {
	ExpressLaneController common.Address
	Amount                *big.Int
	Signature             []byte
}

// ELCRound is an auto generated low-level Go binding around an user-defined struct.
type ELCRound struct {
	ExpressLaneController common.Address
	Round                 uint64
}

// InitArgs is an auto generated low-level Go binding around an user-defined struct.
type InitArgs struct {
	Auctioneer              common.Address
	BiddingToken            common.Address
	Beneficiary             common.Address
	RoundTimingInfo         RoundTimingInfo
	MinReservePrice         *big.Int
	AuctioneerAdmin         common.Address
	MinReservePriceSetter   common.Address
	ReservePriceSetter      common.Address
	ReservePriceSetterAdmin common.Address
	BeneficiarySetter       common.Address
	RoundTimingSetter       common.Address
	MasterAdmin             common.Address
}

// RoundTimingInfo is an auto generated low-level Go binding around an user-defined struct.
type RoundTimingInfo struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}

// Transferor is an auto generated low-level Go binding around an user-defined struct.
type Transferor struct {
	Addr            common.Address
	FixedUntilRound uint64
}

// BalanceLibMetaData contains all meta data concerning the BalanceLib contract.
var BalanceLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203976f8013546a49c16ae80f8cf4dcdb350ea9590e88e6a740926f444c9fd65a764736f6c63430008090033",
}

// BalanceLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceLibMetaData.ABI instead.
var BalanceLibABI = BalanceLibMetaData.ABI

// BalanceLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BalanceLibMetaData.Bin instead.
var BalanceLibBin = BalanceLibMetaData.Bin

// DeployBalanceLib deploys a new Ethereum contract, binding an instance of BalanceLib to it.
func DeployBalanceLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceLib, error) {
	parsed, err := BalanceLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BalanceLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceLib{BalanceLibCaller: BalanceLibCaller{contract: contract}, BalanceLibTransactor: BalanceLibTransactor{contract: contract}, BalanceLibFilterer: BalanceLibFilterer{contract: contract}}, nil
}

// BalanceLib is an auto generated Go binding around an Ethereum contract.
type BalanceLib struct {
	BalanceLibCaller     // Read-only binding to the contract
	BalanceLibTransactor // Write-only binding to the contract
	BalanceLibFilterer   // Log filterer for contract events
}

// BalanceLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceLibSession struct {
	Contract     *BalanceLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceLibCallerSession struct {
	Contract *BalanceLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BalanceLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceLibTransactorSession struct {
	Contract     *BalanceLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BalanceLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceLibRaw struct {
	Contract *BalanceLib // Generic contract binding to access the raw methods on
}

// BalanceLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceLibCallerRaw struct {
	Contract *BalanceLibCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceLibTransactorRaw struct {
	Contract *BalanceLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceLib creates a new instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLib(address common.Address, backend bind.ContractBackend) (*BalanceLib, error) {
	contract, err := bindBalanceLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceLib{BalanceLibCaller: BalanceLibCaller{contract: contract}, BalanceLibTransactor: BalanceLibTransactor{contract: contract}, BalanceLibFilterer: BalanceLibFilterer{contract: contract}}, nil
}

// NewBalanceLibCaller creates a new read-only instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLibCaller(address common.Address, caller bind.ContractCaller) (*BalanceLibCaller, error) {
	contract, err := bindBalanceLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceLibCaller{contract: contract}, nil
}

// NewBalanceLibTransactor creates a new write-only instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceLibTransactor, error) {
	contract, err := bindBalanceLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceLibTransactor{contract: contract}, nil
}

// NewBalanceLibFilterer creates a new log filterer instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceLibFilterer, error) {
	contract, err := bindBalanceLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceLibFilterer{contract: contract}, nil
}

// bindBalanceLib binds a generic wrapper to an already deployed contract.
func bindBalanceLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceLib *BalanceLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceLib.Contract.BalanceLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceLib *BalanceLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceLib.Contract.BalanceLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceLib *BalanceLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceLib.Contract.BalanceLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceLib *BalanceLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceLib *BalanceLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceLib *BalanceLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceLib.Contract.contract.Transact(opts, method, params...)
}

// BurnerMetaData contains all meta data concerning the Burner contract.
var BurnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20BurnableUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161027f38038061027f83398101604081905261002f91610067565b6001600160a01b0381166100565760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0316608052610097565b60006020828403121561007957600080fd5b81516001600160a01b038116811461009057600080fd5b9392505050565b6080516101c86100b760003960008181604a0152609d01526101c86000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806344df8e701461003b578063fc0c546a14610045575b600080fd5b610043610088565b005b61006c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b6040516370a0823160e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906342966c689082906370a082319060240160206040518083038186803b1580156100ef57600080fd5b505afa158015610103573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101279190610179565b6040518263ffffffff1660e01b815260040161014591815260200190565b600060405180830381600087803b15801561015f57600080fd5b505af1158015610173573d6000803e3d6000fd5b50505050565b60006020828403121561018b57600080fd5b505191905056fea2646970667358221220ca6e7e7252f18b975fabecfd6292b273ac99e25dadbd2c28b771fdc262e8fb9864736f6c63430008090033",
}

// BurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnerMetaData.ABI instead.
var BurnerABI = BurnerMetaData.ABI

// BurnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BurnerMetaData.Bin instead.
var BurnerBin = BurnerMetaData.Bin

// DeployBurner deploys a new Ethereum contract, binding an instance of Burner to it.
func DeployBurner(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Burner, error) {
	parsed, err := BurnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnerBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Burner{BurnerCaller: BurnerCaller{contract: contract}, BurnerTransactor: BurnerTransactor{contract: contract}, BurnerFilterer: BurnerFilterer{contract: contract}}, nil
}

// Burner is an auto generated Go binding around an Ethereum contract.
type Burner struct {
	BurnerCaller     // Read-only binding to the contract
	BurnerTransactor // Write-only binding to the contract
	BurnerFilterer   // Log filterer for contract events
}

// BurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnerSession struct {
	Contract     *Burner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnerCallerSession struct {
	Contract *BurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnerTransactorSession struct {
	Contract     *BurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnerRaw struct {
	Contract *Burner // Generic contract binding to access the raw methods on
}

// BurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnerCallerRaw struct {
	Contract *BurnerCaller // Generic read-only contract binding to access the raw methods on
}

// BurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnerTransactorRaw struct {
	Contract *BurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurner creates a new instance of Burner, bound to a specific deployed contract.
func NewBurner(address common.Address, backend bind.ContractBackend) (*Burner, error) {
	contract, err := bindBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Burner{BurnerCaller: BurnerCaller{contract: contract}, BurnerTransactor: BurnerTransactor{contract: contract}, BurnerFilterer: BurnerFilterer{contract: contract}}, nil
}

// NewBurnerCaller creates a new read-only instance of Burner, bound to a specific deployed contract.
func NewBurnerCaller(address common.Address, caller bind.ContractCaller) (*BurnerCaller, error) {
	contract, err := bindBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerCaller{contract: contract}, nil
}

// NewBurnerTransactor creates a new write-only instance of Burner, bound to a specific deployed contract.
func NewBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnerTransactor, error) {
	contract, err := bindBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerTransactor{contract: contract}, nil
}

// NewBurnerFilterer creates a new log filterer instance of Burner, bound to a specific deployed contract.
func NewBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnerFilterer, error) {
	contract, err := bindBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnerFilterer{contract: contract}, nil
}

// bindBurner binds a generic wrapper to an already deployed contract.
func bindBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burner *BurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Burner.Contract.BurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burner *BurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burner.Contract.BurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burner *BurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burner.Contract.BurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burner *BurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Burner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burner *BurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burner *BurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burner.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Burner *BurnerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Burner.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Burner *BurnerSession) Token() (common.Address, error) {
	return _Burner.Contract.Token(&_Burner.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Burner *BurnerCallerSession) Token() (common.Address, error) {
	return _Burner.Contract.Token(&_Burner.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_Burner *BurnerTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burner.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_Burner *BurnerSession) Burn() (*types.Transaction, error) {
	return _Burner.Contract.Burn(&_Burner.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_Burner *BurnerTransactorSession) Burn() (*types.Transaction, error) {
	return _Burner.Contract.Burn(&_Burner.TransactOpts)
}

// ExpressLaneAuctionMetaData contains all meta data concerning the ExpressLaneAuction contract.
var ExpressLaneAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionNotClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidsWrongOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"name\":\"FixedTransferor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountRequested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountRequested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalanceAcc\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newRound\",\"type\":\"uint64\"}],\"name\":\"InvalidNewRound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"currentStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newStart\",\"type\":\"uint64\"}],\"name\":\"InvalidNewStart\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NegativeOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"roundStart\",\"type\":\"int64\"}],\"name\":\"NegativeRoundStart\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotExpressLaneController\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"expectedTransferor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"NotTransferor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReserveBlackout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"ReservePriceNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReservePrice\",\"type\":\"uint256\"}],\"name\":\"ReservePriceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"RoundAlreadyResolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoundDurationTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"RoundNotResolved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"RoundTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"}],\"name\":\"RoundTooOld\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameBidder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TieBidsWrongOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalMaxRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAuctionClosingSeconds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBiddingToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isMultiBidAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceExpressLaneController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstPriceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundStartTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundEndTimestamp\",\"type\":\"uint64\"}],\"name\":\"AuctionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"SetBeneficiary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"}],\"name\":\"SetExpressLaneController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"SetMinReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"SetReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"name\":\"SetRoundTimingInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"name\":\"SetTransferor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundWithdrawable\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BENEFICIARY_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_TIMING_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"balanceOfAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiaryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushBeneficiaryBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBidHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_auctioneer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_biddingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"_roundTimingInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_minReservePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_auctioneerAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minReservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetterAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiarySetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_roundTimingSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterAdmin\",\"type\":\"address\"}],\"internalType\":\"structInitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAuctionRoundClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReserveBlackout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"secondPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveMultiBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveSingleBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolvedRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"roundTimestamps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundTimingInfo\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinReservePrice\",\"type\":\"uint256\"}],\"name\":\"setMinReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"newRoundTimingInfo\",\"type\":\"tuple\"}],\"name\":\"setRoundTimingInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"internalType\":\"structTransferor\",\"name\":\"transferor\",\"type\":\"tuple\"}],\"name\":\"setTransferor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"}],\"name\":\"transferExpressLaneController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"withdrawableBalanceAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051613ffc61003060003960006114220152613ffc6000f3fe608060405234801561001057600080fd5b50600436106102945760003560e01c80637b617f9411610167578063c5b6aa2f116100ce578063e2fc6f6811610087578063e2fc6f681461069f578063e3f7bb55146106a9578063e460d2c5146106d0578063e4d20c1d146106d8578063f698da25146106eb578063fed87be8146106f357600080fd5b8063c5b6aa2f1461063f578063ca15c87314610647578063ce9c7c0d1461065a578063cfe9232b1461066d578063d547741f14610682578063db2e1eed1461069557600080fd5b80639a1fadd3116101205780639a1fadd3146105e1578063a217fddf146105f4578063b3ee252f146105fc578063b51d1d4f14610611578063b6b55f2514610619578063bef0ec741461062c57600080fd5b80637b617f941461054957806383af0a1f1461057c5780638948cc4e146105865780638a19c8bc1461059b5780639010d07c146105bb57806391d14854146105ce57600080fd5b80632f2ff15d1161020b578063639d7566116101c4578063639d7566146104915780636a514beb146104a55780636ad72517146105085780636dc4fc4e146105105780636e8cace51461052357806370a082311461053657600080fd5b80632f2ff15d14610404578063336a5b5e1461041757806336568abe1461042c57806338af3eed1461043f578063447a709e1461046b5780635633c3371461047e57600080fd5b80630d253fbe1161025d5780630d253fbe1461037457806314d963161461038a5780631682e50b146103b15780631c31f710146103c6578063248a9ca3146103d95780632d668ce7146103fc57600080fd5b80627be2fe146102995780630152682d146102ae57806301ffc9a71461031d57806302b629381461034057806304c584ad14610361575b600080fd5b6102ac6102a73660046135c1565b610706565b005b610104546102e490600781900b906001600160401b03600160401b8204811691600160801b8104821691600160c01b9091041684565b6040805160079590950b85526001600160401b039384166020860152918316918401919091521660608201526080015b60405180910390f35b61033061032b3660046135fa565b610917565b6040519015158152602001610314565b61035361034e366004613624565b610942565b604051908152602001610314565b61035361036f366004613641565b6109ba565b61037c610a3a565b6040516103149291906136a3565b6103537f3fb9f0655b78e8eabe9e0f51d65db56c7690d4329012c3faf1fbd6d43f65826181565b610353600080516020613f4783398151915281565b6102ac6103d4366004613624565b610aea565b6103536103e73660046136be565b60009081526065602052604090206001015490565b610330610b6f565b6102ac6104123660046136d7565b610bca565b610353600080516020613f6783398151915281565b6102ac61043a3660046136d7565b610bf5565b61010054610453906001600160a01b031681565b6040516001600160a01b039091168152602001610314565b6102ac610479366004613714565b610c73565b61035361048c366004613777565b610e9a565b61010154610453906001600160a01b031681565b6104e16104b3366004613624565b610106602052600090815260409020546001600160a01b03811690600160a01b90046001600160401b031682565b604080516001600160a01b0390931683526001600160401b03909116602083015201610314565b6102ac610fac565b6102ac61051e3660046137a5565b610ff6565b610353610531366004613777565b61111a565b610353610544366004613624565b611200565b61055c6105573660046137e1565b611278565b604080516001600160401b03938416815292909116602083015201610314565b6103536101035481565b610353600080516020613fa783398151915281565b6105a36112da565b6040516001600160401b039091168152602001610314565b6104536105c93660046137fe565b611330565b6103306105dc3660046136d7565b611348565b6102ac6105ef366004613820565b611373565b610353600081565b610353600080516020613f2783398151915281565b6102ac611803565b6102ac6106273660046136be565b6118c9565b6102ac61063a366004613833565b611934565b6102ac611a7d565b6103536106553660046136be565b611b39565b6102ac6106683660046136be565b611b50565b610353600080516020613f8783398151915281565b6102ac6106903660046136d7565b611bfd565b6103536101025481565b6103536101055481565b6103537fa8131bb4589277d6866d942849029b416b39e61eb3969a32787130bbdd292a9681565b610330611c23565b6102ac6106e63660046136be565b611c97565b610353611d07565b6102ac610701366004613845565b611d11565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152600061075b82611e7a565b9050806001600160401b0316846001600160401b031610156107a857604051631cafa7eb60e11b81526001600160401b038086166004830152821660248201526044015b60405180910390fd5b60006107b560fe86611eb0565b80546001600160a01b039081166000818152610106602052604090205492935091168015610811576001600160a01b038116331461080c57868133604051633b10eca560e11b815260040161079f93929190613857565b610840565b6001600160a01b0382163314610840578682336040516333057b6960e11b815260040161079f93929190613857565b82546001600160a01b0319166001600160a01b038716178355600080610866878a611f33565b90925090506001600160a01b03831661087f5783610881565b825b6001600160a01b0316886001600160a01b0316856001600160a01b03167fb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b8c426001600160401b0316876001600160401b0316106108df57866108e1565b425b604080516001600160401b03938416815291831660208301529187168183015290519081900360600190a4505050505050505050565b60006001600160e01b03198216635a05180f60e01b148061093c575061093c82611fb2565b92915050565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b9004909116606082015260009061093c9061099c90611e7a565b6001600160a01b038416600090815260fd6020526040902090611fe7565b604080517f0358b2b705d5c5ef47651be44f418326852a390f3b4c933661a5f4f0d8fa1ee360208201526001600160401b038516918101919091526001600160a01b038316606082015260808101829052600090610a309060a00160405160208183030381529060405280519060200120612010565b90505b9392505050565b6040805180820190915260008082526020820152604080518082019091526000808252602082015260ff5460fe546001600160401b03600160a01b928390048116929091041611610a8e5760ff60fe610a93565b60fe60ff5b60408051808201825292546001600160a01b0380821685526001600160401b03600160a01b928390048116602080880191909152845180860190955294549182168452919004169181019190915290939092509050565b600080516020613f67833981519152610b03813361205e565b61010054604080516001600160a01b03928316815291841660208301527f8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6910160405180910390a15061010080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152600090610bc5906120c2565b905090565b600082815260656020526040902060010154610be6813361205e565b610bf0838361212d565b505050565b6001600160a01b0381163314610c655760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161079f565b610c6f828261214f565b5050565b600080516020613f87833981519152610c8c813361205e565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152610cdf816120c2565b610cfc5760405163b9adeefd60e01b815260040160405180910390fd5b826020013584602001351015610d255760405163a234cb1960e01b815260040160405180910390fd5b6101025483602001351015610d5f5761010254604051632b7cdbad60e11b815260208501356004820152602481019190915260440161079f565b6000610d6a82611e7a565b90506000610d798260016138ad565b9050600080610d90610d8a89613946565b84612171565b91509150600080610daa89610da490613946565b86612171565b91509150816001600160a01b0316846001600160a01b03161415610de15760405163f4a3e48560e01b815260040160405180910390fd5b88602001358a60200135148015610e4957508181604051602001610e069291906139fe565b60408051601f1981840301815290829052805160209182012091610e2e9187918791016139fe565b6040516020818303038152906040528051906020012060001c105b15610e67576040516348c2d05760e11b815260040160405180910390fd5b600080610e748988611f33565b91509150610e8c60018d888e602001358c878761224f565b505050505050505050505050565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152600090610ef090611e7a565b6001600160401b0316826001600160401b03161015610f8a576040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b900490911660608201528290610f5e90611e7a565b604051631cafa7eb60e11b81526001600160401b0392831660048201529116602482015260440161079f565b6001600160a01b038316600090815260fd60205260409020610a3390836123a5565b6101055480610fce57604051631f2a200560e01b815260040160405180910390fd5b6000610105556101005461010154610ff3916001600160a01b039182169116836123bd565b50565b600080516020613f8783398151915261100f813361205e565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152611062816120c2565b61107f5760405163b9adeefd60e01b815260040160405180910390fd5b61010254836020013510156110b95761010254604051632b7cdbad60e11b815260208501356004820152602481019190915260440161079f565b60006110c482611e7a565b905060006110d38260016138ad565b905060006110e96110e387613946565b83612171565b5090506000806110f98685611f33565b91509150611110600089856101025489878761224f565b5050505050505050565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b9004909116606082015260009061117090611e7a565b6001600160401b0316826001600160401b031610156111de576040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b900490911660608201528290610f5e90611e7a565b6001600160a01b038316600090815260fd60205260409020610a339083611fe7565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b9004909116606082015260009061093c9061125a90611e7a565b6001600160a01b038416600090815260fd60205260409020906123a5565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b9004909116606082015260009081906112d19084611f33565b91509150915091565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152600090610bc590611e7a565b6000828152609760205260408120610a339083612420565b60009182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600054610100900460ff1661138e5760005460ff1615611392565b303b155b6113f55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161079f565b600054610100900460ff16158015611417576000805461ffff19166101011790555b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156114a55760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b606482015260840161079f565b6114ad61242c565b6114fa6040518060400160405280601281526020017122bc383932b9b9a630b732a0bab1ba34b7b760711b815250604051806040016040528060018152602001603160f81b815250612455565b600061150c6040840160208501613624565b6001600160a01b0316141561153457604051633fb3c7af60e01b815260040160405180910390fd5b6115446040830160208401613624565b61010180546001600160a01b0319166001600160a01b03929092169190911790556115756060830160408401613624565b61010080546001600160a01b0319166001600160a01b03929092169190911790557f8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e660006115c96060850160408601613624565b604080516001600160a01b0393841681529290911660208301520160405180910390a160e0820135610103819055604080516000815260208101929092527f5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0910160405180910390a160e0820135610102819055604080516000815260208101929092527f9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794910160405180910390a1600061168a6080840160608501613a2f565b60070b12156116ac57604051630b7a36ff60e11b815260040160405180910390fd5b6116b882606001612486565b6116d460006116cf6101e085016101c08601613624565b61212d565b6116f8600080516020613fa78339815191526116cf61014085016101208601613624565b61171c600080516020613f678339815191526116cf6101a085016101808601613624565b611773600080516020613f8783398151915261173b6020850185613624565b7f3fb9f0655b78e8eabe9e0f51d65db56c7690d4329012c3faf1fbd6d43f65826161176e61012087016101008801613624565b612641565b6117ca600080516020613f2783398151915261179761016085016101408601613624565b7fa8131bb4589277d6866d942849029b416b39e61eb3969a32787130bbdd292a9661176e61018087016101608801613624565b6117ee600080516020613f478339815191526116cf6101c085016101a08601613624565b8015610c6f576000805461ff00191690555050565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b9004909116606082015260009061185990611e7a565b6118649060026138ad565b33600090815260fd6020526040902080549192506118829083612665565b604080518281526001600160401b038416602082015233917f31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401910160405180910390a25050565b33600090815260fd602052604090206118e29082612702565b610101546118fb906001600160a01b031633308461276f565b60405181815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c906020015b60405180910390a250565b3360009081526101066020526040902080546001600160a01b0316158015906119c557506040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b900490911660608201526119ab90611e7a565b81546001600160401b03918216600160a01b909104909116115b156119f7578054604051633aec4cf960e11b8152600160a01b9091046001600160401b0316600482015260240161079f565b336000908152610106602052604090208290611a138282613a4c565b50611a2390506020830183613624565b6001600160a01b0316337ff6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0611a5e60408601602087016137e1565b6040516001600160401b03909116815260200160405180910390a35050565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b90049091166060820152600090611aec90611ad790611e7a565b33600090815260fd60205260409020906127a7565b61010154909150611b07906001600160a01b031633836123bd565b60405181815233907f9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda890602001611929565b600081815260976020526040812061093c9061280b565b600080516020613f27833981519152611b69813361205e565b6000611b7560fe612815565b5080546040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b900482166060820152929350611bd69291600160a01b900416612858565b15611bf4576040516309e00d2f60e31b815260040160405180910390fd5b610bf08361290e565b600082815260656020526040902060010154611c19813361205e565b610bf0838361214f565b600080611c3060fe612815565b5080546040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b900482166060820152929350611c919291600160a01b900416612858565b91505090565b600080516020613fa7833981519152611cb0813361205e565b6101035460408051918252602082018490527f5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0910160405180910390a161010382905561010254821115610c6f57610c6f8261290e565b6000610bc5612984565b600080516020613f47833981519152611d2a813361205e565b6040805160808101825261010454600781900b82526001600160401b03600160401b820481166020840152600160801b8204811693830193909352600160c01b900490911660608201526000611d7f82611e7a565b90506000611d9a611d9536879003870187613aac565b611e7a565b9050806001600160401b0316826001600160401b031614611de1576040516368c18ca960e01b81526001600160401b0380841660048301528216602482015260440161079f565b6000611df8611df18460016138ad565b8590611f33565b5090506000611e20611e0b8460016138ad565b611e1a368a90038a018a613aac565b90611f33565b509050806001600160401b0316826001600160401b031614611e685760405163141c4d3b60e31b81526001600160401b0380841660048301528216602482015260440161079f565b611e7187612486565b50505050505050565b8051600090600790810b4290910b1215611e9657506000919050565b60208201518251611ea6906129ff565b61093c9190613b46565b60006001600160401b03821683820154600160a01b90046001600160401b03161415611ee2578260005b01905061093c565b6001600160401b0382168360010154600160a01b90046001600160401b03161415611f0f57826001611eda565b604051631f760a5b60e31b81526001600160401b038316600482015260240161079f565b6000806000838560200151611f489190613b6c565b8551611f549190613b9b565b905060008160070b1215611f815760405163f160ad7960e01b8152600782900b600482015260240161079f565b60208501518190600090600190611f9890846138ad565b611fa29190613bec565b91945090925050505b9250929050565b60006001600160e01b03198216637965db0b60e01b148061093c57506301ffc9a760e01b6001600160e01b031983161461093c565b60018201546000906001600160401b03908116908316101561200a576000610a33565b50505490565b600061093c61201d612984565b8360405161190160f01b6020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b6120688282611348565b610c6f57612080816001600160a01b03166014612a0b565b61208b836020612a0b565b60405160200161209c929190613c40565b60408051601f198184030181529082905262461bcd60e51b825261079f91600401613caf565b8051600090600790810b4290910b12156120de57506000919050565b60006120ed83600001516129ff565b905060008360200151826121019190613ce2565b9050836040015184602001516121179190613bec565b6001600160401b03908116911610159392505050565b6121378282612ba6565b6000828152609760205260409020610bf09082612c2c565b6121598282612c41565b6000828152609760205260409020610bf09082612ca8565b600080600061218984866000015187602001516109ba565b905060006121a4866040015183612cbd90919063ffffffff16565b905060006121b3600187613bec565b6020808901516001600160a01b038516600090815260fd9092526040909120919250906121e090836123a5565b1015612244576020808801516001600160a01b038416600090815260fd909252604090912083919061221290846123a5565b604051630dac930560e21b81526001600160a01b0390931660048401526024830191909152604482015260640161079f565b509590945092505050565b600061225c8460016138ad565b90506122788161226f60208a018a613624565b60fe9190612ce1565b6001600160a01b038616600090815260fd6020526040902061229b908686612db4565b8461010560008282546122ae9190613d08565b90915550600090506122c36020890189613624565b604080516001600160401b038581168252878116602083015286168183015290516001600160a01b0392909216916000917fb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b919081900360600190a461232c6020880188613624565b604080516001600160401b03848116825260208b8101359083015281830189905286811660608301528516608082015290516001600160a01b03928316928916918b1515917f7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d479181900360a00190a45050505050505050565b60006123b18383611fe7565b8354610a339190613d20565b6040516001600160a01b038316602482015260448101829052610bf090849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152612e15565b6000610a338383612ee7565b600054610100900460ff166124535760405162461bcd60e51b815260040161079f90613d37565b565b600054610100900460ff1661247c5760405162461bcd60e51b815260040161079f90613d37565b610c6f8282612f11565b61249660608201604083016137e1565b6001600160401b03166124bc5760405163023dd6a960e11b815260040160405180910390fd5b620151806124d060408301602084016137e1565b6001600160401b03161115612514576124ef60408201602083016137e1565b60405163c34a76cf60e01b81526001600160401b03909116600482015260240161079f565b61252460408201602083016137e1565b6001600160401b031661253d60608301604084016137e1565b61254d60808401606085016137e1565b61255791906138ad565b6001600160401b0316111561257f576040516301936f1b60e51b815260040160405180910390fd5b8061010461258d8282613d82565b507f982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd4790506125c3611d9536849003840184613aac565b6125d06020840184613a2f565b6125e060408501602086016137e1565b6125f060608601604087016137e1565b61260060808701606088016137e1565b604080516001600160401b03968716815260079590950b6020860152928516848401529084166060840152909216608082015290519081900360a00190a150565b61264b848461212d565b612655828261212d565b61265f8483612f52565b50505050565b815461268457604051631f2a200560e01b815260040160405180910390fd5b6001600160401b0381811614156126ae57604051631ec4eeef60e11b815260040160405180910390fd5b60018201546001600160401b03908116146126dc576040516304eb6b3f60e01b815260040160405180910390fd5b600191909101805467ffffffffffffffff19166001600160401b03909216919091179055565b8061272057604051631f2a200560e01b815260040160405180910390fd5b60018201546001600160401b03908116146127525760018201805467ffffffffffffffff19166001600160401b031790555b808260000160008282546127669190613d08565b90915550505050565b6040516001600160a01b038085166024830152831660448201526064810182905261265f9085906323b872dd60e01b906084016123e9565b60006001600160401b0382811614156127d357604051631ec4eeef60e11b815260040160405180910390fd5b60006127df8484611fe7565b9050806127ff57604051630686827b60e51b815260040160405180910390fd5b60008455905092915050565b600061093c825490565b6000808083810190506000846001015482546001600160401b03600160a01b9283900481169290910416101561284e5750506001808401905b9094909350915050565b8151600090600790810b4290910b12156128745750600061093c565b600061287f84611e7a565b905061288c8160016138ad565b6001600160401b0316836001600160401b0316106128ae57600091505061093c565b60006128bd85600001516129ff565b905060008560200151826128d19190613ce2565b90508560600151866040015187602001516128ec9190613bec565b6128f69190613bec565b6001600160401b039081169116101595945050505050565b610103548110156129415761010354604051636d27939760e11b815261079f918391600401918252602082015260400190565b6101025460408051918252602082018390527f9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794910160405180910390a161010255565b6000610bc57f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6129b360c95490565b60ca546040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b600061093c8242613e34565b60606000612a1a836002613e86565b612a25906002613d08565b6001600160401b03811115612a3c57612a3c6138d8565b6040519080825280601f01601f191660200182016040528015612a66576020820181803683370190505b509050600360fc1b81600081518110612a8157612a81613881565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110612ab057612ab0613881565b60200101906001600160f81b031916908160001a9053506000612ad4846002613e86565b612adf906001613d08565b90505b6001811115612b57576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110612b1357612b13613881565b1a60f81b828281518110612b2957612b29613881565b60200101906001600160f81b031916908160001a90535060049490941c93612b5081613ea5565b9050612ae2565b508315610a335760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161079f565b612bb08282611348565b610c6f5760008281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055612be83390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000610a33836001600160a01b038416612f9d565b612c4b8282611348565b15610c6f5760008281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6000610a33836001600160a01b038416612fec565b6000806000612ccc85856130df565b91509150612cd98161314c565b509392505050565b600080612ced85612815565b815491935091506001600160401b03808616600160a01b9092041610612d3157604051631147e1cd60e21b81526001600160401b038516600482015260240161079f565b604080518082019091526001600160a01b03841681526001600160401b038516602082015260018218908660ff831660028110612d7057612d70613881565b8251910180546020909301516001600160401b0316600160a01b026001600160e01b03199093166001600160a01b0390921691909117919091179055505050505050565b6000612dc084836123a5565b9050801580612dce57508281105b15612df65760405163cf47918160e01b8152600481018490526024810182905260440161079f565b82846000016000828254612e0a9190613d20565b909155505050505050565b6000612e6a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166133029092919063ffffffff16565b805190915015610bf05780806020019051810190612e889190613ebc565b610bf05760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161079f565b6000826000018281548110612efe57612efe613881565b9060005260206000200154905092915050565b600054610100900460ff16612f385760405162461bcd60e51b815260040161079f90613d37565b81516020928301208151919092012060c99190915560ca55565b600082815260656020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6000818152600183016020526040812054612fe45750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561093c565b50600061093c565b600081815260018301602052604081205480156130d5576000613010600183613d20565b855490915060009061302490600190613d20565b905081811461308957600086600001828154811061304457613044613881565b906000526020600020015490508087600001848154811061306757613067613881565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061309a5761309a613ede565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061093c565b600091505061093c565b6000808251604114156131165760208301516040840151606085015160001a61310a87828585613311565b94509450505050611fab565b82516040141561314057602083015160408401516131358683836133f4565b935093505050611fab565b50600090506002611fab565b600081600481111561316057613160613ef4565b14156131695750565b600181600481111561317d5761317d613ef4565b14156131c65760405162461bcd60e51b815260206004820152601860248201527745434453413a20696e76616c6964207369676e617475726560401b604482015260640161079f565b60028160048111156131da576131da613ef4565b14156132285760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161079f565b600381600481111561323c5761323c613ef4565b14156132955760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161079f565b60048160048111156132a9576132a9613ef4565b1415610ff35760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161079f565b6060610a30848460008561342d565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b0383111561333e57506000905060036133eb565b8460ff16601b1415801561335657508460ff16601c14155b1561336757506000905060046133eb565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156133bb573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166133e4576000600192509250506133eb565b9150600090505b94509492505050565b6000806001600160ff1b0383168161341160ff86901c601b613d08565b905061341f87828885613311565b935093505050935093915050565b60608247101561348e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161079f565b6001600160a01b0385163b6134e55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161079f565b600080866001600160a01b031685876040516135019190613f0a565b60006040518083038185875af1925050503d806000811461353e576040519150601f19603f3d011682016040523d82523d6000602084013e613543565b606091505b509150915061355382828661355e565b979650505050505050565b6060831561356d575081610a33565b82511561357d5782518084602001fd5b8160405162461bcd60e51b815260040161079f9190613caf565b6001600160401b0381168114610ff357600080fd5b6001600160a01b0381168114610ff357600080fd5b600080604083850312156135d457600080fd5b82356135df81613597565b915060208301356135ef816135ac565b809150509250929050565b60006020828403121561360c57600080fd5b81356001600160e01b031981168114610a3357600080fd5b60006020828403121561363657600080fd5b8135610a33816135ac565b60008060006060848603121561365657600080fd5b833561366181613597565b92506020840135613671816135ac565b929592945050506040919091013590565b80516001600160a01b031682526020908101516001600160401b0316910152565b608081016136b18285613682565b610a336040830184613682565b6000602082840312156136d057600080fd5b5035919050565b600080604083850312156136ea57600080fd5b8235915060208301356135ef816135ac565b60006060828403121561370e57600080fd5b50919050565b6000806040838503121561372757600080fd5b82356001600160401b038082111561373e57600080fd5b61374a868387016136fc565b9350602085013591508082111561376057600080fd5b5061376d858286016136fc565b9150509250929050565b6000806040838503121561378a57600080fd5b8235613795816135ac565b915060208301356135ef81613597565b6000602082840312156137b757600080fd5b81356001600160401b038111156137cd57600080fd5b6137d9848285016136fc565b949350505050565b6000602082840312156137f357600080fd5b8135610a3381613597565b6000806040838503121561381157600080fd5b50508035926020909101359150565b60006101e0828403121561370e57600080fd5b60006040828403121561370e57600080fd5b60006080828403121561370e57600080fd5b6001600160401b039390931683526001600160a01b03918216602084015216604082015260600190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160401b038083168185168083038211156138cf576138cf613897565b01949350505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715613910576139106138d8565b60405290565b604051601f8201601f191681016001600160401b038111828210171561393e5761393e6138d8565b604052919050565b60006060823603121561395857600080fd5b6139606138ee565b823561396b816135ac565b81526020838101358183015260408401356001600160401b038082111561399157600080fd5b9085019036601f8301126139a457600080fd5b8135818111156139b6576139b66138d8565b6139c8601f8201601f19168501613916565b915080825236848285010111156139de57600080fd5b808484018584013760009082019093019290925250604082015292915050565b60609290921b6bffffffffffffffffffffffff19168252601482015260340190565b8060070b8114610ff357600080fd5b600060208284031215613a4157600080fd5b8135610a3381613a20565b8135613a57816135ac565b81546001600160a01b031981166001600160a01b039290921691821783556020840135613a8381613597565b6001600160e01b03199190911690911760a09190911b67ffffffffffffffff60a01b1617905550565b600060808284031215613abe57600080fd5b604051608081018181106001600160401b0382111715613ae057613ae06138d8565b6040528235613aee81613a20565b81526020830135613afe81613597565b60208201526040830135613b1181613597565b60408201526060830135613b2481613597565b60608201529392505050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b0380841680613b6057613b60613b30565b92169190910492915050565b60006001600160401b0380831681851681830481118215151615613b9257613b92613897565b02949350505050565b60008160070b8360070b6000821282677fffffffffffffff03821381151615613bc657613bc6613897565b82677fffffffffffffff19038212811615613be357613be3613897565b50019392505050565b60006001600160401b0383811690831681811015613c0c57613c0c613897565b039392505050565b60005b83811015613c2f578181015183820152602001613c17565b8381111561265f5750506000910152565b76020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b815260008351613c72816017850160208801613c14565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351613ca3816028840160208801613c14565b01602801949350505050565b6020815260008251806020840152613cce816040850160208701613c14565b601f01601f19169190910160400192915050565b60006001600160401b0380841680613cfc57613cfc613b30565b92169190910692915050565b60008219821115613d1b57613d1b613897565b500190565b600082821015613d3257613d32613897565b500390565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b8135613d8d81613a20565b81546001600160401b038281166001600160401b031983161784556020850135613db681613597565b67ffffffffffffffff60401b604091821b16919093166001600160801b031983168117821785559285013590613deb82613597565b6001600160c01b03199283168417811760809290921b67ffffffffffffffff60801b169182178555606086013592613e2284613597565b93171760c09190911b90911617905550565b60008160070b8360070b6000811281677fffffffffffffff1901831281151615613e6057613e60613897565b81677fffffffffffffff018313811615613e7c57613e7c613897565b5090039392505050565b6000816000190483118215151615613ea057613ea0613897565b500290565b600081613eb457613eb4613897565b506000190190565b600060208284031215613ece57600080fd5b81518015158114610a3357600080fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b60008251613f1c818460208701613c14565b919091019291505056fe19e6f23df7275b48d1c33822c6ad041a743378552246ac819f578ae1d6709cf96d8dad7188c7ed005c55bf77fbf589583d8668b0dad30a9b9dd016321a5c256fc1b97c934675624ef2089089ac12ae8922988c11dc8a578dfbac10d9eecf47611d693f62a755e2b3c6494da41af454605b9006057cb3c79b6adda1378f2a50a7b07567e7223e21f7dce4c0a89131ce9c32d0d3484085f3f331dea8caef56d141a26469706673582212203ff054de2ed394c09d544b1ddffde4f5174ad5ce2ff5a513a7b970c46193484264736f6c63430008090033",
}

// ExpressLaneAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use ExpressLaneAuctionMetaData.ABI instead.
var ExpressLaneAuctionABI = ExpressLaneAuctionMetaData.ABI

// ExpressLaneAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExpressLaneAuctionMetaData.Bin instead.
var ExpressLaneAuctionBin = ExpressLaneAuctionMetaData.Bin

// DeployExpressLaneAuction deploys a new Ethereum contract, binding an instance of ExpressLaneAuction to it.
func DeployExpressLaneAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExpressLaneAuction, error) {
	parsed, err := ExpressLaneAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExpressLaneAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExpressLaneAuction{ExpressLaneAuctionCaller: ExpressLaneAuctionCaller{contract: contract}, ExpressLaneAuctionTransactor: ExpressLaneAuctionTransactor{contract: contract}, ExpressLaneAuctionFilterer: ExpressLaneAuctionFilterer{contract: contract}}, nil
}

// ExpressLaneAuction is an auto generated Go binding around an Ethereum contract.
type ExpressLaneAuction struct {
	ExpressLaneAuctionCaller     // Read-only binding to the contract
	ExpressLaneAuctionTransactor // Write-only binding to the contract
	ExpressLaneAuctionFilterer   // Log filterer for contract events
}

// ExpressLaneAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExpressLaneAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpressLaneAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExpressLaneAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpressLaneAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExpressLaneAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpressLaneAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExpressLaneAuctionSession struct {
	Contract     *ExpressLaneAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExpressLaneAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExpressLaneAuctionCallerSession struct {
	Contract *ExpressLaneAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExpressLaneAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExpressLaneAuctionTransactorSession struct {
	Contract     *ExpressLaneAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExpressLaneAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExpressLaneAuctionRaw struct {
	Contract *ExpressLaneAuction // Generic contract binding to access the raw methods on
}

// ExpressLaneAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExpressLaneAuctionCallerRaw struct {
	Contract *ExpressLaneAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// ExpressLaneAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExpressLaneAuctionTransactorRaw struct {
	Contract *ExpressLaneAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExpressLaneAuction creates a new instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuction(address common.Address, backend bind.ContractBackend) (*ExpressLaneAuction, error) {
	contract, err := bindExpressLaneAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuction{ExpressLaneAuctionCaller: ExpressLaneAuctionCaller{contract: contract}, ExpressLaneAuctionTransactor: ExpressLaneAuctionTransactor{contract: contract}, ExpressLaneAuctionFilterer: ExpressLaneAuctionFilterer{contract: contract}}, nil
}

// NewExpressLaneAuctionCaller creates a new read-only instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuctionCaller(address common.Address, caller bind.ContractCaller) (*ExpressLaneAuctionCaller, error) {
	contract, err := bindExpressLaneAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionCaller{contract: contract}, nil
}

// NewExpressLaneAuctionTransactor creates a new write-only instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*ExpressLaneAuctionTransactor, error) {
	contract, err := bindExpressLaneAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionTransactor{contract: contract}, nil
}

// NewExpressLaneAuctionFilterer creates a new log filterer instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*ExpressLaneAuctionFilterer, error) {
	contract, err := bindExpressLaneAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionFilterer{contract: contract}, nil
}

// bindExpressLaneAuction binds a generic wrapper to an already deployed contract.
func bindExpressLaneAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExpressLaneAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpressLaneAuction *ExpressLaneAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExpressLaneAuction.Contract.ExpressLaneAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpressLaneAuction *ExpressLaneAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ExpressLaneAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpressLaneAuction *ExpressLaneAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ExpressLaneAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpressLaneAuction *ExpressLaneAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExpressLaneAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.contract.Transact(opts, method, params...)
}

// AUCTIONEERADMINROLE is a free data retrieval call binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) AUCTIONEERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "AUCTIONEER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUCTIONEERADMINROLE is a free data retrieval call binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) AUCTIONEERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// AUCTIONEERADMINROLE is a free data retrieval call binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) AUCTIONEERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// AUCTIONEERROLE is a free data retrieval call binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) AUCTIONEERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "AUCTIONEER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUCTIONEERROLE is a free data retrieval call binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) AUCTIONEERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERROLE(&_ExpressLaneAuction.CallOpts)
}

// AUCTIONEERROLE is a free data retrieval call binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) AUCTIONEERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERROLE(&_ExpressLaneAuction.CallOpts)
}

// BENEFICIARYSETTERROLE is a free data retrieval call binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BENEFICIARYSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "BENEFICIARY_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BENEFICIARYSETTERROLE is a free data retrieval call binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BENEFICIARYSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// BENEFICIARYSETTERROLE is a free data retrieval call binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BENEFICIARYSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DEFAULTADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DEFAULTADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// MINRESERVESETTERROLE is a free data retrieval call binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) MINRESERVESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "MIN_RESERVE_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINRESERVESETTERROLE is a free data retrieval call binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) MINRESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// MINRESERVESETTERROLE is a free data retrieval call binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) MINRESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERADMINROLE is a free data retrieval call binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RESERVESETTERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "RESERVE_SETTER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESERVESETTERADMINROLE is a free data retrieval call binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RESERVESETTERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERADMINROLE is a free data retrieval call binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RESERVESETTERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERROLE is a free data retrieval call binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RESERVESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "RESERVE_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESERVESETTERROLE is a free data retrieval call binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERROLE is a free data retrieval call binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// ROUNDTIMINGSETTERROLE is a free data retrieval call binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) ROUNDTIMINGSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "ROUND_TIMING_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUNDTIMINGSETTERROLE is a free data retrieval call binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ROUNDTIMINGSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// ROUNDTIMINGSETTERROLE is a free data retrieval call binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) ROUNDTIMINGSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOf(&_ExpressLaneAuction.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOf(&_ExpressLaneAuction.CallOpts, account)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BalanceOfAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "balanceOfAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOfAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOfAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) Beneficiary() (common.Address, error) {
	return _ExpressLaneAuction.Contract.Beneficiary(&_ExpressLaneAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _ExpressLaneAuction.Contract.Beneficiary(&_ExpressLaneAuction.CallOpts)
}

// BeneficiaryBalance is a free data retrieval call binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BeneficiaryBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "beneficiaryBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeneficiaryBalance is a free data retrieval call binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BeneficiaryBalance() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BeneficiaryBalance(&_ExpressLaneAuction.CallOpts)
}

// BeneficiaryBalance is a free data retrieval call binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BeneficiaryBalance() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BeneficiaryBalance(&_ExpressLaneAuction.CallOpts)
}

// BiddingToken is a free data retrieval call binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BiddingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "biddingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BiddingToken is a free data retrieval call binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BiddingToken() (common.Address, error) {
	return _ExpressLaneAuction.Contract.BiddingToken(&_ExpressLaneAuction.CallOpts)
}

// BiddingToken is a free data retrieval call binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BiddingToken() (common.Address, error) {
	return _ExpressLaneAuction.Contract.BiddingToken(&_ExpressLaneAuction.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) CurrentRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "currentRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) CurrentRound() (uint64, error) {
	return _ExpressLaneAuction.Contract.CurrentRound(&_ExpressLaneAuction.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) CurrentRound() (uint64, error) {
	return _ExpressLaneAuction.Contract.CurrentRound(&_ExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) DomainSeparator() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DomainSeparator(&_ExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) DomainSeparator() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DomainSeparator(&_ExpressLaneAuction.CallOpts)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetBidHash(opts *bind.CallOpts, round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getBidHash", round, expressLaneController, amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetBidHash(&_ExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetBidHash(&_ExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetRoleAdmin(&_ExpressLaneAuction.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetRoleAdmin(&_ExpressLaneAuction.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ExpressLaneAuction.Contract.GetRoleMember(&_ExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ExpressLaneAuction.Contract.GetRoleMember(&_ExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.GetRoleMemberCount(&_ExpressLaneAuction.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.GetRoleMemberCount(&_ExpressLaneAuction.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExpressLaneAuction.Contract.HasRole(&_ExpressLaneAuction.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExpressLaneAuction.Contract.HasRole(&_ExpressLaneAuction.CallOpts, role, account)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) IsAuctionRoundClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "isAuctionRoundClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) IsAuctionRoundClosed() (bool, error) {
	return _ExpressLaneAuction.Contract.IsAuctionRoundClosed(&_ExpressLaneAuction.CallOpts)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) IsAuctionRoundClosed() (bool, error) {
	return _ExpressLaneAuction.Contract.IsAuctionRoundClosed(&_ExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) IsReserveBlackout(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "isReserveBlackout")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) IsReserveBlackout() (bool, error) {
	return _ExpressLaneAuction.Contract.IsReserveBlackout(&_ExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) IsReserveBlackout() (bool, error) {
	return _ExpressLaneAuction.Contract.IsReserveBlackout(&_ExpressLaneAuction.CallOpts)
}

// MinReservePrice is a free data retrieval call binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) MinReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "minReservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReservePrice is a free data retrieval call binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) MinReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.MinReservePrice(&_ExpressLaneAuction.CallOpts)
}

// MinReservePrice is a free data retrieval call binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) MinReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.MinReservePrice(&_ExpressLaneAuction.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.ReservePrice(&_ExpressLaneAuction.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) ReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.ReservePrice(&_ExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) ResolvedRounds(opts *bind.CallOpts) (ELCRound, ELCRound, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "resolvedRounds")

	if err != nil {
		return *new(ELCRound), *new(ELCRound), err
	}

	out0 := *abi.ConvertType(out[0], new(ELCRound)).(*ELCRound)
	out1 := *abi.ConvertType(out[1], new(ELCRound)).(*ELCRound)

	return out0, out1, err

}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _ExpressLaneAuction.Contract.ResolvedRounds(&_ExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _ExpressLaneAuction.Contract.ResolvedRounds(&_ExpressLaneAuction.CallOpts)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64, uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RoundTimestamps(opts *bind.CallOpts, round uint64) (uint64, uint64, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "roundTimestamps", round)

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64, uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RoundTimestamps(round uint64) (uint64, uint64, error) {
	return _ExpressLaneAuction.Contract.RoundTimestamps(&_ExpressLaneAuction.CallOpts, round)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64, uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RoundTimestamps(round uint64) (uint64, uint64, error) {
	return _ExpressLaneAuction.Contract.RoundTimestamps(&_ExpressLaneAuction.CallOpts, round)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RoundTimingInfo(opts *bind.CallOpts) (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "roundTimingInfo")

	outstruct := new(struct {
		OffsetTimestamp          int64
		RoundDurationSeconds     uint64
		AuctionClosingSeconds    uint64
		ReserveSubmissionSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OffsetTimestamp = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.RoundDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AuctionClosingSeconds = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ReserveSubmissionSeconds = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _ExpressLaneAuction.Contract.RoundTimingInfo(&_ExpressLaneAuction.CallOpts)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _ExpressLaneAuction.Contract.RoundTimingInfo(&_ExpressLaneAuction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExpressLaneAuction.Contract.SupportsInterface(&_ExpressLaneAuction.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExpressLaneAuction.Contract.SupportsInterface(&_ExpressLaneAuction.CallOpts, interfaceId)
}

// TransferorOf is a free data retrieval call binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address ) view returns(address addr, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) TransferorOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr            common.Address
	FixedUntilRound uint64
}, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "transferorOf", arg0)

	outstruct := new(struct {
		Addr            common.Address
		FixedUntilRound uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FixedUntilRound = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// TransferorOf is a free data retrieval call binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address ) view returns(address addr, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) TransferorOf(arg0 common.Address) (struct {
	Addr            common.Address
	FixedUntilRound uint64
}, error) {
	return _ExpressLaneAuction.Contract.TransferorOf(&_ExpressLaneAuction.CallOpts, arg0)
}

// TransferorOf is a free data retrieval call binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address ) view returns(address addr, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) TransferorOf(arg0 common.Address) (struct {
	Addr            common.Address
	FixedUntilRound uint64
}, error) {
	return _ExpressLaneAuction.Contract.TransferorOf(&_ExpressLaneAuction.CallOpts, arg0)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) WithdrawableBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalance(&_ExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalance(&_ExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) WithdrawableBalanceAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalanceAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Deposit(&_ExpressLaneAuction.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Deposit(&_ExpressLaneAuction.TransactOpts, amount)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) FinalizeWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "finalizeWithdrawal")
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FinalizeWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FinalizeWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) FlushBeneficiaryBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "flushBeneficiaryBalance")
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_ExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_ExpressLaneAuction.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.GrantRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.GrantRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) Initialize(opts *bind.TransactOpts, args InitArgs) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "initialize", args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Initialize(&_ExpressLaneAuction.TransactOpts, args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Initialize(&_ExpressLaneAuction.TransactOpts, args)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) InitiateWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "initiateWithdrawal")
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.InitiateWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.InitiateWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RenounceRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RenounceRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) ResolveMultiBidAuction(opts *bind.TransactOpts, firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "resolveMultiBidAuction", firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveMultiBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveMultiBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) ResolveSingleBidAuction(opts *bind.TransactOpts, firstPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "resolveSingleBidAuction", firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveSingleBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveSingleBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RevokeRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RevokeRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setBeneficiary", newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetBeneficiary(&_ExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetBeneficiary(&_ExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetMinReservePrice(opts *bind.TransactOpts, newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setMinReservePrice", newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetMinReservePrice(&_ExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetMinReservePrice(&_ExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetReservePrice(opts *bind.TransactOpts, newReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setReservePrice", newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetReservePrice(&_ExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetReservePrice(&_ExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetRoundTimingInfo(opts *bind.TransactOpts, newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setRoundTimingInfo", newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetRoundTimingInfo(&_ExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetRoundTimingInfo(&_ExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetTransferor(opts *bind.TransactOpts, transferor Transferor) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setTransferor", transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetTransferor(&_ExpressLaneAuction.TransactOpts, transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetTransferor(&_ExpressLaneAuction.TransactOpts, transferor)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) TransferExpressLaneController(opts *bind.TransactOpts, round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "transferExpressLaneController", round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.TransferExpressLaneController(&_ExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.TransferExpressLaneController(&_ExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// ExpressLaneAuctionAuctionResolvedIterator is returned from FilterAuctionResolved and is used to iterate over the raw logs and unpacked data for AuctionResolved events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionAuctionResolvedIterator struct {
	Event *ExpressLaneAuctionAuctionResolved // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionAuctionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionAuctionResolved)
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
		it.Event = new(ExpressLaneAuctionAuctionResolved)
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
func (it *ExpressLaneAuctionAuctionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionAuctionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionAuctionResolved represents a AuctionResolved event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionAuctionResolved struct {
	IsMultiBidAuction               bool
	Round                           uint64
	FirstPriceBidder                common.Address
	FirstPriceExpressLaneController common.Address
	FirstPriceAmount                *big.Int
	Price                           *big.Int
	RoundStartTimestamp             uint64
	RoundEndTimestamp               uint64
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterAuctionResolved is a free log retrieval operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterAuctionResolved(opts *bind.FilterOpts, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (*ExpressLaneAuctionAuctionResolvedIterator, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionAuctionResolvedIterator{contract: _ExpressLaneAuction.contract, event: "AuctionResolved", logs: logs, sub: sub}, nil
}

// WatchAuctionResolved is a free log subscription operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchAuctionResolved(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionAuctionResolved, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (event.Subscription, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionAuctionResolved)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
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

// ParseAuctionResolved is a log parse operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseAuctionResolved(log types.Log) (*ExpressLaneAuctionAuctionResolved, error) {
	event := new(ExpressLaneAuctionAuctionResolved)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionDepositIterator struct {
	Event *ExpressLaneAuctionDeposit // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionDeposit)
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
		it.Event = new(ExpressLaneAuctionDeposit)
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
func (it *ExpressLaneAuctionDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionDeposit represents a Deposit event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*ExpressLaneAuctionDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionDepositIterator{contract: _ExpressLaneAuction.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionDeposit)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseDeposit(log types.Log) (*ExpressLaneAuctionDeposit, error) {
	event := new(ExpressLaneAuctionDeposit)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleAdminChangedIterator struct {
	Event *ExpressLaneAuctionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionRoleAdminChanged)
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
		it.Event = new(ExpressLaneAuctionRoleAdminChanged)
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
func (it *ExpressLaneAuctionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionRoleAdminChanged represents a RoleAdminChanged event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ExpressLaneAuctionRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionRoleAdminChangedIterator{contract: _ExpressLaneAuction.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionRoleAdminChanged)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseRoleAdminChanged(log types.Log) (*ExpressLaneAuctionRoleAdminChanged, error) {
	event := new(ExpressLaneAuctionRoleAdminChanged)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleGrantedIterator struct {
	Event *ExpressLaneAuctionRoleGranted // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionRoleGranted)
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
		it.Event = new(ExpressLaneAuctionRoleGranted)
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
func (it *ExpressLaneAuctionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionRoleGranted represents a RoleGranted event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExpressLaneAuctionRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionRoleGrantedIterator{contract: _ExpressLaneAuction.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionRoleGranted)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseRoleGranted(log types.Log) (*ExpressLaneAuctionRoleGranted, error) {
	event := new(ExpressLaneAuctionRoleGranted)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleRevokedIterator struct {
	Event *ExpressLaneAuctionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionRoleRevoked)
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
		it.Event = new(ExpressLaneAuctionRoleRevoked)
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
func (it *ExpressLaneAuctionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionRoleRevoked represents a RoleRevoked event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExpressLaneAuctionRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionRoleRevokedIterator{contract: _ExpressLaneAuction.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionRoleRevoked)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseRoleRevoked(log types.Log) (*ExpressLaneAuctionRoleRevoked, error) {
	event := new(ExpressLaneAuctionRoleRevoked)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetBeneficiaryIterator is returned from FilterSetBeneficiary and is used to iterate over the raw logs and unpacked data for SetBeneficiary events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetBeneficiaryIterator struct {
	Event *ExpressLaneAuctionSetBeneficiary // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetBeneficiaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetBeneficiary)
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
		it.Event = new(ExpressLaneAuctionSetBeneficiary)
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
func (it *ExpressLaneAuctionSetBeneficiaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetBeneficiaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetBeneficiary represents a SetBeneficiary event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetBeneficiary struct {
	OldBeneficiary common.Address
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetBeneficiary is a free log retrieval operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetBeneficiary(opts *bind.FilterOpts) (*ExpressLaneAuctionSetBeneficiaryIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetBeneficiaryIterator{contract: _ExpressLaneAuction.contract, event: "SetBeneficiary", logs: logs, sub: sub}, nil
}

// WatchSetBeneficiary is a free log subscription operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetBeneficiary(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetBeneficiary) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetBeneficiary)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
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

// ParseSetBeneficiary is a log parse operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetBeneficiary(log types.Log) (*ExpressLaneAuctionSetBeneficiary, error) {
	event := new(ExpressLaneAuctionSetBeneficiary)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetExpressLaneControllerIterator is returned from FilterSetExpressLaneController and is used to iterate over the raw logs and unpacked data for SetExpressLaneController events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetExpressLaneControllerIterator struct {
	Event *ExpressLaneAuctionSetExpressLaneController // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetExpressLaneControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetExpressLaneController)
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
		it.Event = new(ExpressLaneAuctionSetExpressLaneController)
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
func (it *ExpressLaneAuctionSetExpressLaneControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetExpressLaneControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetExpressLaneController represents a SetExpressLaneController event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetExpressLaneController struct {
	Round                         uint64
	PreviousExpressLaneController common.Address
	NewExpressLaneController      common.Address
	Transferor                    common.Address
	StartTimestamp                uint64
	EndTimestamp                  uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetExpressLaneController is a free log retrieval operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetExpressLaneController(opts *bind.FilterOpts, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (*ExpressLaneAuctionSetExpressLaneControllerIterator, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetExpressLaneControllerIterator{contract: _ExpressLaneAuction.contract, event: "SetExpressLaneController", logs: logs, sub: sub}, nil
}

// WatchSetExpressLaneController is a free log subscription operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetExpressLaneController(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetExpressLaneController, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetExpressLaneController)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
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

// ParseSetExpressLaneController is a log parse operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetExpressLaneController(log types.Log) (*ExpressLaneAuctionSetExpressLaneController, error) {
	event := new(ExpressLaneAuctionSetExpressLaneController)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetMinReservePriceIterator is returned from FilterSetMinReservePrice and is used to iterate over the raw logs and unpacked data for SetMinReservePrice events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetMinReservePriceIterator struct {
	Event *ExpressLaneAuctionSetMinReservePrice // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetMinReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetMinReservePrice)
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
		it.Event = new(ExpressLaneAuctionSetMinReservePrice)
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
func (it *ExpressLaneAuctionSetMinReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetMinReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetMinReservePrice represents a SetMinReservePrice event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetMinReservePrice struct {
	OldPrice *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMinReservePrice is a free log retrieval operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetMinReservePrice(opts *bind.FilterOpts) (*ExpressLaneAuctionSetMinReservePriceIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetMinReservePriceIterator{contract: _ExpressLaneAuction.contract, event: "SetMinReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetMinReservePrice is a free log subscription operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetMinReservePrice(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetMinReservePrice) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetMinReservePrice)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
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

// ParseSetMinReservePrice is a log parse operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetMinReservePrice(log types.Log) (*ExpressLaneAuctionSetMinReservePrice, error) {
	event := new(ExpressLaneAuctionSetMinReservePrice)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetReservePriceIterator is returned from FilterSetReservePrice and is used to iterate over the raw logs and unpacked data for SetReservePrice events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetReservePriceIterator struct {
	Event *ExpressLaneAuctionSetReservePrice // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetReservePrice)
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
		it.Event = new(ExpressLaneAuctionSetReservePrice)
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
func (it *ExpressLaneAuctionSetReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetReservePrice represents a SetReservePrice event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetReservePrice struct {
	OldReservePrice *big.Int
	NewReservePrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReservePrice is a free log retrieval operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetReservePrice(opts *bind.FilterOpts) (*ExpressLaneAuctionSetReservePriceIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetReservePriceIterator{contract: _ExpressLaneAuction.contract, event: "SetReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetReservePrice is a free log subscription operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetReservePrice(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetReservePrice) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetReservePrice)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
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

// ParseSetReservePrice is a log parse operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetReservePrice(log types.Log) (*ExpressLaneAuctionSetReservePrice, error) {
	event := new(ExpressLaneAuctionSetReservePrice)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetRoundTimingInfoIterator is returned from FilterSetRoundTimingInfo and is used to iterate over the raw logs and unpacked data for SetRoundTimingInfo events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetRoundTimingInfoIterator struct {
	Event *ExpressLaneAuctionSetRoundTimingInfo // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetRoundTimingInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetRoundTimingInfo)
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
		it.Event = new(ExpressLaneAuctionSetRoundTimingInfo)
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
func (it *ExpressLaneAuctionSetRoundTimingInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetRoundTimingInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetRoundTimingInfo represents a SetRoundTimingInfo event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetRoundTimingInfo struct {
	CurrentRound             uint64
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetRoundTimingInfo is a free log retrieval operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetRoundTimingInfo(opts *bind.FilterOpts) (*ExpressLaneAuctionSetRoundTimingInfoIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetRoundTimingInfoIterator{contract: _ExpressLaneAuction.contract, event: "SetRoundTimingInfo", logs: logs, sub: sub}, nil
}

// WatchSetRoundTimingInfo is a free log subscription operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetRoundTimingInfo(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetRoundTimingInfo) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetRoundTimingInfo)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
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

// ParseSetRoundTimingInfo is a log parse operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetRoundTimingInfo(log types.Log) (*ExpressLaneAuctionSetRoundTimingInfo, error) {
	event := new(ExpressLaneAuctionSetRoundTimingInfo)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetTransferorIterator is returned from FilterSetTransferor and is used to iterate over the raw logs and unpacked data for SetTransferor events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetTransferorIterator struct {
	Event *ExpressLaneAuctionSetTransferor // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetTransferorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetTransferor)
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
		it.Event = new(ExpressLaneAuctionSetTransferor)
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
func (it *ExpressLaneAuctionSetTransferorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetTransferorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetTransferor represents a SetTransferor event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetTransferor struct {
	ExpressLaneController common.Address
	Transferor            common.Address
	FixedUntilRound       uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetTransferor is a free log retrieval operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetTransferor(opts *bind.FilterOpts, expressLaneController []common.Address, transferor []common.Address) (*ExpressLaneAuctionSetTransferorIterator, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetTransferorIterator{contract: _ExpressLaneAuction.contract, event: "SetTransferor", logs: logs, sub: sub}, nil
}

// WatchSetTransferor is a free log subscription operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetTransferor(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetTransferor, expressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetTransferor)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
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

// ParseSetTransferor is a log parse operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetTransferor(log types.Log) (*ExpressLaneAuctionSetTransferor, error) {
	event := new(ExpressLaneAuctionSetTransferor)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalFinalizedIterator struct {
	Event *ExpressLaneAuctionWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionWithdrawalFinalized)
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
		it.Event = new(ExpressLaneAuctionWithdrawalFinalized)
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
func (it *ExpressLaneAuctionWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionWithdrawalFinalized represents a WithdrawalFinalized event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalFinalized struct {
	Account          common.Address
	WithdrawalAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, account []common.Address) (*ExpressLaneAuctionWithdrawalFinalizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionWithdrawalFinalizedIterator{contract: _ExpressLaneAuction.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionWithdrawalFinalized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionWithdrawalFinalized)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseWithdrawalFinalized(log types.Log) (*ExpressLaneAuctionWithdrawalFinalized, error) {
	event := new(ExpressLaneAuctionWithdrawalFinalized)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalInitiatedIterator struct {
	Event *ExpressLaneAuctionWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionWithdrawalInitiated)
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
		it.Event = new(ExpressLaneAuctionWithdrawalInitiated)
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
func (it *ExpressLaneAuctionWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionWithdrawalInitiated represents a WithdrawalInitiated event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalInitiated struct {
	Account           common.Address
	WithdrawalAmount  *big.Int
	RoundWithdrawable *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, account []common.Address) (*ExpressLaneAuctionWithdrawalInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionWithdrawalInitiatedIterator{contract: _ExpressLaneAuction.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionWithdrawalInitiated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionWithdrawalInitiated)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseWithdrawalInitiated(log types.Log) (*ExpressLaneAuctionWithdrawalInitiated, error) {
	event := new(ExpressLaneAuctionWithdrawalInitiated)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionMetaData contains all meta data concerning the IExpressLaneAuction contract.
var IExpressLaneAuctionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isMultiBidAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceExpressLaneController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstPriceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundStartTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundEndTimestamp\",\"type\":\"uint64\"}],\"name\":\"AuctionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"SetBeneficiary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"}],\"name\":\"SetExpressLaneController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"SetMinReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"SetReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"name\":\"SetRoundTimingInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"name\":\"SetTransferor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundWithdrawable\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BENEFICIARY_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_TIMING_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"balanceOfAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiaryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushBeneficiaryBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBidHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_auctioneer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_biddingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"_roundTimingInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_minReservePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_auctioneerAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minReservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetterAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiarySetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_roundTimingSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterAdmin\",\"type\":\"address\"}],\"internalType\":\"structInitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAuctionRoundClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReserveBlackout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"secondPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveMultiBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveSingleBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolvedRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"roundTimestamps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundTimingInfo\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinReservePrice\",\"type\":\"uint256\"}],\"name\":\"setMinReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"newRoundTimingInfo\",\"type\":\"tuple\"}],\"name\":\"setRoundTimingInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"internalType\":\"structTransferor\",\"name\":\"transferor\",\"type\":\"tuple\"}],\"name\":\"setTransferor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"}],\"name\":\"transferExpressLaneController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"}],\"name\":\"transferorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntil\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"withdrawableBalanceAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IExpressLaneAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IExpressLaneAuctionMetaData.ABI instead.
var IExpressLaneAuctionABI = IExpressLaneAuctionMetaData.ABI

// IExpressLaneAuction is an auto generated Go binding around an Ethereum contract.
type IExpressLaneAuction struct {
	IExpressLaneAuctionCaller     // Read-only binding to the contract
	IExpressLaneAuctionTransactor // Write-only binding to the contract
	IExpressLaneAuctionFilterer   // Log filterer for contract events
}

// IExpressLaneAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExpressLaneAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExpressLaneAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExpressLaneAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExpressLaneAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExpressLaneAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExpressLaneAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExpressLaneAuctionSession struct {
	Contract     *IExpressLaneAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IExpressLaneAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExpressLaneAuctionCallerSession struct {
	Contract *IExpressLaneAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IExpressLaneAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExpressLaneAuctionTransactorSession struct {
	Contract     *IExpressLaneAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IExpressLaneAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExpressLaneAuctionRaw struct {
	Contract *IExpressLaneAuction // Generic contract binding to access the raw methods on
}

// IExpressLaneAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExpressLaneAuctionCallerRaw struct {
	Contract *IExpressLaneAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IExpressLaneAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExpressLaneAuctionTransactorRaw struct {
	Contract *IExpressLaneAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExpressLaneAuction creates a new instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuction(address common.Address, backend bind.ContractBackend) (*IExpressLaneAuction, error) {
	contract, err := bindIExpressLaneAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuction{IExpressLaneAuctionCaller: IExpressLaneAuctionCaller{contract: contract}, IExpressLaneAuctionTransactor: IExpressLaneAuctionTransactor{contract: contract}, IExpressLaneAuctionFilterer: IExpressLaneAuctionFilterer{contract: contract}}, nil
}

// NewIExpressLaneAuctionCaller creates a new read-only instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuctionCaller(address common.Address, caller bind.ContractCaller) (*IExpressLaneAuctionCaller, error) {
	contract, err := bindIExpressLaneAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionCaller{contract: contract}, nil
}

// NewIExpressLaneAuctionTransactor creates a new write-only instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IExpressLaneAuctionTransactor, error) {
	contract, err := bindIExpressLaneAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionTransactor{contract: contract}, nil
}

// NewIExpressLaneAuctionFilterer creates a new log filterer instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IExpressLaneAuctionFilterer, error) {
	contract, err := bindIExpressLaneAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionFilterer{contract: contract}, nil
}

// bindIExpressLaneAuction binds a generic wrapper to an already deployed contract.
func bindIExpressLaneAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExpressLaneAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExpressLaneAuction *IExpressLaneAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExpressLaneAuction.Contract.IExpressLaneAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExpressLaneAuction *IExpressLaneAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.IExpressLaneAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExpressLaneAuction *IExpressLaneAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.IExpressLaneAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExpressLaneAuction *IExpressLaneAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExpressLaneAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOf(&_IExpressLaneAuction.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOf(&_IExpressLaneAuction.CallOpts, account)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) BalanceOfAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "balanceOfAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOfAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOfAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) CurrentRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "currentRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) CurrentRound() (uint64, error) {
	return _IExpressLaneAuction.Contract.CurrentRound(&_IExpressLaneAuction.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) CurrentRound() (uint64, error) {
	return _IExpressLaneAuction.Contract.CurrentRound(&_IExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) DomainSeparator() ([32]byte, error) {
	return _IExpressLaneAuction.Contract.DomainSeparator(&_IExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) DomainSeparator() ([32]byte, error) {
	return _IExpressLaneAuction.Contract.DomainSeparator(&_IExpressLaneAuction.CallOpts)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetBidHash(opts *bind.CallOpts, round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getBidHash", round, expressLaneController, amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetBidHash(&_IExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetBidHash(&_IExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetRoleAdmin(&_IExpressLaneAuction.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetRoleAdmin(&_IExpressLaneAuction.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IExpressLaneAuction.Contract.GetRoleMember(&_IExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IExpressLaneAuction.Contract.GetRoleMember(&_IExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.GetRoleMemberCount(&_IExpressLaneAuction.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.GetRoleMemberCount(&_IExpressLaneAuction.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IExpressLaneAuction.Contract.HasRole(&_IExpressLaneAuction.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IExpressLaneAuction.Contract.HasRole(&_IExpressLaneAuction.CallOpts, role, account)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) IsAuctionRoundClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "isAuctionRoundClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) IsAuctionRoundClosed() (bool, error) {
	return _IExpressLaneAuction.Contract.IsAuctionRoundClosed(&_IExpressLaneAuction.CallOpts)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) IsAuctionRoundClosed() (bool, error) {
	return _IExpressLaneAuction.Contract.IsAuctionRoundClosed(&_IExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) IsReserveBlackout(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "isReserveBlackout")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) IsReserveBlackout() (bool, error) {
	return _IExpressLaneAuction.Contract.IsReserveBlackout(&_IExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) IsReserveBlackout() (bool, error) {
	return _IExpressLaneAuction.Contract.IsReserveBlackout(&_IExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) ResolvedRounds(opts *bind.CallOpts) (ELCRound, ELCRound, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "resolvedRounds")

	if err != nil {
		return *new(ELCRound), *new(ELCRound), err
	}

	out0 := *abi.ConvertType(out[0], new(ELCRound)).(*ELCRound)
	out1 := *abi.ConvertType(out[1], new(ELCRound)).(*ELCRound)

	return out0, out1, err

}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _IExpressLaneAuction.Contract.ResolvedRounds(&_IExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _IExpressLaneAuction.Contract.ResolvedRounds(&_IExpressLaneAuction.CallOpts)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64 start, uint64 end)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) RoundTimestamps(opts *bind.CallOpts, round uint64) (struct {
	Start uint64
	End   uint64
}, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "roundTimestamps", round)

	outstruct := new(struct {
		Start uint64
		End   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.End = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64 start, uint64 end)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RoundTimestamps(round uint64) (struct {
	Start uint64
	End   uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimestamps(&_IExpressLaneAuction.CallOpts, round)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64 start, uint64 end)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) RoundTimestamps(round uint64) (struct {
	Start uint64
	End   uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimestamps(&_IExpressLaneAuction.CallOpts, round)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) RoundTimingInfo(opts *bind.CallOpts) (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "roundTimingInfo")

	outstruct := new(struct {
		OffsetTimestamp          int64
		RoundDurationSeconds     uint64
		AuctionClosingSeconds    uint64
		ReserveSubmissionSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OffsetTimestamp = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.RoundDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AuctionClosingSeconds = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ReserveSubmissionSeconds = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimingInfo(&_IExpressLaneAuction.CallOpts)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimingInfo(&_IExpressLaneAuction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IExpressLaneAuction.Contract.SupportsInterface(&_IExpressLaneAuction.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IExpressLaneAuction.Contract.SupportsInterface(&_IExpressLaneAuction.CallOpts, interfaceId)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) WithdrawableBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalance(&_IExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalance(&_IExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) WithdrawableBalanceAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalanceAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// AUCTIONEERADMINROLE is a paid mutator transaction binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) AUCTIONEERADMINROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "AUCTIONEER_ADMIN_ROLE")
}

// AUCTIONEERADMINROLE is a paid mutator transaction binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) AUCTIONEERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// AUCTIONEERADMINROLE is a paid mutator transaction binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) AUCTIONEERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// AUCTIONEERROLE is a paid mutator transaction binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) AUCTIONEERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "AUCTIONEER_ROLE")
}

// AUCTIONEERROLE is a paid mutator transaction binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) AUCTIONEERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERROLE(&_IExpressLaneAuction.TransactOpts)
}

// AUCTIONEERROLE is a paid mutator transaction binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) AUCTIONEERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERROLE(&_IExpressLaneAuction.TransactOpts)
}

// BENEFICIARYSETTERROLE is a paid mutator transaction binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) BENEFICIARYSETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "BENEFICIARY_SETTER_ROLE")
}

// BENEFICIARYSETTERROLE is a paid mutator transaction binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BENEFICIARYSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// BENEFICIARYSETTERROLE is a paid mutator transaction binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) BENEFICIARYSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// MINRESERVESETTERROLE is a paid mutator transaction binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) MINRESERVESETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "MIN_RESERVE_SETTER_ROLE")
}

// MINRESERVESETTERROLE is a paid mutator transaction binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) MINRESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// MINRESERVESETTERROLE is a paid mutator transaction binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) MINRESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERADMINROLE is a paid mutator transaction binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RESERVESETTERADMINROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "RESERVE_SETTER_ADMIN_ROLE")
}

// RESERVESETTERADMINROLE is a paid mutator transaction binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RESERVESETTERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERADMINROLE is a paid mutator transaction binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RESERVESETTERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERROLE is a paid mutator transaction binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RESERVESETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "RESERVE_SETTER_ROLE")
}

// RESERVESETTERROLE is a paid mutator transaction binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERROLE is a paid mutator transaction binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// ROUNDTIMINGSETTERROLE is a paid mutator transaction binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ROUNDTIMINGSETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "ROUND_TIMING_SETTER_ROLE")
}

// ROUNDTIMINGSETTERROLE is a paid mutator transaction binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ROUNDTIMINGSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// ROUNDTIMINGSETTERROLE is a paid mutator transaction binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ROUNDTIMINGSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// Beneficiary is a paid mutator transaction binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) Beneficiary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "beneficiary")
}

// Beneficiary is a paid mutator transaction binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) Beneficiary() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Beneficiary(&_IExpressLaneAuction.TransactOpts)
}

// Beneficiary is a paid mutator transaction binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) Beneficiary() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Beneficiary(&_IExpressLaneAuction.TransactOpts)
}

// BeneficiaryBalance is a paid mutator transaction binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) BeneficiaryBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "beneficiaryBalance")
}

// BeneficiaryBalance is a paid mutator transaction binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// BeneficiaryBalance is a paid mutator transaction binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) BeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// BiddingToken is a paid mutator transaction binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) BiddingToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "biddingToken")
}

// BiddingToken is a paid mutator transaction binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BiddingToken() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BiddingToken(&_IExpressLaneAuction.TransactOpts)
}

// BiddingToken is a paid mutator transaction binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) BiddingToken() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BiddingToken(&_IExpressLaneAuction.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Deposit(&_IExpressLaneAuction.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Deposit(&_IExpressLaneAuction.TransactOpts, amount)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) FinalizeWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "finalizeWithdrawal")
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FinalizeWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FinalizeWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) FlushBeneficiaryBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "flushBeneficiaryBalance")
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.GrantRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.GrantRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) Initialize(opts *bind.TransactOpts, args InitArgs) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "initialize", args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Initialize(&_IExpressLaneAuction.TransactOpts, args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Initialize(&_IExpressLaneAuction.TransactOpts, args)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) InitiateWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "initiateWithdrawal")
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.InitiateWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.InitiateWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// MinReservePrice is a paid mutator transaction binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) MinReservePrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "minReservePrice")
}

// MinReservePrice is a paid mutator transaction binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) MinReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MinReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// MinReservePrice is a paid mutator transaction binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) MinReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MinReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RenounceRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RenounceRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// ReservePrice is a paid mutator transaction binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ReservePrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "reservePrice")
}

// ReservePrice is a paid mutator transaction binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// ReservePrice is a paid mutator transaction binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ResolveMultiBidAuction(opts *bind.TransactOpts, firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "resolveMultiBidAuction", firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveMultiBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveMultiBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ResolveSingleBidAuction(opts *bind.TransactOpts, firstPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "resolveSingleBidAuction", firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveSingleBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveSingleBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RevokeRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RevokeRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setBeneficiary", newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetBeneficiary(&_IExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetBeneficiary(&_IExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetMinReservePrice(opts *bind.TransactOpts, newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setMinReservePrice", newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetMinReservePrice(&_IExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetMinReservePrice(&_IExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetReservePrice(opts *bind.TransactOpts, newReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setReservePrice", newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetReservePrice(&_IExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetReservePrice(&_IExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetRoundTimingInfo(opts *bind.TransactOpts, newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setRoundTimingInfo", newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetRoundTimingInfo(&_IExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetRoundTimingInfo(&_IExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetTransferor(opts *bind.TransactOpts, transferor Transferor) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setTransferor", transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetTransferor(&_IExpressLaneAuction.TransactOpts, transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetTransferor(&_IExpressLaneAuction.TransactOpts, transferor)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) TransferExpressLaneController(opts *bind.TransactOpts, round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "transferExpressLaneController", round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferExpressLaneController(&_IExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferExpressLaneController(&_IExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// TransferorOf is a paid mutator transaction binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address expressLaneController) returns(address addr, uint64 fixedUntil)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) TransferorOf(opts *bind.TransactOpts, expressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "transferorOf", expressLaneController)
}

// TransferorOf is a paid mutator transaction binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address expressLaneController) returns(address addr, uint64 fixedUntil)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) TransferorOf(expressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferorOf(&_IExpressLaneAuction.TransactOpts, expressLaneController)
}

// TransferorOf is a paid mutator transaction binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address expressLaneController) returns(address addr, uint64 fixedUntil)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) TransferorOf(expressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferorOf(&_IExpressLaneAuction.TransactOpts, expressLaneController)
}

// IExpressLaneAuctionAuctionResolvedIterator is returned from FilterAuctionResolved and is used to iterate over the raw logs and unpacked data for AuctionResolved events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionAuctionResolvedIterator struct {
	Event *IExpressLaneAuctionAuctionResolved // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionAuctionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionAuctionResolved)
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
		it.Event = new(IExpressLaneAuctionAuctionResolved)
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
func (it *IExpressLaneAuctionAuctionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionAuctionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionAuctionResolved represents a AuctionResolved event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionAuctionResolved struct {
	IsMultiBidAuction               bool
	Round                           uint64
	FirstPriceBidder                common.Address
	FirstPriceExpressLaneController common.Address
	FirstPriceAmount                *big.Int
	Price                           *big.Int
	RoundStartTimestamp             uint64
	RoundEndTimestamp               uint64
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterAuctionResolved is a free log retrieval operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterAuctionResolved(opts *bind.FilterOpts, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (*IExpressLaneAuctionAuctionResolvedIterator, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionAuctionResolvedIterator{contract: _IExpressLaneAuction.contract, event: "AuctionResolved", logs: logs, sub: sub}, nil
}

// WatchAuctionResolved is a free log subscription operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchAuctionResolved(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionAuctionResolved, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (event.Subscription, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionAuctionResolved)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
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

// ParseAuctionResolved is a log parse operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseAuctionResolved(log types.Log) (*IExpressLaneAuctionAuctionResolved, error) {
	event := new(IExpressLaneAuctionAuctionResolved)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionDepositIterator struct {
	Event *IExpressLaneAuctionDeposit // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionDeposit)
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
		it.Event = new(IExpressLaneAuctionDeposit)
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
func (it *IExpressLaneAuctionDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionDeposit represents a Deposit event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*IExpressLaneAuctionDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionDepositIterator{contract: _IExpressLaneAuction.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionDeposit)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseDeposit(log types.Log) (*IExpressLaneAuctionDeposit, error) {
	event := new(IExpressLaneAuctionDeposit)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleAdminChangedIterator struct {
	Event *IExpressLaneAuctionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionRoleAdminChanged)
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
		it.Event = new(IExpressLaneAuctionRoleAdminChanged)
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
func (it *IExpressLaneAuctionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionRoleAdminChanged represents a RoleAdminChanged event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IExpressLaneAuctionRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionRoleAdminChangedIterator{contract: _IExpressLaneAuction.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionRoleAdminChanged)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseRoleAdminChanged(log types.Log) (*IExpressLaneAuctionRoleAdminChanged, error) {
	event := new(IExpressLaneAuctionRoleAdminChanged)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleGrantedIterator struct {
	Event *IExpressLaneAuctionRoleGranted // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionRoleGranted)
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
		it.Event = new(IExpressLaneAuctionRoleGranted)
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
func (it *IExpressLaneAuctionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionRoleGranted represents a RoleGranted event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IExpressLaneAuctionRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionRoleGrantedIterator{contract: _IExpressLaneAuction.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionRoleGranted)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseRoleGranted(log types.Log) (*IExpressLaneAuctionRoleGranted, error) {
	event := new(IExpressLaneAuctionRoleGranted)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleRevokedIterator struct {
	Event *IExpressLaneAuctionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionRoleRevoked)
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
		it.Event = new(IExpressLaneAuctionRoleRevoked)
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
func (it *IExpressLaneAuctionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionRoleRevoked represents a RoleRevoked event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IExpressLaneAuctionRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionRoleRevokedIterator{contract: _IExpressLaneAuction.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionRoleRevoked)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseRoleRevoked(log types.Log) (*IExpressLaneAuctionRoleRevoked, error) {
	event := new(IExpressLaneAuctionRoleRevoked)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetBeneficiaryIterator is returned from FilterSetBeneficiary and is used to iterate over the raw logs and unpacked data for SetBeneficiary events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetBeneficiaryIterator struct {
	Event *IExpressLaneAuctionSetBeneficiary // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetBeneficiaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetBeneficiary)
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
		it.Event = new(IExpressLaneAuctionSetBeneficiary)
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
func (it *IExpressLaneAuctionSetBeneficiaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetBeneficiaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetBeneficiary represents a SetBeneficiary event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetBeneficiary struct {
	OldBeneficiary common.Address
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetBeneficiary is a free log retrieval operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetBeneficiary(opts *bind.FilterOpts) (*IExpressLaneAuctionSetBeneficiaryIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetBeneficiaryIterator{contract: _IExpressLaneAuction.contract, event: "SetBeneficiary", logs: logs, sub: sub}, nil
}

// WatchSetBeneficiary is a free log subscription operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetBeneficiary(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetBeneficiary) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetBeneficiary)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
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

// ParseSetBeneficiary is a log parse operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetBeneficiary(log types.Log) (*IExpressLaneAuctionSetBeneficiary, error) {
	event := new(IExpressLaneAuctionSetBeneficiary)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetExpressLaneControllerIterator is returned from FilterSetExpressLaneController and is used to iterate over the raw logs and unpacked data for SetExpressLaneController events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetExpressLaneControllerIterator struct {
	Event *IExpressLaneAuctionSetExpressLaneController // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetExpressLaneControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetExpressLaneController)
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
		it.Event = new(IExpressLaneAuctionSetExpressLaneController)
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
func (it *IExpressLaneAuctionSetExpressLaneControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetExpressLaneControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetExpressLaneController represents a SetExpressLaneController event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetExpressLaneController struct {
	Round                         uint64
	PreviousExpressLaneController common.Address
	NewExpressLaneController      common.Address
	Transferor                    common.Address
	StartTimestamp                uint64
	EndTimestamp                  uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetExpressLaneController is a free log retrieval operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetExpressLaneController(opts *bind.FilterOpts, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (*IExpressLaneAuctionSetExpressLaneControllerIterator, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetExpressLaneControllerIterator{contract: _IExpressLaneAuction.contract, event: "SetExpressLaneController", logs: logs, sub: sub}, nil
}

// WatchSetExpressLaneController is a free log subscription operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetExpressLaneController(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetExpressLaneController, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetExpressLaneController)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
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

// ParseSetExpressLaneController is a log parse operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetExpressLaneController(log types.Log) (*IExpressLaneAuctionSetExpressLaneController, error) {
	event := new(IExpressLaneAuctionSetExpressLaneController)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetMinReservePriceIterator is returned from FilterSetMinReservePrice and is used to iterate over the raw logs and unpacked data for SetMinReservePrice events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetMinReservePriceIterator struct {
	Event *IExpressLaneAuctionSetMinReservePrice // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetMinReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetMinReservePrice)
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
		it.Event = new(IExpressLaneAuctionSetMinReservePrice)
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
func (it *IExpressLaneAuctionSetMinReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetMinReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetMinReservePrice represents a SetMinReservePrice event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetMinReservePrice struct {
	OldPrice *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMinReservePrice is a free log retrieval operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetMinReservePrice(opts *bind.FilterOpts) (*IExpressLaneAuctionSetMinReservePriceIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetMinReservePriceIterator{contract: _IExpressLaneAuction.contract, event: "SetMinReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetMinReservePrice is a free log subscription operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetMinReservePrice(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetMinReservePrice) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetMinReservePrice)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
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

// ParseSetMinReservePrice is a log parse operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetMinReservePrice(log types.Log) (*IExpressLaneAuctionSetMinReservePrice, error) {
	event := new(IExpressLaneAuctionSetMinReservePrice)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetReservePriceIterator is returned from FilterSetReservePrice and is used to iterate over the raw logs and unpacked data for SetReservePrice events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetReservePriceIterator struct {
	Event *IExpressLaneAuctionSetReservePrice // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetReservePrice)
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
		it.Event = new(IExpressLaneAuctionSetReservePrice)
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
func (it *IExpressLaneAuctionSetReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetReservePrice represents a SetReservePrice event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetReservePrice struct {
	OldReservePrice *big.Int
	NewReservePrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReservePrice is a free log retrieval operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetReservePrice(opts *bind.FilterOpts) (*IExpressLaneAuctionSetReservePriceIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetReservePriceIterator{contract: _IExpressLaneAuction.contract, event: "SetReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetReservePrice is a free log subscription operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetReservePrice(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetReservePrice) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetReservePrice)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
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

// ParseSetReservePrice is a log parse operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetReservePrice(log types.Log) (*IExpressLaneAuctionSetReservePrice, error) {
	event := new(IExpressLaneAuctionSetReservePrice)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetRoundTimingInfoIterator is returned from FilterSetRoundTimingInfo and is used to iterate over the raw logs and unpacked data for SetRoundTimingInfo events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetRoundTimingInfoIterator struct {
	Event *IExpressLaneAuctionSetRoundTimingInfo // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetRoundTimingInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetRoundTimingInfo)
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
		it.Event = new(IExpressLaneAuctionSetRoundTimingInfo)
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
func (it *IExpressLaneAuctionSetRoundTimingInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetRoundTimingInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetRoundTimingInfo represents a SetRoundTimingInfo event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetRoundTimingInfo struct {
	CurrentRound             uint64
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetRoundTimingInfo is a free log retrieval operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetRoundTimingInfo(opts *bind.FilterOpts) (*IExpressLaneAuctionSetRoundTimingInfoIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetRoundTimingInfoIterator{contract: _IExpressLaneAuction.contract, event: "SetRoundTimingInfo", logs: logs, sub: sub}, nil
}

// WatchSetRoundTimingInfo is a free log subscription operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetRoundTimingInfo(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetRoundTimingInfo) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetRoundTimingInfo)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
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

// ParseSetRoundTimingInfo is a log parse operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetRoundTimingInfo(log types.Log) (*IExpressLaneAuctionSetRoundTimingInfo, error) {
	event := new(IExpressLaneAuctionSetRoundTimingInfo)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetTransferorIterator is returned from FilterSetTransferor and is used to iterate over the raw logs and unpacked data for SetTransferor events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetTransferorIterator struct {
	Event *IExpressLaneAuctionSetTransferor // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetTransferorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetTransferor)
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
		it.Event = new(IExpressLaneAuctionSetTransferor)
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
func (it *IExpressLaneAuctionSetTransferorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetTransferorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetTransferor represents a SetTransferor event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetTransferor struct {
	ExpressLaneController common.Address
	Transferor            common.Address
	FixedUntilRound       uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetTransferor is a free log retrieval operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetTransferor(opts *bind.FilterOpts, expressLaneController []common.Address, transferor []common.Address) (*IExpressLaneAuctionSetTransferorIterator, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetTransferorIterator{contract: _IExpressLaneAuction.contract, event: "SetTransferor", logs: logs, sub: sub}, nil
}

// WatchSetTransferor is a free log subscription operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetTransferor(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetTransferor, expressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetTransferor)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
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

// ParseSetTransferor is a log parse operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetTransferor(log types.Log) (*IExpressLaneAuctionSetTransferor, error) {
	event := new(IExpressLaneAuctionSetTransferor)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalFinalizedIterator struct {
	Event *IExpressLaneAuctionWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionWithdrawalFinalized)
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
		it.Event = new(IExpressLaneAuctionWithdrawalFinalized)
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
func (it *IExpressLaneAuctionWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionWithdrawalFinalized represents a WithdrawalFinalized event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalFinalized struct {
	Account          common.Address
	WithdrawalAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, account []common.Address) (*IExpressLaneAuctionWithdrawalFinalizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionWithdrawalFinalizedIterator{contract: _IExpressLaneAuction.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionWithdrawalFinalized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionWithdrawalFinalized)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseWithdrawalFinalized(log types.Log) (*IExpressLaneAuctionWithdrawalFinalized, error) {
	event := new(IExpressLaneAuctionWithdrawalFinalized)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalInitiatedIterator struct {
	Event *IExpressLaneAuctionWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionWithdrawalInitiated)
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
		it.Event = new(IExpressLaneAuctionWithdrawalInitiated)
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
func (it *IExpressLaneAuctionWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionWithdrawalInitiated represents a WithdrawalInitiated event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalInitiated struct {
	Account           common.Address
	WithdrawalAmount  *big.Int
	RoundWithdrawable *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, account []common.Address) (*IExpressLaneAuctionWithdrawalInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionWithdrawalInitiatedIterator{contract: _IExpressLaneAuction.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionWithdrawalInitiated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionWithdrawalInitiated)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseWithdrawalInitiated(log types.Log) (*IExpressLaneAuctionWithdrawalInitiated, error) {
	event := new(IExpressLaneAuctionWithdrawalInitiated)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LatestELCRoundsLibMetaData contains all meta data concerning the LatestELCRoundsLib contract.
var LatestELCRoundsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220746dd8de3dfb6ffeb56880b36509d1fc866276e62adfe3b467451f0306b63fce64736f6c63430008090033",
}

// LatestELCRoundsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use LatestELCRoundsLibMetaData.ABI instead.
var LatestELCRoundsLibABI = LatestELCRoundsLibMetaData.ABI

// LatestELCRoundsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LatestELCRoundsLibMetaData.Bin instead.
var LatestELCRoundsLibBin = LatestELCRoundsLibMetaData.Bin

// DeployLatestELCRoundsLib deploys a new Ethereum contract, binding an instance of LatestELCRoundsLib to it.
func DeployLatestELCRoundsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LatestELCRoundsLib, error) {
	parsed, err := LatestELCRoundsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LatestELCRoundsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LatestELCRoundsLib{LatestELCRoundsLibCaller: LatestELCRoundsLibCaller{contract: contract}, LatestELCRoundsLibTransactor: LatestELCRoundsLibTransactor{contract: contract}, LatestELCRoundsLibFilterer: LatestELCRoundsLibFilterer{contract: contract}}, nil
}

// LatestELCRoundsLib is an auto generated Go binding around an Ethereum contract.
type LatestELCRoundsLib struct {
	LatestELCRoundsLibCaller     // Read-only binding to the contract
	LatestELCRoundsLibTransactor // Write-only binding to the contract
	LatestELCRoundsLibFilterer   // Log filterer for contract events
}

// LatestELCRoundsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type LatestELCRoundsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LatestELCRoundsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LatestELCRoundsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LatestELCRoundsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LatestELCRoundsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LatestELCRoundsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LatestELCRoundsLibSession struct {
	Contract     *LatestELCRoundsLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LatestELCRoundsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LatestELCRoundsLibCallerSession struct {
	Contract *LatestELCRoundsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LatestELCRoundsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LatestELCRoundsLibTransactorSession struct {
	Contract     *LatestELCRoundsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LatestELCRoundsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type LatestELCRoundsLibRaw struct {
	Contract *LatestELCRoundsLib // Generic contract binding to access the raw methods on
}

// LatestELCRoundsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LatestELCRoundsLibCallerRaw struct {
	Contract *LatestELCRoundsLibCaller // Generic read-only contract binding to access the raw methods on
}

// LatestELCRoundsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LatestELCRoundsLibTransactorRaw struct {
	Contract *LatestELCRoundsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLatestELCRoundsLib creates a new instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLib(address common.Address, backend bind.ContractBackend) (*LatestELCRoundsLib, error) {
	contract, err := bindLatestELCRoundsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLib{LatestELCRoundsLibCaller: LatestELCRoundsLibCaller{contract: contract}, LatestELCRoundsLibTransactor: LatestELCRoundsLibTransactor{contract: contract}, LatestELCRoundsLibFilterer: LatestELCRoundsLibFilterer{contract: contract}}, nil
}

// NewLatestELCRoundsLibCaller creates a new read-only instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLibCaller(address common.Address, caller bind.ContractCaller) (*LatestELCRoundsLibCaller, error) {
	contract, err := bindLatestELCRoundsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLibCaller{contract: contract}, nil
}

// NewLatestELCRoundsLibTransactor creates a new write-only instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*LatestELCRoundsLibTransactor, error) {
	contract, err := bindLatestELCRoundsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLibTransactor{contract: contract}, nil
}

// NewLatestELCRoundsLibFilterer creates a new log filterer instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*LatestELCRoundsLibFilterer, error) {
	contract, err := bindLatestELCRoundsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLibFilterer{contract: contract}, nil
}

// bindLatestELCRoundsLib binds a generic wrapper to an already deployed contract.
func bindLatestELCRoundsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LatestELCRoundsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LatestELCRoundsLib *LatestELCRoundsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LatestELCRoundsLib.Contract.LatestELCRoundsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LatestELCRoundsLib *LatestELCRoundsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.LatestELCRoundsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LatestELCRoundsLib *LatestELCRoundsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.LatestELCRoundsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LatestELCRoundsLib *LatestELCRoundsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LatestELCRoundsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LatestELCRoundsLib *LatestELCRoundsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LatestELCRoundsLib *LatestELCRoundsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.contract.Transact(opts, method, params...)
}

// RoundTimingInfoLibMetaData contains all meta data concerning the RoundTimingInfoLib contract.
var RoundTimingInfoLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c90e0cb1f7c8a2d579b5818a2a2e47b21d335afc5dfc8222540cbd9f03d93dac64736f6c63430008090033",
}

// RoundTimingInfoLibABI is the input ABI used to generate the binding from.
// Deprecated: Use RoundTimingInfoLibMetaData.ABI instead.
var RoundTimingInfoLibABI = RoundTimingInfoLibMetaData.ABI

// RoundTimingInfoLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoundTimingInfoLibMetaData.Bin instead.
var RoundTimingInfoLibBin = RoundTimingInfoLibMetaData.Bin

// DeployRoundTimingInfoLib deploys a new Ethereum contract, binding an instance of RoundTimingInfoLib to it.
func DeployRoundTimingInfoLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoundTimingInfoLib, error) {
	parsed, err := RoundTimingInfoLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoundTimingInfoLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoundTimingInfoLib{RoundTimingInfoLibCaller: RoundTimingInfoLibCaller{contract: contract}, RoundTimingInfoLibTransactor: RoundTimingInfoLibTransactor{contract: contract}, RoundTimingInfoLibFilterer: RoundTimingInfoLibFilterer{contract: contract}}, nil
}

// RoundTimingInfoLib is an auto generated Go binding around an Ethereum contract.
type RoundTimingInfoLib struct {
	RoundTimingInfoLibCaller     // Read-only binding to the contract
	RoundTimingInfoLibTransactor // Write-only binding to the contract
	RoundTimingInfoLibFilterer   // Log filterer for contract events
}

// RoundTimingInfoLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoundTimingInfoLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTimingInfoLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoundTimingInfoLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTimingInfoLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoundTimingInfoLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTimingInfoLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoundTimingInfoLibSession struct {
	Contract     *RoundTimingInfoLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RoundTimingInfoLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoundTimingInfoLibCallerSession struct {
	Contract *RoundTimingInfoLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RoundTimingInfoLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoundTimingInfoLibTransactorSession struct {
	Contract     *RoundTimingInfoLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RoundTimingInfoLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoundTimingInfoLibRaw struct {
	Contract *RoundTimingInfoLib // Generic contract binding to access the raw methods on
}

// RoundTimingInfoLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoundTimingInfoLibCallerRaw struct {
	Contract *RoundTimingInfoLibCaller // Generic read-only contract binding to access the raw methods on
}

// RoundTimingInfoLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoundTimingInfoLibTransactorRaw struct {
	Contract *RoundTimingInfoLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoundTimingInfoLib creates a new instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLib(address common.Address, backend bind.ContractBackend) (*RoundTimingInfoLib, error) {
	contract, err := bindRoundTimingInfoLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLib{RoundTimingInfoLibCaller: RoundTimingInfoLibCaller{contract: contract}, RoundTimingInfoLibTransactor: RoundTimingInfoLibTransactor{contract: contract}, RoundTimingInfoLibFilterer: RoundTimingInfoLibFilterer{contract: contract}}, nil
}

// NewRoundTimingInfoLibCaller creates a new read-only instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLibCaller(address common.Address, caller bind.ContractCaller) (*RoundTimingInfoLibCaller, error) {
	contract, err := bindRoundTimingInfoLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLibCaller{contract: contract}, nil
}

// NewRoundTimingInfoLibTransactor creates a new write-only instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RoundTimingInfoLibTransactor, error) {
	contract, err := bindRoundTimingInfoLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLibTransactor{contract: contract}, nil
}

// NewRoundTimingInfoLibFilterer creates a new log filterer instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RoundTimingInfoLibFilterer, error) {
	contract, err := bindRoundTimingInfoLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLibFilterer{contract: contract}, nil
}

// bindRoundTimingInfoLib binds a generic wrapper to an already deployed contract.
func bindRoundTimingInfoLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoundTimingInfoLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoundTimingInfoLib *RoundTimingInfoLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoundTimingInfoLib.Contract.RoundTimingInfoLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoundTimingInfoLib *RoundTimingInfoLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.RoundTimingInfoLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoundTimingInfoLib *RoundTimingInfoLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.RoundTimingInfoLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoundTimingInfoLib *RoundTimingInfoLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoundTimingInfoLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoundTimingInfoLib *RoundTimingInfoLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoundTimingInfoLib *RoundTimingInfoLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.contract.Transact(opts, method, params...)
}
