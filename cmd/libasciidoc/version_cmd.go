package main

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc"
	"github.com/spf13/cobra"
)

// NewVersionCmd returns the root command
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and build info",
		Run: func(cmd *cobra.Command, args []string) {
			if libasciidoc.BuildTag != "" {
				fmt.Fprintf(cmd.OutOrStdout(), "version:    %s\n", libasciidoc.BuildTag)
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "commit:     %s\n", libasciidoc.BuildCommit)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "build time: %s\n", libasciidoc.BuildTime)
		},
	}
}
