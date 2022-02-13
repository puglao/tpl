/*
Copyright © 2022 puglao <puglao@cloudemo.site>

*/
package cmd

import (
	"github.com/puglao/tpl/internal"
	"github.com/spf13/cobra"
	"log"
)

var defaultTpl = `{{ .Values.doc }}
{{- range .Values.memberlist }}
({{ .id }}) {{with .personalInfo }}{{ .name }}, age {{ .age }}
{{ if .smoke }}
{{ if eq .group "adult" }}
audlt
{{ else if eq .group "elder" }}
oldman {{- end }} drinks {{ range .drinks -}}
{{ .}},
{{- end }}
{{- end }}
{{- end }}
{{- end }}
`

var defaultVal = `doc: memberlist
memberlist:
  - id: 00001
    personalInfo:
      name: doggo
      group: "adult"
      age: 7
      drinks:
      - water
      - beer
      - coffee
      smoke: false
  - id: 00002
    personalInfo:
      name: kitty
      group: elder
      age: 6
      drinks:
      - juice
      - milk
      smoke: true
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize template and values file",
	Long:  `Initialize template and values file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if internal.IsFileExist(cfg.TplFile) {
			log.Fatalf("Detect %s exist. Stop initalizing template.\n", cfg.TplFile)
		}
		if internal.IsFileExist(cfg.ValFile) {
			log.Fatalf("Detect %s exist. Stop initalizing template.\n", cfg.ValFile)
		}
		// var data []byte
		// data, _ = f.ReadFile("template/template.tpl")
		internal.CreateFile(cfg.TplFile, []byte(defaultTpl))
		// data, _ = f.ReadFile("template/values.yaml")
		internal.CreateFile(cfg.ValFile, []byte(defaultVal))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
