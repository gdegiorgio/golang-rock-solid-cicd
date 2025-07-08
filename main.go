package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello-cli",
	Short: "A simple CLI that says hello",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Hello, world!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
