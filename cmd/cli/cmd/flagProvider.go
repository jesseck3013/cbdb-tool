package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CobraFlagProvider struct {
	cmd *cobra.Command
}

func (c CobraFlagProvider) GetBool(name string) bool {
	v, err := c.cmd.Flags().GetBool(name)
	if err != nil {
		err := fmt.Errorf("Developer Error: Register an undefined flag %q: %v", name, err)
		panic(err)
	}
	return v
}
