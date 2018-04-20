package main

import (
	"context"
	"fmt"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/renderer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logLevel string

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {
	var source string
	rootCmd := &cobra.Command{
		Use:   "libasciidoc",
		Short: "libasciidoc is a tool to generate an html output from an asciidoc file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flag("source").Value.String() == "" {
				return fmt.Errorf("flag 'source' is required")
			}
			source := cmd.Flag("source").Value.String()
			_, err := libasciidoc.ConvertFileToHTML(context.Background(), source, cmd.OutOrStdout(), renderer.IncludeHeaderFooter(true)) //renderer.IncludeHeaderFooter(true)
			if err != nil {
				return err
			}
			return nil
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			lvl, err := log.ParseLevel(logLevel)
			if err != nil {
				log.Fatalf("unable to parse log level %v", err)
			}
			log.Debug("Setting log level to %v", lvl)
			log.SetLevel(lvl)
		},
	}
	flags := rootCmd.Flags()
	flags.StringVarP(&source, "source", "s", "", "the path to the asciidoc source to process")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log", "error", "log level to set {debug, info, warning, error, fatal, panic}")
	return rootCmd
}
