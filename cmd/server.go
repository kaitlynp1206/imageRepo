package cmd

import (
	"context"
	"fmt"

	"github.com/kaitlynp1206/imageRepo/pkg/server"
	"github.com/spf13/cobra"
)

//func init() {
//	rootCmd.AddCommand(serverCmd)
//	serverCmd.PersistentFlags().StringP()
//}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Server",
	Long:  `Start Server`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(context.Background())
		fmt.Println("Starting")
		s.Start()
	},
}
