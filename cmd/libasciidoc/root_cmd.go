package main

import (
	"context"
	"os"
	"strings"

	"path/filepath"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/renderer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {
	var noHeaderFooter bool
	var logLevel string
	rootCmd := &cobra.Command{
		Use: "libasciidoc FILE...",
		Short: `libasciidoc is a tool to generate an html output from an asciidoc file

Positional args:
If no files are specified, input is read from STDIN
If more than 1 file is specified, then output is written to ".html" file alongside the source file
`,
		Args: cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if len(args) == 0 {
				_, err = libasciidoc.ConvertToHTML(context.Background(), os.Stdin, cmd.OutOrStdout(), renderer.IncludeHeaderFooter(!noHeaderFooter))
			} else if len(args) == 1 {
				_, err = libasciidoc.ConvertFileToHTML(context.Background(), args[0], cmd.OutOrStdout(), renderer.IncludeHeaderFooter(!noHeaderFooter)) //renderer.IncludeHeaderFooter(true)
			} else {
				for _, source := range args {
					path, _ := filepath.Abs(source)
					log.Debugf("Starting to process file %v", path)
					outname := strings.TrimSuffix(path, filepath.Ext(path)) + ".html"
					outfile, e := os.Create(outname)
					if e != nil {
						log.Warnf("Cannot create output file - %v, skipping", outname)
						continue
					}
					defer func() {
						e = outfile.Close()
						if e != nil {
							log.Errorf("Cannot close output file: %v", outname)
						}
					}()
					_, e = libasciidoc.ConvertFileToHTML(context.Background(), source, outfile, renderer.IncludeHeaderFooter(!noHeaderFooter)) //renderer.IncludeHeaderFooter(true)
					if e == nil {
						log.Infof("File %v created", outname)
					} else {
						err = e
					}
				}
			}
			return err
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			lvl, err := log.ParseLevel(logLevel)
			if err != nil {
				log.Errorf("unable to parse log level %v", err)
				return err
			}
			log.Debug("Setting log level to %v", lvl)
			log.SetLevel(lvl)
			return nil
		},
	}
	flags := rootCmd.Flags()
	flags.BoolVarP(&noHeaderFooter, "no-header-footer", "s", false, "Do not render header/footer (Default: false)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log", "warning", "log level to set {debug, info, warning, error, fatal, panic}")
	return rootCmd
}
