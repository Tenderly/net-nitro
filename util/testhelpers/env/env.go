// Copyright 2024, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package env

import (
	"os"

	"github.com/tenderly/net-nitro/go-ethereum/core/rawdb"
	"github.com/tenderly/net-nitro/go-ethereum/log"
)

// There are two CI steps, one to run tests using the path state scheme, and one to run tests using the hash state scheme.
// An environment variable controls that behavior.
func GetTestStateScheme() string {
	envTestStateScheme := os.Getenv("TEST_STATE_SCHEME")
	stateScheme := rawdb.PathScheme
	if envTestStateScheme == rawdb.PathScheme || envTestStateScheme == rawdb.HashScheme {
		stateScheme = envTestStateScheme
	}
	log.Debug("test state scheme", "testStateScheme", stateScheme)
	return stateScheme
}
