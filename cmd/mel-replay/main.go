// Copyright 2025-2026, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package main

import (
	"github.com/tenderly/net-nitro/go-ethereum/common"

	"github.com/tenderly/net-nitro/arbutil"
)

type preimageResolver interface {
	ResolveTypedPreimage(preimageType arbutil.PreimageType, hash common.Hash) ([]byte, error)
}

func main() {
}
