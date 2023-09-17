package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "Git is a distributed version control system.",
	Long: `Git is a free and open source distributed version control system
designed to handle everything from small to very large projects 
with speed and efficiency.`,
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
