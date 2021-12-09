package constants

import (
	proto "github.com/Ankr-network/ankrscan-proto-contract"
	"github.com/pkg/errors"

)

func AllBlockchainNames() []string {
	names := make([]string, 0)
	for _, chain := range blockchains {
		names = append(names, chain.BlockchainName)
	}
	return names
}

func AllBlockchains() []*proto.BlockchainConstants {
	return blockchains
}

func Blockchain(name string) (*proto.BlockchainConstants, error) {
	var result *proto.BlockchainConstants
	for _, blockchain := range blockchains {
		if blockchain.BlockchainName == name {
			result = blockchain
		}
	}
	if result == nil {
		return nil, errors.Errorf("no blockchain with name %s", name)
	}
	return result, nil
}

var blockchains = []*proto.BlockchainConstants{
	{
		BlockchainName: "ARBITRUM",
		VerboseName:    "Arbitrum",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "ETH",
		Decimals:       18,
	},
	{
		BlockchainName: "AVAXc",
		VerboseName:    "Avalanche C-chain",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "AVAX",
		Decimals:       18,
	},
	{
		BlockchainName: "BSC",
		VerboseName:    "Binance Smart Chain",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "BNB",
		Decimals:       18,
	},
	{
		BlockchainName: "DOT",
		VerboseName:    "Polkadot",
		ChainType:      proto.ChainType_CHAIN_TYPE_POLKADOT,
		Symbol:         "DOT",
		Decimals:       10,
	},
	{
		BlockchainName: "ETH",
		VerboseName:    "Ethereum",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "ETH",
		Decimals:       18,
	},
	{
		BlockchainName: "FTM",
		VerboseName:    "Fantom",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "FTM",
		Decimals:       18,
	},
	{
		BlockchainName: "HT",
		VerboseName:    "Heco",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "HT",
		Decimals:       18,
	},
	{
		BlockchainName: "KSM",
		VerboseName:    "Kusama",
		ChainType:      proto.ChainType_CHAIN_TYPE_POLKADOT,
		Symbol:         "KSM",
		Decimals:       12,
	},
	{
		BlockchainName: "MATIC",
		VerboseName:    "Polygon",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "MATIC",
		Decimals:       18,
	},
	{
		BlockchainName: "OKEX",
		VerboseName:    "OKEx Chain",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "OKT",
		Decimals:       18,
	},
	{
		BlockchainName: "xDAI",
		VerboseName:    "xDai Stable Chain",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "xDAI",
		Decimals:       18,
	},
}
