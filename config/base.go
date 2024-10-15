package config

import "github.com/18aaddy/selene-practics/consensus/consensus_core"

// base config for a network

type BaseConfig struct {
	RpcBindIp            string               `json:"rpc_bind_ip"`
	RpcPort              uint16               `json:"rpc_port"`
	ConsensusRpc         *string              `json:"consensus_rpc"`
	DefaultCheckpoint    [32]byte             `json:"default_checkpoint"` // In cli.go, checkpoint is currently taken as []byte{}
	Chain                ChainConfig          `json:"chain"`              // but it should be [32]byte as it is a hash
	Forks                consensus_core.Forks `json:"forks"`
	MaxCheckpointAge     uint64               `json:"max_checkpoint_age"`
	DataDir              *string              `json:"data_dir"`
	LoadExternalFallback bool                 `json:"load_external_fallback"`
	StrictCheckpointAge  bool                 `json:"strict_checkpoint_age"`
}

// implement a default method for the above struct
func (b BaseConfig) Default() BaseConfig {
	LOCALHOST := "127.0.0.1" // Using 127.0.0.1 as default host ip address
	return BaseConfig{
		RpcBindIp:            LOCALHOST,
		RpcPort:              0,
		ConsensusRpc:         nil,
		DefaultCheckpoint:    [32]byte{},
		Chain:                ChainConfig{},
		Forks:                consensus_core.Forks{},
		MaxCheckpointAge:     0,
		DataDir:              nil,
		LoadExternalFallback: false,
		StrictCheckpointAge:  false,
	}
}
