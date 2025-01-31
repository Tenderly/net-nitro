package main

import (
	"github.com/tenderly/net-nitro/linters/koanf"
	"github.com/tenderly/net-nitro/linters/pointercheck"
	"github.com/tenderly/net-nitro/linters/rightshift"
	"github.com/tenderly/net-nitro/linters/structinit"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		koanf.Analyzer,
		pointercheck.Analyzer,
		rightshift.Analyzer,
		structinit.Analyzer,
	)
}
