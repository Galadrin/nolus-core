package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	wasmcli "github.com/vmierzhev/custom-wasmd/x/wasm/client/cli"
)

func AddGenesisWasmMsgCmd(defaultNodeHome string) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "add-wasm-genesis-message",
		Short:                      "Wasm genesis subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	genesisIO := wasmcli.NewDefaultGenesisIO()
	txCmd.AddCommand(
		wasmcli.GenesisStoreCodeCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisInstantiateContractCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisExecuteContractCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisListContractsCmd(defaultNodeHome, genesisIO),
		wasmcli.GenesisListCodesCmd(defaultNodeHome, genesisIO),
		wasmcli.ProposalMigrateContractCmd(),
		wasmcli.MigrateContractCmd(),
		wasmcli.ClearContractAdminCmd(),
		wasmcli.ProposalClearContractAdminCmd(),
		wasmcli.ProposalExecuteContractCmd(),
		wasmcli.ProposalInstantiateContractCmd(),
		wasmcli.ProposalStoreCodeCmd(),
		wasmcli.ProposalPinCodesCmd(),
		wasmcli.ExecuteContractCmd(),
		wasmcli.ProposalUpdateContractAdminCmd(),
	)
	return txCmd
}
