package injective_rpc

import (
	"fmt"

	"testing"

	"cosmosmonitor/rpc"
	"cosmosmonitor/types"
	base "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	distribution "cosmossdk.io/api/cosmos/distribution/v1beta1"
	gov "cosmossdk.io/api/cosmos/gov/v1beta1"
	staking "cosmossdk.io/api/cosmos/staking/v1beta1"
)

func TestGetValInfo(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("injective-grpc.polkachu.com:14390")
	if err != nil {
		logger.Error("Failed to create injective gRPC client, err:", err)
	}

	stakingQueryCli := staking.NewQueryClient(grpcConn)
	govQueryCli := gov.NewQueryClient(grpcConn)
	baseCli := base.NewServiceClient(grpcConn)
	distributionCli := distribution.NewQueryClient(grpcConn)

	cc := &InjectiveCli{
		ChainCli: &rpc.ChainCli{
			StakingQueryCli: stakingQueryCli,
			GovQueryCli:     govQueryCli,
			BaseQuaryCli:    baseCli,
			DistributionCli: distributionCli,
		},
	}
	monitorObj := make([]string, 0)
	monitorObj = append(monitorObj, "injvaloper1g4d6dmvnpg7w7yugy6kplndp7jpfmf3krtschp")
	monitors, _ := cc.GetValInfo(monitorObj)
	for _, monitor := range monitors {
		fmt.Println("monitor:", monitor)
	}
}

func TestGetProposal(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("injective-grpc.polkachu.com:14390")
	if err != nil {
		logger.Error("Failed to create injective gRPC client, err:", err)
	}

	stakingQueryCli := staking.NewQueryClient(grpcConn)
	govQueryCli := gov.NewQueryClient(grpcConn)
	baseCli := base.NewServiceClient(grpcConn)
	distributionCli := distribution.NewQueryClient(grpcConn)

	cc := &InjectiveCli{
		ChainCli: &rpc.ChainCli{
			StakingQueryCli: stakingQueryCli,
			GovQueryCli:     govQueryCli,
			BaseQuaryCli:    baseCli,
			DistributionCli: distributionCli,
		},
	}

	monitorObjs := make([]*types.MonitorObj, 0)
	monitorObjs = append(monitorObjs, &types.MonitorObj{
		Moniker:         "Figment",
		OperatorAddr:    "injvaloper1g4d6dmvnpg7w7yugy6kplndp7jpfmf3krtschp",
		OperatorAddrHex: "6087607e1e56f6ee7934abaf65834c92d618104c",
		SelfStakeAddr:   "inj1g4d6dmvnpg7w7yugy6kplndp7jpfmf3k5d9ak9",
	})
	monitors, _ := cc.GetProposal(monitorObjs)
	for _, monitor := range monitors {
		fmt.Println("proposal:", monitor)
	}
}

func TestGetValPerformance(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("injective-grpc.polkachu.com:14390")
	if err != nil {
		logger.Error("Failed to create injective gRPC client, err:", err)
	}

	stakingQueryCli := staking.NewQueryClient(grpcConn)
	govQueryCli := gov.NewQueryClient(grpcConn)
	baseCli := base.NewServiceClient(grpcConn)
	distributionCli := distribution.NewQueryClient(grpcConn)

	cc := &InjectiveCli{
		ChainCli: &rpc.ChainCli{
			StakingQueryCli: stakingQueryCli,
			GovQueryCli:     govQueryCli,
			BaseQuaryCli:    baseCli,
			DistributionCli: distributionCli,
		},
	}

	monitorObjs := make([]*types.MonitorObj, 0)
	monitorObjs = append(monitorObjs, &types.MonitorObj{
		Moniker:         "Figment",
		OperatorAddr:    "injvaloper1g4d6dmvnpg7w7yugy6kplndp7jpfmf3krtschp",
		OperatorAddrHex: "6087607e1e56f6ee7934abaf65834c92d618104c",
		SelfStakeAddr:   "inj1g4d6dmvnpg7w7yugy6kplndp7jpfmf3k5d9ak9",
	})
	proposalAssignment, valSign, valSignMissed, _ := cc.GetValPerformance(19163489, monitorObjs)
	for _, monitor := range proposalAssignment {
		fmt.Println("proposalAssignment:", monitor)
	}

	for _, monitor := range valSign {
		fmt.Println("valSign:", monitor)
	}

	for _, monitor := range valSignMissed {
		fmt.Println("valSignMissed:", monitor)
	}
}
