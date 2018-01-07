// Copyright 2015 The go-ruereum Authors
// This file is part of the go-ruereum library.
//
// The go-ruereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ruereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ruereum library. If not, see <http://www.gnu.org/licenses/>.

package rue

import (
	"context"
	"math/big"

	"github.com/Rue-Foundation/go-rue/accounts"
	"github.com/Rue-Foundation/go-rue/common"
	"github.com/Rue-Foundation/go-rue/common/math"
	"github.com/Rue-Foundation/go-rue/core"
	"github.com/Rue-Foundation/go-rue/core/bloombits"
	"github.com/Rue-Foundation/go-rue/core/state"
	"github.com/Rue-Foundation/go-rue/core/types"
	"github.com/Rue-Foundation/go-rue/core/vm"
	"github.com/Rue-Foundation/go-rue/rue/downloader"
	"github.com/Rue-Foundation/go-rue/rue/gasprice"
	"github.com/Rue-Foundation/go-rue/ruedb"
	"github.com/Rue-Foundation/go-rue/event"
	"github.com/Rue-Foundation/go-rue/params"
	"github.com/Rue-Foundation/go-rue/rpc"
)

// RueApiBackend implements rueapi.Backend for full nodes
type RueApiBackend struct {
rue*Ruereum
	gpo *gasprice.Oracle
}

func (b *RueApiBackend) ChainConfig() *params.ChainConfig {
	return b.rue.chainConfig
}

func (b *RueApiBackend) CurrentBlock() *types.Block {
	return b.rue.blockchain.CurrentBlock()
}

func (b *RueApiBackend) SetHead(number uint64) {
	b.rue.protocolManager.downloader.Cancel()
	b.rue.blockchain.SetHead(number)
}

func (b *RueApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.rue.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.rue.blockchain.CurrentBlock().Header(), nil
	}
	return b.rue.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *RueApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.rue.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.rue.blockchain.CurrentBlock(), nil
	}
	return b.rue.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *RueApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := b.rue.miner.Pending()
		return state, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, err := b.rue.BlockChain().StateAt(header.Root)
	return stateDb, header, err
}

func (b *RueApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.rue.blockchain.GetBlockByHash(blockHash), nil
}

func (b *RueApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.rue.chainDb, blockHash, core.GetBlockNumber(b.rue.chainDb, blockHash)), nil
}

func (b *RueApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.rue.blockchain.GetTdByHash(blockHash)
}

func (b *RueApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.rue.BlockChain(), nil)
	return vm.NewEVM(context, state, b.rue.chainConfig, vmCfg), vmError, nil
}

func (b *RueApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.rue.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *RueApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.rue.BlockChain().SubscribeChainEvent(ch)
}

func (b *RueApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.rue.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *RueApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.rue.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *RueApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.rue.BlockChain().SubscribeLogsEvent(ch)
}

func (b *RueApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.rue.txPool.AddLocal(signedTx)
}

func (b *RueApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.rue.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *RueApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.rue.txPool.Get(hash)
}

func (b *RueApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.rue.txPool.State().GetNonce(addr), nil
}

func (b *RueApiBackend) Stats() (pending int, queued int) {
	return b.rue.txPool.Stats()
}

func (b *RueApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.rue.TxPool().Content()
}

func (b *RueApiBackend) SubscribeTxPreEvent(ch chan<- core.TxPreEvent) event.Subscription {
	return b.rue.TxPool().SubscribeTxPreEvent(ch)
}

func (b *RueApiBackend) Downloader() *downloader.Downloader {
	return b.rue.Downloader()
}

func (b *RueApiBackend) ProtocolVersion() int {
	return b.rue.RueVersion()
}

func (b *RueApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *RueApiBackend) ChainDb() ruedb.Database {
	return b.rue.ChainDb()
}

func (b *RueApiBackend) EventMux() *event.TypeMux {
	return b.rue.EventMux()
}

func (b *RueApiBackend) AccountManager() *accounts.Manager {
	return b.rue.AccountManager()
}

func (b *RueApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.rue.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *RueApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.rue.bloomRequests)
	}
}
