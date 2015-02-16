package core

import (
	"github.com/jonasnick/go-ethereum/crypto"
	"github.com/jonasnick/go-ethereum/ethutil"
	"github.com/jonasnick/go-ethereum/event"
	"github.com/jonasnick/go-ethereum/p2p"
)

type EthManager interface {
	BlockProcessor() *BlockProcessor
	ChainManager() *ChainManager
	TxPool() *TxPool
	PeerCount() int
	IsMining() bool
	IsListening() bool
	Peers() []*p2p.Peer
	KeyManager() *crypto.KeyManager
	Db() ethutil.Database
	EventMux() *event.TypeMux
}
