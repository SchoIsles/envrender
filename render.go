package envtpl

import (
	"bytes"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

type renderer struct {
	values map[string]interface{}
}

func NewRenderer() *renderer {
	in := make(map[string]interface{})
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		in[pair[0]] = pair[1]
	}
	return &renderer{values: in}
}

func (r *renderer) render(w io.Writer, tpl string) error {
	t := template.New("gotpl")
	t.Funcs(sprig.TxtFuncMap())
	tmpl, err := t.Parse(tpl)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, r.values); err != nil {
		return err
	}
	return nil
}

func Render(tpl string) (string, error) {
	var buf bytes.Buffer
	if err := RenderWithWriter(&buf, tpl); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderFile(filePath string) (string, error) {
	var buf bytes.Buffer
	if err := RenderFileWithWriter(&buf, filePath); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderWithWriter(w io.Writer, tpl string) error {
	if err := NewRenderer().render(w, tpl); err != nil {
		return err
	}
	return nil
}

func RenderFileWithWriter(w io.Writer, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = RenderWithWriter(w, string(content)); err != nil {
		return err
	}
	return nil
}
