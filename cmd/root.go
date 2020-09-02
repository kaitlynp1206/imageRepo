package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imageRepo",
	Short: "Image Repository Service",
	Long:  `Image Repository Service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Welcome to the Image Repository Service")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
