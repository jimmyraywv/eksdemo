package resource

import (
	"eksdemo/pkg/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

type Resource struct {
	cmd.Command
	cmd.Flags
	Options

	Manager
}

func (r *Resource) NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     r.Command.Name + " NAME",
		Short:   r.Description,
		Long:    "Create " + r.Description,
		Aliases: r.Aliases,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := r.Flags.ValidateFlags(); err != nil {
				return err
			}

			if err := r.Options.Validate(); err != nil {
				return err
			}

			cmd.SilenceUsage = true
			r.SetName(args[0])

			if err := r.PreCreate(); err != nil {
				return err
			}

			if r.Manager == nil {
				return fmt.Errorf("feature not yet implemented")
			}

			return r.Create(r.Options)
		},
	}
	r.Flags = r.Options.AddCreateFlags(cmd, r.Flags)

	return cmd
}

func (r *Resource) NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     r.Command.Name + " NAME",
		Short:   r.Description,
		Long:    "Delete " + r.Description,
		Aliases: r.Aliases,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := r.Flags.ValidateFlags(); err != nil {
				return err
			}
			cmd.SilenceUsage = true

			r.SetName(args[0])

			if r.Manager == nil {
				return fmt.Errorf("feature not yet implemented")
			}

			return r.Delete(r.Options)
		},
	}
	r.Flags = r.Options.AddDeleteFlags(cmd, r.Flags)

	return cmd
}
