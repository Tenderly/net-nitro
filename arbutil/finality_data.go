// Copyright 2021-2025, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package arbutil

import "github.com/tenderly/net-nitro/go-ethereum/common"

type FinalityData struct {
	MsgIdx    MessageIndex
	BlockHash common.Hash
}
