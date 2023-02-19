package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCosmosCmd)
}

var runCosmosCmd = &cobra.Command{
	Use:   "runCosmos",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := Run(); err != nil {
			return err
		}

		return nil
	},
}
