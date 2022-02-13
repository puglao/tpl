/*
Copyright © 2022 puglao <puglao@cloudemo.site>

*/

package engine

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/puglao/tpl/internal"
	"io/ioutil"
	"log"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"text/template"
)

type valueData struct {
	Values interface{}
}

func GenerateTpl(cfg *internal.Config, out *bytes.Buffer) {
	data, err := ioutil.ReadFile(cfg.ValFile)
	if err != nil {
		log.Fatal(err)
	}
	v := new(valueData)
	yaml.Unmarshal(data, &v.Values)
	tmpl, err := template.New("test").Funcs(sprig.TxtFuncMap()).ParseFiles(cfg.TplFile)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(out, filepath.Base(cfg.TplFile), v)
	if err != nil {
		log.Fatal(err)
	}
	return
}
