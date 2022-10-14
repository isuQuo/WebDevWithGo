package views

import (
	"html/template"
	"io/fs"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		panic(err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}
