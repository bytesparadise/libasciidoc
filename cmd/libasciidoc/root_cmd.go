package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"path/filepath"

	"github.com/bytesparadise/libasciidoc"
	logsupport "github.com/bytesparadise/libasciidoc/pkg/log"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {

	var noHeaderFooter bool
	var outputName string
	var logLevel string

	rootCmd := &cobra.Command{
		Use:   "libasciidoc [flags] FILE",
		Short: `libasciidoc is a tool to convert from Asciidoc to HTML`,
		Args:  cobra.ArbitraryArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			lvl, err := log.ParseLevel(logLevel)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "unable to parse log level '%v'", logLevel)
				return err
			}
			logsupport.Setup()
			log.SetLevel(lvl)
			log.SetOutput(cmd.OutOrStdout())
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helpCommand.RunE(cmd, args)
			}
			for _, source := range args {
				out, close := getOut(cmd, source, outputName)
				if out != nil {
					defer close()
					path, _ := filepath.Abs(source)
					log.Debugf("Starting to process file %v", path)
					_, err := libasciidoc.ConvertFileToHTML(context.Background(), source, out, renderer.IncludeHeaderFooter(!noHeaderFooter)) //renderer.IncludeHeaderFooter(true)
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	flags := rootCmd.Flags()
	flags.BoolVarP(&noHeaderFooter, "no-header-footer", "s", false, "do not render header/footer (default: false)")
	flags.StringVarP(&outputName, "out-file", "o", "", "output file (default: based on path of input file); use - to output to STDOUT")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log", "warning", "log level to set [debug|info|warning|error|fatal|panic]")
	return rootCmd
}

type closeFunc func() error

func defaultCloseFunc() closeFunc {
	return func() error { return nil }
}

func newCloseFileFunc(c io.Closer) closeFunc {
	return func() error {
		return c.Close()
	}
}

func getOut(cmd *cobra.Command, source, outputName string) (io.Writer, closeFunc) {
	if outputName == "-" {
		// outfile is STDOUT
		return cmd.OutOrStdout(), defaultCloseFunc()
	} else if outputName != "" {
		// outfile is specified in the command line
		outfile, e := os.Create(outputName)
		if e != nil {
			log.Warnf("Cannot create output file - %v, skipping", outputName)
		}
		return outfile, newCloseFileFunc(outfile)
	} else if source != "" {
		// outfile is based on source
		path, _ := filepath.Abs(source)
		outname := strings.TrimSuffix(path, filepath.Ext(path)) + ".html"
		outfile, err := os.Create(outname)
		if err != nil {
			log.Warnf("Cannot create output file - %v, skipping", outname)
			return nil, nil
		}
		return outfile, newCloseFileFunc(outfile)
	}
	return cmd.OutOrStdout(), defaultCloseFunc()
}
