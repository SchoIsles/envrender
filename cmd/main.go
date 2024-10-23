package main

import (
	"fmt"

	"github.com/SchoIsles/envtpl"
)

func main() {
	tpl := `
echo PATH: {{ .PATH }}
current: {{ now | date "2006-01-02" }}
`
	res, err := envtpl.Render(tpl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
