// Copyright 2023-2026, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md
package server_common

import (
	"github.com/tenderly/net-nitro/go-ethereum/common"

	"github.com/tenderly/net-nitro/util/containers"
	"github.com/tenderly/net-nitro/validator"
)

type ValRun struct {
	containers.PromiseInterface[validator.GoGlobalState]
	root common.Hash
}

func (r *ValRun) WasmModuleRoot() common.Hash {
	return r.root
}

func NewValRun(promise containers.PromiseInterface[validator.GoGlobalState], root common.Hash) *ValRun {
	return &ValRun{
		PromiseInterface: promise,
		root:             root,
	}
}
