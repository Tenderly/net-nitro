package util

import (
	"context"
	"errors"
	"math/big"

	"github.com/tenderly/net-nitro/go-ethereum/ethclient"
	"github.com/tenderly/net-nitro/go-ethereum/rpc"

	"github.com/tenderly/net-nitro/bold/chain-abstraction"
)

var (
	_ protocol.ChainBackend = &BackendWrapper{
		desiredBlockNum: rpc.LatestBlockNumber,
	}
)

type ethClient = ethclient.Client
type BackendWrapper struct {
	*ethClient
	desiredBlockNum rpc.BlockNumber
}

func NewBackendWrapper(client *ethclient.Client, desiredBlockNum rpc.BlockNumber) *BackendWrapper {
	return &BackendWrapper{client, desiredBlockNum}
}

func (b BackendWrapper) HeaderU64(ctx context.Context) (uint64, error) {
	header, err := b.HeaderByNumber(ctx, big.NewInt(int64(b.desiredBlockNum)))
	if err != nil {
		return 0, err
	}
	if !header.Number.IsUint64() {
		return 0, errors.New("block number is not uint64")
	}
	return header.Number.Uint64(), nil
}
