package validator

import (
	"github.com/tenderly/net-nitro/go-ethereum/common"
	"github.com/tenderly/net-nitro/go-ethereum/core/rawdb"

	"github.com/tenderly/net-nitro/daprovider"
)

type BatchInfo struct {
	Number uint64
	Data   []byte
}

// lint:require-exhaustive-initialization
type ValidationInput struct {
	Id            uint64
	HasDelayedMsg bool
	DelayedMsgNr  uint64
	Preimages     daprovider.PreimagesMap
	UserWasms     map[rawdb.WasmTarget]map[common.Hash][]byte
	BatchInfo     []BatchInfo
	DelayedMsg    []byte
	StartState    GoGlobalState
	DebugChain    bool
}
