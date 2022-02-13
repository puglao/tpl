/*
Copyright © 2022 puglao <puglao@cloudemo.site>

*/
package cmd

import (
	"bytes"
	"github.com/puglao/tpl/internal"
	"github.com/puglao/tpl/internal/engine"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var outputFile string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long:  `A brief description of your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !internal.IsFileExist(cfg.TplFile) {
			log.Fatalf("%s is not exist. Run `tpl init` to generate it.\n", cfg.TplFile)
		}
		if !internal.IsFileExist(cfg.ValFile) {
			log.Fatalf("%s is not exist. Run `tpl init` to generate it.\n", cfg.ValFile)
		}
		var buff bytes.Buffer
		engine.GenerateTpl(cfg, &buff)
		if outputFile == "" {
			buff.WriteTo(os.Stdout)
		} else {
			internal.CreateFile(outputFile, buff.Bytes())
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output File name, output to stdout if omitted")
}
