package execution

import (
	"math/big"

	Common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	//"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/params"
	//"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/triedb"	
	"github.com/ethereum/go-ethereum/core/stateless"

	"github.com/BlocSoc-iitr/selene/common"
)
type B256 = Common.Hash
type U256 = big.Int

type Evm struct {
	execution *ExecutionClient
	chainID   uint64
	tag       common.BlockTag
}
func NewEvm(execution *ExecutionClient, chainID uint64, tag common.BlockTag) *Evm {
	return &Evm{
		execution: execution,
		chainID:   chainID,
		tag:       tag,
	}
}

// TODO: Call and EstimateGas function for Evm struct
func (e *Evm) Call(opts *CallOpts) ([]byte,error){	}


func (e *Evm) CallInner(opts *CallOpts) ([]byte, error) {
	txContext := vm.TxContext{
		Origin:   *opts.From,
		GasPrice: opts.GasPrice,
	}
	tag:= e.tag
	block, err := e.execution.GetBlock(tag, false)
	if err != nil {
		return nil, err
	}

	blockContext := vm.BlockContext{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash: func(n uint64) B256 {
			return B256{} // You might want to implement this properly
		},
		Coinbase:    Common.BytesToAddress(block.Miner.Addr[:]),
		BlockNumber: new(U256).SetUint64(block.Number),
		Time:        block.Timestamp,
		Difficulty:  block.Difficulty.ToBig(),
		GasLimit:    block.GasLimit,
		BaseFee:     block.BaseFeePerGas.ToBig(),
	}
	
	db:= rawdb.NewMemoryDatabase()
	tdb:= triedb.NewDatabase(db, nil)
	sdb:= state.NewDatabase(tdb)
	//root:= trie.NewSecure(common.Hash{}, trie.NewDatabase(sdb))
	state, err := state.New(types.EmptyRootHash, sdb)
	witness:=stateless.NewWitness(block,)
	state.StartPrefetcher("hello",witness)
	// Create a new vm object
	var chainConfig *params.ChainConfig
	chainID:=e.chainID
	switch (int64(chainID)) {
		case MainnetID:
			chainConfig = params.MainnetChainConfig
		case HoleskyID:
			chainConfig = params.HoleskyChainConfig
		case SepoliaID:
			chainConfig = params.SepoliaChainConfig
		case LocalDevID:
			chainConfig = params.AllEthashProtocolChanges
		default:
			// Handle unknown chain ID
			chainConfig = nil
		}
		//Note other chainids not implemented(local testing)
		//	"github.com/ethereum/go-ethereum/params"

	config:= vm.Config{}
	evm := vm.NewEVM(blockContext,txContext,state,chainConfig,config)

	msg := core.Message{
        From:     *opts.From,
        To:       opts.To,
        Value:    opts.Value,
        GasLimit: opts.Gas.Uint64(),
        GasPrice: opts.GasPrice,
        Data:     opts.Data,
    }
	gp := new(core.GasPool).AddGas(opts.Gas.Uint64())
	for {
        // Check if state needs update (you might need to implement this method)
        if needsUpdate, err := e.stateNeedsUpdate(state); err != nil {
            return nil, err
        } else if needsUpdate {
            if err := e.updateState(state); err != nil {
                return nil, err
            }
        }

        // Apply the message
        ret, err := core.ApplyMessage(evm, &msg, gp)
        if err != nil {
            return nil, err
        }



        // If the transaction failed, we continue the loop to retry
        // You might want to add a max retry count to prevent infinite loops
    }
}

type stateDatabase struct {
	db ethdb.Database
}
type ResultAndState struct{
	Result ExecutionResult
	State 
}