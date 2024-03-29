package cmd

import (
	"eksdemo/pkg/resource/cluster"
	"eksdemo/pkg/resource/irsa"
	"eksdemo/pkg/resource/nodegroup"

	"github.com/spf13/cobra"
)

func newCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete resource(s)",
	}

	// Don't show flag errors for delete without a subcommand
	cmd.DisableFlagParsing = true

	cmd.AddCommand(cluster.NewResource().NewDeleteCmd())
	cmd.AddCommand(irsa.NewResource().NewDeleteCmd())
	cmd.AddCommand(nodegroup.NewResource().NewDeleteCmd())

	return cmd
}
