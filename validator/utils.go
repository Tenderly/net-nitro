package validator

import (
	"github.com/tenderly/net-nitro/go-ethereum/common"
	"github.com/tenderly/net-nitro/go-ethereum/log"
)

func SpawnerSupportsModule(spawner ValidationSpawner, requested common.Hash) bool {
	supported, err := spawner.WasmModuleRoots()
	if err != nil {
		log.Warn("WasmModuleRoots returned error", "err", err)
		return false
	}
	for _, root := range supported {
		if root == requested {
			return true
		}
	}
	return false
}
