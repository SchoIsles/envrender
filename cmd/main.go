package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/SchoIsles/envtpl"
)

type Options struct {
	File string
	Out  string
}

func main() {
	opt := &Options{}
	flags := flag.NewFlagSet("envtpl", flag.ExitOnError)
	flags.StringVar(&opt.File, "f", "", "Input file path")
	flags.StringVar(&opt.Out, "o", "", "Output file path")
	if err := flags.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var output io.Writer
	if opt.Out != "" {
		file, err := os.OpenFile(opt.Out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			exitErr(err)
		}
		output = file
		defer file.Close()
	} else {
		output = os.Stdout
	}

	if opt.File != "" && opt.File != "-" {
		err := envtpl.RenderFileWithWriter(output, opt.File)
		if err != nil {
			exitErr(err)
		}
	} else {
		var input bytes.Buffer
		_, err := io.Copy(&input, os.Stdin)
		if err != nil {
			exitErr(err)
		}
		err = envtpl.RenderWithWriter(output, input.String())
		if err != nil {
			exitErr(err)
		}
	}

}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
