package proxy

import (
<<<<<<< HEAD
	cmn "github.com/tendermint/tendermint/libs/common"
=======
	"github.com/tendermint/tendermint/libs/bytes"
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	lrpc "github.com/tendermint/tendermint/lite2/rpc"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpcserver "github.com/tendermint/tendermint/rpc/lib/server"
	rpctypes "github.com/tendermint/tendermint/rpc/lib/types"
	"github.com/tendermint/tendermint/types"
)

func RPCRoutes(c *lrpc.Client) map[string]*rpcserver.RPCFunc {
	return map[string]*rpcserver.RPCFunc{
		// Subscribe/unsubscribe are reserved for websocket events.
		"subscribe":       rpcserver.NewWSRPCFunc(c.SubscribeWS, "query"),
		"unsubscribe":     rpcserver.NewWSRPCFunc(c.UnsubscribeWS, "query"),
		"unsubscribe_all": rpcserver.NewWSRPCFunc(c.UnsubscribeAllWS, ""),

		// info API
		"health":               rpcserver.NewRPCFunc(makeHealthFunc(c), ""),
		"status":               rpcserver.NewRPCFunc(makeStatusFunc(c), ""),
		"net_info":             rpcserver.NewRPCFunc(makeNetInfoFunc(c), ""),
		"blockchain":           rpcserver.NewRPCFunc(makeBlockchainInfoFunc(c), "minHeight,maxHeight"),
		"genesis":              rpcserver.NewRPCFunc(makeGenesisFunc(c), ""),
		"block":                rpcserver.NewRPCFunc(makeBlockFunc(c), "height"),
		"block_results":        rpcserver.NewRPCFunc(makeBlockResultsFunc(c), "height"),
		"commit":               rpcserver.NewRPCFunc(makeCommitFunc(c), "height"),
		"tx":                   rpcserver.NewRPCFunc(makeTxFunc(c), "hash,prove"),
<<<<<<< HEAD
		"tx_search":            rpcserver.NewRPCFunc(makeTxSearchFunc(c), "query,prove,page,per_page"),
		"validators":           rpcserver.NewRPCFunc(makeValidatorsFunc(c), "height"),
=======
		"tx_search":            rpcserver.NewRPCFunc(makeTxSearchFunc(c), "query,prove,page,per_page,order_by"),
		"validators":           rpcserver.NewRPCFunc(makeValidatorsFunc(c), "height,page,per_page"),
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
		"dump_consensus_state": rpcserver.NewRPCFunc(makeDumpConsensusStateFunc(c), ""),
		"consensus_state":      rpcserver.NewRPCFunc(makeConsensusStateFunc(c), ""),
		"consensus_params":     rpcserver.NewRPCFunc(makeConsensusParamsFunc(c), "height"),
		"unconfirmed_txs":      rpcserver.NewRPCFunc(makeUnconfirmedTxsFunc(c), "limit"),
		"num_unconfirmed_txs":  rpcserver.NewRPCFunc(makeNumUnconfirmedTxsFunc(c), ""),

		// tx broadcast API
		"broadcast_tx_commit": rpcserver.NewRPCFunc(makeBroadcastTxCommitFunc(c), "tx"),
		"broadcast_tx_sync":   rpcserver.NewRPCFunc(makeBroadcastTxSyncFunc(c), "tx"),
		"broadcast_tx_async":  rpcserver.NewRPCFunc(makeBroadcastTxAsyncFunc(c), "tx"),

		// abci API
		"abci_query": rpcserver.NewRPCFunc(makeABCIQueryFunc(c), "path,data,height,prove"),
		"abci_info":  rpcserver.NewRPCFunc(makeABCIInfoFunc(c), ""),

		// evidence API
		"broadcast_evidence": rpcserver.NewRPCFunc(makeBroadcastEvidenceFunc(c), "evidence"),
	}
}

<<<<<<< HEAD
func makeHealthFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
=======
type rpcHealthFunc func(ctx *rpctypes.Context) (*ctypes.ResultHealth, error)

func makeHealthFunc(c *lrpc.Client) rpcHealthFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
		return c.Health()
	}
}

<<<<<<< HEAD
func makeStatusFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultStatus, error) {
=======
type rpcStatusFunc func(ctx *rpctypes.Context) (*ctypes.ResultStatus, error)

// nolint: interfacer
func makeStatusFunc(c *lrpc.Client) rpcStatusFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultStatus, error) {
		return c.Status()
	}
}

<<<<<<< HEAD
func makeNetInfoFunc(c *lrpc.Client) func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultNetInfo, error) {
=======
type rpcNetInfoFunc func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultNetInfo, error)

func makeNetInfoFunc(c *lrpc.Client) rpcNetInfoFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultNetInfo, error) {
		return c.NetInfo()
	}
}

<<<<<<< HEAD
func makeBlockchainInfoFunc(c *lrpc.Client) func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
=======
type rpcBlockchainInfoFunc func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error)

func makeBlockchainInfoFunc(c *lrpc.Client) rpcBlockchainInfoFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
		return c.BlockchainInfo(minHeight, maxHeight)
	}
}

<<<<<<< HEAD
func makeGenesisFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultGenesis, error) {
=======
type rpcGenesisFunc func(ctx *rpctypes.Context) (*ctypes.ResultGenesis, error)

func makeGenesisFunc(c *lrpc.Client) rpcGenesisFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultGenesis, error) {
		return c.Genesis()
	}
}

<<<<<<< HEAD
func makeBlockFunc(c *lrpc.Client) func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlock, error) {
=======
type rpcBlockFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlock, error)

func makeBlockFunc(c *lrpc.Client) rpcBlockFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlock, error) {
		return c.Block(height)
	}
}

<<<<<<< HEAD
func makeBlockResultsFunc(c *lrpc.Client) func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlockResults, error) {
=======
type rpcBlockResultsFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlockResults, error)

func makeBlockResultsFunc(c *lrpc.Client) rpcBlockResultsFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultBlockResults, error) {
		return c.BlockResults(height)
	}
}

<<<<<<< HEAD
func makeCommitFunc(c *lrpc.Client) func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultCommit, error) {
=======
type rpcCommitFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultCommit, error)

func makeCommitFunc(c *lrpc.Client) rpcCommitFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultCommit, error) {
		return c.Commit(height)
	}
}

<<<<<<< HEAD
func makeTxFunc(c *lrpc.Client) func(ctx *rpctypes.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
=======
type rpcTxFunc func(ctx *rpctypes.Context, hash []byte, prove bool) (*ctypes.ResultTx, error)

func makeTxFunc(c *lrpc.Client) rpcTxFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
		return c.Tx(hash, prove)
	}
}

<<<<<<< HEAD
func makeTxSearchFunc(c *lrpc.Client) func(ctx *rpctypes.Context, query string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
	return func(ctx *rpctypes.Context, query string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
		return c.TxSearch(query, prove, page, perPage)
	}
}

func makeValidatorsFunc(c *lrpc.Client) func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultValidators, error) {
	return func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultValidators, error) {
		return c.Validators(height)
	}
}

func makeDumpConsensusStateFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultDumpConsensusState, error) {
=======
type rpcTxSearchFunc func(ctx *rpctypes.Context, query string, prove bool,
	page, perPage int, orderBy string) (*ctypes.ResultTxSearch, error)

func makeTxSearchFunc(c *lrpc.Client) rpcTxSearchFunc {
	return func(ctx *rpctypes.Context, query string, prove bool, page, perPage int, orderBy string) (
		*ctypes.ResultTxSearch, error) {
		return c.TxSearch(query, prove, page, perPage, orderBy)
	}
}

type rpcValidatorsFunc func(ctx *rpctypes.Context, height *int64,
	page, perPage int) (*ctypes.ResultValidators, error)

func makeValidatorsFunc(c *lrpc.Client) rpcValidatorsFunc {
	return func(ctx *rpctypes.Context, height *int64, page, perPage int) (*ctypes.ResultValidators, error) {
		return c.Validators(height, page, perPage)
	}
}

type rpcDumpConsensusStateFunc func(ctx *rpctypes.Context) (*ctypes.ResultDumpConsensusState, error)

func makeDumpConsensusStateFunc(c *lrpc.Client) rpcDumpConsensusStateFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultDumpConsensusState, error) {
		return c.DumpConsensusState()
	}
}

<<<<<<< HEAD
func makeConsensusStateFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultConsensusState, error) {
=======
type rpcConsensusStateFunc func(ctx *rpctypes.Context) (*ctypes.ResultConsensusState, error)

func makeConsensusStateFunc(c *lrpc.Client) rpcConsensusStateFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultConsensusState, error) {
		return c.ConsensusState()
	}
}

<<<<<<< HEAD
func makeConsensusParamsFunc(c *lrpc.Client) func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
	return func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
		return c.ConsensusParams()
	}
}

func makeUnconfirmedTxsFunc(c *lrpc.Client) func(ctx *rpctypes.Context, limit int) (*ctypes.ResultUnconfirmedTxs, error) {
=======
type rpcConsensusParamsFunc func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultConsensusParams, error)

func makeConsensusParamsFunc(c *lrpc.Client) rpcConsensusParamsFunc {
	return func(ctx *rpctypes.Context, height *int64) (*ctypes.ResultConsensusParams, error) {
		return c.ConsensusParams(height)
	}
}

type rpcUnconfirmedTxsFunc func(ctx *rpctypes.Context, limit int) (*ctypes.ResultUnconfirmedTxs, error)

func makeUnconfirmedTxsFunc(c *lrpc.Client) rpcUnconfirmedTxsFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, limit int) (*ctypes.ResultUnconfirmedTxs, error) {
		return c.UnconfirmedTxs(limit)
	}
}

<<<<<<< HEAD
func makeNumUnconfirmedTxsFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultUnconfirmedTxs, error) {
=======
type rpcNumUnconfirmedTxsFunc func(ctx *rpctypes.Context) (*ctypes.ResultUnconfirmedTxs, error)

func makeNumUnconfirmedTxsFunc(c *lrpc.Client) rpcNumUnconfirmedTxsFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultUnconfirmedTxs, error) {
		return c.NumUnconfirmedTxs()
	}
}

<<<<<<< HEAD
func makeBroadcastTxCommitFunc(c *lrpc.Client) func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
=======
type rpcBroadcastTxCommitFunc func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error)

func makeBroadcastTxCommitFunc(c *lrpc.Client) rpcBroadcastTxCommitFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
		return c.BroadcastTxCommit(tx)
	}
}

<<<<<<< HEAD
func makeBroadcastTxSyncFunc(c *lrpc.Client) func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
=======
type rpcBroadcastTxSyncFunc func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error)

func makeBroadcastTxSyncFunc(c *lrpc.Client) rpcBroadcastTxSyncFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
		return c.BroadcastTxSync(tx)
	}
}

<<<<<<< HEAD
func makeBroadcastTxAsyncFunc(c *lrpc.Client) func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
=======
type rpcBroadcastTxAsyncFunc func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error)

func makeBroadcastTxAsyncFunc(c *lrpc.Client) rpcBroadcastTxAsyncFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
		return c.BroadcastTxAsync(tx)
	}
}

<<<<<<< HEAD
func makeABCIQueryFunc(c *lrpc.Client) func(ctx *rpctypes.Context, path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error) {
	return func(ctx *rpctypes.Context, path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error) {
=======
type rpcABCIQueryFunc func(ctx *rpctypes.Context, path string, data bytes.HexBytes) (*ctypes.ResultABCIQuery, error)

func makeABCIQueryFunc(c *lrpc.Client) rpcABCIQueryFunc {
	return func(ctx *rpctypes.Context, path string, data bytes.HexBytes) (*ctypes.ResultABCIQuery, error) {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
		return c.ABCIQuery(path, data)
	}
}

<<<<<<< HEAD
func makeABCIInfoFunc(c *lrpc.Client) func(ctx *rpctypes.Context) (*ctypes.ResultABCIInfo, error) {
=======
type rpcABCIInfoFunc func(ctx *rpctypes.Context) (*ctypes.ResultABCIInfo, error)

func makeABCIInfoFunc(c *lrpc.Client) rpcABCIInfoFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context) (*ctypes.ResultABCIInfo, error) {
		return c.ABCIInfo()
	}
}

<<<<<<< HEAD
func makeBroadcastEvidenceFunc(c *lrpc.Client) func(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
=======
type rpcBroadcastEvidenceFunc func(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error)

// nolint: interfacer
func makeBroadcastEvidenceFunc(c *lrpc.Client) rpcBroadcastEvidenceFunc {
>>>>>>> df3eee455c9d2a4a9698a35aa0dfe6d5d2efd53d
	return func(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
		return c.BroadcastEvidence(ev)
	}
}
