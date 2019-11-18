package cmd

import (
	"fmt"

	"github.com/jazztong/csla/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of csla",
	Long:  `All software has versions. This is csla`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version.Version)
	},
}
