// Copyright 2023-2024, Offchain Labs, Inc.
// For license information, see:
// https://github.com/tenderly/net-nitro/blob/master/LICENSE.md

package solimpl

import "github.com/tenderly/net-nitro/bold/chain-abstraction"

func (a *AssertionChain) SetBackend(b protocol.ChainBackend) {
	a.backend = b
}
