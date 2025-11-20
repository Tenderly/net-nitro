// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package handlergen

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

// CompatibilityFallbackHandlerMetaData contains all meta data concerning the CompatibilityFallbackHandler contract.
var CompatibilityFallbackHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"encodeMessageDataForSafe\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHashForSafe\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611896806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b35760003560e01c80636ac24784116100715780636ac2478414610622578063b2494df314610711578063bc197c8114610770578063bd61951d14610906578063e86637db14610a18578063f23a6e6114610baa576100b3565b806223de29146100b857806301ffc9a7146101f05780630a1028c414610253578063150b7a02146103225780631626ba7e1461041857806323031640146104ce575b600080fd5b6101ee600480360360c08110156100ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561015557600080fd5b82018360208201111561016757600080fd5b8035906020019184600183028401116401000000008311171561018957600080fd5b9091929391929390803590602001906401000000008111156101aa57600080fd5b8201836020820111156101bc57600080fd5b803590602001918460018302840111640100000000831117156101de57600080fd5b9091929391929390505050610caa565b005b61023b6004803603602081101561020657600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610cb4565b60405180821515815260200191505060405180910390f35b61030c6004803603602081101561026957600080fd5b810190808035906020019064010000000081111561028657600080fd5b82018360208201111561029857600080fd5b803590602001918460018302840111640100000000831117156102ba57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610dee565b6040518082815260200191505060405180910390f35b6103e36004803603608081101561033857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561039f57600080fd5b8201836020820111156103b157600080fd5b803590602001918460018302840111640100000000831117156103d357600080fd5b9091929391929390505050610e01565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b6104996004803603604081101561042e57600080fd5b81019080803590602001909291908035906020019064010000000081111561045557600080fd5b82018360208201111561046757600080fd5b8035906020019184600183028401116401000000008311171561048957600080fd5b9091929391929390505050610e1e565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b6105a7600480360360408110156104e457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561052157600080fd5b82018360208201111561053357600080fd5b8035906020019184600183028401116401000000008311171561055557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611041565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105e75780820151818401526020810190506105cc565b50505050905090810190601f1680156106145780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106fb6004803603604081101561063857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561067557600080fd5b82018360208201111561068757600080fd5b803590602001918460018302840111640100000000831117156106a957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506111ad565b6040518082815260200191505060405180910390f35b6107196111c8565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561075c578082015181840152602081019050610741565b505050509050019250505060405180910390f35b6108d1600480360360a081101561078657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156107e357600080fd5b8201836020820111156107f557600080fd5b8035906020019184602083028401116401000000008311171561081757600080fd5b90919293919293908035906020019064010000000081111561083857600080fd5b82018360208201111561084a57600080fd5b8035906020019184602083028401116401000000008311171561086c57600080fd5b90919293919293908035906020019064010000000081111561088d57600080fd5b82018360208201111561089f57600080fd5b803590602001918460018302840111640100000000831117156108c157600080fd5b909192939192939050505061132f565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b61099d6004803603604081101561091c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561095957600080fd5b82018360208201111561096b57600080fd5b8035906020019184600183028401116401000000008311171561098d57600080fd5b909192939192939050505061134f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109dd5780820151818401526020810190506109c2565b50505050905090810190601f168015610a0a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610b2f6004803603610140811015610a2f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190640100000000811115610a7657600080fd5b820183602082011115610a8857600080fd5b80359060200191846001830284011164010000000083111715610aaa57600080fd5b9091929391929390803560ff169060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506113c9565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b6f578082015181840152602081019050610b54565b50505050905090810190601f168015610b9c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610c75600480360360a0811015610bc057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190640100000000811115610c3157600080fd5b820183602082011115610c4357600080fd5b80359060200191846001830284011164010000000083111715610c6557600080fd5b90919293919293905050506115f6565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b5050505050505050565b60007f4e2312e0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610d7f57507f150b7a02000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610de757507f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b6000610dfa33836111ad565b9050919050565b6000610e0b611614565b63150b7a0260e01b905095945050505050565b6000803390506000610e4f828760405160200180828152602001915050604051602081830303815290604052611041565b90506000818051906020012090506000868690501415610f705760008373ffffffffffffffffffffffffffffffffffffffff16635ae6bd37836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610ebc57600080fd5b505afa158015610ed0573d6000803e3d6000fd5b505050506040513d6020811015610ee657600080fd5b81019080805190602001909291905050501415610f6b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f48617368206e6f7420617070726f76656400000000000000000000000000000081525060200191505060405180910390fd5b61102d565b8273ffffffffffffffffffffffffffffffffffffffff1663f855438b60008389896040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509550505050505060006040518083038186803b15801561101457600080fd5b505afa158015611028573d6000803e3d6000fd5b505050505b631626ba7e60e01b93505050509392505050565b606060007f60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca60001b83805190602001206040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050601960f81b600160f81b8573ffffffffffffffffffffffffffffffffffffffff1663f698da256040518163ffffffff1660e01b815260040160206040518083038186803b1580156110f057600080fd5b505afa158015611104573d6000803e3d6000fd5b505050506040513d602081101561111a57600080fd5b81019080805190602001909291905050508360405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101847effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260010183815260200182815260200194505050505060405160208183030381529060405291505092915050565b60006111b98383611041565b80519060200120905092915050565b6060600033905060008173ffffffffffffffffffffffffffffffffffffffff1663cc2f84526001600a6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060006040518083038186803b15801561124257600080fd5b505afa158015611256573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250604081101561128057600080fd5b81019080805160405193929190846401000000008211156112a057600080fd5b838201915060208201858111156112b657600080fd5b82518660208202830111640100000000821117156112d357600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561130a5780820151818401526020810190506112ef565b5050505090500160405260200180519060200190929190505050509050809250505090565b6000611339611614565b63bc197c8160e01b905098975050505050505050565b60606040517fb4faba09000000000000000000000000000000000000000000000000000000008152600436036004808301376040600036836000335af160403d1081171561139c57600080fd5b60208051016040519350808401604052806020853e6000516113bf578060208501fd5b5050509392505050565b6060600033905060008173ffffffffffffffffffffffffffffffffffffffff1663f698da256040518163ffffffff1660e01b815260040160206040518083038186803b15801561141857600080fd5b505afa15801561142c573d6000803e3d6000fd5b505050506040513d602081101561144257600080fd5b8101908080519060200190929190505050905060007fbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d860001b8f8f8f8f60405180838380828437808301925050509250505060405180910390208e8e8e8e8e8e8e604051602001808c81526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018a81526020018981526020018860018111156114e457fe5b81526020018781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019b505050505050505050505050604051602081830303815290604052805190602001209050601960f81b600160f81b838360405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101847effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260010183815260200182815260200194505050505060405160208183030381529060405293505050509b9a5050505050505050505050565b6000611600611614565b63f23a6e6160e01b90509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff16635624b25b7f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d560001b60001c60016040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561169657600080fd5b505afa1580156116aa573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156116d457600080fd5b81019080805160405193929190846401000000008211156116f457600080fd5b8382019150602082018581111561170a57600080fd5b825186600182028301116401000000008211171561172757600080fd5b8083526020830192505050908051906020019080838360005b8381101561175b578082015181840152602081019050611740565b50505050905090810190601f1680156117885780820380516001836020036101000a031916815260200191505b50604052505050905060008180602001905160208110156117a857600080fd5b810190808051906020019092919050505090503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461185c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6e6f7420612066616c6c6261636b2063616c6c0000000000000000000000000081525060200191505060405180910390fd5b505056fea2646970667358221220b726ac1476e8c5db8701e0875ac21d3d24771ab2d3c61762057654d411c9390064736f6c63430007060033",
}

// CompatibilityFallbackHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use CompatibilityFallbackHandlerMetaData.ABI instead.
var CompatibilityFallbackHandlerABI = CompatibilityFallbackHandlerMetaData.ABI

// CompatibilityFallbackHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompatibilityFallbackHandlerMetaData.Bin instead.
var CompatibilityFallbackHandlerBin = CompatibilityFallbackHandlerMetaData.Bin

// DeployCompatibilityFallbackHandler deploys a new Ethereum contract, binding an instance of CompatibilityFallbackHandler to it.
func DeployCompatibilityFallbackHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompatibilityFallbackHandler, error) {
	parsed, err := CompatibilityFallbackHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompatibilityFallbackHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompatibilityFallbackHandler{CompatibilityFallbackHandlerCaller: CompatibilityFallbackHandlerCaller{contract: contract}, CompatibilityFallbackHandlerTransactor: CompatibilityFallbackHandlerTransactor{contract: contract}, CompatibilityFallbackHandlerFilterer: CompatibilityFallbackHandlerFilterer{contract: contract}}, nil
}

// CompatibilityFallbackHandler is an auto generated Go binding around an Ethereum contract.
type CompatibilityFallbackHandler struct {
	CompatibilityFallbackHandlerCaller     // Read-only binding to the contract
	CompatibilityFallbackHandlerTransactor // Write-only binding to the contract
	CompatibilityFallbackHandlerFilterer   // Log filterer for contract events
}

// CompatibilityFallbackHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompatibilityFallbackHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibilityFallbackHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompatibilityFallbackHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibilityFallbackHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompatibilityFallbackHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibilityFallbackHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompatibilityFallbackHandlerSession struct {
	Contract     *CompatibilityFallbackHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CompatibilityFallbackHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompatibilityFallbackHandlerCallerSession struct {
	Contract *CompatibilityFallbackHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CompatibilityFallbackHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompatibilityFallbackHandlerTransactorSession struct {
	Contract     *CompatibilityFallbackHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CompatibilityFallbackHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompatibilityFallbackHandlerRaw struct {
	Contract *CompatibilityFallbackHandler // Generic contract binding to access the raw methods on
}

// CompatibilityFallbackHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompatibilityFallbackHandlerCallerRaw struct {
	Contract *CompatibilityFallbackHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// CompatibilityFallbackHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompatibilityFallbackHandlerTransactorRaw struct {
	Contract *CompatibilityFallbackHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompatibilityFallbackHandler creates a new instance of CompatibilityFallbackHandler, bound to a specific deployed contract.
func NewCompatibilityFallbackHandler(address common.Address, backend bind.ContractBackend) (*CompatibilityFallbackHandler, error) {
	contract, err := bindCompatibilityFallbackHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompatibilityFallbackHandler{CompatibilityFallbackHandlerCaller: CompatibilityFallbackHandlerCaller{contract: contract}, CompatibilityFallbackHandlerTransactor: CompatibilityFallbackHandlerTransactor{contract: contract}, CompatibilityFallbackHandlerFilterer: CompatibilityFallbackHandlerFilterer{contract: contract}}, nil
}

// NewCompatibilityFallbackHandlerCaller creates a new read-only instance of CompatibilityFallbackHandler, bound to a specific deployed contract.
func NewCompatibilityFallbackHandlerCaller(address common.Address, caller bind.ContractCaller) (*CompatibilityFallbackHandlerCaller, error) {
	contract, err := bindCompatibilityFallbackHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibilityFallbackHandlerCaller{contract: contract}, nil
}

// NewCompatibilityFallbackHandlerTransactor creates a new write-only instance of CompatibilityFallbackHandler, bound to a specific deployed contract.
func NewCompatibilityFallbackHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*CompatibilityFallbackHandlerTransactor, error) {
	contract, err := bindCompatibilityFallbackHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibilityFallbackHandlerTransactor{contract: contract}, nil
}

// NewCompatibilityFallbackHandlerFilterer creates a new log filterer instance of CompatibilityFallbackHandler, bound to a specific deployed contract.
func NewCompatibilityFallbackHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*CompatibilityFallbackHandlerFilterer, error) {
	contract, err := bindCompatibilityFallbackHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompatibilityFallbackHandlerFilterer{contract: contract}, nil
}

// bindCompatibilityFallbackHandler binds a generic wrapper to an already deployed contract.
func bindCompatibilityFallbackHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompatibilityFallbackHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompatibilityFallbackHandler.Contract.CompatibilityFallbackHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.Contract.CompatibilityFallbackHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.Contract.CompatibilityFallbackHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompatibilityFallbackHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.Contract.contract.Transact(opts, method, params...)
}

// EncodeMessageDataForSafe is a free data retrieval call binding the contract method 0x23031640.
//
// Solidity: function encodeMessageDataForSafe(address safe, bytes message) view returns(bytes)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) EncodeMessageDataForSafe(opts *bind.CallOpts, safe common.Address, message []byte) ([]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "encodeMessageDataForSafe", safe, message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMessageDataForSafe is a free data retrieval call binding the contract method 0x23031640.
//
// Solidity: function encodeMessageDataForSafe(address safe, bytes message) view returns(bytes)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) EncodeMessageDataForSafe(safe common.Address, message []byte) ([]byte, error) {
	return _CompatibilityFallbackHandler.Contract.EncodeMessageDataForSafe(&_CompatibilityFallbackHandler.CallOpts, safe, message)
}

// EncodeMessageDataForSafe is a free data retrieval call binding the contract method 0x23031640.
//
// Solidity: function encodeMessageDataForSafe(address safe, bytes message) view returns(bytes)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) EncodeMessageDataForSafe(safe common.Address, message []byte) ([]byte, error) {
	return _CompatibilityFallbackHandler.Contract.EncodeMessageDataForSafe(&_CompatibilityFallbackHandler.CallOpts, safe, message)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 nonce) view returns(bytes)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "encodeTransactionData", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 nonce) view returns(bytes)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, nonce *big.Int) ([]byte, error) {
	return _CompatibilityFallbackHandler.Contract.EncodeTransactionData(&_CompatibilityFallbackHandler.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 nonce) view returns(bytes)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, nonce *big.Int) ([]byte, error) {
	return _CompatibilityFallbackHandler.Contract.EncodeTransactionData(&_CompatibilityFallbackHandler.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, nonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) GetMessageHash(opts *bind.CallOpts, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "getMessageHash", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _CompatibilityFallbackHandler.Contract.GetMessageHash(&_CompatibilityFallbackHandler.CallOpts, message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _CompatibilityFallbackHandler.Contract.GetMessageHash(&_CompatibilityFallbackHandler.CallOpts, message)
}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) GetMessageHashForSafe(opts *bind.CallOpts, safe common.Address, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "getMessageHashForSafe", safe, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) GetMessageHashForSafe(safe common.Address, message []byte) ([32]byte, error) {
	return _CompatibilityFallbackHandler.Contract.GetMessageHashForSafe(&_CompatibilityFallbackHandler.CallOpts, safe, message)
}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) GetMessageHashForSafe(safe common.Address, message []byte) ([32]byte, error) {
	return _CompatibilityFallbackHandler.Contract.GetMessageHashForSafe(&_CompatibilityFallbackHandler.CallOpts, safe, message)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) GetModules() ([]common.Address, error) {
	return _CompatibilityFallbackHandler.Contract.GetModules(&_CompatibilityFallbackHandler.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) GetModules() ([]common.Address, error) {
	return _CompatibilityFallbackHandler.Contract.GetModules(&_CompatibilityFallbackHandler.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) IsValidSignature(opts *bind.CallOpts, _dataHash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "isValidSignature", _dataHash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) IsValidSignature(_dataHash [32]byte, _signature []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.IsValidSignature(&_CompatibilityFallbackHandler.CallOpts, _dataHash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) IsValidSignature(_dataHash [32]byte, _signature []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.IsValidSignature(&_CompatibilityFallbackHandler.CallOpts, _dataHash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.OnERC1155BatchReceived(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.OnERC1155BatchReceived(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.OnERC1155Received(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.OnERC1155Received(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.OnERC721Received(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CompatibilityFallbackHandler.Contract.OnERC721Received(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CompatibilityFallbackHandler.Contract.SupportsInterface(&_CompatibilityFallbackHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CompatibilityFallbackHandler.Contract.SupportsInterface(&_CompatibilityFallbackHandler.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _CompatibilityFallbackHandler.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _CompatibilityFallbackHandler.Contract.TokensReceived(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _CompatibilityFallbackHandler.Contract.TokensReceived(&_CompatibilityFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerTransactor) Simulate(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.contract.Transact(opts, "simulate", targetContract, calldataPayload)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerSession) Simulate(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.Contract.Simulate(&_CompatibilityFallbackHandler.TransactOpts, targetContract, calldataPayload)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_CompatibilityFallbackHandler *CompatibilityFallbackHandlerTransactorSession) Simulate(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _CompatibilityFallbackHandler.Contract.Simulate(&_CompatibilityFallbackHandler.TransactOpts, targetContract, calldataPayload)
}

// ExtensibleFallbackHandlerMetaData contains all meta data concerning the ExtensibleFallbackHandler contract.
var ExtensibleFallbackHandlerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractISafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"AddedInterface\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractISafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractISafeSignatureVerifier\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISafeSignatureVerifier\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"ChangedDomainVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractISafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oldMethod\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newMethod\",\"type\":\"bytes32\"}],\"name\":\"ChangedSafeMethod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractISafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"RemovedInterface\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32[]\",\"name\":\"handlerWithSelectors\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportedInterfaceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISafe\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"domainVerifiers\",\"outputs\":[{\"internalType\":\"contractISafeSignatureVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"name\":\"removeSupportedInterfaceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISafe\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"safeInterfaces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISafe\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"safeMethods\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"contractISafeSignatureVerifier\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"setDomainVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newMethod\",\"type\":\"bytes32\"}],\"name\":\"setSafeMethod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"name\":\"setSupportedInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612d8b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d35760003560e01c806351cad5ee1161008c5780637f73528b116100665780637f73528b14610ac5578063b435a13b14610b1e578063bc197c8114610ba1578063f23a6e6114610d37576100d4565b806351cad5ee1461095457806361f5401b146109cc57806364f95acc14610a6e576100d4565b806301ffc9a7146105d45780630a3fe05414610637578063150b7a02146106b85780631626ba7e146107ae578063327b533c146108645780633365582c14610906576100d4565b5b6000366060601860003690501015610154576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f696e76616c6964206d6574686f642073656c6563746f7200000000000000000081525060200191505060405180910390fd5b600080600080610162610e37565b9350935093509350600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561020d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6d6574686f642068616e646c6572206e6f74207365740000000000000000000081525060200191505060405180910390fd5b81156103ed578073ffffffffffffffffffffffffffffffffffffffff166325d6803f85856000803660009060146000369050039261024d93929190612d22565b6040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b1580156102ed57600080fd5b505afa158015610301573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561032b57600080fd5b810190808051604051939291908464010000000082111561034b57600080fd5b8382019150602082018581111561036157600080fd5b825186600182028301116401000000008211171561037e57600080fd5b8083526020830192505050908051906020019080838360005b838110156103b2578082015181840152602081019050610397565b50505050905090810190601f1680156103df5780820380516001836020036101000a031916815260200191505b5060405250505094506105c5565b8073ffffffffffffffffffffffffffffffffffffffff166325d6803f85856000803660009060146000369050039261042793929190612d22565b6040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b1580156104c957600080fd5b505af11580156104dd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561050757600080fd5b810190808051604051939291908464010000000082111561052757600080fd5b8382019150602082018581111561053d57600080fd5b825186600182028301116401000000008211171561055a57600080fd5b8083526020830192505050908051906020019080838360005b8381101561058e578082015181840152602081019050610573565b50505050905090810190601f1680156105bb5780820380516001836020036101000a031916815260200191505b5060405250505094505b50505050915050805190602001f35b61061f600480360360208110156105ea57600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610f14565b60405180821515815260200191505060405180910390f35b6106a26004803603604081101561064d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506110a0565b6040518082815260200191505060405180910390f35b610779600480360360808110156106ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561073557600080fd5b82018360208201111561074757600080fd5b8035906020019184600183028401116401000000008311171561076957600080fd5b90919293919293905050506110c5565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b61082f600480360360408110156107c457600080fd5b8101908080359060200190929190803590602001906401000000008111156107eb57600080fd5b8201836020820111156107fd57600080fd5b8035906020019184600183028401116401000000008311171561081f57600080fd5b90919293919293905050506110e2565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b6109046004803603604081101561087a57600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190803590602001906401000000008111156108c057600080fd5b8201836020820111156108d257600080fd5b803590602001918460208302840111640100000000831117156108f457600080fd5b90919293919293905050506115ff565b005b6109526004803603604081101561091c57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117e9565b005b6109a06004803603604081101561096a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611a38565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610a6c600480360360408110156109e257600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919080359060200190640100000000811115610a2857600080fd5b820183602082011115610a3a57600080fd5b80359060200191846020830284011164010000000083111715610a5c57600080fd5b9091929391929390505050611a7a565b005b610ac360048036036040811015610a8457600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919080359060200190929190505050611c97565b005b610b1c60048036036040811015610adb57600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190803515159060200190929190505050611d5c565b005b610b8960048036036040811015610b3457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909291905050506120b9565b60405180821515815260200191505060405180910390f35b610d02600480360360a0811015610bb757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610c1457600080fd5b820183602082011115610c2657600080fd5b80359060200191846020830284011164010000000083111715610c4857600080fd5b909192939192939080359060200190640100000000811115610c6957600080fd5b820183602082011115610c7b57600080fd5b80359060200191846020830284011164010000000083111715610c9d57600080fd5b909192939192939080359060200190640100000000811115610cbe57600080fd5b820183602082011115610cd057600080fd5b80359060200191846001830284011164010000000083111715610cf257600080fd5b90919293919293905050506120e8565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b610e02600480360360a0811015610d4d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190640100000000811115610dbe57600080fd5b820183602082011115610dd057600080fd5b80359060200191846001830284011164010000000083111715610df257600080fd5b9091929391929390505050612108565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b600080600080610e45612126565b8094508195505050610f066000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080357fffffffff00000000000000000000000000000000000000000000000000000000167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002054612141565b809250819350505090919293565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610fdf57507f98c8e097000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610fef5750610fee82612169565b5b806110995750600260006110016123db565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff165b9050919050565b6000602052816000526040600020602052806000526040600020600091509150505481565b60006110cf6123e3565b63150b7a0260e01b905095945050505050565b60008060006110ef612126565b91509150600485859050106115a457600085359050635fd7e97d60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614801561115c575060448686905010155b156115a257600080878760049060449261117893929190612d22565b604081101561118657600080fd5b810190808035906020019092919080359060200190929190505050915091506000600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461159e576000808a8a600490809261126493929190612d22565b608081101561127257600080fd5b810190808035906020019092919080359060200190929190803590602001906401000000008111156112a357600080fd5b8201836020820111156112b557600080fd5b803590602001918460018302840111640100000000831117156112d757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561133a57600080fd5b82018360208201111561134c57600080fd5b8035906020019184600183028401116401000000008311171561136e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050509350935050508b6113cd86868561262f565b80519060200120141561159b578273ffffffffffffffffffffffffffffffffffffffff166353f00b1489898f898988886040518863ffffffff1660e01b8152600401808873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018481526020018060200180602001838103835285818151815260200191508051906020019080838360005b8381101561149a57808201518184015260208101905061147f565b50505050905090810190601f1680156114c75780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156115005780820151818401526020810190506114e5565b50505050905090810190601f16801561152d5780820380516001836020036101000a031916815260200191505b50995050505050505050505060206040518083038186803b15801561155157600080fd5b505afa158015611565573d6000803e3d6000fd5b505050506040513d602081101561157b57600080fd5b8101908080519060200190929190505050985050505050505050506115f8565b50505b5050505b505b6115f3828787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050612737565b925050505b9392505050565b6116076123db565b73ffffffffffffffffffffffffffffffffffffffff16611625612a55565b73ffffffffffffffffffffffffffffffffffffffff16146116ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c7920736166652063616e2063616c6c2074686973206d6574686f64000081525060200191505060405180910390fd5b60006116b8612a55565b905060008060e01b9050600084849050905060005b818110156117225760008060006116f58989868181106116e957fe5b90506020020135612adf565b92509250925061170f878361170a8685612b14565b612b67565b81861895505050508060010190506116cd565b50857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146117d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e74657266616365206964206d69736d61746368000000000000000000000081525060200191505060405180910390fd5b6117e1866001611d5c565b505050505050565b6117f16123db565b73ffffffffffffffffffffffffffffffffffffffff1661180f612a55565b73ffffffffffffffffffffffffffffffffffffffff1614611898576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c7920736166652063616e2063616c6c2074686973206d6574686f64000081525060200191505060405180910390fd5b60006118a2612a55565b90506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167f06341ac2f62eb79165de8c8b54ab89a8825b12746c2b98e01dff1007837d2ba8858386604051808481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a250505050565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611a826123db565b73ffffffffffffffffffffffffffffffffffffffff16611aa0612a55565b73ffffffffffffffffffffffffffffffffffffffff1614611b29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c7920736166652063616e2063616c6c2074686973206d6574686f64000081525060200191505060405180910390fd5b6000611b33612a55565b905060008060e01b9050600084849050905060005b81811015611bd057611b8f84878784818110611b6057fe5b905060200201357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166000801b612b67565b858582818110611b9b57fe5b905060200201357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191683189250806001019050611b48565b50857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614611c84576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e74657266616365206964206d69736d61746368000000000000000000000081525060200191505060405180910390fd5b611c8f866000611d5c565b505050505050565b611c9f6123db565b73ffffffffffffffffffffffffffffffffffffffff16611cbd612a55565b73ffffffffffffffffffffffffffffffffffffffff1614611d46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c7920736166652063616e2063616c6c2074686973206d6574686f64000081525060200191505060405180910390fd5b611d58611d51612a55565b8383612b67565b5050565b611d646123db565b73ffffffffffffffffffffffffffffffffffffffff16611d82612a55565b73ffffffffffffffffffffffffffffffffffffffff1614611e0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6f6e6c7920736166652063616e2063616c6c2074686973206d6574686f64000081525060200191505060405180910390fd5b6000611e156123db565b905063ffffffff60e01b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415611eb3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f696e76616c696420696e7465726661636520696400000000000000000000000081525060200191505060405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050801515841515146120b25783826000877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff0219169083151502179055508315612043578273ffffffffffffffffffffffffffffffffffffffff167f3d5024c13f12fa602dbf52b1979058940c224ebf170c83a4e358725ae50db36d8660405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390a26120b1565b8273ffffffffffffffffffffffffffffffffffffffff167ff159d834f487747c1b3f17e2107743e42e6eecff444d894e511c18943072b29f8660405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390a25b5b5050505050565b60026020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006120f26123e3565b63bc197c8160e01b905098975050505050505050565b60006121126123e3565b63f23a6e6160e01b90509695505050505050565b6000806121316123db565b915061213b612a55565b90509091565b6000808260f81c15915073ffffffffffffffffffffffffffffffffffffffff83169050915091565b60007f150b7a02000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061223457507f4e2312e0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061229c57507f1626ba7e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061230457507f62af8dc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061236c57507f99372930000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806123d457507f64f95acc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b600033905090565b60003373ffffffffffffffffffffffffffffffffffffffff16635624b25b7f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d560001b60001c60016040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561246557600080fd5b505afa158015612479573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156124a357600080fd5b81019080805160405193929190846401000000008211156124c357600080fd5b838201915060208201858111156124d957600080fd5b82518660018202830111640100000000821117156124f657600080fd5b8083526020830192505050908051906020019080838360005b8381101561252a57808201518184015260208101905061250f565b50505050905090810190601f1680156125575780820380516001836020036101000a031916815260200191505b506040525050509050600081806020019051602081101561257757600080fd5b810190808051906020019092919050505090503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461262b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6e6f7420612066616c6c6261636b2063616c6c0000000000000000000000000081525060200191505060405180910390fd5b5050565b6060601960f81b600160f81b8585856040516020018083815260200182805190602001908083835b6020831061267a5780518252602082019150602081019050602083039250612657565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012060405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101847effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260010183815260200182815260200194505050505060405160208183030381529060405290509392505050565b60008061282d8573ffffffffffffffffffffffffffffffffffffffff1663f698da256040518163ffffffff1660e01b815260040160206040518083038186803b15801561278357600080fd5b505afa158015612797573d6000803e3d6000fd5b505050506040513d60208110156127ad57600080fd5b81019080805190602001909291905050507f60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca60001b8660405160200180828152602001915050604051602081830303815290604052805190602001206040516020018082815260200191505060405160208183030381529060405261262f565b905060008180519060200120905060008451141561294c5760008673ffffffffffffffffffffffffffffffffffffffff16635ae6bd37836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561289857600080fd5b505afa1580156128ac573d6000803e3d6000fd5b505050506040513d60208110156128c257600080fd5b81019080805190602001909291905050501415612947576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f48617368206e6f7420617070726f76656400000000000000000000000000000081525060200191505060405180910390fd5b612a42565b8573ffffffffffffffffffffffffffffffffffffffff1663f855438b600083876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156129dd5780820151818401526020810190506129c2565b50505050905090810190601f168015612a0a5780820380516001836020036101000a031916815260200191505b5094505050505060006040518083038186803b158015612a2957600080fd5b505afa158015612a3d573d6000803e3d6000fd5b505050505b631626ba7e60e01b925050509392505050565b6000601460003690501015612ad2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f496e76616c69642063616c6c64617461206c656e67746800000000000000000081525060200191505060405180910390fd5b601436033560601c905090565b60008060008360f81c15925073ffffffffffffffffffffffffffffffffffffffff841690508360a01c60a81b91509193909250565b600082612b41577f0100000000000000000000000000000000000000000000000000000000000000612b44565b60005b8273ffffffffffffffffffffffffffffffffffffffff161760001b905092915050565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000857bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019081526020016000205490506000612c0984612141565b915050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612c48576000801b93505b83836000877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001908152602001600020819055508573ffffffffffffffffffffffffffffffffffffffff167fe6e8ad7e5547dc860775f9f0638e195a4751a4cfbb7fd2ab60a52eb6c07082ec86848760405180847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001838152602001828152602001935050505060405180910390a2505050505050565b60008085851115612d3257600080fd5b83861115612d3f57600080fd5b600185028301915084860390509450949250505056fea26469706673582212205ef3d30ba0400c51cffe4d66020db6f8e0785fe148ba70591b7ba7b501e647ac64736f6c63430007060033",
}

// ExtensibleFallbackHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtensibleFallbackHandlerMetaData.ABI instead.
var ExtensibleFallbackHandlerABI = ExtensibleFallbackHandlerMetaData.ABI

// ExtensibleFallbackHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtensibleFallbackHandlerMetaData.Bin instead.
var ExtensibleFallbackHandlerBin = ExtensibleFallbackHandlerMetaData.Bin

// DeployExtensibleFallbackHandler deploys a new Ethereum contract, binding an instance of ExtensibleFallbackHandler to it.
func DeployExtensibleFallbackHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExtensibleFallbackHandler, error) {
	parsed, err := ExtensibleFallbackHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtensibleFallbackHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExtensibleFallbackHandler{ExtensibleFallbackHandlerCaller: ExtensibleFallbackHandlerCaller{contract: contract}, ExtensibleFallbackHandlerTransactor: ExtensibleFallbackHandlerTransactor{contract: contract}, ExtensibleFallbackHandlerFilterer: ExtensibleFallbackHandlerFilterer{contract: contract}}, nil
}

// ExtensibleFallbackHandler is an auto generated Go binding around an Ethereum contract.
type ExtensibleFallbackHandler struct {
	ExtensibleFallbackHandlerCaller     // Read-only binding to the contract
	ExtensibleFallbackHandlerTransactor // Write-only binding to the contract
	ExtensibleFallbackHandlerFilterer   // Log filterer for contract events
}

// ExtensibleFallbackHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtensibleFallbackHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensibleFallbackHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtensibleFallbackHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensibleFallbackHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtensibleFallbackHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensibleFallbackHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtensibleFallbackHandlerSession struct {
	Contract     *ExtensibleFallbackHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ExtensibleFallbackHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtensibleFallbackHandlerCallerSession struct {
	Contract *ExtensibleFallbackHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ExtensibleFallbackHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtensibleFallbackHandlerTransactorSession struct {
	Contract     *ExtensibleFallbackHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ExtensibleFallbackHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtensibleFallbackHandlerRaw struct {
	Contract *ExtensibleFallbackHandler // Generic contract binding to access the raw methods on
}

// ExtensibleFallbackHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtensibleFallbackHandlerCallerRaw struct {
	Contract *ExtensibleFallbackHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ExtensibleFallbackHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtensibleFallbackHandlerTransactorRaw struct {
	Contract *ExtensibleFallbackHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtensibleFallbackHandler creates a new instance of ExtensibleFallbackHandler, bound to a specific deployed contract.
func NewExtensibleFallbackHandler(address common.Address, backend bind.ContractBackend) (*ExtensibleFallbackHandler, error) {
	contract, err := bindExtensibleFallbackHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandler{ExtensibleFallbackHandlerCaller: ExtensibleFallbackHandlerCaller{contract: contract}, ExtensibleFallbackHandlerTransactor: ExtensibleFallbackHandlerTransactor{contract: contract}, ExtensibleFallbackHandlerFilterer: ExtensibleFallbackHandlerFilterer{contract: contract}}, nil
}

// NewExtensibleFallbackHandlerCaller creates a new read-only instance of ExtensibleFallbackHandler, bound to a specific deployed contract.
func NewExtensibleFallbackHandlerCaller(address common.Address, caller bind.ContractCaller) (*ExtensibleFallbackHandlerCaller, error) {
	contract, err := bindExtensibleFallbackHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerCaller{contract: contract}, nil
}

// NewExtensibleFallbackHandlerTransactor creates a new write-only instance of ExtensibleFallbackHandler, bound to a specific deployed contract.
func NewExtensibleFallbackHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtensibleFallbackHandlerTransactor, error) {
	contract, err := bindExtensibleFallbackHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerTransactor{contract: contract}, nil
}

// NewExtensibleFallbackHandlerFilterer creates a new log filterer instance of ExtensibleFallbackHandler, bound to a specific deployed contract.
func NewExtensibleFallbackHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtensibleFallbackHandlerFilterer, error) {
	contract, err := bindExtensibleFallbackHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerFilterer{contract: contract}, nil
}

// bindExtensibleFallbackHandler binds a generic wrapper to an already deployed contract.
func bindExtensibleFallbackHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtensibleFallbackHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensibleFallbackHandler.Contract.ExtensibleFallbackHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.ExtensibleFallbackHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.ExtensibleFallbackHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensibleFallbackHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.contract.Transact(opts, method, params...)
}

// DomainVerifiers is a free data retrieval call binding the contract method 0x51cad5ee.
//
// Solidity: function domainVerifiers(address , bytes32 ) view returns(address)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) DomainVerifiers(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "domainVerifiers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DomainVerifiers is a free data retrieval call binding the contract method 0x51cad5ee.
//
// Solidity: function domainVerifiers(address , bytes32 ) view returns(address)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) DomainVerifiers(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _ExtensibleFallbackHandler.Contract.DomainVerifiers(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1)
}

// DomainVerifiers is a free data retrieval call binding the contract method 0x51cad5ee.
//
// Solidity: function domainVerifiers(address , bytes32 ) view returns(address)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) DomainVerifiers(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _ExtensibleFallbackHandler.Contract.DomainVerifiers(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes signature) view returns(bytes4 magic)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "isValidSignature", _hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes signature) view returns(bytes4 magic)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) IsValidSignature(_hash [32]byte, signature []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.IsValidSignature(&_ExtensibleFallbackHandler.CallOpts, _hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes signature) view returns(bytes4 magic)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) IsValidSignature(_hash [32]byte, signature []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.IsValidSignature(&_ExtensibleFallbackHandler.CallOpts, _hash, signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.OnERC1155BatchReceived(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.OnERC1155BatchReceived(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.OnERC1155Received(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.OnERC1155Received(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.OnERC721Received(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _ExtensibleFallbackHandler.Contract.OnERC721Received(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// SafeInterfaces is a free data retrieval call binding the contract method 0xb435a13b.
//
// Solidity: function safeInterfaces(address , bytes4 ) view returns(bool)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) SafeInterfaces(opts *bind.CallOpts, arg0 common.Address, arg1 [4]byte) (bool, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "safeInterfaces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SafeInterfaces is a free data retrieval call binding the contract method 0xb435a13b.
//
// Solidity: function safeInterfaces(address , bytes4 ) view returns(bool)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) SafeInterfaces(arg0 common.Address, arg1 [4]byte) (bool, error) {
	return _ExtensibleFallbackHandler.Contract.SafeInterfaces(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1)
}

// SafeInterfaces is a free data retrieval call binding the contract method 0xb435a13b.
//
// Solidity: function safeInterfaces(address , bytes4 ) view returns(bool)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) SafeInterfaces(arg0 common.Address, arg1 [4]byte) (bool, error) {
	return _ExtensibleFallbackHandler.Contract.SafeInterfaces(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1)
}

// SafeMethods is a free data retrieval call binding the contract method 0x0a3fe054.
//
// Solidity: function safeMethods(address , bytes4 ) view returns(bytes32)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) SafeMethods(opts *bind.CallOpts, arg0 common.Address, arg1 [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "safeMethods", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SafeMethods is a free data retrieval call binding the contract method 0x0a3fe054.
//
// Solidity: function safeMethods(address , bytes4 ) view returns(bytes32)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) SafeMethods(arg0 common.Address, arg1 [4]byte) ([32]byte, error) {
	return _ExtensibleFallbackHandler.Contract.SafeMethods(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1)
}

// SafeMethods is a free data retrieval call binding the contract method 0x0a3fe054.
//
// Solidity: function safeMethods(address , bytes4 ) view returns(bytes32)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) SafeMethods(arg0 common.Address, arg1 [4]byte) ([32]byte, error) {
	return _ExtensibleFallbackHandler.Contract.SafeMethods(&_ExtensibleFallbackHandler.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExtensibleFallbackHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExtensibleFallbackHandler.Contract.SupportsInterface(&_ExtensibleFallbackHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExtensibleFallbackHandler.Contract.SupportsInterface(&_ExtensibleFallbackHandler.CallOpts, interfaceId)
}

// AddSupportedInterfaceBatch is a paid mutator transaction binding the contract method 0x327b533c.
//
// Solidity: function addSupportedInterfaceBatch(bytes4 _interfaceId, bytes32[] handlerWithSelectors) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactor) AddSupportedInterfaceBatch(opts *bind.TransactOpts, _interfaceId [4]byte, handlerWithSelectors [][32]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.contract.Transact(opts, "addSupportedInterfaceBatch", _interfaceId, handlerWithSelectors)
}

// AddSupportedInterfaceBatch is a paid mutator transaction binding the contract method 0x327b533c.
//
// Solidity: function addSupportedInterfaceBatch(bytes4 _interfaceId, bytes32[] handlerWithSelectors) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) AddSupportedInterfaceBatch(_interfaceId [4]byte, handlerWithSelectors [][32]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.AddSupportedInterfaceBatch(&_ExtensibleFallbackHandler.TransactOpts, _interfaceId, handlerWithSelectors)
}

// AddSupportedInterfaceBatch is a paid mutator transaction binding the contract method 0x327b533c.
//
// Solidity: function addSupportedInterfaceBatch(bytes4 _interfaceId, bytes32[] handlerWithSelectors) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorSession) AddSupportedInterfaceBatch(_interfaceId [4]byte, handlerWithSelectors [][32]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.AddSupportedInterfaceBatch(&_ExtensibleFallbackHandler.TransactOpts, _interfaceId, handlerWithSelectors)
}

// RemoveSupportedInterfaceBatch is a paid mutator transaction binding the contract method 0x61f5401b.
//
// Solidity: function removeSupportedInterfaceBatch(bytes4 _interfaceId, bytes4[] selectors) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactor) RemoveSupportedInterfaceBatch(opts *bind.TransactOpts, _interfaceId [4]byte, selectors [][4]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.contract.Transact(opts, "removeSupportedInterfaceBatch", _interfaceId, selectors)
}

// RemoveSupportedInterfaceBatch is a paid mutator transaction binding the contract method 0x61f5401b.
//
// Solidity: function removeSupportedInterfaceBatch(bytes4 _interfaceId, bytes4[] selectors) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) RemoveSupportedInterfaceBatch(_interfaceId [4]byte, selectors [][4]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.RemoveSupportedInterfaceBatch(&_ExtensibleFallbackHandler.TransactOpts, _interfaceId, selectors)
}

// RemoveSupportedInterfaceBatch is a paid mutator transaction binding the contract method 0x61f5401b.
//
// Solidity: function removeSupportedInterfaceBatch(bytes4 _interfaceId, bytes4[] selectors) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorSession) RemoveSupportedInterfaceBatch(_interfaceId [4]byte, selectors [][4]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.RemoveSupportedInterfaceBatch(&_ExtensibleFallbackHandler.TransactOpts, _interfaceId, selectors)
}

// SetDomainVerifier is a paid mutator transaction binding the contract method 0x3365582c.
//
// Solidity: function setDomainVerifier(bytes32 domainSeparator, address newVerifier) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactor) SetDomainVerifier(opts *bind.TransactOpts, domainSeparator [32]byte, newVerifier common.Address) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.contract.Transact(opts, "setDomainVerifier", domainSeparator, newVerifier)
}

// SetDomainVerifier is a paid mutator transaction binding the contract method 0x3365582c.
//
// Solidity: function setDomainVerifier(bytes32 domainSeparator, address newVerifier) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) SetDomainVerifier(domainSeparator [32]byte, newVerifier common.Address) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.SetDomainVerifier(&_ExtensibleFallbackHandler.TransactOpts, domainSeparator, newVerifier)
}

// SetDomainVerifier is a paid mutator transaction binding the contract method 0x3365582c.
//
// Solidity: function setDomainVerifier(bytes32 domainSeparator, address newVerifier) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorSession) SetDomainVerifier(domainSeparator [32]byte, newVerifier common.Address) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.SetDomainVerifier(&_ExtensibleFallbackHandler.TransactOpts, domainSeparator, newVerifier)
}

// SetSafeMethod is a paid mutator transaction binding the contract method 0x64f95acc.
//
// Solidity: function setSafeMethod(bytes4 selector, bytes32 newMethod) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactor) SetSafeMethod(opts *bind.TransactOpts, selector [4]byte, newMethod [32]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.contract.Transact(opts, "setSafeMethod", selector, newMethod)
}

// SetSafeMethod is a paid mutator transaction binding the contract method 0x64f95acc.
//
// Solidity: function setSafeMethod(bytes4 selector, bytes32 newMethod) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) SetSafeMethod(selector [4]byte, newMethod [32]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.SetSafeMethod(&_ExtensibleFallbackHandler.TransactOpts, selector, newMethod)
}

// SetSafeMethod is a paid mutator transaction binding the contract method 0x64f95acc.
//
// Solidity: function setSafeMethod(bytes4 selector, bytes32 newMethod) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorSession) SetSafeMethod(selector [4]byte, newMethod [32]byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.SetSafeMethod(&_ExtensibleFallbackHandler.TransactOpts, selector, newMethod)
}

// SetSupportedInterface is a paid mutator transaction binding the contract method 0x7f73528b.
//
// Solidity: function setSupportedInterface(bytes4 interfaceId, bool supported) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactor) SetSupportedInterface(opts *bind.TransactOpts, interfaceId [4]byte, supported bool) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.contract.Transact(opts, "setSupportedInterface", interfaceId, supported)
}

// SetSupportedInterface is a paid mutator transaction binding the contract method 0x7f73528b.
//
// Solidity: function setSupportedInterface(bytes4 interfaceId, bool supported) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) SetSupportedInterface(interfaceId [4]byte, supported bool) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.SetSupportedInterface(&_ExtensibleFallbackHandler.TransactOpts, interfaceId, supported)
}

// SetSupportedInterface is a paid mutator transaction binding the contract method 0x7f73528b.
//
// Solidity: function setSupportedInterface(bytes4 interfaceId, bool supported) returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorSession) SetSupportedInterface(interfaceId [4]byte, supported bool) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.SetSupportedInterface(&_ExtensibleFallbackHandler.TransactOpts, interfaceId, supported)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.Fallback(&_ExtensibleFallbackHandler.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ExtensibleFallbackHandler.Contract.Fallback(&_ExtensibleFallbackHandler.TransactOpts, calldata)
}

// ExtensibleFallbackHandlerAddedInterfaceIterator is returned from FilterAddedInterface and is used to iterate over the raw logs and unpacked data for AddedInterface events raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerAddedInterfaceIterator struct {
	Event *ExtensibleFallbackHandlerAddedInterface // Event containing the contract specifics and raw log

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
func (it *ExtensibleFallbackHandlerAddedInterfaceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensibleFallbackHandlerAddedInterface)
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
		it.Event = new(ExtensibleFallbackHandlerAddedInterface)
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
func (it *ExtensibleFallbackHandlerAddedInterfaceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensibleFallbackHandlerAddedInterfaceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensibleFallbackHandlerAddedInterface represents a AddedInterface event raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerAddedInterface struct {
	Safe        common.Address
	InterfaceId [4]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddedInterface is a free log retrieval operation binding the contract event 0x3d5024c13f12fa602dbf52b1979058940c224ebf170c83a4e358725ae50db36d.
//
// Solidity: event AddedInterface(address indexed safe, bytes4 interfaceId)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) FilterAddedInterface(opts *bind.FilterOpts, safe []common.Address) (*ExtensibleFallbackHandlerAddedInterfaceIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.FilterLogs(opts, "AddedInterface", safeRule)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerAddedInterfaceIterator{contract: _ExtensibleFallbackHandler.contract, event: "AddedInterface", logs: logs, sub: sub}, nil
}

// WatchAddedInterface is a free log subscription operation binding the contract event 0x3d5024c13f12fa602dbf52b1979058940c224ebf170c83a4e358725ae50db36d.
//
// Solidity: event AddedInterface(address indexed safe, bytes4 interfaceId)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) WatchAddedInterface(opts *bind.WatchOpts, sink chan<- *ExtensibleFallbackHandlerAddedInterface, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.WatchLogs(opts, "AddedInterface", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensibleFallbackHandlerAddedInterface)
				if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "AddedInterface", log); err != nil {
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

// ParseAddedInterface is a log parse operation binding the contract event 0x3d5024c13f12fa602dbf52b1979058940c224ebf170c83a4e358725ae50db36d.
//
// Solidity: event AddedInterface(address indexed safe, bytes4 interfaceId)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) ParseAddedInterface(log types.Log) (*ExtensibleFallbackHandlerAddedInterface, error) {
	event := new(ExtensibleFallbackHandlerAddedInterface)
	if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "AddedInterface", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensibleFallbackHandlerChangedDomainVerifierIterator is returned from FilterChangedDomainVerifier and is used to iterate over the raw logs and unpacked data for ChangedDomainVerifier events raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerChangedDomainVerifierIterator struct {
	Event *ExtensibleFallbackHandlerChangedDomainVerifier // Event containing the contract specifics and raw log

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
func (it *ExtensibleFallbackHandlerChangedDomainVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensibleFallbackHandlerChangedDomainVerifier)
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
		it.Event = new(ExtensibleFallbackHandlerChangedDomainVerifier)
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
func (it *ExtensibleFallbackHandlerChangedDomainVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensibleFallbackHandlerChangedDomainVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensibleFallbackHandlerChangedDomainVerifier represents a ChangedDomainVerifier event raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerChangedDomainVerifier struct {
	Safe            common.Address
	DomainSeparator [32]byte
	OldVerifier     common.Address
	NewVerifier     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangedDomainVerifier is a free log retrieval operation binding the contract event 0x06341ac2f62eb79165de8c8b54ab89a8825b12746c2b98e01dff1007837d2ba8.
//
// Solidity: event ChangedDomainVerifier(address indexed safe, bytes32 domainSeparator, address oldVerifier, address newVerifier)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) FilterChangedDomainVerifier(opts *bind.FilterOpts, safe []common.Address) (*ExtensibleFallbackHandlerChangedDomainVerifierIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.FilterLogs(opts, "ChangedDomainVerifier", safeRule)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerChangedDomainVerifierIterator{contract: _ExtensibleFallbackHandler.contract, event: "ChangedDomainVerifier", logs: logs, sub: sub}, nil
}

// WatchChangedDomainVerifier is a free log subscription operation binding the contract event 0x06341ac2f62eb79165de8c8b54ab89a8825b12746c2b98e01dff1007837d2ba8.
//
// Solidity: event ChangedDomainVerifier(address indexed safe, bytes32 domainSeparator, address oldVerifier, address newVerifier)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) WatchChangedDomainVerifier(opts *bind.WatchOpts, sink chan<- *ExtensibleFallbackHandlerChangedDomainVerifier, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.WatchLogs(opts, "ChangedDomainVerifier", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensibleFallbackHandlerChangedDomainVerifier)
				if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "ChangedDomainVerifier", log); err != nil {
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

// ParseChangedDomainVerifier is a log parse operation binding the contract event 0x06341ac2f62eb79165de8c8b54ab89a8825b12746c2b98e01dff1007837d2ba8.
//
// Solidity: event ChangedDomainVerifier(address indexed safe, bytes32 domainSeparator, address oldVerifier, address newVerifier)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) ParseChangedDomainVerifier(log types.Log) (*ExtensibleFallbackHandlerChangedDomainVerifier, error) {
	event := new(ExtensibleFallbackHandlerChangedDomainVerifier)
	if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "ChangedDomainVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensibleFallbackHandlerChangedSafeMethodIterator is returned from FilterChangedSafeMethod and is used to iterate over the raw logs and unpacked data for ChangedSafeMethod events raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerChangedSafeMethodIterator struct {
	Event *ExtensibleFallbackHandlerChangedSafeMethod // Event containing the contract specifics and raw log

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
func (it *ExtensibleFallbackHandlerChangedSafeMethodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensibleFallbackHandlerChangedSafeMethod)
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
		it.Event = new(ExtensibleFallbackHandlerChangedSafeMethod)
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
func (it *ExtensibleFallbackHandlerChangedSafeMethodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensibleFallbackHandlerChangedSafeMethodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensibleFallbackHandlerChangedSafeMethod represents a ChangedSafeMethod event raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerChangedSafeMethod struct {
	Safe      common.Address
	Selector  [4]byte
	OldMethod [32]byte
	NewMethod [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedSafeMethod is a free log retrieval operation binding the contract event 0xe6e8ad7e5547dc860775f9f0638e195a4751a4cfbb7fd2ab60a52eb6c07082ec.
//
// Solidity: event ChangedSafeMethod(address indexed safe, bytes4 selector, bytes32 oldMethod, bytes32 newMethod)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) FilterChangedSafeMethod(opts *bind.FilterOpts, safe []common.Address) (*ExtensibleFallbackHandlerChangedSafeMethodIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.FilterLogs(opts, "ChangedSafeMethod", safeRule)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerChangedSafeMethodIterator{contract: _ExtensibleFallbackHandler.contract, event: "ChangedSafeMethod", logs: logs, sub: sub}, nil
}

// WatchChangedSafeMethod is a free log subscription operation binding the contract event 0xe6e8ad7e5547dc860775f9f0638e195a4751a4cfbb7fd2ab60a52eb6c07082ec.
//
// Solidity: event ChangedSafeMethod(address indexed safe, bytes4 selector, bytes32 oldMethod, bytes32 newMethod)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) WatchChangedSafeMethod(opts *bind.WatchOpts, sink chan<- *ExtensibleFallbackHandlerChangedSafeMethod, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.WatchLogs(opts, "ChangedSafeMethod", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensibleFallbackHandlerChangedSafeMethod)
				if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "ChangedSafeMethod", log); err != nil {
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

// ParseChangedSafeMethod is a log parse operation binding the contract event 0xe6e8ad7e5547dc860775f9f0638e195a4751a4cfbb7fd2ab60a52eb6c07082ec.
//
// Solidity: event ChangedSafeMethod(address indexed safe, bytes4 selector, bytes32 oldMethod, bytes32 newMethod)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) ParseChangedSafeMethod(log types.Log) (*ExtensibleFallbackHandlerChangedSafeMethod, error) {
	event := new(ExtensibleFallbackHandlerChangedSafeMethod)
	if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "ChangedSafeMethod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensibleFallbackHandlerRemovedInterfaceIterator is returned from FilterRemovedInterface and is used to iterate over the raw logs and unpacked data for RemovedInterface events raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerRemovedInterfaceIterator struct {
	Event *ExtensibleFallbackHandlerRemovedInterface // Event containing the contract specifics and raw log

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
func (it *ExtensibleFallbackHandlerRemovedInterfaceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensibleFallbackHandlerRemovedInterface)
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
		it.Event = new(ExtensibleFallbackHandlerRemovedInterface)
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
func (it *ExtensibleFallbackHandlerRemovedInterfaceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensibleFallbackHandlerRemovedInterfaceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensibleFallbackHandlerRemovedInterface represents a RemovedInterface event raised by the ExtensibleFallbackHandler contract.
type ExtensibleFallbackHandlerRemovedInterface struct {
	Safe        common.Address
	InterfaceId [4]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemovedInterface is a free log retrieval operation binding the contract event 0xf159d834f487747c1b3f17e2107743e42e6eecff444d894e511c18943072b29f.
//
// Solidity: event RemovedInterface(address indexed safe, bytes4 interfaceId)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) FilterRemovedInterface(opts *bind.FilterOpts, safe []common.Address) (*ExtensibleFallbackHandlerRemovedInterfaceIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.FilterLogs(opts, "RemovedInterface", safeRule)
	if err != nil {
		return nil, err
	}
	return &ExtensibleFallbackHandlerRemovedInterfaceIterator{contract: _ExtensibleFallbackHandler.contract, event: "RemovedInterface", logs: logs, sub: sub}, nil
}

// WatchRemovedInterface is a free log subscription operation binding the contract event 0xf159d834f487747c1b3f17e2107743e42e6eecff444d894e511c18943072b29f.
//
// Solidity: event RemovedInterface(address indexed safe, bytes4 interfaceId)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) WatchRemovedInterface(opts *bind.WatchOpts, sink chan<- *ExtensibleFallbackHandlerRemovedInterface, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _ExtensibleFallbackHandler.contract.WatchLogs(opts, "RemovedInterface", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensibleFallbackHandlerRemovedInterface)
				if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "RemovedInterface", log); err != nil {
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

// ParseRemovedInterface is a log parse operation binding the contract event 0xf159d834f487747c1b3f17e2107743e42e6eecff444d894e511c18943072b29f.
//
// Solidity: event RemovedInterface(address indexed safe, bytes4 interfaceId)
func (_ExtensibleFallbackHandler *ExtensibleFallbackHandlerFilterer) ParseRemovedInterface(log types.Log) (*ExtensibleFallbackHandlerRemovedInterface, error) {
	event := new(ExtensibleFallbackHandlerRemovedInterface)
	if err := _ExtensibleFallbackHandler.contract.UnpackLog(event, "RemovedInterface", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HandlerContextMetaData contains all meta data concerning the HandlerContext contract.
var HandlerContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// HandlerContextABI is the input ABI used to generate the binding from.
// Deprecated: Use HandlerContextMetaData.ABI instead.
var HandlerContextABI = HandlerContextMetaData.ABI

// HandlerContext is an auto generated Go binding around an Ethereum contract.
type HandlerContext struct {
	HandlerContextCaller     // Read-only binding to the contract
	HandlerContextTransactor // Write-only binding to the contract
	HandlerContextFilterer   // Log filterer for contract events
}

// HandlerContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type HandlerContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HandlerContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HandlerContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HandlerContextSession struct {
	Contract     *HandlerContext   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HandlerContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HandlerContextCallerSession struct {
	Contract *HandlerContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HandlerContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HandlerContextTransactorSession struct {
	Contract     *HandlerContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HandlerContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type HandlerContextRaw struct {
	Contract *HandlerContext // Generic contract binding to access the raw methods on
}

// HandlerContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HandlerContextCallerRaw struct {
	Contract *HandlerContextCaller // Generic read-only contract binding to access the raw methods on
}

// HandlerContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HandlerContextTransactorRaw struct {
	Contract *HandlerContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHandlerContext creates a new instance of HandlerContext, bound to a specific deployed contract.
func NewHandlerContext(address common.Address, backend bind.ContractBackend) (*HandlerContext, error) {
	contract, err := bindHandlerContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HandlerContext{HandlerContextCaller: HandlerContextCaller{contract: contract}, HandlerContextTransactor: HandlerContextTransactor{contract: contract}, HandlerContextFilterer: HandlerContextFilterer{contract: contract}}, nil
}

// NewHandlerContextCaller creates a new read-only instance of HandlerContext, bound to a specific deployed contract.
func NewHandlerContextCaller(address common.Address, caller bind.ContractCaller) (*HandlerContextCaller, error) {
	contract, err := bindHandlerContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerContextCaller{contract: contract}, nil
}

// NewHandlerContextTransactor creates a new write-only instance of HandlerContext, bound to a specific deployed contract.
func NewHandlerContextTransactor(address common.Address, transactor bind.ContractTransactor) (*HandlerContextTransactor, error) {
	contract, err := bindHandlerContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerContextTransactor{contract: contract}, nil
}

// NewHandlerContextFilterer creates a new log filterer instance of HandlerContext, bound to a specific deployed contract.
func NewHandlerContextFilterer(address common.Address, filterer bind.ContractFilterer) (*HandlerContextFilterer, error) {
	contract, err := bindHandlerContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HandlerContextFilterer{contract: contract}, nil
}

// bindHandlerContext binds a generic wrapper to an already deployed contract.
func bindHandlerContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HandlerContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerContext *HandlerContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerContext.Contract.HandlerContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerContext *HandlerContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerContext.Contract.HandlerContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerContext *HandlerContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerContext.Contract.HandlerContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerContext *HandlerContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerContext.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerContext *HandlerContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerContext.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerContext *HandlerContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerContext.Contract.contract.Transact(opts, method, params...)
}

// TokenCallbackHandlerMetaData contains all meta data concerning the TokenCallbackHandler contract.
var TokenCallbackHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109a3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100565760003560e01c806223de291461005b57806301ffc9a714610193578063150b7a02146101f6578063bc197c81146102ec578063f23a6e6114610482575b600080fd5b610191600480360360c081101561007157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156100f857600080fd5b82018360208201111561010a57600080fd5b8035906020019184600183028401116401000000008311171561012c57600080fd5b90919293919293908035906020019064010000000081111561014d57600080fd5b82018360208201111561015f57600080fd5b8035906020019184600183028401116401000000008311171561018157600080fd5b9091929391929390505050610582565b005b6101de600480360360208110156101a957600080fd5b8101908080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916906020019092919050505061058c565b60405180821515815260200191505060405180910390f35b6102b76004803603608081101561020c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561027357600080fd5b82018360208201111561028557600080fd5b803590602001918460018302840111640100000000831117156102a757600080fd5b90919293919293905050506106c6565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b61044d600480360360a081101561030257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561035f57600080fd5b82018360208201111561037157600080fd5b8035906020019184602083028401116401000000008311171561039357600080fd5b9091929391929390803590602001906401000000008111156103b457600080fd5b8201836020820111156103c657600080fd5b803590602001918460208302840111640100000000831117156103e857600080fd5b90919293919293908035906020019064010000000081111561040957600080fd5b82018360208201111561041b57600080fd5b8035906020019184600183028401116401000000008311171561043d57600080fd5b90919293919293905050506106e3565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b61054d600480360360a081101561049857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561050957600080fd5b82018360208201111561051b57600080fd5b8035906020019184600183028401116401000000008311171561053d57600080fd5b9091929391929390505050610703565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b5050505050505050565b60007f4e2312e0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061065757507f150b7a02000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806106bf57507f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b60006106d0610721565b63150b7a0260e01b905095945050505050565b60006106ed610721565b63bc197c8160e01b905098975050505050505050565b600061070d610721565b63f23a6e6160e01b90509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff16635624b25b7f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d560001b60001c60016040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156107a357600080fd5b505afa1580156107b7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156107e157600080fd5b810190808051604051939291908464010000000082111561080157600080fd5b8382019150602082018581111561081757600080fd5b825186600182028301116401000000008211171561083457600080fd5b8083526020830192505050908051906020019080838360005b8381101561086857808201518184015260208101905061084d565b50505050905090810190601f1680156108955780820380516001836020036101000a031916815260200191505b50604052505050905060008180602001905160208110156108b557600080fd5b810190808051906020019092919050505090503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610969576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6e6f7420612066616c6c6261636b2063616c6c0000000000000000000000000081525060200191505060405180910390fd5b505056fea2646970667358221220899bc941ea272892b5792af91da62ec82a4d59648b5ad142153c23d429b0f9ed64736f6c63430007060033",
}

// TokenCallbackHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenCallbackHandlerMetaData.ABI instead.
var TokenCallbackHandlerABI = TokenCallbackHandlerMetaData.ABI

// TokenCallbackHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenCallbackHandlerMetaData.Bin instead.
var TokenCallbackHandlerBin = TokenCallbackHandlerMetaData.Bin

// DeployTokenCallbackHandler deploys a new Ethereum contract, binding an instance of TokenCallbackHandler to it.
func DeployTokenCallbackHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenCallbackHandler, error) {
	parsed, err := TokenCallbackHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenCallbackHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenCallbackHandler{TokenCallbackHandlerCaller: TokenCallbackHandlerCaller{contract: contract}, TokenCallbackHandlerTransactor: TokenCallbackHandlerTransactor{contract: contract}, TokenCallbackHandlerFilterer: TokenCallbackHandlerFilterer{contract: contract}}, nil
}

// TokenCallbackHandler is an auto generated Go binding around an Ethereum contract.
type TokenCallbackHandler struct {
	TokenCallbackHandlerCaller     // Read-only binding to the contract
	TokenCallbackHandlerTransactor // Write-only binding to the contract
	TokenCallbackHandlerFilterer   // Log filterer for contract events
}

// TokenCallbackHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCallbackHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenCallbackHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenCallbackHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenCallbackHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenCallbackHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenCallbackHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenCallbackHandlerSession struct {
	Contract     *TokenCallbackHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenCallbackHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallbackHandlerCallerSession struct {
	Contract *TokenCallbackHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TokenCallbackHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenCallbackHandlerTransactorSession struct {
	Contract     *TokenCallbackHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TokenCallbackHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenCallbackHandlerRaw struct {
	Contract *TokenCallbackHandler // Generic contract binding to access the raw methods on
}

// TokenCallbackHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallbackHandlerCallerRaw struct {
	Contract *TokenCallbackHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenCallbackHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenCallbackHandlerTransactorRaw struct {
	Contract *TokenCallbackHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenCallbackHandler creates a new instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandler(address common.Address, backend bind.ContractBackend) (*TokenCallbackHandler, error) {
	contract, err := bindTokenCallbackHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandler{TokenCallbackHandlerCaller: TokenCallbackHandlerCaller{contract: contract}, TokenCallbackHandlerTransactor: TokenCallbackHandlerTransactor{contract: contract}, TokenCallbackHandlerFilterer: TokenCallbackHandlerFilterer{contract: contract}}, nil
}

// NewTokenCallbackHandlerCaller creates a new read-only instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandlerCaller(address common.Address, caller bind.ContractCaller) (*TokenCallbackHandlerCaller, error) {
	contract, err := bindTokenCallbackHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandlerCaller{contract: contract}, nil
}

// NewTokenCallbackHandlerTransactor creates a new write-only instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenCallbackHandlerTransactor, error) {
	contract, err := bindTokenCallbackHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandlerTransactor{contract: contract}, nil
}

// NewTokenCallbackHandlerFilterer creates a new log filterer instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenCallbackHandlerFilterer, error) {
	contract, err := bindTokenCallbackHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandlerFilterer{contract: contract}, nil
}

// bindTokenCallbackHandler binds a generic wrapper to an already deployed contract.
func bindTokenCallbackHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenCallbackHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenCallbackHandler *TokenCallbackHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenCallbackHandler.Contract.TokenCallbackHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenCallbackHandler *TokenCallbackHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.TokenCallbackHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenCallbackHandler *TokenCallbackHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.TokenCallbackHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenCallbackHandler *TokenCallbackHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenCallbackHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenCallbackHandler *TokenCallbackHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenCallbackHandler *TokenCallbackHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155BatchReceived(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155BatchReceived(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC721Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC721Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenCallbackHandler.Contract.SupportsInterface(&_TokenCallbackHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenCallbackHandler.Contract.SupportsInterface(&_TokenCallbackHandler.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_TokenCallbackHandler *TokenCallbackHandlerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _TokenCallbackHandler.Contract.TokensReceived(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _TokenCallbackHandler.Contract.TokensReceived(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}
