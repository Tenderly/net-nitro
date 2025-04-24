package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/tenderly/net-nitro/linters/koanf"
	"github.com/tenderly/net-nitro/linters/pointercheck"
	"github.com/tenderly/net-nitro/linters/rightshift"
	"github.com/tenderly/net-nitro/linters/structinit"
)

func main() {
	multichecker.Main(
		koanf.Analyzer,
		pointercheck.Analyzer,
		rightshift.Analyzer,
		structinit.Analyzer,
	)
}
