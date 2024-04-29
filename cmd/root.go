package cmd

import (
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = cobra.Command{
	Use: "stori",
	Run: func(cmd *cobra.Command, args []string) {
		settlement()
	},
}

func RootCmd() *cobra.Command {
	rootCmd.AddCommand(&migrateCmd, &settlementCmd, &templateCmd)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file to use. Defaults to .config.yml")

	return &rootCmd
}
