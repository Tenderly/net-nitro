// Copyright 2021-2026, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package gethexec

import (
	"github.com/tenderly/net-nitro/go-ethereum/common"
	"github.com/tenderly/net-nitro/go-ethereum/core/state"
	"github.com/tenderly/net-nitro/go-ethereum/core/types"
	"github.com/tenderly/net-nitro/go-ethereum/core/vm"

	"github.com/tenderly/net-nitro/execution/gethexec/eventfilter"
)

// txFilterer implements core.TxFilterer for address-based transaction filtering
// for node API calls such as eth_estimateGas and eth_call. It wraps ExecutionEngine to resolve the address
// checker lazily, so tests can inject checkers via ExecEngine.SetAddressChecker.
type txFilterer struct {
	execEngine  *ExecutionEngine
	eventFilter *eventfilter.EventFilter
}

func (f *txFilterer) Setup(vmStatedb vm.StateDB) {
	statedb, ok := vmStatedb.(*state.StateDB)
	if !ok {
		// Non-state.StateDB backends (e.g. external simulation wrappers) skip tx filtering setup.
		return
	}
	statedb.SetAddressChecker(f.execEngine.addressChecker)
	statedb.SetTxContext(common.Hash{}, 0)
}

func (f *txFilterer) TouchAddresses(vmStatedb vm.StateDB, tx *types.Transaction, sender common.Address) {
	statedb, ok := vmStatedb.(*state.StateDB)
	if !ok {
		return
	}
	touchAddresses(statedb, tx, sender)
}

func (f *txFilterer) CheckFiltered(vmStatedb vm.StateDB) error {
	statedb, ok := vmStatedb.(*state.StateDB)
	if !ok {
		return nil
	}
	applyEventFilter(f.eventFilter, statedb)
	if filtered, _ := statedb.IsAddressFiltered(); filtered {
		return state.ErrArbTxFilter
	}
	return nil
}
