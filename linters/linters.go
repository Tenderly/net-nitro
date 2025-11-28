// Copyright 2024-2025, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/tenderly/net-nitro/linters/jsonneverempty"
	"github.com/tenderly/net-nitro/linters/koanf"
	"github.com/tenderly/net-nitro/linters/namedfieldsinit"
	"github.com/tenderly/net-nitro/linters/pointercheck"
	"github.com/tenderly/net-nitro/linters/rightshift"
	"github.com/tenderly/net-nitro/linters/structinit"
)

func main() {
	multichecker.Main(
		koanf.Analyzer,
		namedfieldsinit.Analyzer,
		pointercheck.Analyzer,
		rightshift.Analyzer,
		structinit.Analyzer,
		jsonneverempty.Analyzer,
	)
}
