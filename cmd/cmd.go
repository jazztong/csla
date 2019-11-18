package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csla",
	Short: "csla is a project template generator to Create Serverless Application",
	Long: `A simple and easy understand project template generator focus in building Serverless Application
	using terraform as backbone of intra as a code`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You seen new to this tools, you can type `csla -h` to learn more")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newAppCmd)
}

// Execute initialize the commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
