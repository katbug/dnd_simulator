package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "D&DSim",
	Short: "D&DSim is a d&d battle simulator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
