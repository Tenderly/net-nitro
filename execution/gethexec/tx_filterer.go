// Copyright 2021-2026, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package gethexec

import (
	"time"

	"github.com/google/uuid"

	"github.com/tenderly/net-nitro/go-ethereum/arbitrum/filter"
	"github.com/tenderly/net-nitro/go-ethereum/common"
	"github.com/tenderly/net-nitro/go-ethereum/core/state"
	"github.com/tenderly/net-nitro/go-ethereum/core/types"
	"github.com/tenderly/net-nitro/go-ethereum/core/vm"
	"github.com/tenderly/net-nitro/go-ethereum/log"

	"github.com/tenderly/net-nitro/execution/gethexec/addressfilter"
	"github.com/tenderly/net-nitro/execution/gethexec/eventfilter"
)

// txFilterer implements core.TxFilterer for address-based transaction filtering
// for node API calls such as eth_estimateGas and eth_call. It wraps ExecutionEngine to resolve the address
// checker lazily, so tests can inject checkers via ExecEngine.SetAddressChecker.
type txFilterer struct {
	execEngine  *ExecutionEngine
	eventFilter *eventfilter.EventFilter
	// nil disables filtered-tx reporting (e.g. the eth_call backend filterer).
	filteringReportRPCClient *FilteringReportRPCClient
}

func (f *txFilterer) Setup(vmStatedb vm.StateDB) {
	statedb, ok := vmStatedb.(*state.StateDB)
	if !ok {
		// Non-state.StateDB backends (e.g. external simulation wrappers) skip tx filtering setup.
		return
	}
	if f.execEngine.addressChecker != nil {
		statedb.SetAddressCheckerState(f.execEngine.addressChecker.NewTxState())
	}
	statedb.SetTxContext(common.Hash{}, 0)
}

func (f *txFilterer) TouchAddresses(vmStatedb vm.StateDB, tx *types.Transaction, sender common.Address) {
	statedb, ok := vmStatedb.(*state.StateDB)
	if !ok {
		return
	}
	touchAddresses(statedb, tx, sender)
}

func (f *txFilterer) CheckFiltered(vmStatedb vm.StateDB, rootTx *types.Transaction, header *types.Header) error {
	statedb, ok := vmStatedb.(*state.StateDB)
	if !ok {
		return nil
	}
	applyEventFilter(f.eventFilter, statedb)
	if filtered, records := statedb.IsAddressFiltered(); filtered {
		if f.filteringReportRPCClient != nil {
			f.reportFilteredTx(rootTx, header, records)
		}
		return state.ErrArbTxFilter
	}
	return nil
}

func (f *txFilterer) reportFilteredTx(tx *types.Transaction, header *types.Header, filteredAddresses []filter.FilteredAddressRecord) {
	txHash := tx.Hash()
	txRLP, err := tx.MarshalBinary()
	if err != nil {
		log.Error("failed to marshal filtered tx", "txHash", txHash, "err", err)
		return
	}
	report := addressfilter.FilteredTxReport{
		ID:                uuid.Must(uuid.NewV7()).String(),
		TxHash:            txHash,
		TxRLP:             txRLP,
		FilteredAddresses: filteredAddresses,
		ChainID:           f.execEngine.bc.Config().ChainID.Uint64(),
		BlockNumber:       header.Number.Uint64() + 1,
		ParentBlockHash:   header.Hash(),
		PositionInBlock:   0,
		FilteredAt:        time.Now().UTC(),
		IsDelayed:         false,
		DelayedReportData: nil,
	}
	f.filteringReportRPCClient.ReportFilteredTransactions([]addressfilter.FilteredTxReport{report})
}
