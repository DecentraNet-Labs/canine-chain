package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jackal-dao/canine/x/lp/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group lp queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListLPool())
	cmd.AddCommand(CmdShowLPool())
	cmd.AddCommand(CmdShowLProviderRecord())
	cmd.AddCommand(CmdEstimateSwapOut())

	cmd.AddCommand(CmdEstimateSwapIn())

	cmd.AddCommand(CmdEstimateContribution())

	cmd.AddCommand(CmdMakeValidPair())

	cmd.AddCommand(CmdEstimatePoolRemove())

	cmd.AddCommand(CmdListRecordsFromPool())

	// this line is used by starport scaffolding # 1

	return cmd
}
