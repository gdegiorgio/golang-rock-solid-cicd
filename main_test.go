package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestHelloCLI(t *testing.T) {

	var buf bytes.Buffer
	var rootCmd = &cobra.Command{
		Use:   "hello-cli",
		Short: "A simple CLI that says hello",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Hello, world!")
		},
	}

	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{}) // no arguments

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	output := buf.String()
	expected := "Hello, world!\n"

	if strings.TrimSpace(output) != strings.TrimSpace(expected) {
		t.Errorf("unexpected output:\n  got:  %q\n  want: %q", output, expected)
	}
}
