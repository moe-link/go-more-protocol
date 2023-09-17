package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "MoeLink",
	Short: "MoeLink is a multi-protocol service.",
	Long: `MoeLink is a multi-protocol service that supports common video streaming protocols, 
			communication protocols, and IoT protocols, among others.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run moe-link...")
	},
	//Run: func(cmd *cobra.Command, args []string) {
	//	Error(cmd, args, errors.New("unrecognized command"))
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
