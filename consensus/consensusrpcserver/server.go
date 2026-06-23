// Copyright 2025-2026, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md
package consensusrpcserver

import (
	"context"

	"github.com/tenderly/net-nitro/go-ethereum/common"

	"github.com/tenderly/net-nitro/arbos/arbostypes"
	"github.com/tenderly/net-nitro/arbutil"
	"github.com/tenderly/net-nitro/consensus"
	"github.com/tenderly/net-nitro/execution"
)

type ConsensusRPCServer struct {
	consensus consensus.FullConsensusClient
}

func NewConsensusRPCServer(consensus consensus.FullConsensusClient) *ConsensusRPCServer {
	return &ConsensusRPCServer{consensus}
}

func (a *ConsensusRPCServer) GetL1Confirmations(ctx context.Context, msgIdx arbutil.MessageIndex) (uint64, error) {
	return a.consensus.GetL1Confirmations(msgIdx).Await(ctx)
}

func (a *ConsensusRPCServer) FindBatchContainingMessage(ctx context.Context, msgIdx arbutil.MessageIndex) (uint64, error) {
	return a.consensus.FindBatchContainingMessage(msgIdx).Await(ctx)
}

func (a *ConsensusRPCServer) BlockMetadataAtMessageIndex(ctx context.Context, msgIdx arbutil.MessageIndex) (common.BlockMetadata, error) {
	return a.consensus.BlockMetadataAtMessageIndex(msgIdx).Await(ctx)
}

func (a *ConsensusRPCServer) WriteMessageFromSequencer(ctx context.Context, msgIdx arbutil.MessageIndex, msgWithMeta arbostypes.MessageWithMetadata, msgResult execution.MessageResult, blockMetadata common.BlockMetadata) error {
	_, err := a.consensus.WriteMessageFromSequencer(msgIdx, msgWithMeta, msgResult, blockMetadata).Await(ctx)
	return err
}

func (a *ConsensusRPCServer) ExpectChosenSequencer(ctx context.Context) error {
	_, err := a.consensus.ExpectChosenSequencer().Await(ctx)
	return err
}
