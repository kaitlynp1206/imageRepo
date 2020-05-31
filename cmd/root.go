package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imageRepo",
	Short: ":)",
	Long:  `#TODO :)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("HELLO`")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
