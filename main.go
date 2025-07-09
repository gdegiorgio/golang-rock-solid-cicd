package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var name string

var rootCmd = &cobra.Command{
	Use:   "hello-cli",
	Short: "A simple CLI that says hello",
	Run: func(cmd *cobra.Command, args []string) {
		if name != "" {
			cmd.Printf("Hello, %s!\n", name)
		} else {
			cmd.Println("Hello, world!")
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&name, "name", "", "Custom name to include in greeting")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
