package execution

import (
	// "math/big"

	// Common "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core"
	// "github.com/ethereum/go-ethereum/core/state"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/core/vm"
	// "github.com/ethereum/go-ethereum/ethdb"

	// //"github.com/ethereum/go-ethereum/ethdb/memorydb"
	// "github.com/ethereum/go-ethereum/params"
	// //"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/core/rawdb"
	// "github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/triedb"
	// "github.com/ethereum/go-ethereum/rlp"

	// "github.com/18aaddy/selene-practice/common"


// type B256 = Common.Hash
// type U256 = big.Int
// type HeaderReader interface {
// 	GetHeader(hash B256, number uint64) *types.Header
// }
// type Evm struct {
// 	execution *ExecutionClient
// 	chainID   uint64
// 	tag       common.BlockTag
// }

// func NewEvm(execution *ExecutionClient, chainID uint64, tag common.BlockTag) *Evm {
// 	return &Evm{
// 		execution: execution,
// 		chainID:   chainID,
// 		tag:       tag,
// 	}
// }
// func (e *Evm) CallInner(opts *CallOpts) ([]byte, error) {
// 	txContext := vm.TxContext{
// 		Origin:   *opts.From,
// 		GasPrice: opts.GasPrice,
// 	}
// 	tag := e.tag
// 	block, err := e.execution.GetBlock(tag, false)
// 	if err != nil {
// 		return nil, err
// 	}
// 	blockContext := vm.BlockContext{
// 		CanTransfer: core.CanTransfer,
// 		Transfer:    core.Transfer,
// 		GetHash: func(n uint64) B256 {
// 			return B256{} // You might want to implement this properly
// 		},
// 		Coinbase:    block.Miner.Addr,
// 		BlockNumber: new(U256).SetUint64(block.Number),
// 		Time:        block.Timestamp,
// 		Difficulty:  block.Difficulty.ToBig(),
// 		GasLimit:    block.GasLimit,
// 		BaseFee:     block.BaseFeePerGas.ToBig(),
// 	}
// 	db := rawdb.NewMemoryDatabase()
// 	tdb := triedb.NewDatabase(db, nil)
// 	sdb := state.NewDatabase(tdb, nil)
// 	//root:= trie.NewSecure(common.Hash{}, trie.NewDatabase(sdb))
// 	state, err := state.New(types.EmptyRootHash, sdb)
// 	//witness:=stateless.NewWitness(block,)
// 	//state.StartPrefetcher("hello",witness)
// 	// Create a new vm object
// 	var chainConfig *params.ChainConfig
// 	chainID := e.chainID
// 	switch int64(chainID) {
// 	case MainnetID:
// 		chainConfig = params.MainnetChainConfig
// 	case HoleskyID:
// 		chainConfig = params.HoleskyChainConfig
// 	case SepoliaID:
// 		chainConfig = params.SepoliaChainConfig
// 	case LocalDevID:
// 		chainConfig = params.AllEthashProtocolChanges
// 	default:
// 		// Handle unknown chain ID
// 		chainConfig = nil
// 	}
// 	//Note other chainids not implemented(local testing)
// 	//	"github.com/ethereum/go-ethereum/params"

// 	config := vm.Config{}
// 	//Prefetch database:
// 	var witness *stateless.Witness
// 	var reader BlockHeaderReader
// 	var header types.Header
// 	witness, err = stateless.NewWitness(&header, &reader)
// 	state.StartPrefetcher("evm", witness)
// 	evm := vm.NewEVM(blockContext, txContext, state, chainConfig, config)

// 	/*
// 	   	msg := core.Message{
// 	           From:     *opts.From,
// 	           To:       opts.To,
// 	           Value:    opts.Value,
// 	           GasLimit: opts.Gas.Uint64(),
// 	           GasPrice: opts.GasPrice,
// 	           Data:     opts.Data,
// 	       }
// 	   	gp := new(core.GasPool).AddGas(opts.Gas.Uint64())
// 	   	for {
// 	           // Check if state needs update (you might need to implement this method)
// 	           if needsUpdate, err := e.stateNeedsUpdate(state); err != nil {
// 	               return nil, err
// 	           } else if needsUpdate {
// 	               if err := e.updateState(state); err != nil {
// 	                   return nil, err
// 	               }
// 	           }

// 	           // Apply the message
// 	           ret, err := core.ApplyMessage(evm, msg, gp)
// 	           if err != nil {
// 	               return nil, err
// 	           }



// 	           // If the transaction failed, we continue the loop to retry
// 	           // You might want to add a max retry count to prevent infinite loops
// 	       }*/
// }


// type BlockHeaderReader struct {
//     db ethdb.Database // LevelDB or other database
// }

// func (m *BlockHeaderReader) GetHeader(hash Common.Hash, number uint64) *types.Header {
//     // Use the database to retrieve the header by hash or number
//     key := makeKeyForHeader(hash, number) // Implement key generation logic based on your storage strategy
//     data, err := m.db.Get(key)
//     if err != nil {
//         return nil // Handle error appropriately, e.g., return nil if not found
//     }

//     var header types.Header
//     if err := rlp.DecodeBytes(data, &header); err != nil {
//         return nil // Handle decoding error
//     }
    
//     return &header
// }

// // Example function to create a key for the header (implementation depends on your design)
// func makeKeyForHeader(hash Common.Hash, number uint64) []byte {
//     // Create a key based on the hash and number
//     // This is a placeholder; you need to define your own logic
//     return append(hash[:], byte(number))
// }

// func (e *Evm) Call(opts *CallOpts) ([]byte, error) {}

// type stateDatabase struct {
// 	db ethdb.Database
// }
// type ResultAndState struct {
// 	Result ExecutionResult
// 	State
// }


	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	// "github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)


type EVM struct {
	execution *ExecutionClient
	chainID   *big.Int
	tag       BlockTag
	statedb   *state.StateDB
}

type BlockTag string

const (
	Latest  BlockTag = "latest"
	Pending BlockTag = "pending"
	// Add other block tags as needed
)

func NewEVM(execution *ExecutionClient, chainID *big.Int, tag BlockTag) (*EVM, error) {
	// Initialize stateDB
	db := rawdb.NewMemoryDatabase()
	tdb := triedb.NewDatabase(db, nil)
	sdb := state.NewDatabase(tdb, nil)
	// db := ethdb.New()
	statedb, err := state.New(common.Hash{}, sdb)
	if err != nil {
		return nil, err
	}

	return &EVM{
		execution: execution,
		chainID:   chainID,
		tag:       tag,
		statedb:   statedb,
	}, nil
}

func (e *EVM) Call(tx *types.Transaction) ([]byte, error) {
	msg, err := tx.AsMessage(types.NewEIP155Signer(e.chainID), nil)
	if err != nil {
		return nil, err
	}

	// Create EVM context
	blockchain := e.execution.getBlockchain() // Implement this method in ExecutionClient
	header := blockchain.CurrentHeader()
	
	vmConfig := vm.Config{}
	evmContext := core.NewEVMContext(msg, header, blockchain, nil)
	evm := vm.NewEVM(evmContext, e.statedb, params.MainnetChainConfig, vmConfig)

	// Set gas limit to max uint64 if not specified
	gasLimit := uint64(tx.Gas())
	if gasLimit == 0 {
		gasLimit = uint64(0xffffffffffffffff)
	}

	// Execute the transaction
	ret, _, err := core.ApplyMessage(evm, msg, new(core.GasPool).AddGas(gasLimit))
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (e *EVM) EstimateGas(tx *types.Transaction) (uint64, error) {
	msg, err := tx.AsMessage(types.NewEIP155Signer(e.chainID), nil)
	if err != nil {
		return 0, err
	}

	// Create EVM context
	blockchain := e.execution.getBlockchain() // Implement this method in ExecutionClient
	header := blockchain.CurrentHeader()
	
	vmConfig := vm.Config{}
	evmContext := core.NewEVMContext(msg, header, blockchain, nil)
	evm := vm.NewEVM(evmContext, e.statedb, params.MainnetChainConfig, vmConfig)

	// Set gas limit to max uint64
	gasLimit := uint64(0xffffffffffffffff)

	// Execute the transaction
	_, leftOverGas, err := core.ApplyMessage(evm, msg, new(core.GasPool).AddGas(gasLimit))
	if err != nil {
		return 0, err
	}

	return gasLimit - leftOverGas, nil
}

func (e *EVM) prefetchState(tx *types.Transaction) error {
	// Implement state prefetching logic
	// This might involve populating the statedb with necessary account data
	// You may need to use e.execution.rpc to fetch data from the blockchain
	return nil
}

// Custom error types
var (
	ErrRevert = errors.New("execution reverted")
	ErrOutOfGas = errors.New("out of gas")
	// Add other custom error types as needed
)

// Implement other necessary methods and helpers
// ...

func (ec *ExecutionClient) getBlockchain() *core.BlockChain {
	// Implement this method to return a *core.BlockChain instance
	// This might involve creating a simulated backend or connecting to a real node
	panic("not implemented")
}