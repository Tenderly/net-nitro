// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractsgen

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

// SafeMetaData contains all meta data concerning the Safe contract.
var SafeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"moduleGuard\",\"type\":\"address\"}],\"name\":\"ChangedModuleGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moduleGuard\",\"type\":\"address\"}],\"name\":\"setModuleGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060016004819055506153cb806100286000396000f3fe6080604052600436106101dc5760003560e01c8063affed0e011610102578063e19a9dd911610095578063f698da2511610064578063f698da25146114c1578063f855438b146114ec578063f8dc5dd9146115de578063ffa1ad741461165957610231565b8063e19a9dd914611363578063e318b52b146113b4578063e75235b814611445578063f08a03231461147057610231565b8063d4d9bdcd116100d1578063d4d9bdcd1461112c578063d8d11f7814611167578063e009cfde146112a1578063e068df371461131257610231565b8063affed0e014610ddc578063b4faba0914610e07578063b63e800d14610eef578063cc2f84521461105f57610231565b80635624b25b1161017a5780636a761202116101495780636a76120214610a1e5780637d83297414610bda578063934f3a1114610c49578063a0e67e2b14610d7057610231565b80635624b25b146108855780635ae6bd3714610943578063610b592514610992578063694e80c3146109e357610231565b80632d9ad53d116101b65780632d9ad53d146105215780632f54bf6e14610588578063468721a7146105ef5780635229073f1461070457610231565b80630d582f131461029957806312fb68e0146102f45780631fcac7f31461042557610231565b36610231573373ffffffffffffffffffffffffffffffffffffffff167f3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d346040518082815260200191505060405180910390a2005b34801561023d57600080fd5b507f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d5548061026a57600080f35b60405136600082373360601b3682015260008060143601836000865af13d6000833e80610295573d82fd5b3d82f35b3480156102a557600080fd5b506102f2600480360360408110156102bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116e9565b005b34801561030057600080fd5b506104236004803603608081101561031757600080fd5b81019080803590602001909291908035906020019064010000000081111561033e57600080fd5b82018360208201111561035057600080fd5b8035906020019184600183028401116401000000008311171561037257600080fd5b90919293919293908035906020019064010000000081111561039357600080fd5b8201836020820111156103a557600080fd5b803590602001918460018302840111640100000000831117156103c757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611a41565b005b34801561043157600080fd5b5061051f6004803603608081101561044857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561048f57600080fd5b8201836020820111156104a157600080fd5b803590602001918460018302840111640100000000831117156104c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611a54565b005b34801561052d57600080fd5b506105706004803603602081101561054457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e76565b60405180821515815260200191505060405180910390f35b34801561059457600080fd5b506105d7600480360360208110156105ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f48565b60405180821515815260200191505060405180910390f35b3480156105fb57600080fd5b506106ec6004803603608081101561061257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561065957600080fd5b82018360208201111561066b57600080fd5b8035906020019184600183028401116401000000008311171561068d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff169060200190929190505050612018565b60405180821515815260200191505060405180910390f35b34801561071057600080fd5b506108016004803603608081101561072757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561076e57600080fd5b82018360208201111561078057600080fd5b803590602001918460018302840111640100000000831117156107a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff169060200190929190505050612071565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561084957808201518184015260208101905061082e565b50505050905090810190601f1680156108765780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561089157600080fd5b506108c8600480360360408110156108a857600080fd5b8101908080359060200190929190803590602001909291905050506120e6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109085780820151818401526020810190506108ed565b50505050905090810190601f1680156109355780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561094f57600080fd5b5061097c6004803603602081101561096657600080fd5b810190808035906020019092919050505061216c565b6040518082815260200191505060405180910390f35b34801561099e57600080fd5b506109e1600480360360208110156109b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612184565b005b3480156109ef57600080fd5b50610a1c60048036036020811015610a0657600080fd5b8101908080359060200190929190505050612482565b005b610bc26004803603610140811015610a3557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190640100000000811115610a7c57600080fd5b820183602082011115610a8e57600080fd5b80359060200191846001830284011164010000000083111715610ab057600080fd5b9091929391929390803560ff169060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610b3c57600080fd5b820183602082011115610b4e57600080fd5b80359060200191846001830284011164010000000083111715610b7057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050612532565b60405180821515815260200191505060405180910390f35b348015610be657600080fd5b50610c3360048036036040811015610bfd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506129c1565b6040518082815260200191505060405180910390f35b348015610c5557600080fd5b50610d6e60048036036060811015610c6c57600080fd5b810190808035906020019092919080359060200190640100000000811115610c9357600080fd5b820183602082011115610ca557600080fd5b80359060200191846001830284011164010000000083111715610cc757600080fd5b909192939192939080359060200190640100000000811115610ce857600080fd5b820183602082011115610cfa57600080fd5b80359060200191846001830284011164010000000083111715610d1c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506129e6565b005b348015610d7c57600080fd5b50610d856129f7565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610dc8578082015181840152602081019050610dad565b505050509050019250505060405180910390f35b348015610de857600080fd5b50610df1612b9e565b6040518082815260200191505060405180910390f35b348015610e1357600080fd5b50610eed60048036036040811015610e2a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610e6757600080fd5b820183602082011115610e7957600080fd5b80359060200191846001830284011164010000000083111715610e9b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050612ba4565b005b348015610efb57600080fd5b5061105d6004803603610100811015610f1357600080fd5b8101908080359060200190640100000000811115610f3057600080fd5b820183602082011115610f4257600080fd5b80359060200191846020830284011164010000000083111715610f6457600080fd5b909192939192939080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610faf57600080fd5b820183602082011115610fc157600080fd5b80359060200191846001830284011164010000000083111715610fe357600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612bcb565b005b34801561106b57600080fd5b506110b86004803603604081101561108257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612d89565b60405180806020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019060200280838360005b838110156111175780820151818401526020810190506110fc565b50505050905001935050505060405180910390f35b34801561113857600080fd5b506111656004803603602081101561114f57600080fd5b8101908080359060200190929190505050613067565b005b34801561117357600080fd5b5061128b600480360361014081101561118b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156111d257600080fd5b8201836020820111156111e457600080fd5b8035906020019184600183028401116401000000008311171561120657600080fd5b9091929391929390803560ff169060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506131c2565b6040518082815260200191505060405180910390f35b3480156112ad57600080fd5b50611310600480360360408110156112c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061326b565b005b34801561131e57600080fd5b506113616004803603602081101561133557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613568565b005b34801561136f57600080fd5b506113b26004803603602081101561138657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613716565b005b3480156113c057600080fd5b50611443600480360360608110156113d757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506138bb565b005b34801561145157600080fd5b5061145a613e03565b6040518082815260200191505060405180910390f35b34801561147c57600080fd5b506114bf6004803603602081101561149357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613e0d565b005b3480156114cd57600080fd5b506114d6613e64565b6040518082815260200191505060405180910390f35b3480156114f857600080fd5b506115dc6004803603606081101561150f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561155657600080fd5b82018360208201111561156857600080fd5b8035906020019184600183028401116401000000008311171561158a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050613ee0565b005b3480156115ea57600080fd5b506116576004803603606081101561160157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613f2c565b005b34801561166557600080fd5b5061166e61427e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156116ae578082015181840152602081019050611693565b50505050905090810190601f1680156116db5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6116f16142b7565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806117585750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b8061178e57503073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b156117bd576117bc7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461187a576118797f4753323034000000000000000000000000000000000000000000000000000000614316565b5b60026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360008154600101919050819055508173ffffffffffffffffffffffffffffffffffffffff167f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2660405160405180910390a28060045414611a3d57611a3c81612482565b5b5050565b611a4d33868484611a54565b5050505050565b611a6860418261435590919063ffffffff16565b82511015611a9a57611a997f4753303230000000000000000000000000000000000000000000000000000000614316565b5b6000808060008060005b86811015611e6a57611ab6888261438f565b8260ff1692508094508195508296505050506000841415611b33578260001c9450611aeb60418861435590919063ffffffff16565b8260001c1015611b1f57611b1e7f4753303231000000000000000000000000000000000000000000000000000000614316565b5b611b2e858a8a8560001c6143be565b611d30565b6001841415611c02578260001c94508473ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1614158015611bce57506000600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b815260200190815260200160002054145b15611bfd57611bfc7f4753303235000000000000000000000000000000000000000000000000000000614316565b5b611d2f565b601e841115611cc75760018960405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c018281526020019150506040516020818303038152906040528051906020012060048603858560405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611cb6573d6000803e3d6000fd5b505050602060405103519450611d2e565b60018985858560405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611d21573d6000803e3d6000fd5b5050506020604051035194505b5b5b8573ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16111580611df65750600073ffffffffffffffffffffffffffffffffffffffff16600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80611e2d5750600173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b15611e5c57611e5b7f4753303236000000000000000000000000000000000000000000000000000000614316565b5b849550806001019050611aa4565b50505050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff16600173ffffffffffffffffffffffffffffffffffffffff1614158015611f415750600073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b9050919050565b6000600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806120105750600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b159050919050565b6000806000612029878787876145e7565b9150915061205a878787877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61487f565b92506120678282856148cb565b5050949350505050565b60006060600080612084888888886145e7565b915091506120b5888888887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61487f565b9350604051925060203d0183016040523d83523d6000602085013e6120db8282866148cb565b505094509492505050565b60606000600583901b67ffffffffffffffff8111801561210557600080fd5b506040519080825280601f01601f1916602001820160405280156121385781602001600182028036833780820191505090505b50905060005b83811015612161578085015480602083026020850101525080600101905061213e565b508091505092915050565b60076020528060005260406000206000915090505481565b61218c6142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806121f35750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15612222576122217f4753313031000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146122df576122de7f4753313032000000000000000000000000000000000000000000000000000000614316565b5b60016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f844060405160405180910390a250565b61248a6142b7565b6003548111156124be576124bd7f4753323031000000000000000000000000000000000000000000000000000000614316565b5b60008114156124f1576124f07f4753323032000000000000000000000000000000000000000000000000000000614316565b5b806004819055507f610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93816040518082815260200191505060405180910390a150565b60006125478c8c8c8c8c8c8c8c8c8c8c614a0c565b600061256c8d8d8d8d8d8d8d8d8d8d60056000815480929190600101919050556131c2565b9050612579338285613ee0565b6000612583614a19565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612769578073ffffffffffffffffffffffffffffffffffffffff166375f0bb528f8f8f8f8f8f8f8f8f8f8f336040518d63ffffffff1660e01b8152600401808d73ffffffffffffffffffffffffffffffffffffffff1681526020018c8152602001806020018a600181111561262657fe5b81526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001806020018473ffffffffffffffffffffffffffffffffffffffff16815260200183810383528d8d82818152602001925080828437600081840152601f19601f820116905080830192505050838103825285818151815260200191508051906020019080838360005b838110156126f85780820151818401526020810190506126dd565b50505050905090810190601f1680156127255780820380516001836020036101000a031916815260200191505b509e505050505050505050505050505050600060405180830381600087803b15801561275057600080fd5b505af1158015612764573d6000803e3d6000fd5b505050505b6101f46127916109c48b01603f60068d901b8161278257fe5b04614a4290919063ffffffff16565b015a10156127c3576127c27f4753303130000000000000000000000000000000000000000000000000000000614316565b5b60005a905061282c8f8f8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e60008d14612821578e612827565b6109c45a035b61487f565b93506128415a82614a5c90919063ffffffff16565b905083158015612851575060008a145b801561285d5750600088145b1561286e576040513d6000823e3d81fd5b60008089111561288857612885828b8b8b8b614a7c565b90505b84156128cb57837f442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e826040518082815260200191505060405180910390a2612904565b837f23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23826040518082815260200191505060405180910390a25b5050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146129b0578073ffffffffffffffffffffffffffffffffffffffff16639327136883856040518363ffffffff1660e01b815260040180838152602001821515815260200192505050600060405180830381600087803b15801561299757600080fd5b505af11580156129ab573d6000803e3d6000fd5b505050505b50509b9a5050505050505050505050565b6008602052816000526040600020602052806000526040600020600091509150505481565b6129f1338583613ee0565b50505050565b6060600060035467ffffffffffffffff81118015612a1457600080fd5b50604051908082528060200260200182016040528015612a435781602001602082028036833780820191505090505b50905060008060026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612b955780838381518110612aee57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816001019150612aad565b82935050505090565b60055481565b600080825160208401855af46040518181523d60208201523d6000604083013e60403d0181fd5b3373ffffffffffffffffffffffffffffffffffffffff167f141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a88b8b8b8b8960405180806020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281038252878782818152602001925060200280828437600081840152601f19601f820116905080830192505050965050505050505060405180910390a2612cd78a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505089614c30565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614612d1557612d1484614fda565b5b612d638787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061505e565b6000821115612d7d57612d7b82600060018685614a7c565b505b50505050505050505050565b60606000600173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614158015612dd05750612dce84611e76565b155b15612dff57612dfe7f4753313035000000000000000000000000000000000000000000000000000000614316565b5b6000831415612e3257612e317f4753313036000000000000000000000000000000000000000000000000000000614316565b5b8267ffffffffffffffff81118015612e4957600080fd5b50604051908082528060200260200182016040528015612e785781602001602082028036833780820191505090505b5091506000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015612f4a5750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b8015612f5557508381105b1561300e5781838281518110612f6757fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150806001019050612ee0565b600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461305c5782600182038151811061305157fe5b602002602001015191505b808352509250929050565b600073ffffffffffffffffffffffffffffffffffffffff16600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613125576131247f4753303330000000000000000000000000000000000000000000000000000000614316565b5b6001600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16817ff2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c60405160405180910390a350565b6000806131cd613e64565b90506040518a8c82378a81207fbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d882528e60208301528d60408301528060608301528a60808301528960a08301528860c08301528760e08301528661010083015285610120830152846101408301526101608220604083015261190182528260208301526042601e83012093505050509b9a5050505050505050505050565b6132736142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806132da5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15613309576133087f4753313031000000000000000000000000000000000000000000000000000000614316565b5b8073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146133c5576133c47f4753313033000000000000000000000000000000000000000000000000000000614316565b5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167faab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace405427660405160405180910390a25050565b6135706142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561367557508073ffffffffffffffffffffffffffffffffffffffff166301ffc9a77f58401ed8000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b815260040180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060206040518083038186803b15801561363857600080fd5b505afa15801561364c573d6000803e3d6000fd5b505050506040513d602081101561366257600080fd5b8101908080519060200190929190505050155b156136a4576136a37f4753333031000000000000000000000000000000000000000000000000000000614316565b5b60007fb104e0b93118902c651344349b610029d694cfdec91c589c91ebafbcd028994760001b90508181558173ffffffffffffffffffffffffffffffffffffffff167fcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb860405160405180910390a25050565b61371e6142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561382357508073ffffffffffffffffffffffffffffffffffffffff166301ffc9a77fe6d7a83a000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b815260040180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060206040518083038186803b1580156137e657600080fd5b505afa1580156137fa573d6000803e3d6000fd5b505050506040513d602081101561381057600080fd5b8101908080519060200190929190505050155b15613852576138517f4753333030000000000000000000000000000000000000000000000000000000614316565b5b807f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c8558073ffffffffffffffffffffffffffffffffffffffff167f1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa260405160405180910390a250565b6138c36142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16148061392a5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b8061396057503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1561398f5761398e7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613a4c57613a4b7f4753323034000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480613ab35750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b15613ae257613ae17f4753323033000000000000000000000000000000000000000000000000000000614316565b5b8173ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613b9e57613b9d7f4753323035000000000000000000000000000000000000000000000000000000614316565b5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf60405160405180910390a28073ffffffffffffffffffffffffffffffffffffffff167f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2660405160405180910390a2505050565b6000600454905090565b613e156142b7565b613e1e81614fda565b8073ffffffffffffffffffffffffffffffffffffffff167f5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b060405160405180910390a250565b6000804690507f47e79534a245952e8b16893a336b85a3d9ea9fa8c573f3d803afb92a7946921860001b8130604051602001808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff16815260200193505050506040516020818303038152906040528051906020012091505090565b600060045490506000811415613f1a57613f197f4753303031000000000000000000000000000000000000000000000000000000614316565b5b613f2684848484611a54565b50505050565b613f346142b7565b80600360008154600190039190508190551015613f7557613f747f4753323031000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480613fdc5750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b1561400b5761400a7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b8173ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146140c7576140c67f4753323035000000000000000000000000000000000000000000000000000000614316565b5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf60405160405180910390a280600454146142795761427881612482565b5b505050565b6040518060400160405280600581526020017f312e352e3000000000000000000000000000000000000000000000000000000081525081565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614314576143137f4753303331000000000000000000000000000000000000000000000000000000614316565b5b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260056024820152816044820152606481fd5b6000808314156143685760009050614389565b600082840290508284828161437957fe5b041461438457600080fd5b809150505b92915050565b6000806000836041026020810186015192506040810186015191506060810186015160001a9350509250925092565b81516143d460208361526890919063ffffffff16565b1115614404576144037f4753303232000000000000000000000000000000000000000000000000000000614316565b5b600060208284010151905082516144378261442960208661526890919063ffffffff16565b61526890919063ffffffff16565b1115614467576144667f4753303233000000000000000000000000000000000000000000000000000000614316565b5b60606020838501019050631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168673ffffffffffffffffffffffffffffffffffffffff16631626ba7e87846040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561450b5780820151818401526020810190506144f0565b50505050905090810190601f1680156145385780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b15801561455657600080fd5b505afa15801561456a573d6000803e3d6000fd5b505050506040513d602081101561458057600080fd5b81019080805190602001909291905050507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146145df576145de7f4753303234000000000000000000000000000000000000000000000000000000614316565b5b505050505050565b6000806145f686868686615287565b6145fe61528d565b9150600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806146c65750600073ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b156146f5576146f47f4753313034000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614614876578173ffffffffffffffffffffffffffffffffffffffff1663728c297287878787336040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018060200184600181111561478f57fe5b81526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b838110156147e85780820151818401526020810190506147cd565b50505050905090810190601f1680156148155780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561483857600080fd5b505af115801561484c573d6000803e3d6000fd5b505050506040513d602081101561486257600080fd5b810190808051906020019092919050505090505b94509492505050565b600060018081111561488d57fe5b83600181111561489957fe5b14156148b2576000808551602087018986f490506148c2565b600080855160208701888a87f190505b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614614975578273ffffffffffffffffffffffffffffffffffffffff16632acc37aa83836040518363ffffffff1660e01b815260040180838152602001821515815260200192505050600060405180830381600087803b15801561495c57600080fd5b505af1158015614970573d6000803e3d6000fd5b505050505b80156149c3573373ffffffffffffffffffffffffffffffffffffffff167f6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb860405160405180910390a2614a07565b3373ffffffffffffffffffffffffffffffffffffffff167facd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd37560405160405180910390a25b505050565b5050505050505050505050565b60007f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c854905090565b600081831015614a525781614a54565b825b905092915050565b600082821115614a6b57600080fd5b600082840390508091505092915050565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614614ab95782614abb565b325b9050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415614bc557614b253a8610614b02573a614b04565b855b614b17888a61526890919063ffffffff16565b61435590919063ffffffff16565b915060008173ffffffffffffffffffffffffffffffffffffffff168360405180600001905060006040518083038185875af1925050503d8060008114614b87576040519150601f19603f3d011682016040523d82523d6000602084013e614b8c565b606091505b5050905080614bbf57614bbe7f4753303131000000000000000000000000000000000000000000000000000000614316565b5b50614c26565b614bea85614bdc888a61526890919063ffffffff16565b61435590919063ffffffff16565b9150614bf78482846152be565b614c2557614c247f4753303132000000000000000000000000000000000000000000000000000000614316565b5b5b5095945050505050565b60006004541115614c6557614c647f4753323030000000000000000000000000000000000000000000000000000000614316565b5b8151811115614c9857614c977f4753323031000000000000000000000000000000000000000000000000000000614316565b5b6000811415614ccb57614cca7f4753323032000000000000000000000000000000000000000000000000000000614316565b5b60006001905060008351905060005b81811015614f46576000858281518110614cf057fe5b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480614d615750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80614d9757503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80614dcd57508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b15614dfc57614dfb7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614eb957614eb87f4753323034000000000000000000000000000000000000000000000000000000614316565b5b80600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080935050806001019050614cda565b506001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806003819055508260048190555050505050565b3073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615038576150377f4753343030000000000000000000000000000000000000000000000000000000614316565b5b807f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d55550565b600073ffffffffffffffffffffffffffffffffffffffff1660016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461511c5761511b7f4753313030000000000000000000000000000000000000000000000000000000614316565b5b6001806000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614615264576151d882615382565b615206576152057f4753303032000000000000000000000000000000000000000000000000000000614316565b5b6152358260008360017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61487f565b615263576152627f4753303030000000000000000000000000000000000000000000000000000000614316565b5b5b5050565b60008082840190508381101561527d57600080fd5b8091505092915050565b50505050565b6000807fb104e0b93118902c651344349b610029d694cfdec91c589c91ebafbcd028994760001b9050805491505090565b60008063a9059cbb8484604051602401808373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050602060008251602084016000896127105a03f13d60008114615365576020811461536d5760009350615378565b819350615378565b600051158215171593505b5050509392505050565b600080823b90506000811191505091905056fea26469706673582212209a2e21511303be354f644bd2370dafe8bffb1ce6653c5075c9142ccc42d1534464736f6c63430007060033",
}

// SafeABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMetaData.ABI instead.
var SafeABI = SafeMetaData.ABI

// SafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMetaData.Bin instead.
var SafeBin = SafeMetaData.Bin

// DeploySafe deploys a new Ethereum contract, binding an instance of Safe to it.
func DeploySafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Safe, error) {
	parsed, err := SafeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Safe{SafeCaller: SafeCaller{contract: contract}, SafeTransactor: SafeTransactor{contract: contract}, SafeFilterer: SafeFilterer{contract: contract}}, nil
}

// Safe is an auto generated Go binding around an Ethereum contract.
type Safe struct {
	SafeCaller     // Read-only binding to the contract
	SafeTransactor // Write-only binding to the contract
	SafeFilterer   // Log filterer for contract events
}

// SafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeSession struct {
	Contract     *Safe             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCallerSession struct {
	Contract *SafeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeTransactorSession struct {
	Contract     *SafeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeRaw struct {
	Contract *Safe // Generic contract binding to access the raw methods on
}

// SafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCallerRaw struct {
	Contract *SafeCaller // Generic read-only contract binding to access the raw methods on
}

// SafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeTransactorRaw struct {
	Contract *SafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafe creates a new instance of Safe, bound to a specific deployed contract.
func NewSafe(address common.Address, backend bind.ContractBackend) (*Safe, error) {
	contract, err := bindSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safe{SafeCaller: SafeCaller{contract: contract}, SafeTransactor: SafeTransactor{contract: contract}, SafeFilterer: SafeFilterer{contract: contract}}, nil
}

// NewSafeCaller creates a new read-only instance of Safe, bound to a specific deployed contract.
func NewSafeCaller(address common.Address, caller bind.ContractCaller) (*SafeCaller, error) {
	contract, err := bindSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCaller{contract: contract}, nil
}

// NewSafeTransactor creates a new write-only instance of Safe, bound to a specific deployed contract.
func NewSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeTransactor, error) {
	contract, err := bindSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeTransactor{contract: contract}, nil
}

// NewSafeFilterer creates a new log filterer instance of Safe, bound to a specific deployed contract.
func NewSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeFilterer, error) {
	contract, err := bindSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeFilterer{contract: contract}, nil
}

// bindSafe binds a generic wrapper to an already deployed contract.
func bindSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.SafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.SafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safe *SafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safe *SafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safe *SafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safe.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Safe *SafeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Safe *SafeSession) VERSION() (string, error) {
	return _Safe.Contract.VERSION(&_Safe.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Safe *SafeCallerSession) VERSION() (string, error) {
	return _Safe.Contract.VERSION(&_Safe.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_Safe *SafeCaller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_Safe *SafeSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _Safe.Contract.ApprovedHashes(&_Safe.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_Safe *SafeCallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _Safe.Contract.ApprovedHashes(&_Safe.CallOpts, arg0, arg1)
}

// CheckNSignatures12fb68e0 is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_Safe *SafeCaller) CheckNSignatures12fb68e0(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures12fb68e0 is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_Safe *SafeSession) CheckNSignatures12fb68e0(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _Safe.Contract.CheckNSignatures12fb68e0(&_Safe.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures12fb68e0 is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_Safe *SafeCallerSession) CheckNSignatures12fb68e0(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _Safe.Contract.CheckNSignatures12fb68e0(&_Safe.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures1fcac7f3 is a free data retrieval call binding the contract method 0x1fcac7f3.
//
// Solidity: function checkNSignatures(address executor, bytes32 dataHash, bytes signatures, uint256 requiredSignatures) view returns()
func (_Safe *SafeCaller) CheckNSignatures1fcac7f3(opts *bind.CallOpts, executor common.Address, dataHash [32]byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "checkNSignatures0", executor, dataHash, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures1fcac7f3 is a free data retrieval call binding the contract method 0x1fcac7f3.
//
// Solidity: function checkNSignatures(address executor, bytes32 dataHash, bytes signatures, uint256 requiredSignatures) view returns()
func (_Safe *SafeSession) CheckNSignatures1fcac7f3(executor common.Address, dataHash [32]byte, signatures []byte, requiredSignatures *big.Int) error {
	return _Safe.Contract.CheckNSignatures1fcac7f3(&_Safe.CallOpts, executor, dataHash, signatures, requiredSignatures)
}

// CheckNSignatures1fcac7f3 is a free data retrieval call binding the contract method 0x1fcac7f3.
//
// Solidity: function checkNSignatures(address executor, bytes32 dataHash, bytes signatures, uint256 requiredSignatures) view returns()
func (_Safe *SafeCallerSession) CheckNSignatures1fcac7f3(executor common.Address, dataHash [32]byte, signatures []byte, requiredSignatures *big.Int) error {
	return _Safe.Contract.CheckNSignatures1fcac7f3(&_Safe.CallOpts, executor, dataHash, signatures, requiredSignatures)
}

// CheckSignatures934f3a11 is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_Safe *SafeCaller) CheckSignatures934f3a11(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures934f3a11 is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_Safe *SafeSession) CheckSignatures934f3a11(dataHash [32]byte, data []byte, signatures []byte) error {
	return _Safe.Contract.CheckSignatures934f3a11(&_Safe.CallOpts, dataHash, data, signatures)
}

// CheckSignatures934f3a11 is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_Safe *SafeCallerSession) CheckSignatures934f3a11(dataHash [32]byte, data []byte, signatures []byte) error {
	return _Safe.Contract.CheckSignatures934f3a11(&_Safe.CallOpts, dataHash, data, signatures)
}

// CheckSignaturesf855438b is a free data retrieval call binding the contract method 0xf855438b.
//
// Solidity: function checkSignatures(address executor, bytes32 dataHash, bytes signatures) view returns()
func (_Safe *SafeCaller) CheckSignaturesf855438b(opts *bind.CallOpts, executor common.Address, dataHash [32]byte, signatures []byte) error {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "checkSignatures0", executor, dataHash, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignaturesf855438b is a free data retrieval call binding the contract method 0xf855438b.
//
// Solidity: function checkSignatures(address executor, bytes32 dataHash, bytes signatures) view returns()
func (_Safe *SafeSession) CheckSignaturesf855438b(executor common.Address, dataHash [32]byte, signatures []byte) error {
	return _Safe.Contract.CheckSignaturesf855438b(&_Safe.CallOpts, executor, dataHash, signatures)
}

// CheckSignaturesf855438b is a free data retrieval call binding the contract method 0xf855438b.
//
// Solidity: function checkSignatures(address executor, bytes32 dataHash, bytes signatures) view returns()
func (_Safe *SafeCallerSession) CheckSignaturesf855438b(executor common.Address, dataHash [32]byte, signatures []byte) error {
	return _Safe.Contract.CheckSignaturesf855438b(&_Safe.CallOpts, executor, dataHash, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Safe *SafeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Safe *SafeSession) DomainSeparator() ([32]byte, error) {
	return _Safe.Contract.DomainSeparator(&_Safe.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Safe *SafeCallerSession) DomainSeparator() ([32]byte, error) {
	return _Safe.Contract.DomainSeparator(&_Safe.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_Safe *SafeCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_Safe *SafeSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _Safe.Contract.GetModulesPaginated(&_Safe.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_Safe *SafeCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _Safe.Contract.GetModulesPaginated(&_Safe.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Safe *SafeCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Safe *SafeSession) GetOwners() ([]common.Address, error) {
	return _Safe.Contract.GetOwners(&_Safe.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Safe *SafeCallerSession) GetOwners() ([]common.Address, error) {
	return _Safe.Contract.GetOwners(&_Safe.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Safe *SafeCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Safe *SafeSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _Safe.Contract.GetStorageAt(&_Safe.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Safe *SafeCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _Safe.Contract.GetStorageAt(&_Safe.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Safe *SafeCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Safe *SafeSession) GetThreshold() (*big.Int, error) {
	return _Safe.Contract.GetThreshold(&_Safe.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Safe *SafeCallerSession) GetThreshold() (*big.Int, error) {
	return _Safe.Contract.GetThreshold(&_Safe.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32 txHash)
func (_Safe *SafeCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32 txHash)
func (_Safe *SafeSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _Safe.Contract.GetTransactionHash(&_Safe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32 txHash)
func (_Safe *SafeCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _Safe.Contract.GetTransactionHash(&_Safe.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_Safe *SafeCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_Safe *SafeSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _Safe.Contract.IsModuleEnabled(&_Safe.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_Safe *SafeCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _Safe.Contract.IsModuleEnabled(&_Safe.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_Safe *SafeCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_Safe *SafeSession) IsOwner(owner common.Address) (bool, error) {
	return _Safe.Contract.IsOwner(&_Safe.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_Safe *SafeCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _Safe.Contract.IsOwner(&_Safe.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Safe *SafeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Safe *SafeSession) Nonce() (*big.Int, error) {
	return _Safe.Contract.Nonce(&_Safe.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Safe *SafeCallerSession) Nonce() (*big.Int, error) {
	return _Safe.Contract.Nonce(&_Safe.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_Safe *SafeCaller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Safe.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_Safe *SafeSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _Safe.Contract.SignedMessages(&_Safe.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_Safe *SafeCallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _Safe.Contract.SignedMessages(&_Safe.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_Safe *SafeTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_Safe *SafeSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.AddOwnerWithThreshold(&_Safe.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_Safe *SafeTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.AddOwnerWithThreshold(&_Safe.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_Safe *SafeTransactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_Safe *SafeSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _Safe.Contract.ApproveHash(&_Safe.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_Safe *SafeTransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _Safe.Contract.ApproveHash(&_Safe.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_Safe *SafeTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_Safe *SafeSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.ChangeThreshold(&_Safe.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_Safe *SafeTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.ChangeThreshold(&_Safe.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_Safe *SafeTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_Safe *SafeSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.DisableModule(&_Safe.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_Safe *SafeTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.DisableModule(&_Safe.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_Safe *SafeTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_Safe *SafeSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.EnableModule(&_Safe.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_Safe *SafeTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _Safe.Contract.EnableModule(&_Safe.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_Safe *SafeTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_Safe *SafeSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransaction(&_Safe.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_Safe *SafeTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransaction(&_Safe.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_Safe *SafeTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_Safe *SafeSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransactionFromModule(&_Safe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_Safe *SafeTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransactionFromModule(&_Safe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_Safe *SafeTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_Safe *SafeSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransactionFromModuleReturnData(&_Safe.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_Safe *SafeTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _Safe.Contract.ExecTransactionFromModuleReturnData(&_Safe.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_Safe *SafeTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_Safe *SafeSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.RemoveOwner(&_Safe.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_Safe *SafeTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Safe.Contract.RemoveOwner(&_Safe.TransactOpts, prevOwner, owner, _threshold)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_Safe *SafeTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_Safe *SafeSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SetFallbackHandler(&_Safe.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_Safe *SafeTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SetFallbackHandler(&_Safe.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_Safe *SafeTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_Safe *SafeSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SetGuard(&_Safe.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_Safe *SafeTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SetGuard(&_Safe.TransactOpts, guard)
}

// SetModuleGuard is a paid mutator transaction binding the contract method 0xe068df37.
//
// Solidity: function setModuleGuard(address moduleGuard) returns()
func (_Safe *SafeTransactor) SetModuleGuard(opts *bind.TransactOpts, moduleGuard common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "setModuleGuard", moduleGuard)
}

// SetModuleGuard is a paid mutator transaction binding the contract method 0xe068df37.
//
// Solidity: function setModuleGuard(address moduleGuard) returns()
func (_Safe *SafeSession) SetModuleGuard(moduleGuard common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SetModuleGuard(&_Safe.TransactOpts, moduleGuard)
}

// SetModuleGuard is a paid mutator transaction binding the contract method 0xe068df37.
//
// Solidity: function setModuleGuard(address moduleGuard) returns()
func (_Safe *SafeTransactorSession) SetModuleGuard(moduleGuard common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SetModuleGuard(&_Safe.TransactOpts, moduleGuard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_Safe *SafeTransactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_Safe *SafeSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _Safe.Contract.Setup(&_Safe.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_Safe *SafeTransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _Safe.Contract.Setup(&_Safe.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_Safe *SafeTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_Safe *SafeSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Safe.Contract.SimulateAndRevert(&_Safe.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_Safe *SafeTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Safe.Contract.SimulateAndRevert(&_Safe.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_Safe *SafeTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Safe.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_Safe *SafeSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SwapOwner(&_Safe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_Safe *SafeTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Safe.Contract.SwapOwner(&_Safe.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Safe *SafeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Safe.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Safe *SafeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Safe.Contract.Fallback(&_Safe.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Safe *SafeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Safe.Contract.Fallback(&_Safe.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Safe *SafeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Safe *SafeSession) Receive() (*types.Transaction, error) {
	return _Safe.Contract.Receive(&_Safe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Safe *SafeTransactorSession) Receive() (*types.Transaction, error) {
	return _Safe.Contract.Receive(&_Safe.TransactOpts)
}

// SafeAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the Safe contract.
type SafeAddedOwnerIterator struct {
	Event *SafeAddedOwner // Event containing the contract specifics and raw log

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
func (it *SafeAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeAddedOwner)
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
		it.Event = new(SafeAddedOwner)
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
func (it *SafeAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeAddedOwner represents a AddedOwner event raised by the Safe contract.
type SafeAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_Safe *SafeFilterer) FilterAddedOwner(opts *bind.FilterOpts, owner []common.Address) (*SafeAddedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "AddedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeAddedOwnerIterator{contract: _Safe.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_Safe *SafeFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *SafeAddedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "AddedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeAddedOwner)
				if err := _Safe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_Safe *SafeFilterer) ParseAddedOwner(log types.Log) (*SafeAddedOwner, error) {
	event := new(SafeAddedOwner)
	if err := _Safe.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the Safe contract.
type SafeApproveHashIterator struct {
	Event *SafeApproveHash // Event containing the contract specifics and raw log

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
func (it *SafeApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeApproveHash)
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
		it.Event = new(SafeApproveHash)
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
func (it *SafeApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeApproveHash represents a ApproveHash event raised by the Safe contract.
type SafeApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_Safe *SafeFilterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*SafeApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeApproveHashIterator{contract: _Safe.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_Safe *SafeFilterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *SafeApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeApproveHash)
				if err := _Safe.contract.UnpackLog(event, "ApproveHash", log); err != nil {
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

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_Safe *SafeFilterer) ParseApproveHash(log types.Log) (*SafeApproveHash, error) {
	event := new(SafeApproveHash)
	if err := _Safe.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the Safe contract.
type SafeChangedFallbackHandlerIterator struct {
	Event *SafeChangedFallbackHandler // Event containing the contract specifics and raw log

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
func (it *SafeChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeChangedFallbackHandler)
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
		it.Event = new(SafeChangedFallbackHandler)
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
func (it *SafeChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeChangedFallbackHandler represents a ChangedFallbackHandler event raised by the Safe contract.
type SafeChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_Safe *SafeFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts, handler []common.Address) (*SafeChangedFallbackHandlerIterator, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ChangedFallbackHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return &SafeChangedFallbackHandlerIterator{contract: _Safe.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_Safe *SafeFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *SafeChangedFallbackHandler, handler []common.Address) (event.Subscription, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ChangedFallbackHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeChangedFallbackHandler)
				if err := _Safe.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
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

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_Safe *SafeFilterer) ParseChangedFallbackHandler(log types.Log) (*SafeChangedFallbackHandler, error) {
	event := new(SafeChangedFallbackHandler)
	if err := _Safe.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the Safe contract.
type SafeChangedGuardIterator struct {
	Event *SafeChangedGuard // Event containing the contract specifics and raw log

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
func (it *SafeChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeChangedGuard)
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
		it.Event = new(SafeChangedGuard)
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
func (it *SafeChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeChangedGuard represents a ChangedGuard event raised by the Safe contract.
type SafeChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_Safe *SafeFilterer) FilterChangedGuard(opts *bind.FilterOpts, guard []common.Address) (*SafeChangedGuardIterator, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ChangedGuard", guardRule)
	if err != nil {
		return nil, err
	}
	return &SafeChangedGuardIterator{contract: _Safe.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_Safe *SafeFilterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *SafeChangedGuard, guard []common.Address) (event.Subscription, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ChangedGuard", guardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeChangedGuard)
				if err := _Safe.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
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

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_Safe *SafeFilterer) ParseChangedGuard(log types.Log) (*SafeChangedGuard, error) {
	event := new(SafeChangedGuard)
	if err := _Safe.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeChangedModuleGuardIterator is returned from FilterChangedModuleGuard and is used to iterate over the raw logs and unpacked data for ChangedModuleGuard events raised by the Safe contract.
type SafeChangedModuleGuardIterator struct {
	Event *SafeChangedModuleGuard // Event containing the contract specifics and raw log

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
func (it *SafeChangedModuleGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeChangedModuleGuard)
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
		it.Event = new(SafeChangedModuleGuard)
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
func (it *SafeChangedModuleGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeChangedModuleGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeChangedModuleGuard represents a ChangedModuleGuard event raised by the Safe contract.
type SafeChangedModuleGuard struct {
	ModuleGuard common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedModuleGuard is a free log retrieval operation binding the contract event 0xcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb8.
//
// Solidity: event ChangedModuleGuard(address indexed moduleGuard)
func (_Safe *SafeFilterer) FilterChangedModuleGuard(opts *bind.FilterOpts, moduleGuard []common.Address) (*SafeChangedModuleGuardIterator, error) {

	var moduleGuardRule []interface{}
	for _, moduleGuardItem := range moduleGuard {
		moduleGuardRule = append(moduleGuardRule, moduleGuardItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ChangedModuleGuard", moduleGuardRule)
	if err != nil {
		return nil, err
	}
	return &SafeChangedModuleGuardIterator{contract: _Safe.contract, event: "ChangedModuleGuard", logs: logs, sub: sub}, nil
}

// WatchChangedModuleGuard is a free log subscription operation binding the contract event 0xcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb8.
//
// Solidity: event ChangedModuleGuard(address indexed moduleGuard)
func (_Safe *SafeFilterer) WatchChangedModuleGuard(opts *bind.WatchOpts, sink chan<- *SafeChangedModuleGuard, moduleGuard []common.Address) (event.Subscription, error) {

	var moduleGuardRule []interface{}
	for _, moduleGuardItem := range moduleGuard {
		moduleGuardRule = append(moduleGuardRule, moduleGuardItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ChangedModuleGuard", moduleGuardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeChangedModuleGuard)
				if err := _Safe.contract.UnpackLog(event, "ChangedModuleGuard", log); err != nil {
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

// ParseChangedModuleGuard is a log parse operation binding the contract event 0xcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb8.
//
// Solidity: event ChangedModuleGuard(address indexed moduleGuard)
func (_Safe *SafeFilterer) ParseChangedModuleGuard(log types.Log) (*SafeChangedModuleGuard, error) {
	event := new(SafeChangedModuleGuard)
	if err := _Safe.contract.UnpackLog(event, "ChangedModuleGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the Safe contract.
type SafeChangedThresholdIterator struct {
	Event *SafeChangedThreshold // Event containing the contract specifics and raw log

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
func (it *SafeChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeChangedThreshold)
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
		it.Event = new(SafeChangedThreshold)
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
func (it *SafeChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeChangedThreshold represents a ChangedThreshold event raised by the Safe contract.
type SafeChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_Safe *SafeFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*SafeChangedThresholdIterator, error) {

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &SafeChangedThresholdIterator{contract: _Safe.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_Safe *SafeFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *SafeChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeChangedThreshold)
				if err := _Safe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_Safe *SafeFilterer) ParseChangedThreshold(log types.Log) (*SafeChangedThreshold, error) {
	event := new(SafeChangedThreshold)
	if err := _Safe.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the Safe contract.
type SafeDisabledModuleIterator struct {
	Event *SafeDisabledModule // Event containing the contract specifics and raw log

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
func (it *SafeDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeDisabledModule)
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
		it.Event = new(SafeDisabledModule)
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
func (it *SafeDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeDisabledModule represents a DisabledModule event raised by the Safe contract.
type SafeDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_Safe *SafeFilterer) FilterDisabledModule(opts *bind.FilterOpts, module []common.Address) (*SafeDisabledModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "DisabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeDisabledModuleIterator{contract: _Safe.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_Safe *SafeFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *SafeDisabledModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "DisabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeDisabledModule)
				if err := _Safe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_Safe *SafeFilterer) ParseDisabledModule(log types.Log) (*SafeDisabledModule, error) {
	event := new(SafeDisabledModule)
	if err := _Safe.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the Safe contract.
type SafeEnabledModuleIterator struct {
	Event *SafeEnabledModule // Event containing the contract specifics and raw log

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
func (it *SafeEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeEnabledModule)
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
		it.Event = new(SafeEnabledModule)
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
func (it *SafeEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeEnabledModule represents a EnabledModule event raised by the Safe contract.
type SafeEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_Safe *SafeFilterer) FilterEnabledModule(opts *bind.FilterOpts, module []common.Address) (*SafeEnabledModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "EnabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeEnabledModuleIterator{contract: _Safe.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_Safe *SafeFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *SafeEnabledModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "EnabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeEnabledModule)
				if err := _Safe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_Safe *SafeFilterer) ParseEnabledModule(log types.Log) (*SafeEnabledModule, error) {
	event := new(SafeEnabledModule)
	if err := _Safe.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the Safe contract.
type SafeExecutionFailureIterator struct {
	Event *SafeExecutionFailure // Event containing the contract specifics and raw log

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
func (it *SafeExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeExecutionFailure)
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
		it.Event = new(SafeExecutionFailure)
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
func (it *SafeExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeExecutionFailure represents a ExecutionFailure event raised by the Safe contract.
type SafeExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_Safe *SafeFilterer) FilterExecutionFailure(opts *bind.FilterOpts, txHash [][32]byte) (*SafeExecutionFailureIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ExecutionFailure", txHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeExecutionFailureIterator{contract: _Safe.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_Safe *SafeFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *SafeExecutionFailure, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ExecutionFailure", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeExecutionFailure)
				if err := _Safe.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_Safe *SafeFilterer) ParseExecutionFailure(log types.Log) (*SafeExecutionFailure, error) {
	event := new(SafeExecutionFailure)
	if err := _Safe.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the Safe contract.
type SafeExecutionFromModuleFailureIterator struct {
	Event *SafeExecutionFromModuleFailure // Event containing the contract specifics and raw log

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
func (it *SafeExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeExecutionFromModuleFailure)
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
		it.Event = new(SafeExecutionFromModuleFailure)
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
func (it *SafeExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the Safe contract.
type SafeExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_Safe *SafeFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*SafeExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeExecutionFromModuleFailureIterator{contract: _Safe.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_Safe *SafeFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *SafeExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeExecutionFromModuleFailure)
				if err := _Safe.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
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

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_Safe *SafeFilterer) ParseExecutionFromModuleFailure(log types.Log) (*SafeExecutionFromModuleFailure, error) {
	event := new(SafeExecutionFromModuleFailure)
	if err := _Safe.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the Safe contract.
type SafeExecutionFromModuleSuccessIterator struct {
	Event *SafeExecutionFromModuleSuccess // Event containing the contract specifics and raw log

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
func (it *SafeExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeExecutionFromModuleSuccess)
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
		it.Event = new(SafeExecutionFromModuleSuccess)
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
func (it *SafeExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the Safe contract.
type SafeExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_Safe *SafeFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*SafeExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeExecutionFromModuleSuccessIterator{contract: _Safe.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_Safe *SafeFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *SafeExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeExecutionFromModuleSuccess)
				if err := _Safe.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
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

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_Safe *SafeFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*SafeExecutionFromModuleSuccess, error) {
	event := new(SafeExecutionFromModuleSuccess)
	if err := _Safe.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the Safe contract.
type SafeExecutionSuccessIterator struct {
	Event *SafeExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *SafeExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeExecutionSuccess)
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
		it.Event = new(SafeExecutionSuccess)
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
func (it *SafeExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeExecutionSuccess represents a ExecutionSuccess event raised by the Safe contract.
type SafeExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_Safe *SafeFilterer) FilterExecutionSuccess(opts *bind.FilterOpts, txHash [][32]byte) (*SafeExecutionSuccessIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "ExecutionSuccess", txHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeExecutionSuccessIterator{contract: _Safe.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_Safe *SafeFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *SafeExecutionSuccess, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "ExecutionSuccess", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeExecutionSuccess)
				if err := _Safe.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_Safe *SafeFilterer) ParseExecutionSuccess(log types.Log) (*SafeExecutionSuccess, error) {
	event := new(SafeExecutionSuccess)
	if err := _Safe.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the Safe contract.
type SafeRemovedOwnerIterator struct {
	Event *SafeRemovedOwner // Event containing the contract specifics and raw log

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
func (it *SafeRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeRemovedOwner)
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
		it.Event = new(SafeRemovedOwner)
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
func (it *SafeRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeRemovedOwner represents a RemovedOwner event raised by the Safe contract.
type SafeRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_Safe *SafeFilterer) FilterRemovedOwner(opts *bind.FilterOpts, owner []common.Address) (*SafeRemovedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "RemovedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeRemovedOwnerIterator{contract: _Safe.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_Safe *SafeFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *SafeRemovedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "RemovedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeRemovedOwner)
				if err := _Safe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_Safe *SafeFilterer) ParseRemovedOwner(log types.Log) (*SafeRemovedOwner, error) {
	event := new(SafeRemovedOwner)
	if err := _Safe.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the Safe contract.
type SafeSafeReceivedIterator struct {
	Event *SafeSafeReceived // Event containing the contract specifics and raw log

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
func (it *SafeSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeSafeReceived)
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
		it.Event = new(SafeSafeReceived)
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
func (it *SafeSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeSafeReceived represents a SafeReceived event raised by the Safe contract.
type SafeSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_Safe *SafeFilterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*SafeSafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &SafeSafeReceivedIterator{contract: _Safe.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_Safe *SafeFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *SafeSafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeSafeReceived)
				if err := _Safe.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_Safe *SafeFilterer) ParseSafeReceived(log types.Log) (*SafeSafeReceived, error) {
	event := new(SafeSafeReceived)
	if err := _Safe.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeSafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the Safe contract.
type SafeSafeSetupIterator struct {
	Event *SafeSafeSetup // Event containing the contract specifics and raw log

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
func (it *SafeSafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeSafeSetup)
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
		it.Event = new(SafeSafeSetup)
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
func (it *SafeSafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeSafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeSafeSetup represents a SafeSetup event raised by the Safe contract.
type SafeSafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_Safe *SafeFilterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*SafeSafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &SafeSafeSetupIterator{contract: _Safe.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_Safe *SafeFilterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *SafeSafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeSafeSetup)
				if err := _Safe.contract.UnpackLog(event, "SafeSetup", log); err != nil {
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

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_Safe *SafeFilterer) ParseSafeSetup(log types.Log) (*SafeSafeSetup, error) {
	event := new(SafeSafeSetup)
	if err := _Safe.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeSignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the Safe contract.
type SafeSignMsgIterator struct {
	Event *SafeSignMsg // Event containing the contract specifics and raw log

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
func (it *SafeSignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeSignMsg)
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
		it.Event = new(SafeSignMsg)
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
func (it *SafeSignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeSignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeSignMsg represents a SignMsg event raised by the Safe contract.
type SafeSignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_Safe *SafeFilterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*SafeSignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Safe.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeSignMsgIterator{contract: _Safe.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_Safe *SafeFilterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *SafeSignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _Safe.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeSignMsg)
				if err := _Safe.contract.UnpackLog(event, "SignMsg", log); err != nil {
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

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_Safe *SafeFilterer) ParseSignMsg(log types.Log) (*SafeSignMsg, error) {
	event := new(SafeSignMsg)
	if err := _Safe.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2MetaData contains all meta data concerning the SafeL2 contract.
var SafeL2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"moduleGuard\",\"type\":\"address\"}],\"name\":\"ChangedModuleGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"SafeModuleTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"additionalInfo\",\"type\":\"bytes\"}],\"name\":\"SafeMultiSigTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moduleGuard\",\"type\":\"address\"}],\"name\":\"setModuleGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060016004819055506156d780620000296000396000f3fe6080604052600436106101dc5760003560e01c8063affed0e011610102578063e19a9dd911610095578063f698da2511610064578063f698da25146114c1578063f855438b146114ec578063f8dc5dd9146115de578063ffa1ad741461165957610231565b8063e19a9dd914611363578063e318b52b146113b4578063e75235b814611445578063f08a03231461147057610231565b8063d4d9bdcd116100d1578063d4d9bdcd1461112c578063d8d11f7814611167578063e009cfde146112a1578063e068df371461131257610231565b8063affed0e014610ddc578063b4faba0914610e07578063b63e800d14610eef578063cc2f84521461105f57610231565b80635624b25b1161017a5780636a761202116101495780636a76120214610a1e5780637d83297414610bda578063934f3a1114610c49578063a0e67e2b14610d7057610231565b80635624b25b146108855780635ae6bd3714610943578063610b592514610992578063694e80c3146109e357610231565b80632d9ad53d116101b65780632d9ad53d146105215780632f54bf6e14610588578063468721a7146105ef5780635229073f1461070457610231565b80630d582f131461029957806312fb68e0146102f45780631fcac7f31461042557610231565b36610231573373ffffffffffffffffffffffffffffffffffffffff167f3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d346040518082815260200191505060405180910390a2005b34801561023d57600080fd5b507f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d5548061026a57600080f35b60405136600082373360601b3682015260008060143601836000865af13d6000833e80610295573d82fd5b3d82f35b3480156102a557600080fd5b506102f2600480360360408110156102bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116e9565b005b34801561030057600080fd5b506104236004803603608081101561031757600080fd5b81019080803590602001909291908035906020019064010000000081111561033e57600080fd5b82018360208201111561035057600080fd5b8035906020019184600183028401116401000000008311171561037257600080fd5b90919293919293908035906020019064010000000081111561039357600080fd5b8201836020820111156103a557600080fd5b803590602001918460018302840111640100000000831117156103c757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611a41565b005b34801561043157600080fd5b5061051f6004803603608081101561044857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561048f57600080fd5b8201836020820111156104a157600080fd5b803590602001918460018302840111640100000000831117156104c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611a54565b005b34801561052d57600080fd5b506105706004803603602081101561054457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e76565b60405180821515815260200191505060405180910390f35b34801561059457600080fd5b506105d7600480360360208110156105ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f48565b60405180821515815260200191505060405180910390f35b3480156105fb57600080fd5b506106ec6004803603608081101561061257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561065957600080fd5b82018360208201111561066b57600080fd5b8035906020019184600183028401116401000000008311171561068d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff169060200190929190505050612018565b60405180821515815260200191505060405180910390f35b34801561071057600080fd5b506108016004803603608081101561072757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561076e57600080fd5b82018360208201111561078057600080fd5b803590602001918460018302840111640100000000831117156107a257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff169060200190929190505050612071565b60405180831515815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561084957808201518184015260208101905061082e565b50505050905090810190601f1680156108765780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34801561089157600080fd5b506108c8600480360360408110156108a857600080fd5b8101908080359060200190929190803590602001909291905050506120e6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109085780820151818401526020810190506108ed565b50505050905090810190601f1680156109355780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561094f57600080fd5b5061097c6004803603602081101561096657600080fd5b810190808035906020019092919050505061216c565b6040518082815260200191505060405180910390f35b34801561099e57600080fd5b506109e1600480360360208110156109b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612184565b005b3480156109ef57600080fd5b50610a1c60048036036020811015610a0657600080fd5b8101908080359060200190929190505050612482565b005b610bc26004803603610140811015610a3557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190640100000000811115610a7c57600080fd5b820183602082011115610a8e57600080fd5b80359060200191846001830284011164010000000083111715610ab057600080fd5b9091929391929390803560ff169060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610b3c57600080fd5b820183602082011115610b4e57600080fd5b80359060200191846001830284011164010000000083111715610b7057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050612532565b60405180821515815260200191505060405180910390f35b348015610be657600080fd5b50610c3360048036036040811015610bfd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506129c1565b6040518082815260200191505060405180910390f35b348015610c5557600080fd5b50610d6e60048036036060811015610c6c57600080fd5b810190808035906020019092919080359060200190640100000000811115610c9357600080fd5b820183602082011115610ca557600080fd5b80359060200191846001830284011164010000000083111715610cc757600080fd5b909192939192939080359060200190640100000000811115610ce857600080fd5b820183602082011115610cfa57600080fd5b80359060200191846001830284011164010000000083111715610d1c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506129e6565b005b348015610d7c57600080fd5b50610d856129f7565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610dc8578082015181840152602081019050610dad565b505050509050019250505060405180910390f35b348015610de857600080fd5b50610df1612b9e565b6040518082815260200191505060405180910390f35b348015610e1357600080fd5b50610eed60048036036040811015610e2a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610e6757600080fd5b820183602082011115610e7957600080fd5b80359060200191846001830284011164010000000083111715610e9b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050612ba4565b005b348015610efb57600080fd5b5061105d6004803603610100811015610f1357600080fd5b8101908080359060200190640100000000811115610f3057600080fd5b820183602082011115610f4257600080fd5b80359060200191846020830284011164010000000083111715610f6457600080fd5b909192939192939080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610faf57600080fd5b820183602082011115610fc157600080fd5b80359060200191846001830284011164010000000083111715610fe357600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612bcb565b005b34801561106b57600080fd5b506110b86004803603604081101561108257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612d89565b60405180806020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019060200280838360005b838110156111175780820151818401526020810190506110fc565b50505050905001935050505060405180910390f35b34801561113857600080fd5b506111656004803603602081101561114f57600080fd5b8101908080359060200190929190505050613067565b005b34801561117357600080fd5b5061128b600480360361014081101561118b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156111d257600080fd5b8201836020820111156111e457600080fd5b8035906020019184600183028401116401000000008311171561120657600080fd5b9091929391929390803560ff169060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506131c2565b6040518082815260200191505060405180910390f35b3480156112ad57600080fd5b50611310600480360360408110156112c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061326b565b005b34801561131e57600080fd5b506113616004803603602081101561133557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613568565b005b34801561136f57600080fd5b506113b26004803603602081101561138657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613716565b005b3480156113c057600080fd5b50611443600480360360608110156113d757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506138bb565b005b34801561145157600080fd5b5061145a613e03565b6040518082815260200191505060405180910390f35b34801561147c57600080fd5b506114bf6004803603602081101561149357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050613e0d565b005b3480156114cd57600080fd5b506114d6613e64565b6040518082815260200191505060405180910390f35b3480156114f857600080fd5b506115dc6004803603606081101561150f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561155657600080fd5b82018360208201111561156857600080fd5b8035906020019184600183028401116401000000008311171561158a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050613ee0565b005b3480156115ea57600080fd5b506116576004803603606081101561160157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050613f2c565b005b34801561166557600080fd5b5061166e61427e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156116ae578082015181840152602081019050611693565b50505050905090810190601f1680156116db5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6116f16142b7565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806117585750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b8061178e57503073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b156117bd576117bc7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461187a576118797f4753323034000000000000000000000000000000000000000000000000000000614316565b5b60026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360008154600101919050819055508173ffffffffffffffffffffffffffffffffffffffff167f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2660405160405180910390a28060045414611a3d57611a3c81612482565b5b5050565b611a4d33868484611a54565b5050505050565b611a6860418261435590919063ffffffff16565b82511015611a9a57611a997f4753303230000000000000000000000000000000000000000000000000000000614316565b5b6000808060008060005b86811015611e6a57611ab6888261438f565b8260ff1692508094508195508296505050506000841415611b33578260001c9450611aeb60418861435590919063ffffffff16565b8260001c1015611b1f57611b1e7f4753303231000000000000000000000000000000000000000000000000000000614316565b5b611b2e858a8a8560001c6143be565b611d30565b6001841415611c02578260001c94508473ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1614158015611bce57506000600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b815260200190815260200160002054145b15611bfd57611bfc7f4753303235000000000000000000000000000000000000000000000000000000614316565b5b611d2f565b601e841115611cc75760018960405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c018281526020019150506040516020818303038152906040528051906020012060048603858560405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611cb6573d6000803e3d6000fd5b505050602060405103519450611d2e565b60018985858560405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611d21573d6000803e3d6000fd5b5050506020604051035194505b5b5b8573ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16111580611df65750600073ffffffffffffffffffffffffffffffffffffffff16600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80611e2d5750600173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b15611e5c57611e5b7f4753303236000000000000000000000000000000000000000000000000000000614316565b5b849550806001019050611aa4565b50505050505050505050565b60008173ffffffffffffffffffffffffffffffffffffffff16600173ffffffffffffffffffffffffffffffffffffffff1614158015611f415750600073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b9050919050565b6000600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806120105750600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b159050919050565b6000806000612029878787876145e7565b9150915061205a878787877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61487f565b92506120678282856148cb565b5050949350505050565b60006060600080612084888888886145e7565b915091506120b5888888887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61487f565b9350604051925060203d0183016040523d83523d6000602085013e6120db8282866148cb565b505094509492505050565b60606000600583901b67ffffffffffffffff8111801561210557600080fd5b506040519080825280601f01601f1916602001820160405280156121385781602001600182028036833780820191505090505b50905060005b83811015612161578085015480602083026020850101525080600101905061213e565b508091505092915050565b60076020528060005260406000206000915090505481565b61218c6142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806121f35750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15612222576122217f4753313031000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146122df576122de7f4753313032000000000000000000000000000000000000000000000000000000614316565b5b60016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167fecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f844060405160405180910390a250565b61248a6142b7565b6003548111156124be576124bd7f4753323031000000000000000000000000000000000000000000000000000000614316565b5b60008114156124f1576124f07f4753323032000000000000000000000000000000000000000000000000000000614316565b5b806004819055507f610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93816040518082815260200191505060405180910390a150565b60006125478c8c8c8c8c8c8c8c8c8c8c614a0c565b600061256c8d8d8d8d8d8d8d8d8d8d60056000815480929190600101919050556131c2565b9050612579338285613ee0565b6000612583614c32565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612769578073ffffffffffffffffffffffffffffffffffffffff166375f0bb528f8f8f8f8f8f8f8f8f8f8f336040518d63ffffffff1660e01b8152600401808d73ffffffffffffffffffffffffffffffffffffffff1681526020018c8152602001806020018a600181111561262657fe5b81526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001806020018473ffffffffffffffffffffffffffffffffffffffff16815260200183810383528d8d82818152602001925080828437600081840152601f19601f820116905080830192505050838103825285818151815260200191508051906020019080838360005b838110156126f85780820151818401526020810190506126dd565b50505050905090810190601f1680156127255780820380516001836020036101000a031916815260200191505b509e505050505050505050505050505050600060405180830381600087803b15801561275057600080fd5b505af1158015612764573d6000803e3d6000fd5b505050505b6101f46127916109c48b01603f60068d901b8161278257fe5b04614c5b90919063ffffffff16565b015a10156127c3576127c27f4753303130000000000000000000000000000000000000000000000000000000614316565b5b60005a905061282c8f8f8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e60008d14612821578e612827565b6109c45a035b61487f565b93506128415a82614c7590919063ffffffff16565b905083158015612851575060008a145b801561285d5750600088145b1561286e576040513d6000823e3d81fd5b60008089111561288857612885828b8b8b8b614c95565b90505b84156128cb57837f442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e826040518082815260200191505060405180910390a2612904565b837f23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23826040518082815260200191505060405180910390a25b5050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146129b0578073ffffffffffffffffffffffffffffffffffffffff16639327136883856040518363ffffffff1660e01b815260040180838152602001821515815260200192505050600060405180830381600087803b15801561299757600080fd5b505af11580156129ab573d6000803e3d6000fd5b505050505b50509b9a5050505050505050505050565b6008602052816000526040600020602052806000526040600020600091509150505481565b6129f1338583613ee0565b50505050565b6060600060035467ffffffffffffffff81118015612a1457600080fd5b50604051908082528060200260200182016040528015612a435781602001602082028036833780820191505090505b50905060008060026000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612b955780838381518110612aee57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816001019150612aad565b82935050505090565b60055481565b600080825160208401855af46040518181523d60208201523d6000604083013e60403d0181fd5b3373ffffffffffffffffffffffffffffffffffffffff167f141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a88b8b8b8b8960405180806020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281038252878782818152602001925060200280828437600081840152601f19601f820116905080830192505050965050505050505060405180910390a2612cd78a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505089614e49565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614612d1557612d14846151f3565b5b612d638787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050615277565b6000821115612d7d57612d7b82600060018685614c95565b505b50505050505050505050565b60606000600173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614158015612dd05750612dce84611e76565b155b15612dff57612dfe7f4753313035000000000000000000000000000000000000000000000000000000614316565b5b6000831415612e3257612e317f4753313036000000000000000000000000000000000000000000000000000000614316565b5b8267ffffffffffffffff81118015612e4957600080fd5b50604051908082528060200260200182016040528015612e785781602001602082028036833780820191505090505b5091506000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015612f4a5750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b8015612f5557508381105b1561300e5781838281518110612f6757fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150806001019050612ee0565b600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461305c5782600182038151811061305157fe5b602002602001015191505b808352509250929050565b600073ffffffffffffffffffffffffffffffffffffffff16600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613125576131247f4753303330000000000000000000000000000000000000000000000000000000614316565b5b6001600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16817ff2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c60405160405180910390a350565b6000806131cd613e64565b90506040518a8c82378a81207fbb8310d486368db6bd6f849402fdd73ad53d316b5a4b2644ad6efe0f941286d882528e60208301528d60408301528060608301528a60808301528960a08301528860c08301528760e08301528661010083015285610120830152846101408301526101608220604083015261190182528260208301526042601e83012093505050509b9a5050505050505050505050565b6132736142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806132da5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15613309576133087f4753313031000000000000000000000000000000000000000000000000000000614316565b5b8073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146133c5576133c47f4753313033000000000000000000000000000000000000000000000000000000614316565b5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167faab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace405427660405160405180910390a25050565b6135706142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561367557508073ffffffffffffffffffffffffffffffffffffffff166301ffc9a77f58401ed8000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b815260040180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060206040518083038186803b15801561363857600080fd5b505afa15801561364c573d6000803e3d6000fd5b505050506040513d602081101561366257600080fd5b8101908080519060200190929190505050155b156136a4576136a37f4753333031000000000000000000000000000000000000000000000000000000614316565b5b60007fb104e0b93118902c651344349b610029d694cfdec91c589c91ebafbcd028994760001b90508181558173ffffffffffffffffffffffffffffffffffffffff167fcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb860405160405180910390a25050565b61371e6142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561382357508073ffffffffffffffffffffffffffffffffffffffff166301ffc9a77fe6d7a83a000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b815260040180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060206040518083038186803b1580156137e657600080fd5b505afa1580156137fa573d6000803e3d6000fd5b505050506040513d602081101561381057600080fd5b8101908080519060200190929190505050155b15613852576138517f4753333030000000000000000000000000000000000000000000000000000000614316565b5b807f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c8558073ffffffffffffffffffffffffffffffffffffffff167f1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa260405160405180910390a250565b6138c36142b7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16148061392a5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b8061396057503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1561398f5761398e7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613a4c57613a4b7f4753323034000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480613ab35750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b15613ae257613ae17f4753323033000000000000000000000000000000000000000000000000000000614316565b5b8173ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613b9e57613b9d7f4753323035000000000000000000000000000000000000000000000000000000614316565b5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf60405160405180910390a28073ffffffffffffffffffffffffffffffffffffffff167f9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea2660405160405180910390a2505050565b6000600454905090565b613e156142b7565b613e1e816151f3565b8073ffffffffffffffffffffffffffffffffffffffff167f5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b060405160405180910390a250565b6000804690507f47e79534a245952e8b16893a336b85a3d9ea9fa8c573f3d803afb92a7946921860001b8130604051602001808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff16815260200193505050506040516020818303038152906040528051906020012091505090565b600060045490506000811415613f1a57613f197f4753303031000000000000000000000000000000000000000000000000000000614316565b5b613f2684848484611a54565b50505050565b613f346142b7565b80600360008154600190039190508190551015613f7557613f747f4753323031000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480613fdc5750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b1561400b5761400a7f4753323033000000000000000000000000000000000000000000000000000000614316565b5b8173ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146140c7576140c67f4753323035000000000000000000000000000000000000000000000000000000614316565b5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff167ff8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf60405160405180910390a280600454146142795761427881612482565b5b505050565b6040518060400160405280600581526020017f312e352e3000000000000000000000000000000000000000000000000000000081525081565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614314576143137f4753303331000000000000000000000000000000000000000000000000000000614316565b5b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260056024820152816044820152606481fd5b6000808314156143685760009050614389565b600082840290508284828161437957fe5b041461438457600080fd5b809150505b92915050565b6000806000836041026020810186015192506040810186015191506060810186015160001a9350509250925092565b81516143d460208361548190919063ffffffff16565b1115614404576144037f4753303232000000000000000000000000000000000000000000000000000000614316565b5b600060208284010151905082516144378261442960208661548190919063ffffffff16565b61548190919063ffffffff16565b1115614467576144667f4753303233000000000000000000000000000000000000000000000000000000614316565b5b60606020838501019050631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168673ffffffffffffffffffffffffffffffffffffffff16631626ba7e87846040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561450b5780820151818401526020810190506144f0565b50505050905090810190601f1680156145385780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b15801561455657600080fd5b505afa15801561456a573d6000803e3d6000fd5b505050506040513d602081101561458057600080fd5b81019080805190602001909291905050507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146145df576145de7f4753303234000000000000000000000000000000000000000000000000000000614316565b5b505050505050565b6000806145f6868686866154a0565b6145fe615599565b9150600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806146c65750600073ffffffffffffffffffffffffffffffffffffffff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b156146f5576146f47f4753313034000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614614876578173ffffffffffffffffffffffffffffffffffffffff1663728c297287878787336040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018060200184600181111561478f57fe5b81526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b838110156147e85780820151818401526020810190506147cd565b50505050905090810190601f1680156148155780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561483857600080fd5b505af115801561484c573d6000803e3d6000fd5b505050506040513d602081101561486257600080fd5b810190808051906020019092919050505090505b94509492505050565b600060018081111561488d57fe5b83600181111561489957fe5b14156148b2576000808551602087018986f490506148c2565b600080855160208701888a87f190505b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614614975578273ffffffffffffffffffffffffffffffffffffffff16632acc37aa83836040518363ffffffff1660e01b815260040180838152602001821515815260200192505050600060405180830381600087803b15801561495c57600080fd5b505af1158015614970573d6000803e3d6000fd5b505050505b80156149c3573373ffffffffffffffffffffffffffffffffffffffff167f6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb860405160405180910390a2614a07565b3373ffffffffffffffffffffffffffffffffffffffff167facd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd37560405160405180910390a25b505050565b606060055433600454604051602001808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405160208183030381529060405290507f66753cd2356569ee081232e3be8909b950e0a76c1f8460c3a5e3c2be32b11bed8c8c8c8c8c8c8c8c8c8c8c8c604051808d73ffffffffffffffffffffffffffffffffffffffff1681526020018c8152602001806020018a6001811115614abd57fe5b81526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001806020018060200184810384528e8e82818152602001925080828437600081840152601f19601f820116905080830192505050848103835286818151815260200191508051906020019080838360005b83811015614b77578082015181840152602081019050614b5c565b50505050905090810190601f168015614ba45780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015614bdd578082015181840152602081019050614bc2565b50505050905090810190601f168015614c0a5780820380516001836020036101000a031916815260200191505b509f5050505050505050505050505050505060405180910390a1505050505050505050505050565b60007f4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c854905090565b600081831015614c6b5781614c6d565b825b905092915050565b600082821115614c8457600080fd5b600082840390508091505092915050565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614614cd25782614cd4565b325b9050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415614dde57614d3e3a8610614d1b573a614d1d565b855b614d30888a61548190919063ffffffff16565b61435590919063ffffffff16565b915060008173ffffffffffffffffffffffffffffffffffffffff168360405180600001905060006040518083038185875af1925050503d8060008114614da0576040519150601f19603f3d011682016040523d82523d6000602084013e614da5565b606091505b5050905080614dd857614dd77f4753303131000000000000000000000000000000000000000000000000000000614316565b5b50614e3f565b614e0385614df5888a61548190919063ffffffff16565b61435590919063ffffffff16565b9150614e108482846155ca565b614e3e57614e3d7f4753303132000000000000000000000000000000000000000000000000000000614316565b5b5b5095945050505050565b60006004541115614e7e57614e7d7f4753323030000000000000000000000000000000000000000000000000000000614316565b5b8151811115614eb157614eb07f4753323031000000000000000000000000000000000000000000000000000000614316565b5b6000811415614ee457614ee37f4753323032000000000000000000000000000000000000000000000000000000614316565b5b60006001905060008351905060005b8181101561515f576000858281518110614f0957fe5b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480614f7a5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80614fb057503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80614fe657508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b15615015576150147f4753323033000000000000000000000000000000000000000000000000000000614316565b5b600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146150d2576150d17f4753323034000000000000000000000000000000000000000000000000000000614316565b5b80600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080935050806001019050614ef3565b506001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806003819055508260048190555050505050565b3073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615251576152507f4753343030000000000000000000000000000000000000000000000000000000614316565b5b807f6c9a6c4a39284e37ed1cf53d337577d14212a4870fb976a4366c693b939918d55550565b600073ffffffffffffffffffffffffffffffffffffffff1660016000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614615335576153347f4753313030000000000000000000000000000000000000000000000000000000614316565b5b6001806000600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461547d576153f18261568e565b61541f5761541e7f4753303032000000000000000000000000000000000000000000000000000000614316565b5b61544e8260008360017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61487f565b61547c5761547b7f4753303030000000000000000000000000000000000000000000000000000000614316565b5b5b5050565b60008082840190508381101561549657600080fd5b8091505092915050565b7fb648d3644f584ed1c2232d53c46d87e693586486ad0d1175f8656013110b714e3385858585604051808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200183600181111561551857fe5b8152602001828103825284818151815260200191508051906020019080838360005b8381101561555557808201518184015260208101905061553a565b50505050905090810190601f1680156155825780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a150505050565b6000807fb104e0b93118902c651344349b610029d694cfdec91c589c91ebafbcd028994760001b9050805491505090565b60008063a9059cbb8484604051602401808373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050602060008251602084016000896127105a03f13d6000811461567157602081146156795760009350615684565b819350615684565b600051158215171593505b5050509392505050565b600080823b90506000811191505091905056fea2646970667358221220ef709aaa9ae1fcd6146759064f557c4d38d3d132d7d4b983d0abdcaa62c3a81564736f6c63430007060033",
}

// SafeL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeL2MetaData.ABI instead.
var SafeL2ABI = SafeL2MetaData.ABI

// SafeL2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeL2MetaData.Bin instead.
var SafeL2Bin = SafeL2MetaData.Bin

// DeploySafeL2 deploys a new Ethereum contract, binding an instance of SafeL2 to it.
func DeploySafeL2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeL2, error) {
	parsed, err := SafeL2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeL2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeL2{SafeL2Caller: SafeL2Caller{contract: contract}, SafeL2Transactor: SafeL2Transactor{contract: contract}, SafeL2Filterer: SafeL2Filterer{contract: contract}}, nil
}

// SafeL2 is an auto generated Go binding around an Ethereum contract.
type SafeL2 struct {
	SafeL2Caller     // Read-only binding to the contract
	SafeL2Transactor // Write-only binding to the contract
	SafeL2Filterer   // Log filterer for contract events
}

// SafeL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeL2Session struct {
	Contract     *SafeL2           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeL2CallerSession struct {
	Contract *SafeL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SafeL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeL2TransactorSession struct {
	Contract     *SafeL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeL2Raw struct {
	Contract *SafeL2 // Generic contract binding to access the raw methods on
}

// SafeL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeL2CallerRaw struct {
	Contract *SafeL2Caller // Generic read-only contract binding to access the raw methods on
}

// SafeL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeL2TransactorRaw struct {
	Contract *SafeL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeL2 creates a new instance of SafeL2, bound to a specific deployed contract.
func NewSafeL2(address common.Address, backend bind.ContractBackend) (*SafeL2, error) {
	contract, err := bindSafeL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeL2{SafeL2Caller: SafeL2Caller{contract: contract}, SafeL2Transactor: SafeL2Transactor{contract: contract}, SafeL2Filterer: SafeL2Filterer{contract: contract}}, nil
}

// NewSafeL2Caller creates a new read-only instance of SafeL2, bound to a specific deployed contract.
func NewSafeL2Caller(address common.Address, caller bind.ContractCaller) (*SafeL2Caller, error) {
	contract, err := bindSafeL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeL2Caller{contract: contract}, nil
}

// NewSafeL2Transactor creates a new write-only instance of SafeL2, bound to a specific deployed contract.
func NewSafeL2Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeL2Transactor, error) {
	contract, err := bindSafeL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeL2Transactor{contract: contract}, nil
}

// NewSafeL2Filterer creates a new log filterer instance of SafeL2, bound to a specific deployed contract.
func NewSafeL2Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeL2Filterer, error) {
	contract, err := bindSafeL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeL2Filterer{contract: contract}, nil
}

// bindSafeL2 binds a generic wrapper to an already deployed contract.
func bindSafeL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeL2 *SafeL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeL2.Contract.SafeL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeL2 *SafeL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeL2.Contract.SafeL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeL2 *SafeL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeL2.Contract.SafeL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeL2 *SafeL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeL2 *SafeL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeL2 *SafeL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeL2.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeL2 *SafeL2Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeL2 *SafeL2Session) VERSION() (string, error) {
	return _SafeL2.Contract.VERSION(&_SafeL2.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeL2 *SafeL2CallerSession) VERSION() (string, error) {
	return _SafeL2.Contract.VERSION(&_SafeL2.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_SafeL2 *SafeL2Caller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_SafeL2 *SafeL2Session) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _SafeL2.Contract.ApprovedHashes(&_SafeL2.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_SafeL2 *SafeL2CallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _SafeL2.Contract.ApprovedHashes(&_SafeL2.CallOpts, arg0, arg1)
}

// CheckNSignatures12fb68e0 is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeL2 *SafeL2Caller) CheckNSignatures12fb68e0(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures12fb68e0 is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeL2 *SafeL2Session) CheckNSignatures12fb68e0(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _SafeL2.Contract.CheckNSignatures12fb68e0(&_SafeL2.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures12fb68e0 is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeL2 *SafeL2CallerSession) CheckNSignatures12fb68e0(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _SafeL2.Contract.CheckNSignatures12fb68e0(&_SafeL2.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures1fcac7f3 is a free data retrieval call binding the contract method 0x1fcac7f3.
//
// Solidity: function checkNSignatures(address executor, bytes32 dataHash, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeL2 *SafeL2Caller) CheckNSignatures1fcac7f3(opts *bind.CallOpts, executor common.Address, dataHash [32]byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "checkNSignatures0", executor, dataHash, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures1fcac7f3 is a free data retrieval call binding the contract method 0x1fcac7f3.
//
// Solidity: function checkNSignatures(address executor, bytes32 dataHash, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeL2 *SafeL2Session) CheckNSignatures1fcac7f3(executor common.Address, dataHash [32]byte, signatures []byte, requiredSignatures *big.Int) error {
	return _SafeL2.Contract.CheckNSignatures1fcac7f3(&_SafeL2.CallOpts, executor, dataHash, signatures, requiredSignatures)
}

// CheckNSignatures1fcac7f3 is a free data retrieval call binding the contract method 0x1fcac7f3.
//
// Solidity: function checkNSignatures(address executor, bytes32 dataHash, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeL2 *SafeL2CallerSession) CheckNSignatures1fcac7f3(executor common.Address, dataHash [32]byte, signatures []byte, requiredSignatures *big.Int) error {
	return _SafeL2.Contract.CheckNSignatures1fcac7f3(&_SafeL2.CallOpts, executor, dataHash, signatures, requiredSignatures)
}

// CheckSignatures934f3a11 is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_SafeL2 *SafeL2Caller) CheckSignatures934f3a11(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures934f3a11 is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_SafeL2 *SafeL2Session) CheckSignatures934f3a11(dataHash [32]byte, data []byte, signatures []byte) error {
	return _SafeL2.Contract.CheckSignatures934f3a11(&_SafeL2.CallOpts, dataHash, data, signatures)
}

// CheckSignatures934f3a11 is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_SafeL2 *SafeL2CallerSession) CheckSignatures934f3a11(dataHash [32]byte, data []byte, signatures []byte) error {
	return _SafeL2.Contract.CheckSignatures934f3a11(&_SafeL2.CallOpts, dataHash, data, signatures)
}

// CheckSignaturesf855438b is a free data retrieval call binding the contract method 0xf855438b.
//
// Solidity: function checkSignatures(address executor, bytes32 dataHash, bytes signatures) view returns()
func (_SafeL2 *SafeL2Caller) CheckSignaturesf855438b(opts *bind.CallOpts, executor common.Address, dataHash [32]byte, signatures []byte) error {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "checkSignatures0", executor, dataHash, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignaturesf855438b is a free data retrieval call binding the contract method 0xf855438b.
//
// Solidity: function checkSignatures(address executor, bytes32 dataHash, bytes signatures) view returns()
func (_SafeL2 *SafeL2Session) CheckSignaturesf855438b(executor common.Address, dataHash [32]byte, signatures []byte) error {
	return _SafeL2.Contract.CheckSignaturesf855438b(&_SafeL2.CallOpts, executor, dataHash, signatures)
}

// CheckSignaturesf855438b is a free data retrieval call binding the contract method 0xf855438b.
//
// Solidity: function checkSignatures(address executor, bytes32 dataHash, bytes signatures) view returns()
func (_SafeL2 *SafeL2CallerSession) CheckSignaturesf855438b(executor common.Address, dataHash [32]byte, signatures []byte) error {
	return _SafeL2.Contract.CheckSignaturesf855438b(&_SafeL2.CallOpts, executor, dataHash, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeL2 *SafeL2Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeL2 *SafeL2Session) DomainSeparator() ([32]byte, error) {
	return _SafeL2.Contract.DomainSeparator(&_SafeL2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeL2 *SafeL2CallerSession) DomainSeparator() ([32]byte, error) {
	return _SafeL2.Contract.DomainSeparator(&_SafeL2.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeL2 *SafeL2Caller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeL2 *SafeL2Session) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SafeL2.Contract.GetModulesPaginated(&_SafeL2.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeL2 *SafeL2CallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SafeL2.Contract.GetModulesPaginated(&_SafeL2.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeL2 *SafeL2Caller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeL2 *SafeL2Session) GetOwners() ([]common.Address, error) {
	return _SafeL2.Contract.GetOwners(&_SafeL2.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeL2 *SafeL2CallerSession) GetOwners() ([]common.Address, error) {
	return _SafeL2.Contract.GetOwners(&_SafeL2.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_SafeL2 *SafeL2Caller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_SafeL2 *SafeL2Session) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _SafeL2.Contract.GetStorageAt(&_SafeL2.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_SafeL2 *SafeL2CallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _SafeL2.Contract.GetStorageAt(&_SafeL2.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeL2 *SafeL2Caller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeL2 *SafeL2Session) GetThreshold() (*big.Int, error) {
	return _SafeL2.Contract.GetThreshold(&_SafeL2.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeL2 *SafeL2CallerSession) GetThreshold() (*big.Int, error) {
	return _SafeL2.Contract.GetThreshold(&_SafeL2.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32 txHash)
func (_SafeL2 *SafeL2Caller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32 txHash)
func (_SafeL2 *SafeL2Session) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _SafeL2.Contract.GetTransactionHash(&_SafeL2.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32 txHash)
func (_SafeL2 *SafeL2CallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _SafeL2.Contract.GetTransactionHash(&_SafeL2.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeL2 *SafeL2Caller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeL2 *SafeL2Session) IsModuleEnabled(module common.Address) (bool, error) {
	return _SafeL2.Contract.IsModuleEnabled(&_SafeL2.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeL2 *SafeL2CallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SafeL2.Contract.IsModuleEnabled(&_SafeL2.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeL2 *SafeL2Caller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeL2 *SafeL2Session) IsOwner(owner common.Address) (bool, error) {
	return _SafeL2.Contract.IsOwner(&_SafeL2.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeL2 *SafeL2CallerSession) IsOwner(owner common.Address) (bool, error) {
	return _SafeL2.Contract.IsOwner(&_SafeL2.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeL2 *SafeL2Caller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeL2 *SafeL2Session) Nonce() (*big.Int, error) {
	return _SafeL2.Contract.Nonce(&_SafeL2.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeL2 *SafeL2CallerSession) Nonce() (*big.Int, error) {
	return _SafeL2.Contract.Nonce(&_SafeL2.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_SafeL2 *SafeL2Caller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SafeL2.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_SafeL2 *SafeL2Session) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _SafeL2.Contract.SignedMessages(&_SafeL2.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_SafeL2 *SafeL2CallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _SafeL2.Contract.SignedMessages(&_SafeL2.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeL2 *SafeL2Transactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeL2 *SafeL2Session) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.Contract.AddOwnerWithThreshold(&_SafeL2.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeL2 *SafeL2TransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.Contract.AddOwnerWithThreshold(&_SafeL2.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_SafeL2 *SafeL2Transactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_SafeL2 *SafeL2Session) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _SafeL2.Contract.ApproveHash(&_SafeL2.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_SafeL2 *SafeL2TransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _SafeL2.Contract.ApproveHash(&_SafeL2.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeL2 *SafeL2Transactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeL2 *SafeL2Session) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.Contract.ChangeThreshold(&_SafeL2.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeL2 *SafeL2TransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.Contract.ChangeThreshold(&_SafeL2.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeL2 *SafeL2Transactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeL2 *SafeL2Session) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.DisableModule(&_SafeL2.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeL2 *SafeL2TransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.DisableModule(&_SafeL2.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeL2 *SafeL2Transactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeL2 *SafeL2Session) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.EnableModule(&_SafeL2.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeL2 *SafeL2TransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.EnableModule(&_SafeL2.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_SafeL2 *SafeL2Transactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_SafeL2 *SafeL2Session) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _SafeL2.Contract.ExecTransaction(&_SafeL2.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_SafeL2 *SafeL2TransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _SafeL2.Contract.ExecTransaction(&_SafeL2.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_SafeL2 *SafeL2Transactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_SafeL2 *SafeL2Session) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeL2.Contract.ExecTransactionFromModule(&_SafeL2.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_SafeL2 *SafeL2TransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeL2.Contract.ExecTransactionFromModule(&_SafeL2.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SafeL2 *SafeL2Transactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SafeL2 *SafeL2Session) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeL2.Contract.ExecTransactionFromModuleReturnData(&_SafeL2.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SafeL2 *SafeL2TransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeL2.Contract.ExecTransactionFromModuleReturnData(&_SafeL2.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeL2 *SafeL2Transactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeL2 *SafeL2Session) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.Contract.RemoveOwner(&_SafeL2.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeL2 *SafeL2TransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeL2.Contract.RemoveOwner(&_SafeL2.TransactOpts, prevOwner, owner, _threshold)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeL2 *SafeL2Transactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeL2 *SafeL2Session) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SetFallbackHandler(&_SafeL2.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeL2 *SafeL2TransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SetFallbackHandler(&_SafeL2.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeL2 *SafeL2Transactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeL2 *SafeL2Session) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SetGuard(&_SafeL2.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeL2 *SafeL2TransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SetGuard(&_SafeL2.TransactOpts, guard)
}

// SetModuleGuard is a paid mutator transaction binding the contract method 0xe068df37.
//
// Solidity: function setModuleGuard(address moduleGuard) returns()
func (_SafeL2 *SafeL2Transactor) SetModuleGuard(opts *bind.TransactOpts, moduleGuard common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "setModuleGuard", moduleGuard)
}

// SetModuleGuard is a paid mutator transaction binding the contract method 0xe068df37.
//
// Solidity: function setModuleGuard(address moduleGuard) returns()
func (_SafeL2 *SafeL2Session) SetModuleGuard(moduleGuard common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SetModuleGuard(&_SafeL2.TransactOpts, moduleGuard)
}

// SetModuleGuard is a paid mutator transaction binding the contract method 0xe068df37.
//
// Solidity: function setModuleGuard(address moduleGuard) returns()
func (_SafeL2 *SafeL2TransactorSession) SetModuleGuard(moduleGuard common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SetModuleGuard(&_SafeL2.TransactOpts, moduleGuard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_SafeL2 *SafeL2Transactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_SafeL2 *SafeL2Session) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.Setup(&_SafeL2.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_SafeL2 *SafeL2TransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.Setup(&_SafeL2.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_SafeL2 *SafeL2Transactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_SafeL2 *SafeL2Session) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _SafeL2.Contract.SimulateAndRevert(&_SafeL2.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_SafeL2 *SafeL2TransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _SafeL2.Contract.SimulateAndRevert(&_SafeL2.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeL2 *SafeL2Transactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeL2.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeL2 *SafeL2Session) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SwapOwner(&_SafeL2.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeL2 *SafeL2TransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeL2.Contract.SwapOwner(&_SafeL2.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SafeL2 *SafeL2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SafeL2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SafeL2 *SafeL2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SafeL2.Contract.Fallback(&_SafeL2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SafeL2 *SafeL2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SafeL2.Contract.Fallback(&_SafeL2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SafeL2 *SafeL2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeL2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SafeL2 *SafeL2Session) Receive() (*types.Transaction, error) {
	return _SafeL2.Contract.Receive(&_SafeL2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SafeL2 *SafeL2TransactorSession) Receive() (*types.Transaction, error) {
	return _SafeL2.Contract.Receive(&_SafeL2.TransactOpts)
}

// SafeL2AddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the SafeL2 contract.
type SafeL2AddedOwnerIterator struct {
	Event *SafeL2AddedOwner // Event containing the contract specifics and raw log

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
func (it *SafeL2AddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2AddedOwner)
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
		it.Event = new(SafeL2AddedOwner)
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
func (it *SafeL2AddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2AddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2AddedOwner represents a AddedOwner event raised by the SafeL2 contract.
type SafeL2AddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_SafeL2 *SafeL2Filterer) FilterAddedOwner(opts *bind.FilterOpts, owner []common.Address) (*SafeL2AddedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "AddedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2AddedOwnerIterator{contract: _SafeL2.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_SafeL2 *SafeL2Filterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *SafeL2AddedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "AddedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2AddedOwner)
				if err := _SafeL2.contract.UnpackLog(event, "AddedOwner", log); err != nil {
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

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address indexed owner)
func (_SafeL2 *SafeL2Filterer) ParseAddedOwner(log types.Log) (*SafeL2AddedOwner, error) {
	event := new(SafeL2AddedOwner)
	if err := _SafeL2.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the SafeL2 contract.
type SafeL2ApproveHashIterator struct {
	Event *SafeL2ApproveHash // Event containing the contract specifics and raw log

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
func (it *SafeL2ApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ApproveHash)
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
		it.Event = new(SafeL2ApproveHash)
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
func (it *SafeL2ApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ApproveHash represents a ApproveHash event raised by the SafeL2 contract.
type SafeL2ApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_SafeL2 *SafeL2Filterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*SafeL2ApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ApproveHashIterator{contract: _SafeL2.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_SafeL2 *SafeL2Filterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *SafeL2ApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ApproveHash)
				if err := _SafeL2.contract.UnpackLog(event, "ApproveHash", log); err != nil {
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

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_SafeL2 *SafeL2Filterer) ParseApproveHash(log types.Log) (*SafeL2ApproveHash, error) {
	event := new(SafeL2ApproveHash)
	if err := _SafeL2.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the SafeL2 contract.
type SafeL2ChangedFallbackHandlerIterator struct {
	Event *SafeL2ChangedFallbackHandler // Event containing the contract specifics and raw log

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
func (it *SafeL2ChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ChangedFallbackHandler)
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
		it.Event = new(SafeL2ChangedFallbackHandler)
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
func (it *SafeL2ChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ChangedFallbackHandler represents a ChangedFallbackHandler event raised by the SafeL2 contract.
type SafeL2ChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_SafeL2 *SafeL2Filterer) FilterChangedFallbackHandler(opts *bind.FilterOpts, handler []common.Address) (*SafeL2ChangedFallbackHandlerIterator, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ChangedFallbackHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ChangedFallbackHandlerIterator{contract: _SafeL2.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_SafeL2 *SafeL2Filterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *SafeL2ChangedFallbackHandler, handler []common.Address) (event.Subscription, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ChangedFallbackHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ChangedFallbackHandler)
				if err := _SafeL2.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
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

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address indexed handler)
func (_SafeL2 *SafeL2Filterer) ParseChangedFallbackHandler(log types.Log) (*SafeL2ChangedFallbackHandler, error) {
	event := new(SafeL2ChangedFallbackHandler)
	if err := _SafeL2.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the SafeL2 contract.
type SafeL2ChangedGuardIterator struct {
	Event *SafeL2ChangedGuard // Event containing the contract specifics and raw log

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
func (it *SafeL2ChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ChangedGuard)
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
		it.Event = new(SafeL2ChangedGuard)
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
func (it *SafeL2ChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ChangedGuard represents a ChangedGuard event raised by the SafeL2 contract.
type SafeL2ChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_SafeL2 *SafeL2Filterer) FilterChangedGuard(opts *bind.FilterOpts, guard []common.Address) (*SafeL2ChangedGuardIterator, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ChangedGuard", guardRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ChangedGuardIterator{contract: _SafeL2.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_SafeL2 *SafeL2Filterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *SafeL2ChangedGuard, guard []common.Address) (event.Subscription, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ChangedGuard", guardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ChangedGuard)
				if err := _SafeL2.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
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

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address indexed guard)
func (_SafeL2 *SafeL2Filterer) ParseChangedGuard(log types.Log) (*SafeL2ChangedGuard, error) {
	event := new(SafeL2ChangedGuard)
	if err := _SafeL2.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ChangedModuleGuardIterator is returned from FilterChangedModuleGuard and is used to iterate over the raw logs and unpacked data for ChangedModuleGuard events raised by the SafeL2 contract.
type SafeL2ChangedModuleGuardIterator struct {
	Event *SafeL2ChangedModuleGuard // Event containing the contract specifics and raw log

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
func (it *SafeL2ChangedModuleGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ChangedModuleGuard)
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
		it.Event = new(SafeL2ChangedModuleGuard)
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
func (it *SafeL2ChangedModuleGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ChangedModuleGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ChangedModuleGuard represents a ChangedModuleGuard event raised by the SafeL2 contract.
type SafeL2ChangedModuleGuard struct {
	ModuleGuard common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedModuleGuard is a free log retrieval operation binding the contract event 0xcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb8.
//
// Solidity: event ChangedModuleGuard(address indexed moduleGuard)
func (_SafeL2 *SafeL2Filterer) FilterChangedModuleGuard(opts *bind.FilterOpts, moduleGuard []common.Address) (*SafeL2ChangedModuleGuardIterator, error) {

	var moduleGuardRule []interface{}
	for _, moduleGuardItem := range moduleGuard {
		moduleGuardRule = append(moduleGuardRule, moduleGuardItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ChangedModuleGuard", moduleGuardRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ChangedModuleGuardIterator{contract: _SafeL2.contract, event: "ChangedModuleGuard", logs: logs, sub: sub}, nil
}

// WatchChangedModuleGuard is a free log subscription operation binding the contract event 0xcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb8.
//
// Solidity: event ChangedModuleGuard(address indexed moduleGuard)
func (_SafeL2 *SafeL2Filterer) WatchChangedModuleGuard(opts *bind.WatchOpts, sink chan<- *SafeL2ChangedModuleGuard, moduleGuard []common.Address) (event.Subscription, error) {

	var moduleGuardRule []interface{}
	for _, moduleGuardItem := range moduleGuard {
		moduleGuardRule = append(moduleGuardRule, moduleGuardItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ChangedModuleGuard", moduleGuardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ChangedModuleGuard)
				if err := _SafeL2.contract.UnpackLog(event, "ChangedModuleGuard", log); err != nil {
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

// ParseChangedModuleGuard is a log parse operation binding the contract event 0xcd1966d6be16bc0c030cc741a06c6e0efaf8d00de2c8b6a9e11827e125de8bb8.
//
// Solidity: event ChangedModuleGuard(address indexed moduleGuard)
func (_SafeL2 *SafeL2Filterer) ParseChangedModuleGuard(log types.Log) (*SafeL2ChangedModuleGuard, error) {
	event := new(SafeL2ChangedModuleGuard)
	if err := _SafeL2.contract.UnpackLog(event, "ChangedModuleGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the SafeL2 contract.
type SafeL2ChangedThresholdIterator struct {
	Event *SafeL2ChangedThreshold // Event containing the contract specifics and raw log

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
func (it *SafeL2ChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ChangedThreshold)
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
		it.Event = new(SafeL2ChangedThreshold)
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
func (it *SafeL2ChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ChangedThreshold represents a ChangedThreshold event raised by the SafeL2 contract.
type SafeL2ChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_SafeL2 *SafeL2Filterer) FilterChangedThreshold(opts *bind.FilterOpts) (*SafeL2ChangedThresholdIterator, error) {

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &SafeL2ChangedThresholdIterator{contract: _SafeL2.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_SafeL2 *SafeL2Filterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *SafeL2ChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ChangedThreshold)
				if err := _SafeL2.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
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

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_SafeL2 *SafeL2Filterer) ParseChangedThreshold(log types.Log) (*SafeL2ChangedThreshold, error) {
	event := new(SafeL2ChangedThreshold)
	if err := _SafeL2.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2DisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the SafeL2 contract.
type SafeL2DisabledModuleIterator struct {
	Event *SafeL2DisabledModule // Event containing the contract specifics and raw log

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
func (it *SafeL2DisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2DisabledModule)
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
		it.Event = new(SafeL2DisabledModule)
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
func (it *SafeL2DisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2DisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2DisabledModule represents a DisabledModule event raised by the SafeL2 contract.
type SafeL2DisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_SafeL2 *SafeL2Filterer) FilterDisabledModule(opts *bind.FilterOpts, module []common.Address) (*SafeL2DisabledModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "DisabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2DisabledModuleIterator{contract: _SafeL2.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_SafeL2 *SafeL2Filterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *SafeL2DisabledModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "DisabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2DisabledModule)
				if err := _SafeL2.contract.UnpackLog(event, "DisabledModule", log); err != nil {
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

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address indexed module)
func (_SafeL2 *SafeL2Filterer) ParseDisabledModule(log types.Log) (*SafeL2DisabledModule, error) {
	event := new(SafeL2DisabledModule)
	if err := _SafeL2.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2EnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the SafeL2 contract.
type SafeL2EnabledModuleIterator struct {
	Event *SafeL2EnabledModule // Event containing the contract specifics and raw log

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
func (it *SafeL2EnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2EnabledModule)
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
		it.Event = new(SafeL2EnabledModule)
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
func (it *SafeL2EnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2EnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2EnabledModule represents a EnabledModule event raised by the SafeL2 contract.
type SafeL2EnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_SafeL2 *SafeL2Filterer) FilterEnabledModule(opts *bind.FilterOpts, module []common.Address) (*SafeL2EnabledModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "EnabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2EnabledModuleIterator{contract: _SafeL2.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_SafeL2 *SafeL2Filterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *SafeL2EnabledModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "EnabledModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2EnabledModule)
				if err := _SafeL2.contract.UnpackLog(event, "EnabledModule", log); err != nil {
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

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address indexed module)
func (_SafeL2 *SafeL2Filterer) ParseEnabledModule(log types.Log) (*SafeL2EnabledModule, error) {
	event := new(SafeL2EnabledModule)
	if err := _SafeL2.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the SafeL2 contract.
type SafeL2ExecutionFailureIterator struct {
	Event *SafeL2ExecutionFailure // Event containing the contract specifics and raw log

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
func (it *SafeL2ExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ExecutionFailure)
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
		it.Event = new(SafeL2ExecutionFailure)
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
func (it *SafeL2ExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ExecutionFailure represents a ExecutionFailure event raised by the SafeL2 contract.
type SafeL2ExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_SafeL2 *SafeL2Filterer) FilterExecutionFailure(opts *bind.FilterOpts, txHash [][32]byte) (*SafeL2ExecutionFailureIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ExecutionFailure", txHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ExecutionFailureIterator{contract: _SafeL2.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_SafeL2 *SafeL2Filterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *SafeL2ExecutionFailure, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ExecutionFailure", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ExecutionFailure)
				if err := _SafeL2.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 indexed txHash, uint256 payment)
func (_SafeL2 *SafeL2Filterer) ParseExecutionFailure(log types.Log) (*SafeL2ExecutionFailure, error) {
	event := new(SafeL2ExecutionFailure)
	if err := _SafeL2.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the SafeL2 contract.
type SafeL2ExecutionFromModuleFailureIterator struct {
	Event *SafeL2ExecutionFromModuleFailure // Event containing the contract specifics and raw log

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
func (it *SafeL2ExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ExecutionFromModuleFailure)
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
		it.Event = new(SafeL2ExecutionFromModuleFailure)
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
func (it *SafeL2ExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the SafeL2 contract.
type SafeL2ExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SafeL2 *SafeL2Filterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*SafeL2ExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ExecutionFromModuleFailureIterator{contract: _SafeL2.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SafeL2 *SafeL2Filterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *SafeL2ExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ExecutionFromModuleFailure)
				if err := _SafeL2.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
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

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SafeL2 *SafeL2Filterer) ParseExecutionFromModuleFailure(log types.Log) (*SafeL2ExecutionFromModuleFailure, error) {
	event := new(SafeL2ExecutionFromModuleFailure)
	if err := _SafeL2.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the SafeL2 contract.
type SafeL2ExecutionFromModuleSuccessIterator struct {
	Event *SafeL2ExecutionFromModuleSuccess // Event containing the contract specifics and raw log

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
func (it *SafeL2ExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ExecutionFromModuleSuccess)
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
		it.Event = new(SafeL2ExecutionFromModuleSuccess)
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
func (it *SafeL2ExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the SafeL2 contract.
type SafeL2ExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SafeL2 *SafeL2Filterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*SafeL2ExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ExecutionFromModuleSuccessIterator{contract: _SafeL2.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SafeL2 *SafeL2Filterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *SafeL2ExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ExecutionFromModuleSuccess)
				if err := _SafeL2.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
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

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SafeL2 *SafeL2Filterer) ParseExecutionFromModuleSuccess(log types.Log) (*SafeL2ExecutionFromModuleSuccess, error) {
	event := new(SafeL2ExecutionFromModuleSuccess)
	if err := _SafeL2.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2ExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the SafeL2 contract.
type SafeL2ExecutionSuccessIterator struct {
	Event *SafeL2ExecutionSuccess // Event containing the contract specifics and raw log

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
func (it *SafeL2ExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2ExecutionSuccess)
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
		it.Event = new(SafeL2ExecutionSuccess)
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
func (it *SafeL2ExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2ExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2ExecutionSuccess represents a ExecutionSuccess event raised by the SafeL2 contract.
type SafeL2ExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_SafeL2 *SafeL2Filterer) FilterExecutionSuccess(opts *bind.FilterOpts, txHash [][32]byte) (*SafeL2ExecutionSuccessIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "ExecutionSuccess", txHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2ExecutionSuccessIterator{contract: _SafeL2.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_SafeL2 *SafeL2Filterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *SafeL2ExecutionSuccess, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "ExecutionSuccess", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2ExecutionSuccess)
				if err := _SafeL2.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
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

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 indexed txHash, uint256 payment)
func (_SafeL2 *SafeL2Filterer) ParseExecutionSuccess(log types.Log) (*SafeL2ExecutionSuccess, error) {
	event := new(SafeL2ExecutionSuccess)
	if err := _SafeL2.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2RemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the SafeL2 contract.
type SafeL2RemovedOwnerIterator struct {
	Event *SafeL2RemovedOwner // Event containing the contract specifics and raw log

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
func (it *SafeL2RemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2RemovedOwner)
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
		it.Event = new(SafeL2RemovedOwner)
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
func (it *SafeL2RemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2RemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2RemovedOwner represents a RemovedOwner event raised by the SafeL2 contract.
type SafeL2RemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_SafeL2 *SafeL2Filterer) FilterRemovedOwner(opts *bind.FilterOpts, owner []common.Address) (*SafeL2RemovedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "RemovedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2RemovedOwnerIterator{contract: _SafeL2.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_SafeL2 *SafeL2Filterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *SafeL2RemovedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "RemovedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2RemovedOwner)
				if err := _SafeL2.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
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

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address indexed owner)
func (_SafeL2 *SafeL2Filterer) ParseRemovedOwner(log types.Log) (*SafeL2RemovedOwner, error) {
	event := new(SafeL2RemovedOwner)
	if err := _SafeL2.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2SafeModuleTransactionIterator is returned from FilterSafeModuleTransaction and is used to iterate over the raw logs and unpacked data for SafeModuleTransaction events raised by the SafeL2 contract.
type SafeL2SafeModuleTransactionIterator struct {
	Event *SafeL2SafeModuleTransaction // Event containing the contract specifics and raw log

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
func (it *SafeL2SafeModuleTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2SafeModuleTransaction)
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
		it.Event = new(SafeL2SafeModuleTransaction)
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
func (it *SafeL2SafeModuleTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2SafeModuleTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2SafeModuleTransaction represents a SafeModuleTransaction event raised by the SafeL2 contract.
type SafeL2SafeModuleTransaction struct {
	Module    common.Address
	To        common.Address
	Value     *big.Int
	Data      []byte
	Operation uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSafeModuleTransaction is a free log retrieval operation binding the contract event 0xb648d3644f584ed1c2232d53c46d87e693586486ad0d1175f8656013110b714e.
//
// Solidity: event SafeModuleTransaction(address module, address to, uint256 value, bytes data, uint8 operation)
func (_SafeL2 *SafeL2Filterer) FilterSafeModuleTransaction(opts *bind.FilterOpts) (*SafeL2SafeModuleTransactionIterator, error) {

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "SafeModuleTransaction")
	if err != nil {
		return nil, err
	}
	return &SafeL2SafeModuleTransactionIterator{contract: _SafeL2.contract, event: "SafeModuleTransaction", logs: logs, sub: sub}, nil
}

// WatchSafeModuleTransaction is a free log subscription operation binding the contract event 0xb648d3644f584ed1c2232d53c46d87e693586486ad0d1175f8656013110b714e.
//
// Solidity: event SafeModuleTransaction(address module, address to, uint256 value, bytes data, uint8 operation)
func (_SafeL2 *SafeL2Filterer) WatchSafeModuleTransaction(opts *bind.WatchOpts, sink chan<- *SafeL2SafeModuleTransaction) (event.Subscription, error) {

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "SafeModuleTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2SafeModuleTransaction)
				if err := _SafeL2.contract.UnpackLog(event, "SafeModuleTransaction", log); err != nil {
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

// ParseSafeModuleTransaction is a log parse operation binding the contract event 0xb648d3644f584ed1c2232d53c46d87e693586486ad0d1175f8656013110b714e.
//
// Solidity: event SafeModuleTransaction(address module, address to, uint256 value, bytes data, uint8 operation)
func (_SafeL2 *SafeL2Filterer) ParseSafeModuleTransaction(log types.Log) (*SafeL2SafeModuleTransaction, error) {
	event := new(SafeL2SafeModuleTransaction)
	if err := _SafeL2.contract.UnpackLog(event, "SafeModuleTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2SafeMultiSigTransactionIterator is returned from FilterSafeMultiSigTransaction and is used to iterate over the raw logs and unpacked data for SafeMultiSigTransaction events raised by the SafeL2 contract.
type SafeL2SafeMultiSigTransactionIterator struct {
	Event *SafeL2SafeMultiSigTransaction // Event containing the contract specifics and raw log

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
func (it *SafeL2SafeMultiSigTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2SafeMultiSigTransaction)
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
		it.Event = new(SafeL2SafeMultiSigTransaction)
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
func (it *SafeL2SafeMultiSigTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2SafeMultiSigTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2SafeMultiSigTransaction represents a SafeMultiSigTransaction event raised by the SafeL2 contract.
type SafeL2SafeMultiSigTransaction struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signatures     []byte
	AdditionalInfo []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSafeMultiSigTransaction is a free log retrieval operation binding the contract event 0x66753cd2356569ee081232e3be8909b950e0a76c1f8460c3a5e3c2be32b11bed.
//
// Solidity: event SafeMultiSigTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, bytes additionalInfo)
func (_SafeL2 *SafeL2Filterer) FilterSafeMultiSigTransaction(opts *bind.FilterOpts) (*SafeL2SafeMultiSigTransactionIterator, error) {

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "SafeMultiSigTransaction")
	if err != nil {
		return nil, err
	}
	return &SafeL2SafeMultiSigTransactionIterator{contract: _SafeL2.contract, event: "SafeMultiSigTransaction", logs: logs, sub: sub}, nil
}

// WatchSafeMultiSigTransaction is a free log subscription operation binding the contract event 0x66753cd2356569ee081232e3be8909b950e0a76c1f8460c3a5e3c2be32b11bed.
//
// Solidity: event SafeMultiSigTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, bytes additionalInfo)
func (_SafeL2 *SafeL2Filterer) WatchSafeMultiSigTransaction(opts *bind.WatchOpts, sink chan<- *SafeL2SafeMultiSigTransaction) (event.Subscription, error) {

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "SafeMultiSigTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2SafeMultiSigTransaction)
				if err := _SafeL2.contract.UnpackLog(event, "SafeMultiSigTransaction", log); err != nil {
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

// ParseSafeMultiSigTransaction is a log parse operation binding the contract event 0x66753cd2356569ee081232e3be8909b950e0a76c1f8460c3a5e3c2be32b11bed.
//
// Solidity: event SafeMultiSigTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures, bytes additionalInfo)
func (_SafeL2 *SafeL2Filterer) ParseSafeMultiSigTransaction(log types.Log) (*SafeL2SafeMultiSigTransaction, error) {
	event := new(SafeL2SafeMultiSigTransaction)
	if err := _SafeL2.contract.UnpackLog(event, "SafeMultiSigTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2SafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the SafeL2 contract.
type SafeL2SafeReceivedIterator struct {
	Event *SafeL2SafeReceived // Event containing the contract specifics and raw log

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
func (it *SafeL2SafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2SafeReceived)
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
		it.Event = new(SafeL2SafeReceived)
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
func (it *SafeL2SafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2SafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2SafeReceived represents a SafeReceived event raised by the SafeL2 contract.
type SafeL2SafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_SafeL2 *SafeL2Filterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*SafeL2SafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2SafeReceivedIterator{contract: _SafeL2.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_SafeL2 *SafeL2Filterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *SafeL2SafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2SafeReceived)
				if err := _SafeL2.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_SafeL2 *SafeL2Filterer) ParseSafeReceived(log types.Log) (*SafeL2SafeReceived, error) {
	event := new(SafeL2SafeReceived)
	if err := _SafeL2.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2SafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the SafeL2 contract.
type SafeL2SafeSetupIterator struct {
	Event *SafeL2SafeSetup // Event containing the contract specifics and raw log

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
func (it *SafeL2SafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2SafeSetup)
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
		it.Event = new(SafeL2SafeSetup)
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
func (it *SafeL2SafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2SafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2SafeSetup represents a SafeSetup event raised by the SafeL2 contract.
type SafeL2SafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_SafeL2 *SafeL2Filterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*SafeL2SafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2SafeSetupIterator{contract: _SafeL2.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_SafeL2 *SafeL2Filterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *SafeL2SafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2SafeSetup)
				if err := _SafeL2.contract.UnpackLog(event, "SafeSetup", log); err != nil {
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

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_SafeL2 *SafeL2Filterer) ParseSafeSetup(log types.Log) (*SafeL2SafeSetup, error) {
	event := new(SafeL2SafeSetup)
	if err := _SafeL2.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeL2SignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the SafeL2 contract.
type SafeL2SignMsgIterator struct {
	Event *SafeL2SignMsg // Event containing the contract specifics and raw log

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
func (it *SafeL2SignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeL2SignMsg)
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
		it.Event = new(SafeL2SignMsg)
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
func (it *SafeL2SignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeL2SignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeL2SignMsg represents a SignMsg event raised by the SafeL2 contract.
type SafeL2SignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SafeL2 *SafeL2Filterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*SafeL2SignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _SafeL2.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeL2SignMsgIterator{contract: _SafeL2.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SafeL2 *SafeL2Filterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *SafeL2SignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _SafeL2.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeL2SignMsg)
				if err := _SafeL2.contract.UnpackLog(event, "SignMsg", log); err != nil {
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

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SafeL2 *SafeL2Filterer) ParseSignMsg(log types.Log) (*SafeL2SignMsg, error) {
	event := new(SafeL2SignMsg)
	if err := _SafeL2.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
